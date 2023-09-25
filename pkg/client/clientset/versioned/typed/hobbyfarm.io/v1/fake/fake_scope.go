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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScopes implements ScopeInterface
type FakeScopes struct {
	Fake *FakeHobbyfarmV1
	ns   string
}

var scopesResource = schema.GroupVersionResource{Group: "hobbyfarm.io", Version: "v1", Resource: "scopes"}

var scopesKind = schema.GroupVersionKind{Group: "hobbyfarm.io", Version: "v1", Kind: "Scope"}

// Get takes name of the scope, and returns the corresponding scope object, and an error if there is any.
func (c *FakeScopes) Get(ctx context.Context, name string, options v1.GetOptions) (result *hobbyfarmiov1.Scope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scopesResource, c.ns, name), &hobbyfarmiov1.Scope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scope), err
}

// List takes label and field selectors, and returns the list of Scopes that match those selectors.
func (c *FakeScopes) List(ctx context.Context, opts v1.ListOptions) (result *hobbyfarmiov1.ScopeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scopesResource, scopesKind, c.ns, opts), &hobbyfarmiov1.ScopeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hobbyfarmiov1.ScopeList{ListMeta: obj.(*hobbyfarmiov1.ScopeList).ListMeta}
	for _, item := range obj.(*hobbyfarmiov1.ScopeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scopes.
func (c *FakeScopes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scopesResource, c.ns, opts))

}

// Create takes the representation of a scope and creates it.  Returns the server's representation of the scope, and an error, if there is any.
func (c *FakeScopes) Create(ctx context.Context, scope *hobbyfarmiov1.Scope, opts v1.CreateOptions) (result *hobbyfarmiov1.Scope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scopesResource, c.ns, scope), &hobbyfarmiov1.Scope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scope), err
}

// Update takes the representation of a scope and updates it. Returns the server's representation of the scope, and an error, if there is any.
func (c *FakeScopes) Update(ctx context.Context, scope *hobbyfarmiov1.Scope, opts v1.UpdateOptions) (result *hobbyfarmiov1.Scope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scopesResource, c.ns, scope), &hobbyfarmiov1.Scope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scope), err
}

// Delete takes name of the scope and deletes it. Returns an error if one occurs.
func (c *FakeScopes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scopesResource, c.ns, name, opts), &hobbyfarmiov1.Scope{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScopes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scopesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &hobbyfarmiov1.ScopeList{})
	return err
}

// Patch applies the patch and returns the patched scope.
func (c *FakeScopes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hobbyfarmiov1.Scope, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scopesResource, c.ns, name, pt, data, subresources...), &hobbyfarmiov1.Scope{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scope), err
}
