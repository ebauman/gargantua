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

// CoursesGetter has a method to return a CourseInterface.
// A group's client should implement this interface.
type CoursesGetter interface {
	Courses() CourseInterface
}

// CourseInterface has methods to work with Course resources.
type CourseInterface interface {
	Create(*v1.Course) (*v1.Course, error)
	Update(*v1.Course) (*v1.Course, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Course, error)
	List(opts metav1.ListOptions) (*v1.CourseList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Course, err error)
	CourseExpansion
}

// courses implements CourseInterface
type courses struct {
	client rest.Interface
}

// newCourses returns a Courses
func newCourses(c *HobbyfarmV1Client) *courses {
	return &courses{
		client: c.RESTClient(),
	}
}

// Get takes name of the course, and returns the corresponding course object, and an error if there is any.
func (c *courses) Get(name string, options metav1.GetOptions) (result *v1.Course, err error) {
	result = &v1.Course{}
	err = c.client.Get().
		Resource("courses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Courses that match those selectors.
func (c *courses) List(opts metav1.ListOptions) (result *v1.CourseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CourseList{}
	err = c.client.Get().
		Resource("courses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested courses.
func (c *courses) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("courses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a course and creates it.  Returns the server's representation of the course, and an error, if there is any.
func (c *courses) Create(course *v1.Course) (result *v1.Course, err error) {
	result = &v1.Course{}
	err = c.client.Post().
		Resource("courses").
		Body(course).
		Do().
		Into(result)
	return
}

// Update takes the representation of a course and updates it. Returns the server's representation of the course, and an error, if there is any.
func (c *courses) Update(course *v1.Course) (result *v1.Course, err error) {
	result = &v1.Course{}
	err = c.client.Put().
		Resource("courses").
		Name(course.Name).
		Body(course).
		Do().
		Into(result)
	return
}

// Delete takes name of the course and deletes it. Returns an error if one occurs.
func (c *courses) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("courses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *courses) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("courses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched course.
func (c *courses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Course, err error) {
	result = &v1.Course{}
	err = c.client.Patch(pt).
		Resource("courses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}