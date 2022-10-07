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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kube-bind/kube-bind/pkg/apis/kubebind/v1alpha1"
)

// FakeServiceBindingRequests implements ServiceBindingRequestInterface
type FakeServiceBindingRequests struct {
	Fake *FakeKubeBindV1alpha1
	ns   string
}

var servicebindingrequestsResource = schema.GroupVersionResource{Group: "kube-bind.io", Version: "v1alpha1", Resource: "servicebindingrequests"}

var servicebindingrequestsKind = schema.GroupVersionKind{Group: "kube-bind.io", Version: "v1alpha1", Kind: "ServiceBindingRequest"}

// Get takes name of the serviceBindingRequest, and returns the corresponding serviceBindingRequest object, and an error if there is any.
func (c *FakeServiceBindingRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceBindingRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicebindingrequestsResource, c.ns, name), &v1alpha1.ServiceBindingRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceBindingRequest), err
}

// List takes label and field selectors, and returns the list of ServiceBindingRequests that match those selectors.
func (c *FakeServiceBindingRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceBindingRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicebindingrequestsResource, servicebindingrequestsKind, c.ns, opts), &v1alpha1.ServiceBindingRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceBindingRequestList{ListMeta: obj.(*v1alpha1.ServiceBindingRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceBindingRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceBindingRequests.
func (c *FakeServiceBindingRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicebindingrequestsResource, c.ns, opts))

}

// Create takes the representation of a serviceBindingRequest and creates it.  Returns the server's representation of the serviceBindingRequest, and an error, if there is any.
func (c *FakeServiceBindingRequests) Create(ctx context.Context, serviceBindingRequest *v1alpha1.ServiceBindingRequest, opts v1.CreateOptions) (result *v1alpha1.ServiceBindingRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicebindingrequestsResource, c.ns, serviceBindingRequest), &v1alpha1.ServiceBindingRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceBindingRequest), err
}

// Update takes the representation of a serviceBindingRequest and updates it. Returns the server's representation of the serviceBindingRequest, and an error, if there is any.
func (c *FakeServiceBindingRequests) Update(ctx context.Context, serviceBindingRequest *v1alpha1.ServiceBindingRequest, opts v1.UpdateOptions) (result *v1alpha1.ServiceBindingRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicebindingrequestsResource, c.ns, serviceBindingRequest), &v1alpha1.ServiceBindingRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceBindingRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceBindingRequests) UpdateStatus(ctx context.Context, serviceBindingRequest *v1alpha1.ServiceBindingRequest, opts v1.UpdateOptions) (*v1alpha1.ServiceBindingRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicebindingrequestsResource, "status", c.ns, serviceBindingRequest), &v1alpha1.ServiceBindingRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceBindingRequest), err
}

// Delete takes name of the serviceBindingRequest and deletes it. Returns an error if one occurs.
func (c *FakeServiceBindingRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(servicebindingrequestsResource, c.ns, name, opts), &v1alpha1.ServiceBindingRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceBindingRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicebindingrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceBindingRequestList{})
	return err
}

// Patch applies the patch and returns the patched serviceBindingRequest.
func (c *FakeServiceBindingRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceBindingRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicebindingrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceBindingRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceBindingRequest), err
}
