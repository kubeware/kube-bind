/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

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

	v1alpha1 "go.bytebuilders.dev/kube-bind/apis/kubebind/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPIServiceExportRequests implements APIServiceExportRequestInterface
type FakeAPIServiceExportRequests struct {
	Fake *FakeKubeBindV1alpha1
	ns   string
}

var apiserviceexportrequestsResource = v1alpha1.SchemeGroupVersion.WithResource("apiserviceexportrequests")

var apiserviceexportrequestsKind = v1alpha1.SchemeGroupVersion.WithKind("APIServiceExportRequest")

// Get takes name of the aPIServiceExportRequest, and returns the corresponding aPIServiceExportRequest object, and an error if there is any.
func (c *FakeAPIServiceExportRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIServiceExportRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apiserviceexportrequestsResource, c.ns, name), &v1alpha1.APIServiceExportRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportRequest), err
}

// List takes label and field selectors, and returns the list of APIServiceExportRequests that match those selectors.
func (c *FakeAPIServiceExportRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIServiceExportRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apiserviceexportrequestsResource, apiserviceexportrequestsKind, c.ns, opts), &v1alpha1.APIServiceExportRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIServiceExportRequestList{ListMeta: obj.(*v1alpha1.APIServiceExportRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIServiceExportRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIServiceExportRequests.
func (c *FakeAPIServiceExportRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apiserviceexportrequestsResource, c.ns, opts))

}

// Create takes the representation of a aPIServiceExportRequest and creates it.  Returns the server's representation of the aPIServiceExportRequest, and an error, if there is any.
func (c *FakeAPIServiceExportRequests) Create(ctx context.Context, aPIServiceExportRequest *v1alpha1.APIServiceExportRequest, opts v1.CreateOptions) (result *v1alpha1.APIServiceExportRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apiserviceexportrequestsResource, c.ns, aPIServiceExportRequest), &v1alpha1.APIServiceExportRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportRequest), err
}

// Update takes the representation of a aPIServiceExportRequest and updates it. Returns the server's representation of the aPIServiceExportRequest, and an error, if there is any.
func (c *FakeAPIServiceExportRequests) Update(ctx context.Context, aPIServiceExportRequest *v1alpha1.APIServiceExportRequest, opts v1.UpdateOptions) (result *v1alpha1.APIServiceExportRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apiserviceexportrequestsResource, c.ns, aPIServiceExportRequest), &v1alpha1.APIServiceExportRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIServiceExportRequests) UpdateStatus(ctx context.Context, aPIServiceExportRequest *v1alpha1.APIServiceExportRequest, opts v1.UpdateOptions) (*v1alpha1.APIServiceExportRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apiserviceexportrequestsResource, "status", c.ns, aPIServiceExportRequest), &v1alpha1.APIServiceExportRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportRequest), err
}

// Delete takes name of the aPIServiceExportRequest and deletes it. Returns an error if one occurs.
func (c *FakeAPIServiceExportRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apiserviceexportrequestsResource, c.ns, name, opts), &v1alpha1.APIServiceExportRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIServiceExportRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apiserviceexportrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIServiceExportRequestList{})
	return err
}

// Patch applies the patch and returns the patched aPIServiceExportRequest.
func (c *FakeAPIServiceExportRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIServiceExportRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apiserviceexportrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.APIServiceExportRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIServiceExportRequest), err
}
