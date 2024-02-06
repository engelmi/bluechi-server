/* SPDX-License-Identifier: LGPL-2.1-or-later */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/engelmi/bluechi-server/pkg/apis/bluechi/v1alpha1"
	bluechiv1alpha1 "github.com/engelmi/bluechi-server/pkg/generated/applyconfiguration/bluechi/v1alpha1"
	scheme "github.com/engelmi/bluechi-server/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BlueChiSystemsGetter has a method to return a BlueChiSystemInterface.
// A group's client should implement this interface.
type BlueChiSystemsGetter interface {
	BlueChiSystems(namespace string) BlueChiSystemInterface
}

// BlueChiSystemInterface has methods to work with BlueChiSystem resources.
type BlueChiSystemInterface interface {
	Create(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.CreateOptions) (*v1alpha1.BlueChiSystem, error)
	Update(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.UpdateOptions) (*v1alpha1.BlueChiSystem, error)
	UpdateStatus(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.UpdateOptions) (*v1alpha1.BlueChiSystem, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BlueChiSystem, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BlueChiSystemList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BlueChiSystem, err error)
	Apply(ctx context.Context, blueChiSystem *bluechiv1alpha1.BlueChiSystemApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.BlueChiSystem, err error)
	ApplyStatus(ctx context.Context, blueChiSystem *bluechiv1alpha1.BlueChiSystemApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.BlueChiSystem, err error)
	BlueChiSystemExpansion
}

// blueChiSystems implements BlueChiSystemInterface
type blueChiSystems struct {
	client rest.Interface
	ns     string
}

// newBlueChiSystems returns a BlueChiSystems
func newBlueChiSystems(c *OrgV1alpha1Client, namespace string) *blueChiSystems {
	return &blueChiSystems{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the blueChiSystem, and returns the corresponding blueChiSystem object, and an error if there is any.
func (c *blueChiSystems) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BlueChiSystem, err error) {
	result = &v1alpha1.BlueChiSystem{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bluechisystems").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BlueChiSystems that match those selectors.
func (c *blueChiSystems) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BlueChiSystemList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BlueChiSystemList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bluechisystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested blueChiSystems.
func (c *blueChiSystems) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bluechisystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a blueChiSystem and creates it.  Returns the server's representation of the blueChiSystem, and an error, if there is any.
func (c *blueChiSystems) Create(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.CreateOptions) (result *v1alpha1.BlueChiSystem, err error) {
	result = &v1alpha1.BlueChiSystem{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bluechisystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blueChiSystem).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a blueChiSystem and updates it. Returns the server's representation of the blueChiSystem, and an error, if there is any.
func (c *blueChiSystems) Update(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.UpdateOptions) (result *v1alpha1.BlueChiSystem, err error) {
	result = &v1alpha1.BlueChiSystem{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bluechisystems").
		Name(blueChiSystem.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blueChiSystem).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *blueChiSystems) UpdateStatus(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.UpdateOptions) (result *v1alpha1.BlueChiSystem, err error) {
	result = &v1alpha1.BlueChiSystem{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bluechisystems").
		Name(blueChiSystem.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(blueChiSystem).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the blueChiSystem and deletes it. Returns an error if one occurs.
func (c *blueChiSystems) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bluechisystems").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *blueChiSystems) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bluechisystems").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched blueChiSystem.
func (c *blueChiSystems) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BlueChiSystem, err error) {
	result = &v1alpha1.BlueChiSystem{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bluechisystems").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied blueChiSystem.
func (c *blueChiSystems) Apply(ctx context.Context, blueChiSystem *bluechiv1alpha1.BlueChiSystemApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.BlueChiSystem, err error) {
	if blueChiSystem == nil {
		return nil, fmt.Errorf("blueChiSystem provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(blueChiSystem)
	if err != nil {
		return nil, err
	}
	name := blueChiSystem.Name
	if name == nil {
		return nil, fmt.Errorf("blueChiSystem.Name must be provided to Apply")
	}
	result = &v1alpha1.BlueChiSystem{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("bluechisystems").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *blueChiSystems) ApplyStatus(ctx context.Context, blueChiSystem *bluechiv1alpha1.BlueChiSystemApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.BlueChiSystem, err error) {
	if blueChiSystem == nil {
		return nil, fmt.Errorf("blueChiSystem provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(blueChiSystem)
	if err != nil {
		return nil, err
	}

	name := blueChiSystem.Name
	if name == nil {
		return nil, fmt.Errorf("blueChiSystem.Name must be provided to Apply")
	}

	result = &v1alpha1.BlueChiSystem{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("bluechisystems").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
