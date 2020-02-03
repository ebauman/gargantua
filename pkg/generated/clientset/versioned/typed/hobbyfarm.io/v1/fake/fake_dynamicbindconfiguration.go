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

// FakeDynamicBindConfigurations implements DynamicBindConfigurationInterface
type FakeDynamicBindConfigurations struct {
	Fake *FakeHobbyfarmV1
}

var dynamicbindconfigurationsResource = schema.GroupVersionResource{Group: "hobbyfarm.io", Version: "v1", Resource: "dynamicbindconfigurations"}

var dynamicbindconfigurationsKind = schema.GroupVersionKind{Group: "hobbyfarm.io", Version: "v1", Kind: "DynamicBindConfiguration"}

// Get takes name of the dynamicBindConfiguration, and returns the corresponding dynamicBindConfiguration object, and an error if there is any.
func (c *FakeDynamicBindConfigurations) Get(name string, options v1.GetOptions) (result *hobbyfarmiov1.DynamicBindConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dynamicbindconfigurationsResource, name), &hobbyfarmiov1.DynamicBindConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindConfiguration), err
}

// List takes label and field selectors, and returns the list of DynamicBindConfigurations that match those selectors.
func (c *FakeDynamicBindConfigurations) List(opts v1.ListOptions) (result *hobbyfarmiov1.DynamicBindConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dynamicbindconfigurationsResource, dynamicbindconfigurationsKind, opts), &hobbyfarmiov1.DynamicBindConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hobbyfarmiov1.DynamicBindConfigurationList{ListMeta: obj.(*hobbyfarmiov1.DynamicBindConfigurationList).ListMeta}
	for _, item := range obj.(*hobbyfarmiov1.DynamicBindConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dynamicBindConfigurations.
func (c *FakeDynamicBindConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dynamicbindconfigurationsResource, opts))
}

// Create takes the representation of a dynamicBindConfiguration and creates it.  Returns the server's representation of the dynamicBindConfiguration, and an error, if there is any.
func (c *FakeDynamicBindConfigurations) Create(dynamicBindConfiguration *hobbyfarmiov1.DynamicBindConfiguration) (result *hobbyfarmiov1.DynamicBindConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dynamicbindconfigurationsResource, dynamicBindConfiguration), &hobbyfarmiov1.DynamicBindConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindConfiguration), err
}

// Update takes the representation of a dynamicBindConfiguration and updates it. Returns the server's representation of the dynamicBindConfiguration, and an error, if there is any.
func (c *FakeDynamicBindConfigurations) Update(dynamicBindConfiguration *hobbyfarmiov1.DynamicBindConfiguration) (result *hobbyfarmiov1.DynamicBindConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dynamicbindconfigurationsResource, dynamicBindConfiguration), &hobbyfarmiov1.DynamicBindConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindConfiguration), err
}

// Delete takes name of the dynamicBindConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeDynamicBindConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(dynamicbindconfigurationsResource, name), &hobbyfarmiov1.DynamicBindConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDynamicBindConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dynamicbindconfigurationsResource, listOptions)

	_, err := c.Fake.Invokes(action, &hobbyfarmiov1.DynamicBindConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched dynamicBindConfiguration.
func (c *FakeDynamicBindConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hobbyfarmiov1.DynamicBindConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dynamicbindconfigurationsResource, name, pt, data, subresources...), &hobbyfarmiov1.DynamicBindConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.DynamicBindConfiguration), err
}
