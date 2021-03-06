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

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/security/v1alpha1"
)

// FakeOrganizations implements OrganizationInterface
type FakeOrganizations struct {
	Fake *FakeSecurityV1alpha1
}

var organizationsResource = schema.GroupVersionResource{Group: "security.giantswarm.io", Version: "v1alpha1", Resource: "organizations"}

var organizationsKind = schema.GroupVersionKind{Group: "security.giantswarm.io", Version: "v1alpha1", Kind: "Organization"}

// Get takes name of the organization, and returns the corresponding organization object, and an error if there is any.
func (c *FakeOrganizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Organization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(organizationsResource, name), &v1alpha1.Organization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Organization), err
}

// List takes label and field selectors, and returns the list of Organizations that match those selectors.
func (c *FakeOrganizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrganizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(organizationsResource, organizationsKind, opts), &v1alpha1.OrganizationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationList{ListMeta: obj.(*v1alpha1.OrganizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizations.
func (c *FakeOrganizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(organizationsResource, opts))
}

// Create takes the representation of a organization and creates it.  Returns the server's representation of the organization, and an error, if there is any.
func (c *FakeOrganizations) Create(ctx context.Context, organization *v1alpha1.Organization, opts v1.CreateOptions) (result *v1alpha1.Organization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(organizationsResource, organization), &v1alpha1.Organization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Organization), err
}

// Update takes the representation of a organization and updates it. Returns the server's representation of the organization, and an error, if there is any.
func (c *FakeOrganizations) Update(ctx context.Context, organization *v1alpha1.Organization, opts v1.UpdateOptions) (result *v1alpha1.Organization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(organizationsResource, organization), &v1alpha1.Organization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Organization), err
}

// Delete takes name of the organization and deletes it. Returns an error if one occurs.
func (c *FakeOrganizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(organizationsResource, name), &v1alpha1.Organization{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(organizationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationList{})
	return err
}

// Patch applies the patch and returns the patched organization.
func (c *FakeOrganizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Organization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(organizationsResource, name, pt, data, subresources...), &v1alpha1.Organization{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Organization), err
}
