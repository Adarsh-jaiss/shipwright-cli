// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/shipwright-io/build/pkg/apis/build/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuildRuns implements BuildRunInterface
type FakeBuildRuns struct {
	Fake *FakeShipwrightV1beta1
	ns   string
}

var buildrunsResource = v1beta1.SchemeGroupVersion.WithResource("buildruns")

var buildrunsKind = v1beta1.SchemeGroupVersion.WithKind("BuildRun")

// Get takes name of the buildRun, and returns the corresponding buildRun object, and an error if there is any.
func (c *FakeBuildRuns) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BuildRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildrunsResource, c.ns, name), &v1beta1.BuildRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildRun), err
}

// List takes label and field selectors, and returns the list of BuildRuns that match those selectors.
func (c *FakeBuildRuns) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BuildRunList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildrunsResource, buildrunsKind, c.ns, opts), &v1beta1.BuildRunList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BuildRunList{ListMeta: obj.(*v1beta1.BuildRunList).ListMeta}
	for _, item := range obj.(*v1beta1.BuildRunList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buildRuns.
func (c *FakeBuildRuns) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildrunsResource, c.ns, opts))

}

// Create takes the representation of a buildRun and creates it.  Returns the server's representation of the buildRun, and an error, if there is any.
func (c *FakeBuildRuns) Create(ctx context.Context, buildRun *v1beta1.BuildRun, opts v1.CreateOptions) (result *v1beta1.BuildRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildrunsResource, c.ns, buildRun), &v1beta1.BuildRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildRun), err
}

// Update takes the representation of a buildRun and updates it. Returns the server's representation of the buildRun, and an error, if there is any.
func (c *FakeBuildRuns) Update(ctx context.Context, buildRun *v1beta1.BuildRun, opts v1.UpdateOptions) (result *v1beta1.BuildRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildrunsResource, c.ns, buildRun), &v1beta1.BuildRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildRun), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuildRuns) UpdateStatus(ctx context.Context, buildRun *v1beta1.BuildRun, opts v1.UpdateOptions) (*v1beta1.BuildRun, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildrunsResource, "status", c.ns, buildRun), &v1beta1.BuildRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildRun), err
}

// Delete takes name of the buildRun and deletes it. Returns an error if one occurs.
func (c *FakeBuildRuns) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(buildrunsResource, c.ns, name, opts), &v1beta1.BuildRun{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuildRuns) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildrunsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BuildRunList{})
	return err
}

// Patch applies the patch and returns the patched buildRun.
func (c *FakeBuildRuns) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BuildRun, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildrunsResource, c.ns, name, pt, data, subresources...), &v1beta1.BuildRun{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BuildRun), err
}
