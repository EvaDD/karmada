// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	configv1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	versioned "github.com/karmada-io/karmada/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/karmada-io/karmada/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/karmada-io/karmada/pkg/generated/listers/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceExploringWebhookConfigurationInformer provides access to a shared informer and lister for
// ResourceExploringWebhookConfigurations.
type ResourceExploringWebhookConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceExploringWebhookConfigurationLister
}

type resourceExploringWebhookConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceExploringWebhookConfigurationInformer constructs a new informer for ResourceExploringWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceExploringWebhookConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceExploringWebhookConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceExploringWebhookConfigurationInformer constructs a new informer for ResourceExploringWebhookConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceExploringWebhookConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ResourceExploringWebhookConfigurations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha1().ResourceExploringWebhookConfigurations(namespace).Watch(context.TODO(), options)
			},
		},
		&configv1alpha1.ResourceExploringWebhookConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceExploringWebhookConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceExploringWebhookConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceExploringWebhookConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha1.ResourceExploringWebhookConfiguration{}, f.defaultInformer)
}

func (f *resourceExploringWebhookConfigurationInformer) Lister() v1alpha1.ResourceExploringWebhookConfigurationLister {
	return v1alpha1.NewResourceExploringWebhookConfigurationLister(f.Informer().GetIndexer())
}
