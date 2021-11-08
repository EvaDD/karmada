// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/karmada-io/karmada/pkg/apis/config/v1alpha1"
	scheme "github.com/karmada-io/karmada/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceExploringWebhookConfigurationsGetter has a method to return a ResourceExploringWebhookConfigurationInterface.
// A group's client should implement this interface.
type ResourceExploringWebhookConfigurationsGetter interface {
	ResourceExploringWebhookConfigurations(namespace string) ResourceExploringWebhookConfigurationInterface
}

// ResourceExploringWebhookConfigurationInterface has methods to work with ResourceExploringWebhookConfiguration resources.
type ResourceExploringWebhookConfigurationInterface interface {
	Create(ctx context.Context, resourceExploringWebhookConfiguration *v1alpha1.ResourceExploringWebhookConfiguration, opts v1.CreateOptions) (*v1alpha1.ResourceExploringWebhookConfiguration, error)
	Update(ctx context.Context, resourceExploringWebhookConfiguration *v1alpha1.ResourceExploringWebhookConfiguration, opts v1.UpdateOptions) (*v1alpha1.ResourceExploringWebhookConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ResourceExploringWebhookConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ResourceExploringWebhookConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceExploringWebhookConfiguration, err error)
	ResourceExploringWebhookConfigurationExpansion
}

// resourceExploringWebhookConfigurations implements ResourceExploringWebhookConfigurationInterface
type resourceExploringWebhookConfigurations struct {
	client rest.Interface
	ns     string
}

// newResourceExploringWebhookConfigurations returns a ResourceExploringWebhookConfigurations
func newResourceExploringWebhookConfigurations(c *ConfigV1alpha1Client, namespace string) *resourceExploringWebhookConfigurations {
	return &resourceExploringWebhookConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceExploringWebhookConfiguration, and returns the corresponding resourceExploringWebhookConfiguration object, and an error if there is any.
func (c *resourceExploringWebhookConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceExploringWebhookConfiguration, err error) {
	result = &v1alpha1.ResourceExploringWebhookConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceExploringWebhookConfigurations that match those selectors.
func (c *resourceExploringWebhookConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceExploringWebhookConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceExploringWebhookConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceExploringWebhookConfigurations.
func (c *resourceExploringWebhookConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceExploringWebhookConfiguration and creates it.  Returns the server's representation of the resourceExploringWebhookConfiguration, and an error, if there is any.
func (c *resourceExploringWebhookConfigurations) Create(ctx context.Context, resourceExploringWebhookConfiguration *v1alpha1.ResourceExploringWebhookConfiguration, opts v1.CreateOptions) (result *v1alpha1.ResourceExploringWebhookConfiguration, err error) {
	result = &v1alpha1.ResourceExploringWebhookConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceExploringWebhookConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceExploringWebhookConfiguration and updates it. Returns the server's representation of the resourceExploringWebhookConfiguration, and an error, if there is any.
func (c *resourceExploringWebhookConfigurations) Update(ctx context.Context, resourceExploringWebhookConfiguration *v1alpha1.ResourceExploringWebhookConfiguration, opts v1.UpdateOptions) (result *v1alpha1.ResourceExploringWebhookConfiguration, err error) {
	result = &v1alpha1.ResourceExploringWebhookConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		Name(resourceExploringWebhookConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceExploringWebhookConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceExploringWebhookConfiguration and deletes it. Returns an error if one occurs.
func (c *resourceExploringWebhookConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceExploringWebhookConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceExploringWebhookConfiguration.
func (c *resourceExploringWebhookConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceExploringWebhookConfiguration, err error) {
	result = &v1alpha1.ResourceExploringWebhookConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourceexploringwebhookconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
