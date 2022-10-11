/*
Copyright The Kube Bind Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
	scheme "github.com/kube-bind/kube-bind/pkg/client/clientset/versioned/scheme"
)

// ServiceProvidersGetter has a method to return a ServiceProviderInterface.
// A group's client should implement this interface.
type ServiceProvidersGetter interface {
	ServiceProviders(namespace string) ServiceProviderInterface
}

// ServiceProviderInterface has methods to work with ServiceProvider resources.
type ServiceProviderInterface interface {
	Create(ctx context.Context, serviceProvider *v1alpha1.ServiceProvider, opts v1.CreateOptions) (*v1alpha1.ServiceProvider, error)
	Update(ctx context.Context, serviceProvider *v1alpha1.ServiceProvider, opts v1.UpdateOptions) (*v1alpha1.ServiceProvider, error)
	UpdateStatus(ctx context.Context, serviceProvider *v1alpha1.ServiceProvider, opts v1.UpdateOptions) (*v1alpha1.ServiceProvider, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceProvider, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceProviderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceProvider, err error)
	ServiceProviderExpansion
}

// serviceProviders implements ServiceProviderInterface
type serviceProviders struct {
	client rest.Interface
	ns     string
}

// newServiceProviders returns a ServiceProviders
func newServiceProviders(c *KubeBindV1alpha1Client, namespace string) *serviceProviders {
	return &serviceProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceProvider, and returns the corresponding serviceProvider object, and an error if there is any.
func (c *serviceProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceProvider, err error) {
	result = &v1alpha1.ServiceProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceProviders that match those selectors.
func (c *serviceProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceProviders.
func (c *serviceProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceProvider and creates it.  Returns the server's representation of the serviceProvider, and an error, if there is any.
func (c *serviceProviders) Create(ctx context.Context, serviceProvider *v1alpha1.ServiceProvider, opts v1.CreateOptions) (result *v1alpha1.ServiceProvider, err error) {
	result = &v1alpha1.ServiceProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceProvider).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceProvider and updates it. Returns the server's representation of the serviceProvider, and an error, if there is any.
func (c *serviceProviders) Update(ctx context.Context, serviceProvider *v1alpha1.ServiceProvider, opts v1.UpdateOptions) (result *v1alpha1.ServiceProvider, err error) {
	result = &v1alpha1.ServiceProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceproviders").
		Name(serviceProvider.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceProvider).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceProviders) UpdateStatus(ctx context.Context, serviceProvider *v1alpha1.ServiceProvider, opts v1.UpdateOptions) (result *v1alpha1.ServiceProvider, err error) {
	result = &v1alpha1.ServiceProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceproviders").
		Name(serviceProvider.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceProvider).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceProvider and deletes it. Returns an error if one occurs.
func (c *serviceProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceproviders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceproviders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceProvider.
func (c *serviceProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceProvider, err error) {
	result = &v1alpha1.ServiceProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceproviders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}