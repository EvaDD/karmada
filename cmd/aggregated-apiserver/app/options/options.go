package options

import (
	"context"
	"fmt"
	"net"

	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/endpoints/openapi"
	"k8s.io/apiserver/pkg/features"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	netutils "k8s.io/utils/net"

	"github.com/karmada-io/karmada/pkg/aggregatedapiserver"
	clusterv1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	clientset "github.com/karmada-io/karmada/pkg/generated/clientset/versioned"
	informers "github.com/karmada-io/karmada/pkg/generated/informers/externalversions"
	generatedopenapi "github.com/karmada-io/karmada/pkg/generated/openapi"
)

const defaultEtcdPathPrefix = "/registry"

// Options contains everything necessary to create and run aggregated-apiserver.
type Options struct {
	RecommendedOptions    *genericoptions.RecommendedOptions
	SharedInformerFactory informers.SharedInformerFactory

	karmadaConfig string
	Master        string
	// KubeAPIQPS is the QPS to use while talking with karmada-apiserver.
	KubeAPIQPS float32
	// KubeAPIBurst is the burst to allow while talking with karmada-apiserver.
	KubeAPIBurst int
}

// NewOptions returns a new Options.
func NewOptions() *Options {
	o := &Options{
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			aggregatedapiserver.Codecs.LegacyCodec(clusterv1alpha1.SchemeGroupVersion)),
	}
	o.RecommendedOptions.Etcd.StorageConfig.EncodeVersioner = runtime.NewMultiGroupVersioner(clusterv1alpha1.SchemeGroupVersion, schema.GroupKind{Group: clusterv1alpha1.GroupName})
	return o
}

// AddFlags adds flags to the specified FlagSet.
func (o *Options) AddFlags(flags *pflag.FlagSet) {
	o.RecommendedOptions.AddFlags(flags)

	flags.StringVar(&o.karmadaConfig, "karmada-config", o.karmadaConfig, "Path to a karmada-apiserver KubeConfig.")
	flags.StringVar(&o.Master, "master", o.Master, "The address of the Karmada API server. Overrides any value in KubeConfig.")
	flags.Float32Var(&o.KubeAPIQPS, "kube-api-qps", 40.0, "QPS to use while talking with karmada-apiserver. Doesn't cover events and node heartbeat apis which rate limiting is controlled by a different set of flags.")
	flags.IntVar(&o.KubeAPIBurst, "kube-api-burst", 60, "Burst to use while talking with karmada-apiserver. Doesn't cover events and node heartbeat apis which rate limiting is controlled by a different set of flags.")
	utilfeature.DefaultMutableFeatureGate.AddFlag(flags)
}

// Complete fills in fields required to have valid data.
func (o *Options) Complete() error {
	return nil
}

// Validate validates Options.
func (o *Options) Validate() error {
	var errs []error
	errs = append(errs, o.RecommendedOptions.Validate()...)
	return utilerrors.NewAggregate(errs)
}

// Run runs the aggregated-apiserver with options. This should never exit.
func (o *Options) Run(ctx context.Context) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	restConfig, err := clientcmd.BuildConfigFromFlags(o.Master, o.karmadaConfig)
	if err != nil {
		return fmt.Errorf("error building kubeconfig: %w", err)
	}
	restConfig.QPS, restConfig.Burst = o.KubeAPIQPS, o.KubeAPIBurst
	kubeClientSet := kubernetes.NewForConfigOrDie(restConfig)

	server, err := config.Complete().New(kubeClientSet)
	if err != nil {
		return err
	}

	server.GenericAPIServer.AddPostStartHookOrDie("start-aggregated-server-informers", func(context genericapiserver.PostStartHookContext) error {
		config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		o.SharedInformerFactory.Start(context.StopCh)
		return nil
	})

	return server.GenericAPIServer.PrepareRun().Run(ctx.Done())
}

// Config returns config for the api server given Options
func (o *Options) Config() (*aggregatedapiserver.Config, error) {
	// TODO have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	o.RecommendedOptions.Etcd.StorageConfig.Paging = utilfeature.DefaultFeatureGate.Enabled(features.APIListChunking)

	o.RecommendedOptions.ExtraAdmissionInitializers = func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
		client, err := clientset.NewForConfig(c.LoopbackClientConfig)
		if err != nil {
			return nil, err
		}
		informerFactory := informers.NewSharedInformerFactory(client, c.LoopbackClientConfig.Timeout)
		o.SharedInformerFactory = informerFactory
		return []admission.PluginInitializer{}, nil
	}

	serverConfig := genericapiserver.NewRecommendedConfig(aggregatedapiserver.Codecs)
	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(generatedopenapi.GetOpenAPIDefinitions, openapi.NewDefinitionNamer(aggregatedapiserver.Scheme))
	serverConfig.OpenAPIConfig.Info.Title = "Karmada"
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &aggregatedapiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   aggregatedapiserver.ExtraConfig{},
	}
	return config, nil
}
