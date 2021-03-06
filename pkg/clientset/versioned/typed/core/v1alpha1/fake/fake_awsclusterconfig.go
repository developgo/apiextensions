/*
Copyright 2021 Giant Swarm GmbH.

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

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/core/v1alpha1"
)

// FakeAWSClusterConfigs implements AWSClusterConfigInterface
type FakeAWSClusterConfigs struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var awsclusterconfigsResource = schema.GroupVersionResource{Group: "core.giantswarm.io", Version: "v1alpha1", Resource: "awsclusterconfigs"}

var awsclusterconfigsKind = schema.GroupVersionKind{Group: "core.giantswarm.io", Version: "v1alpha1", Kind: "AWSClusterConfig"}

// Get takes name of the aWSClusterConfig, and returns the corresponding aWSClusterConfig object, and an error if there is any.
func (c *FakeAWSClusterConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsclusterconfigsResource, c.ns, name), &v1alpha1.AWSClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSClusterConfig), err
}

// List takes label and field selectors, and returns the list of AWSClusterConfigs that match those selectors.
func (c *FakeAWSClusterConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSClusterConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsclusterconfigsResource, awsclusterconfigsKind, c.ns, opts), &v1alpha1.AWSClusterConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSClusterConfigList{ListMeta: obj.(*v1alpha1.AWSClusterConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSClusterConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSClusterConfigs.
func (c *FakeAWSClusterConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsclusterconfigsResource, c.ns, opts))

}

// Create takes the representation of a aWSClusterConfig and creates it.  Returns the server's representation of the aWSClusterConfig, and an error, if there is any.
func (c *FakeAWSClusterConfigs) Create(ctx context.Context, aWSClusterConfig *v1alpha1.AWSClusterConfig, opts v1.CreateOptions) (result *v1alpha1.AWSClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsclusterconfigsResource, c.ns, aWSClusterConfig), &v1alpha1.AWSClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSClusterConfig), err
}

// Update takes the representation of a aWSClusterConfig and updates it. Returns the server's representation of the aWSClusterConfig, and an error, if there is any.
func (c *FakeAWSClusterConfigs) Update(ctx context.Context, aWSClusterConfig *v1alpha1.AWSClusterConfig, opts v1.UpdateOptions) (result *v1alpha1.AWSClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsclusterconfigsResource, c.ns, aWSClusterConfig), &v1alpha1.AWSClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSClusterConfig), err
}

// Delete takes name of the aWSClusterConfig and deletes it. Returns an error if one occurs.
func (c *FakeAWSClusterConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsclusterconfigsResource, c.ns, name), &v1alpha1.AWSClusterConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSClusterConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsclusterconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSClusterConfigList{})
	return err
}

// Patch applies the patch and returns the patched aWSClusterConfig.
func (c *FakeAWSClusterConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSClusterConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsclusterconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSClusterConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSClusterConfig), err
}
