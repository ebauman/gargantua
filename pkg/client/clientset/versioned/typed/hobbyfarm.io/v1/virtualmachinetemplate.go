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

package v1

import (
	"context"
	"time"

	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	scheme "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualMachineTemplatesGetter has a method to return a VirtualMachineTemplateInterface.
// A group's client should implement this interface.
type VirtualMachineTemplatesGetter interface {
	VirtualMachineTemplates() VirtualMachineTemplateInterface
}

// VirtualMachineTemplateInterface has methods to work with VirtualMachineTemplate resources.
type VirtualMachineTemplateInterface interface {
	Create(ctx context.Context, virtualMachineTemplate *v1.VirtualMachineTemplate, opts metav1.CreateOptions) (*v1.VirtualMachineTemplate, error)
	Update(ctx context.Context, virtualMachineTemplate *v1.VirtualMachineTemplate, opts metav1.UpdateOptions) (*v1.VirtualMachineTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualMachineTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualMachineTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineTemplate, err error)
	VirtualMachineTemplateExpansion
}

// virtualMachineTemplates implements VirtualMachineTemplateInterface
type virtualMachineTemplates struct {
	client rest.Interface
}

// newVirtualMachineTemplates returns a VirtualMachineTemplates
func newVirtualMachineTemplates(c *HobbyfarmV1Client) *virtualMachineTemplates {
	return &virtualMachineTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the virtualMachineTemplate, and returns the corresponding virtualMachineTemplate object, and an error if there is any.
func (c *virtualMachineTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineTemplate, err error) {
	result = &v1.VirtualMachineTemplate{}
	err = c.client.Get().
		Resource("virtualmachinetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineTemplates that match those selectors.
func (c *virtualMachineTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineTemplateList{}
	err = c.client.Get().
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineTemplates.
func (c *virtualMachineTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineTemplate and creates it.  Returns the server's representation of the virtualMachineTemplate, and an error, if there is any.
func (c *virtualMachineTemplates) Create(ctx context.Context, virtualMachineTemplate *v1.VirtualMachineTemplate, opts metav1.CreateOptions) (result *v1.VirtualMachineTemplate, err error) {
	result = &v1.VirtualMachineTemplate{}
	err = c.client.Post().
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineTemplate and updates it. Returns the server's representation of the virtualMachineTemplate, and an error, if there is any.
func (c *virtualMachineTemplates) Update(ctx context.Context, virtualMachineTemplate *v1.VirtualMachineTemplate, opts metav1.UpdateOptions) (result *v1.VirtualMachineTemplate, err error) {
	result = &v1.VirtualMachineTemplate{}
	err = c.client.Put().
		Resource("virtualmachinetemplates").
		Name(virtualMachineTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineTemplate and deletes it. Returns an error if one occurs.
func (c *virtualMachineTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("virtualmachinetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("virtualmachinetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineTemplate.
func (c *virtualMachineTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineTemplate, err error) {
	result = &v1.VirtualMachineTemplate{}
	err = c.client.Patch(pt).
		Resource("virtualmachinetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
