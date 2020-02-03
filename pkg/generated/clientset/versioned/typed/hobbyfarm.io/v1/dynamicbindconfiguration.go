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

package v1

import (
	"time"

	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	scheme "github.com/hobbyfarm/gargantua/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DynamicBindConfigurationsGetter has a method to return a DynamicBindConfigurationInterface.
// A group's client should implement this interface.
type DynamicBindConfigurationsGetter interface {
	DynamicBindConfigurations() DynamicBindConfigurationInterface
}

// DynamicBindConfigurationInterface has methods to work with DynamicBindConfiguration resources.
type DynamicBindConfigurationInterface interface {
	Create(*v1.DynamicBindConfiguration) (*v1.DynamicBindConfiguration, error)
	Update(*v1.DynamicBindConfiguration) (*v1.DynamicBindConfiguration, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.DynamicBindConfiguration, error)
	List(opts metav1.ListOptions) (*v1.DynamicBindConfigurationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DynamicBindConfiguration, err error)
	DynamicBindConfigurationExpansion
}

// dynamicBindConfigurations implements DynamicBindConfigurationInterface
type dynamicBindConfigurations struct {
	client rest.Interface
}

// newDynamicBindConfigurations returns a DynamicBindConfigurations
func newDynamicBindConfigurations(c *HobbyfarmV1Client) *dynamicBindConfigurations {
	return &dynamicBindConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the dynamicBindConfiguration, and returns the corresponding dynamicBindConfiguration object, and an error if there is any.
func (c *dynamicBindConfigurations) Get(name string, options metav1.GetOptions) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Get().
		Resource("dynamicbindconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DynamicBindConfigurations that match those selectors.
func (c *dynamicBindConfigurations) List(opts metav1.ListOptions) (result *v1.DynamicBindConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DynamicBindConfigurationList{}
	err = c.client.Get().
		Resource("dynamicbindconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dynamicBindConfigurations.
func (c *dynamicBindConfigurations) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("dynamicbindconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dynamicBindConfiguration and creates it.  Returns the server's representation of the dynamicBindConfiguration, and an error, if there is any.
func (c *dynamicBindConfigurations) Create(dynamicBindConfiguration *v1.DynamicBindConfiguration) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Post().
		Resource("dynamicbindconfigurations").
		Body(dynamicBindConfiguration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dynamicBindConfiguration and updates it. Returns the server's representation of the dynamicBindConfiguration, and an error, if there is any.
func (c *dynamicBindConfigurations) Update(dynamicBindConfiguration *v1.DynamicBindConfiguration) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Put().
		Resource("dynamicbindconfigurations").
		Name(dynamicBindConfiguration.Name).
		Body(dynamicBindConfiguration).
		Do().
		Into(result)
	return
}

// Delete takes name of the dynamicBindConfiguration and deletes it. Returns an error if one occurs.
func (c *dynamicBindConfigurations) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("dynamicbindconfigurations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dynamicBindConfigurations) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("dynamicbindconfigurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dynamicBindConfiguration.
func (c *dynamicBindConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Patch(pt).
		Resource("dynamicbindconfigurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
