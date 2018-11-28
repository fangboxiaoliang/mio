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

package v1alpha1

import (
	v1alpha1 "hidevops.io/mio/pkg/apis/mio/v1alpha1"
	scheme "hidevops.io/mio/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NotifiesGetter has a method to return a NotifyInterface.
// A group's client should implement this interface.
type NotifiesGetter interface {
	Notifies(namespace string) NotifyInterface
}

// NotifyInterface has methods to work with Notify resources.
type NotifyInterface interface {
	Create(*v1alpha1.Notify) (*v1alpha1.Notify, error)
	Update(*v1alpha1.Notify) (*v1alpha1.Notify, error)
	UpdateStatus(*v1alpha1.Notify) (*v1alpha1.Notify, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Notify, error)
	List(opts v1.ListOptions) (*v1alpha1.NotifyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Notify, err error)
	NotifyExpansion
}

// notifies implements NotifyInterface
type notifies struct {
	client rest.Interface
	ns     string
}

// newNotifies returns a Notifies
func newNotifies(c *MioV1alpha1Client, namespace string) *notifies {
	return &notifies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the notify, and returns the corresponding notify object, and an error if there is any.
func (c *notifies) Get(name string, options v1.GetOptions) (result *v1alpha1.Notify, err error) {
	result = &v1alpha1.Notify{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notifies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Notifies that match those selectors.
func (c *notifies) List(opts v1.ListOptions) (result *v1alpha1.NotifyList, err error) {
	result = &v1alpha1.NotifyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notifies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested notifies.
func (c *notifies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("notifies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a notify and creates it.  Returns the server's representation of the notify, and an error, if there is any.
func (c *notifies) Create(notify *v1alpha1.Notify) (result *v1alpha1.Notify, err error) {
	result = &v1alpha1.Notify{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("notifies").
		Body(notify).
		Do().
		Into(result)
	return
}

// Update takes the representation of a notify and updates it. Returns the server's representation of the notify, and an error, if there is any.
func (c *notifies) Update(notify *v1alpha1.Notify) (result *v1alpha1.Notify, err error) {
	result = &v1alpha1.Notify{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notifies").
		Name(notify.Name).
		Body(notify).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *notifies) UpdateStatus(notify *v1alpha1.Notify) (result *v1alpha1.Notify, err error) {
	result = &v1alpha1.Notify{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notifies").
		Name(notify.Name).
		SubResource("status").
		Body(notify).
		Do().
		Into(result)
	return
}

// Delete takes name of the notify and deletes it. Returns an error if one occurs.
func (c *notifies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notifies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *notifies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notifies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched notify.
func (c *notifies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Notify, err error) {
	result = &v1alpha1.Notify{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("notifies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
