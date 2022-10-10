/*
Copyright The Kubectl Bind contributors.

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

// ServiceExportResourcesGetter has a method to return a ServiceExportResourceInterface.
// A group's client should implement this interface.
type ServiceExportResourcesGetter interface {
	ServiceExportResources(namespace string) ServiceExportResourceInterface
}

// ServiceExportResourceInterface has methods to work with ServiceExportResource resources.
type ServiceExportResourceInterface interface {
	Create(ctx context.Context, serviceExportResource *v1alpha1.ServiceExportResource, opts v1.CreateOptions) (*v1alpha1.ServiceExportResource, error)
	Update(ctx context.Context, serviceExportResource *v1alpha1.ServiceExportResource, opts v1.UpdateOptions) (*v1alpha1.ServiceExportResource, error)
	UpdateStatus(ctx context.Context, serviceExportResource *v1alpha1.ServiceExportResource, opts v1.UpdateOptions) (*v1alpha1.ServiceExportResource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceExportResource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceExportResourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceExportResource, err error)
	ServiceExportResourceExpansion
}

// serviceExportResources implements ServiceExportResourceInterface
type serviceExportResources struct {
	client rest.Interface
	ns     string
}

// newServiceExportResources returns a ServiceExportResources
func newServiceExportResources(c *KubeBindV1alpha1Client, namespace string) *serviceExportResources {
	return &serviceExportResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceExportResource, and returns the corresponding serviceExportResource object, and an error if there is any.
func (c *serviceExportResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceExportResource, err error) {
	result = &v1alpha1.ServiceExportResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceexportresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceExportResources that match those selectors.
func (c *serviceExportResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceExportResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceExportResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceexportresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceExportResources.
func (c *serviceExportResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceexportresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceExportResource and creates it.  Returns the server's representation of the serviceExportResource, and an error, if there is any.
func (c *serviceExportResources) Create(ctx context.Context, serviceExportResource *v1alpha1.ServiceExportResource, opts v1.CreateOptions) (result *v1alpha1.ServiceExportResource, err error) {
	result = &v1alpha1.ServiceExportResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceexportresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceExportResource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceExportResource and updates it. Returns the server's representation of the serviceExportResource, and an error, if there is any.
func (c *serviceExportResources) Update(ctx context.Context, serviceExportResource *v1alpha1.ServiceExportResource, opts v1.UpdateOptions) (result *v1alpha1.ServiceExportResource, err error) {
	result = &v1alpha1.ServiceExportResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceexportresources").
		Name(serviceExportResource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceExportResource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *serviceExportResources) UpdateStatus(ctx context.Context, serviceExportResource *v1alpha1.ServiceExportResource, opts v1.UpdateOptions) (result *v1alpha1.ServiceExportResource, err error) {
	result = &v1alpha1.ServiceExportResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceexportresources").
		Name(serviceExportResource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceExportResource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceExportResource and deletes it. Returns an error if one occurs.
func (c *serviceExportResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceexportresources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceExportResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceexportresources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceExportResource.
func (c *serviceExportResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceExportResource, err error) {
	result = &v1alpha1.ServiceExportResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceexportresources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
