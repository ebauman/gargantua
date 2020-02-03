/*
Copyright The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDynamicBindRequests implements DynamicBindRequestInterface
type FakeDynamicBindRequests struct {
	Fake *FakeHobbyfarmV1
}

var dynamicbindrequestsResource = schema.GroupVersionResource{Group: "hobbyfarm.io", Version: "v1", Resource: "dynamicbindrequests"}

var dynamicbindrequestsKind = schema.GroupVersionKind{Group: "hobbyfarm.io", Version: "v1", Kind: "DynamicBindRequest"}

// Get takes name of the dynamicBindRequest, and returns the corresponding dynamicBindRequest object, and an error if there is any.
func (c *FakeDynamicBindRequests) Get(name string, options v1.GetOptions) (result *hobbyfarmiov1.DynamicBindRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dynamicbindrequestsResource, name), &hobbyfarmiov1.DynamicBindRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindRequest), err
}

// List takes label and field selectors, and returns the list of DynamicBindRequests that match those selectors.
func (c *FakeDynamicBindRequests) List(opts v1.ListOptions) (result *hobbyfarmiov1.DynamicBindRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dynamicbindrequestsResource, dynamicbindrequestsKind, opts), &hobbyfarmiov1.DynamicBindRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hobbyfarmiov1.DynamicBindRequestList{ListMeta: obj.(*hobbyfarmiov1.DynamicBindRequestList).ListMeta}
	for _, item := range obj.(*hobbyfarmiov1.DynamicBindRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dynamicBindRequests.
func (c *FakeDynamicBindRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dynamicbindrequestsResource, opts))
}

// Create takes the representation of a dynamicBindRequest and creates it.  Returns the server's representation of the dynamicBindRequest, and an error, if there is any.
func (c *FakeDynamicBindRequests) Create(dynamicBindRequest *hobbyfarmiov1.DynamicBindRequest) (result *hobbyfarmiov1.DynamicBindRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dynamicbindrequestsResource, dynamicBindRequest), &hobbyfarmiov1.DynamicBindRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindRequest), err
}

// Update takes the representation of a dynamicBindRequest and updates it. Returns the server's representation of the dynamicBindRequest, and an error, if there is any.
func (c *FakeDynamicBindRequests) Update(dynamicBindRequest *hobbyfarmiov1.DynamicBindRequest) (result *hobbyfarmiov1.DynamicBindRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dynamicbindrequestsResource, dynamicBindRequest), &hobbyfarmiov1.DynamicBindRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDynamicBindRequests) UpdateStatus(dynamicBindRequest *hobbyfarmiov1.DynamicBindRequest) (*hobbyfarmiov1.DynamicBindRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dynamicbindrequestsResource, "status", dynamicBindRequest), &hobbyfarmiov1.DynamicBindRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindRequest), err
}

// Delete takes name of the dynamicBindRequest and deletes it. Returns an error if one occurs.
func (c *FakeDynamicBindRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dynamicbindrequestsResource, name), &hobbyfarmiov1.DynamicBindRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDynamicBindRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dynamicbindrequestsResource, listOptions)

	_, err := c.Fake.Invokes(action, &hobbyfarmiov1.DynamicBindRequestList{})
	return err
}

// Patch applies the patch and returns the patched dynamicBindRequest.
func (c *FakeDynamicBindRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hobbyfarmiov1.DynamicBindRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dynamicbindrequestsResource, name, pt, data, subresources...), &hobbyfarmiov1.DynamicBindRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindRequest), err
}
