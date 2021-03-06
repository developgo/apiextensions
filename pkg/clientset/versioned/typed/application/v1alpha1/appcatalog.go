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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/application/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/v3/pkg/clientset/versioned/scheme"
)

// AppCatalogsGetter has a method to return a AppCatalogInterface.
// A group's client should implement this interface.
type AppCatalogsGetter interface {
	AppCatalogs() AppCatalogInterface
}

// AppCatalogInterface has methods to work with AppCatalog resources.
type AppCatalogInterface interface {
	Create(ctx context.Context, appCatalog *v1alpha1.AppCatalog, opts v1.CreateOptions) (*v1alpha1.AppCatalog, error)
	Update(ctx context.Context, appCatalog *v1alpha1.AppCatalog, opts v1.UpdateOptions) (*v1alpha1.AppCatalog, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AppCatalog, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AppCatalogList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppCatalog, err error)
	AppCatalogExpansion
}

// appCatalogs implements AppCatalogInterface
type appCatalogs struct {
	client rest.Interface
}

// newAppCatalogs returns a AppCatalogs
func newAppCatalogs(c *ApplicationV1alpha1Client) *appCatalogs {
	return &appCatalogs{
		client: c.RESTClient(),
	}
}

// Get takes name of the appCatalog, and returns the corresponding appCatalog object, and an error if there is any.
func (c *appCatalogs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AppCatalog, err error) {
	result = &v1alpha1.AppCatalog{}
	err = c.client.Get().
		Resource("appcatalogs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppCatalogs that match those selectors.
func (c *appCatalogs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AppCatalogList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppCatalogList{}
	err = c.client.Get().
		Resource("appcatalogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appCatalogs.
func (c *appCatalogs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("appcatalogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a appCatalog and creates it.  Returns the server's representation of the appCatalog, and an error, if there is any.
func (c *appCatalogs) Create(ctx context.Context, appCatalog *v1alpha1.AppCatalog, opts v1.CreateOptions) (result *v1alpha1.AppCatalog, err error) {
	result = &v1alpha1.AppCatalog{}
	err = c.client.Post().
		Resource("appcatalogs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appCatalog).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a appCatalog and updates it. Returns the server's representation of the appCatalog, and an error, if there is any.
func (c *appCatalogs) Update(ctx context.Context, appCatalog *v1alpha1.AppCatalog, opts v1.UpdateOptions) (result *v1alpha1.AppCatalog, err error) {
	result = &v1alpha1.AppCatalog{}
	err = c.client.Put().
		Resource("appcatalogs").
		Name(appCatalog.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(appCatalog).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the appCatalog and deletes it. Returns an error if one occurs.
func (c *appCatalogs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("appcatalogs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appCatalogs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("appcatalogs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched appCatalog.
func (c *appCatalogs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AppCatalog, err error) {
	result = &v1alpha1.AppCatalog{}
	err = c.client.Patch(pt).
		Resource("appcatalogs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
