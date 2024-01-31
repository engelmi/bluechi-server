/* SPDX-License-Identifier: LGPL-2.1-or-later */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1alpha1 "k8s.io/bluechi-server/pkg/apis/bluechi/v1alpha1"
	bluechiv1alpha1 "k8s.io/bluechi-server/pkg/generated/applyconfiguration/bluechi/v1alpha1"
	testing "k8s.io/client-go/testing"
)

// FakeBlueChiSystems implements BlueChiSystemInterface
type FakeBlueChiSystems struct {
	Fake *FakeOrgV1alpha1
	ns   string
}

var bluechisystemsResource = v1alpha1.SchemeGroupVersion.WithResource("bluechisystems")

var bluechisystemsKind = v1alpha1.SchemeGroupVersion.WithKind("BlueChiSystem")

// Get takes name of the blueChiSystem, and returns the corresponding blueChiSystem object, and an error if there is any.
func (c *FakeBlueChiSystems) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BlueChiSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bluechisystemsResource, c.ns, name), &v1alpha1.BlueChiSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueChiSystem), err
}

// List takes label and field selectors, and returns the list of BlueChiSystems that match those selectors.
func (c *FakeBlueChiSystems) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BlueChiSystemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bluechisystemsResource, bluechisystemsKind, c.ns, opts), &v1alpha1.BlueChiSystemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BlueChiSystemList{ListMeta: obj.(*v1alpha1.BlueChiSystemList).ListMeta}
	for _, item := range obj.(*v1alpha1.BlueChiSystemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested blueChiSystems.
func (c *FakeBlueChiSystems) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bluechisystemsResource, c.ns, opts))

}

// Create takes the representation of a blueChiSystem and creates it.  Returns the server's representation of the blueChiSystem, and an error, if there is any.
func (c *FakeBlueChiSystems) Create(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.CreateOptions) (result *v1alpha1.BlueChiSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bluechisystemsResource, c.ns, blueChiSystem), &v1alpha1.BlueChiSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueChiSystem), err
}

// Update takes the representation of a blueChiSystem and updates it. Returns the server's representation of the blueChiSystem, and an error, if there is any.
func (c *FakeBlueChiSystems) Update(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.UpdateOptions) (result *v1alpha1.BlueChiSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bluechisystemsResource, c.ns, blueChiSystem), &v1alpha1.BlueChiSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueChiSystem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBlueChiSystems) UpdateStatus(ctx context.Context, blueChiSystem *v1alpha1.BlueChiSystem, opts v1.UpdateOptions) (*v1alpha1.BlueChiSystem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bluechisystemsResource, "status", c.ns, blueChiSystem), &v1alpha1.BlueChiSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueChiSystem), err
}

// Delete takes name of the blueChiSystem and deletes it. Returns an error if one occurs.
func (c *FakeBlueChiSystems) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(bluechisystemsResource, c.ns, name, opts), &v1alpha1.BlueChiSystem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBlueChiSystems) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bluechisystemsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BlueChiSystemList{})
	return err
}

// Patch applies the patch and returns the patched blueChiSystem.
func (c *FakeBlueChiSystems) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BlueChiSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bluechisystemsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BlueChiSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueChiSystem), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied blueChiSystem.
func (c *FakeBlueChiSystems) Apply(ctx context.Context, blueChiSystem *bluechiv1alpha1.BlueChiSystemApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.BlueChiSystem, err error) {
	if blueChiSystem == nil {
		return nil, fmt.Errorf("blueChiSystem provided to Apply must not be nil")
	}
	data, err := json.Marshal(blueChiSystem)
	if err != nil {
		return nil, err
	}
	name := blueChiSystem.Name
	if name == nil {
		return nil, fmt.Errorf("blueChiSystem.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bluechisystemsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.BlueChiSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueChiSystem), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeBlueChiSystems) ApplyStatus(ctx context.Context, blueChiSystem *bluechiv1alpha1.BlueChiSystemApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.BlueChiSystem, err error) {
	if blueChiSystem == nil {
		return nil, fmt.Errorf("blueChiSystem provided to Apply must not be nil")
	}
	data, err := json.Marshal(blueChiSystem)
	if err != nil {
		return nil, err
	}
	name := blueChiSystem.Name
	if name == nil {
		return nil, fmt.Errorf("blueChiSystem.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bluechisystemsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.BlueChiSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BlueChiSystem), err
}
