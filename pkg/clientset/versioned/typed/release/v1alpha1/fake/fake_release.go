/*
Copyright 2020 Giant Swarm GmbH.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
)

// FakeReleases implements ReleaseInterface
type FakeReleases struct {
	Fake *FakeReleaseV1alpha1
}

var releasesResource = schema.GroupVersionResource{Group: "release.giantswarm.io", Version: "v1alpha1", Resource: "releases"}

var releasesKind = schema.GroupVersionKind{Group: "release.giantswarm.io", Version: "v1alpha1", Kind: "Release"}

// Get takes name of the release, and returns the corresponding release object, and an error if there is any.
func (c *FakeReleases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(releasesResource, name), &v1alpha1.Release{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Release), err
}

// List takes label and field selectors, and returns the list of Releases that match those selectors.
func (c *FakeReleases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReleaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(releasesResource, releasesKind, opts), &v1alpha1.ReleaseList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReleaseList{ListMeta: obj.(*v1alpha1.ReleaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReleaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested releases.
func (c *FakeReleases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(releasesResource, opts))
}

// Create takes the representation of a release and creates it.  Returns the server's representation of the release, and an error, if there is any.
func (c *FakeReleases) Create(ctx context.Context, release *v1alpha1.Release, opts v1.CreateOptions) (result *v1alpha1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(releasesResource, release), &v1alpha1.Release{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Release), err
}

// Update takes the representation of a release and updates it. Returns the server's representation of the release, and an error, if there is any.
func (c *FakeReleases) Update(ctx context.Context, release *v1alpha1.Release, opts v1.UpdateOptions) (result *v1alpha1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(releasesResource, release), &v1alpha1.Release{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Release), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReleases) UpdateStatus(ctx context.Context, release *v1alpha1.Release, opts v1.UpdateOptions) (*v1alpha1.Release, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(releasesResource, "status", release), &v1alpha1.Release{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Release), err
}

// Delete takes name of the release and deletes it. Returns an error if one occurs.
func (c *FakeReleases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(releasesResource, name), &v1alpha1.Release{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReleases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(releasesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReleaseList{})
	return err
}

// Patch applies the patch and returns the patched release.
func (c *FakeReleases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Release, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(releasesResource, name, pt, data, subresources...), &v1alpha1.Release{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Release), err
}
