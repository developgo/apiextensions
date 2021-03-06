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

	v1alpha2 "github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha2"
)

// FakeAWSClusters implements AWSClusterInterface
type FakeAWSClusters struct {
	Fake *FakeInfrastructureV1alpha2
	ns   string
}

var awsclustersResource = schema.GroupVersionResource{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Resource: "awsclusters"}

var awsclustersKind = schema.GroupVersionKind{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Kind: "AWSCluster"}

// Get takes name of the aWSCluster, and returns the corresponding aWSCluster object, and an error if there is any.
func (c *FakeAWSClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsclustersResource, c.ns, name), &v1alpha2.AWSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// List takes label and field selectors, and returns the list of AWSClusters that match those selectors.
func (c *FakeAWSClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.AWSClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsclustersResource, awsclustersKind, c.ns, opts), &v1alpha2.AWSClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.AWSClusterList{ListMeta: obj.(*v1alpha2.AWSClusterList).ListMeta}
	for _, item := range obj.(*v1alpha2.AWSClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSClusters.
func (c *FakeAWSClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsclustersResource, c.ns, opts))

}

// Create takes the representation of a aWSCluster and creates it.  Returns the server's representation of the aWSCluster, and an error, if there is any.
func (c *FakeAWSClusters) Create(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.CreateOptions) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsclustersResource, c.ns, aWSCluster), &v1alpha2.AWSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// Update takes the representation of a aWSCluster and updates it. Returns the server's representation of the aWSCluster, and an error, if there is any.
func (c *FakeAWSClusters) Update(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.UpdateOptions) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsclustersResource, c.ns, aWSCluster), &v1alpha2.AWSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSClusters) UpdateStatus(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.UpdateOptions) (*v1alpha2.AWSCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awsclustersResource, "status", c.ns, aWSCluster), &v1alpha2.AWSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// Delete takes name of the aWSCluster and deletes it. Returns an error if one occurs.
func (c *FakeAWSClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsclustersResource, c.ns, name), &v1alpha2.AWSCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.AWSClusterList{})
	return err
}

// Patch applies the patch and returns the patched aWSCluster.
func (c *FakeAWSClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsclustersResource, c.ns, name, pt, data, subresources...), &v1alpha2.AWSCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}
