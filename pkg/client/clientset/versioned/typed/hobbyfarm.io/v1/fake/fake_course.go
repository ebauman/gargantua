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
	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCourses implements CourseInterface
type FakeCourses struct {
	Fake *FakeHobbyfarmV1
}

var coursesResource = schema.GroupVersionResource{Group: "hobbyfarm.io", Version: "v1", Resource: "courses"}

var coursesKind = schema.GroupVersionKind{Group: "hobbyfarm.io", Version: "v1", Kind: "Course"}

// Get takes name of the course, and returns the corresponding course object, and an error if there is any.
func (c *FakeCourses) Get(name string, options v1.GetOptions) (result *hobbyfarmiov1.Course, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(coursesResource, name), &hobbyfarmiov1.Course{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Course), err
}

// List takes label and field selectors, and returns the list of Courses that match those selectors.
func (c *FakeCourses) List(opts v1.ListOptions) (result *hobbyfarmiov1.CourseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(coursesResource, coursesKind, opts), &hobbyfarmiov1.CourseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hobbyfarmiov1.CourseList{ListMeta: obj.(*hobbyfarmiov1.CourseList).ListMeta}
	for _, item := range obj.(*hobbyfarmiov1.CourseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested courses.
func (c *FakeCourses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(coursesResource, opts))
}

// Create takes the representation of a course and creates it.  Returns the server's representation of the course, and an error, if there is any.
func (c *FakeCourses) Create(course *hobbyfarmiov1.Course) (result *hobbyfarmiov1.Course, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(coursesResource, course), &hobbyfarmiov1.Course{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Course), err
}

// Update takes the representation of a course and updates it. Returns the server's representation of the course, and an error, if there is any.
func (c *FakeCourses) Update(course *hobbyfarmiov1.Course) (result *hobbyfarmiov1.Course, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(coursesResource, course), &hobbyfarmiov1.Course{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Course), err
}

// Delete takes name of the course and deletes it. Returns an error if one occurs.
func (c *FakeCourses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(coursesResource, name), &hobbyfarmiov1.Course{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCourses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(coursesResource, listOptions)

	_, err := c.Fake.Invokes(action, &hobbyfarmiov1.CourseList{})
	return err
}

// Patch applies the patch and returns the patched course.
func (c *FakeCourses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hobbyfarmiov1.Course, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(coursesResource, name, pt, data, subresources...), &hobbyfarmiov1.Course{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Course), err
}