/*
Copyright 2019, 2021, 2022, 2023 The Multi-Cluster App Dispatcher Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/apis/controller/v1beta1"
	scheme "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppWrappersGetter has a method to return a AppWrapperInterface.
// A group's client should implement this interface.
type AppWrappersGetter interface {
	AppWrappers(namespace string) AppWrapperInterface
}

// AppWrapperInterface has methods to work with AppWrapper resources.
type AppWrapperInterface interface {
	Create(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.CreateOptions) (*v1beta1.AppWrapper, error)
	Update(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.UpdateOptions) (*v1beta1.AppWrapper, error)
	UpdateStatus(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.UpdateOptions) (*v1beta1.AppWrapper, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.AppWrapper, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.AppWrapperList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AppWrapper, err error)
	AppWrapperExpansion
}

// appWrappers implements AppWrapperInterface
type appWrappers struct {
	client rest.Interface
	ns     string
}

// newAppWrappers returns a AppWrappers
func newAppWrappers(c *McadV1beta1Client, namespace string) *appWrappers {
	return &appWrappers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appWrapper, and returns the corresponding appWrapper object, and an error if there is any.
func (c *appWrappers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AppWrapper, err error) {
	result = &v1beta1.AppWrapper{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appwrappers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppWrappers that match those selectors.
func (c *appWrappers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AppWrapperList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.AppWrapperList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appwrappers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appWrappers.
func (c *appWrappers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appwrappers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a appWrapper and creates it.  Returns the server's representation of the appWrapper, and an error, if there is any.
func (c *appWrappers) Create(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.CreateOptions) (result *v1beta1.AppWrapper, err error) {
	result = &v1beta1.AppWrapper{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appwrappers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appWrapper).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a appWrapper and updates it. Returns the server's representation of the appWrapper, and an error, if there is any.
func (c *appWrappers) Update(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.UpdateOptions) (result *v1beta1.AppWrapper, err error) {
	result = &v1beta1.AppWrapper{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appwrappers").
		Name(appWrapper.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appWrapper).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *appWrappers) UpdateStatus(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.UpdateOptions) (result *v1beta1.AppWrapper, err error) {
	result = &v1beta1.AppWrapper{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appwrappers").
		Name(appWrapper.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appWrapper).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the appWrapper and deletes it. Returns an error if one occurs.
func (c *appWrappers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appwrappers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appWrappers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appwrappers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched appWrapper.
func (c *appWrappers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AppWrapper, err error) {
	result = &v1beta1.AppWrapper{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appwrappers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
