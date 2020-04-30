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
	"time"

	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	scheme "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ContentRepositoriesGetter has a method to return a ContentRepositoryInterface.
// A group's client should implement this interface.
type ContentRepositoriesGetter interface {
	ContentRepositories() ContentRepositoryInterface
}

// ContentRepositoryInterface has methods to work with ContentRepository resources.
type ContentRepositoryInterface interface {
	Create(*v1.ContentRepository) (*v1.ContentRepository, error)
	Update(*v1.ContentRepository) (*v1.ContentRepository, error)
	UpdateStatus(*v1.ContentRepository) (*v1.ContentRepository, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ContentRepository, error)
	List(opts metav1.ListOptions) (*v1.ContentRepositoryList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ContentRepository, err error)
	ContentRepositoryExpansion
}

// contentRepositories implements ContentRepositoryInterface
type contentRepositories struct {
	client rest.Interface
}

// newContentRepositories returns a ContentRepositories
func newContentRepositories(c *HobbyfarmV1Client) *contentRepositories {
	return &contentRepositories{
		client: c.RESTClient(),
	}
}

// Get takes name of the contentRepository, and returns the corresponding contentRepository object, and an error if there is any.
func (c *contentRepositories) Get(name string, options metav1.GetOptions) (result *v1.ContentRepository, err error) {
	result = &v1.ContentRepository{}
	err = c.client.Get().
		Resource("contentrepositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ContentRepositories that match those selectors.
func (c *contentRepositories) List(opts metav1.ListOptions) (result *v1.ContentRepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ContentRepositoryList{}
	err = c.client.Get().
		Resource("contentrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested contentRepositories.
func (c *contentRepositories) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("contentrepositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a contentRepository and creates it.  Returns the server's representation of the contentRepository, and an error, if there is any.
func (c *contentRepositories) Create(contentRepository *v1.ContentRepository) (result *v1.ContentRepository, err error) {
	result = &v1.ContentRepository{}
	err = c.client.Post().
		Resource("contentrepositories").
		Body(contentRepository).
		Do().
		Into(result)
	return
}

// Update takes the representation of a contentRepository and updates it. Returns the server's representation of the contentRepository, and an error, if there is any.
func (c *contentRepositories) Update(contentRepository *v1.ContentRepository) (result *v1.ContentRepository, err error) {
	result = &v1.ContentRepository{}
	err = c.client.Put().
		Resource("contentrepositories").
		Name(contentRepository.Name).
		Body(contentRepository).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *contentRepositories) UpdateStatus(contentRepository *v1.ContentRepository) (result *v1.ContentRepository, err error) {
	result = &v1.ContentRepository{}
	err = c.client.Put().
		Resource("contentrepositories").
		Name(contentRepository.Name).
		SubResource("status").
		Body(contentRepository).
		Do().
		Into(result)
	return
}

// Delete takes name of the contentRepository and deletes it. Returns an error if one occurs.
func (c *contentRepositories) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("contentrepositories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *contentRepositories) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("contentrepositories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched contentRepository.
func (c *contentRepositories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ContentRepository, err error) {
	result = &v1.ContentRepository{}
	err = c.client.Patch(pt).
		Resource("contentrepositories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
