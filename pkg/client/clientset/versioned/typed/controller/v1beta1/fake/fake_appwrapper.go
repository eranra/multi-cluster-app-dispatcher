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

package fake

import (
	"context"

	v1beta1 "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/apis/controller/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppWrappers implements AppWrapperInterface
type FakeAppWrappers struct {
	Fake *FakeMcadV1beta1
	ns   string
}

var appwrappersResource = schema.GroupVersionResource{Group: "mcad.ibm.com", Version: "v1beta1", Resource: "appwrappers"}

var appwrappersKind = schema.GroupVersionKind{Group: "mcad.ibm.com", Version: "v1beta1", Kind: "AppWrapper"}

// Get takes name of the appWrapper, and returns the corresponding appWrapper object, and an error if there is any.
func (c *FakeAppWrappers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AppWrapper, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appwrappersResource, c.ns, name), &v1beta1.AppWrapper{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppWrapper), err
}

// List takes label and field selectors, and returns the list of AppWrappers that match those selectors.
func (c *FakeAppWrappers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AppWrapperList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appwrappersResource, appwrappersKind, c.ns, opts), &v1beta1.AppWrapperList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AppWrapperList{ListMeta: obj.(*v1beta1.AppWrapperList).ListMeta}
	for _, item := range obj.(*v1beta1.AppWrapperList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appWrappers.
func (c *FakeAppWrappers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appwrappersResource, c.ns, opts))

}

// Create takes the representation of a appWrapper and creates it.  Returns the server's representation of the appWrapper, and an error, if there is any.
func (c *FakeAppWrappers) Create(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.CreateOptions) (result *v1beta1.AppWrapper, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appwrappersResource, c.ns, appWrapper), &v1beta1.AppWrapper{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppWrapper), err
}

// Update takes the representation of a appWrapper and updates it. Returns the server's representation of the appWrapper, and an error, if there is any.
func (c *FakeAppWrappers) Update(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.UpdateOptions) (result *v1beta1.AppWrapper, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appwrappersResource, c.ns, appWrapper), &v1beta1.AppWrapper{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppWrapper), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppWrappers) UpdateStatus(ctx context.Context, appWrapper *v1beta1.AppWrapper, opts v1.UpdateOptions) (*v1beta1.AppWrapper, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appwrappersResource, "status", c.ns, appWrapper), &v1beta1.AppWrapper{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppWrapper), err
}

// Delete takes name of the appWrapper and deletes it. Returns an error if one occurs.
func (c *FakeAppWrappers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appwrappersResource, c.ns, name), &v1beta1.AppWrapper{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppWrappers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appwrappersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AppWrapperList{})
	return err
}

// Patch applies the patch and returns the patched appWrapper.
func (c *FakeAppWrappers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AppWrapper, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appwrappersResource, c.ns, name, pt, data, subresources...), &v1beta1.AppWrapper{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AppWrapper), err
}
