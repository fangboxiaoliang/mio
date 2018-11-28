/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "hidevops.io/mio/pkg/apis/mio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTestses implements TestsInterface
type FakeTestses struct {
	Fake *FakeMioV1alpha1
	ns   string
}

var testsesResource = schema.GroupVersionResource{Group: "mio.io", Version: "v1alpha1", Resource: "testses"}

var testsesKind = schema.GroupVersionKind{Group: "mio.io", Version: "v1alpha1", Kind: "Tests"}

// Get takes name of the tests, and returns the corresponding tests object, and an error if there is any.
func (c *FakeTestses) Get(name string, options v1.GetOptions) (result *v1alpha1.Tests, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(testsesResource, c.ns, name), &v1alpha1.Tests{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tests), err
}

// List takes label and field selectors, and returns the list of Testses that match those selectors.
func (c *FakeTestses) List(opts v1.ListOptions) (result *v1alpha1.TestsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(testsesResource, testsesKind, c.ns, opts), &v1alpha1.TestsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TestsList{}
	for _, item := range obj.(*v1alpha1.TestsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testses.
func (c *FakeTestses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(testsesResource, c.ns, opts))

}

// Create takes the representation of a tests and creates it.  Returns the server's representation of the tests, and an error, if there is any.
func (c *FakeTestses) Create(tests *v1alpha1.Tests) (result *v1alpha1.Tests, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(testsesResource, c.ns, tests), &v1alpha1.Tests{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tests), err
}

// Update takes the representation of a tests and updates it. Returns the server's representation of the tests, and an error, if there is any.
func (c *FakeTestses) Update(tests *v1alpha1.Tests) (result *v1alpha1.Tests, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(testsesResource, c.ns, tests), &v1alpha1.Tests{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tests), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTestses) UpdateStatus(tests *v1alpha1.Tests) (*v1alpha1.Tests, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(testsesResource, "status", c.ns, tests), &v1alpha1.Tests{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tests), err
}

// Delete takes name of the tests and deletes it. Returns an error if one occurs.
func (c *FakeTestses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(testsesResource, c.ns, name), &v1alpha1.Tests{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTestses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(testsesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TestsList{})
	return err
}

// Patch applies the patch and returns the patched tests.
func (c *FakeTestses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Tests, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(testsesResource, c.ns, name, data, subresources...), &v1alpha1.Tests{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Tests), err
}