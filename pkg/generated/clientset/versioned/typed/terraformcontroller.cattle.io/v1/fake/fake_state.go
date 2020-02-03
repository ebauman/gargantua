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
	terraformcontrollercattleiov1 "github.com/hobbyfarm/gargantua/pkg/apis/terraformcontroller.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStates implements StateInterface
type FakeStates struct {
	Fake *FakeTerraformcontrollerV1
	ns   string
}

var statesResource = schema.GroupVersionResource{Group: "terraformcontroller.cattle.io", Version: "v1", Resource: "states"}

var statesKind = schema.GroupVersionKind{Group: "terraformcontroller.cattle.io", Version: "v1", Kind: "State"}

// Get takes name of the state, and returns the corresponding state object, and an error if there is any.
func (c *FakeStates) Get(name string, options v1.GetOptions) (result *terraformcontrollercattleiov1.State, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(statesResource, c.ns, name), &terraformcontrollercattleiov1.State{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.State), err
}

// List takes label and field selectors, and returns the list of States that match those selectors.
func (c *FakeStates) List(opts v1.ListOptions) (result *terraformcontrollercattleiov1.StateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(statesResource, statesKind, c.ns, opts), &terraformcontrollercattleiov1.StateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &terraformcontrollercattleiov1.StateList{ListMeta: obj.(*terraformcontrollercattleiov1.StateList).ListMeta}
	for _, item := range obj.(*terraformcontrollercattleiov1.StateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested states.
func (c *FakeStates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(statesResource, c.ns, opts))

}

// Create takes the representation of a state and creates it.  Returns the server's representation of the state, and an error, if there is any.
func (c *FakeStates) Create(state *terraformcontrollercattleiov1.State) (result *terraformcontrollercattleiov1.State, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(statesResource, c.ns, state), &terraformcontrollercattleiov1.State{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.State), err
}

// Update takes the representation of a state and updates it. Returns the server's representation of the state, and an error, if there is any.
func (c *FakeStates) Update(state *terraformcontrollercattleiov1.State) (result *terraformcontrollercattleiov1.State, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(statesResource, c.ns, state), &terraformcontrollercattleiov1.State{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.State), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStates) UpdateStatus(state *terraformcontrollercattleiov1.State) (*terraformcontrollercattleiov1.State, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(statesResource, "status", c.ns, state), &terraformcontrollercattleiov1.State{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.State), err
}

// Delete takes name of the state and deletes it. Returns an error if one occurs.
func (c *FakeStates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(statesResource, c.ns, name), &terraformcontrollercattleiov1.State{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(statesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &terraformcontrollercattleiov1.StateList{})
	return err
}

// Patch applies the patch and returns the patched state.
func (c *FakeStates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *terraformcontrollercattleiov1.State, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statesResource, c.ns, name, pt, data, subresources...), &terraformcontrollercattleiov1.State{})

	if obj == nil {
		return nil, err
	}
	return obj.(*terraformcontrollercattleiov1.State), err
}
