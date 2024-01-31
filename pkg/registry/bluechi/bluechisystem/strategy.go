/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bluechisystem

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/bluechi-server/pkg/apis/bluechi/validation"

	"k8s.io/bluechi-server/pkg/apis/bluechi"
)

// NewStrategy creates and returns a bluechiStrategy instance
func NewStrategy(typer runtime.ObjectTyper) bluechiStrategy {
	return bluechiStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, and error in case the given runtime.Object is not a BlueChiSystem
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*bluechi.BlueChiSystem)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a BlueChiSystem")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchBlueChiSystem is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchBlueChiSystem(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *bluechi.BlueChiSystem) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type bluechiStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (bluechiStrategy) NamespaceScoped() bool {
	return true
}

func (bluechiStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (bluechiStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

func (bluechiStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	bluechi := obj.(*bluechi.BlueChiSystem)
	return validation.ValidateBlueChiSystem(bluechi)
}

// WarningsOnCreate returns warnings for the creation of the given object.
func (bluechiStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string { return nil }

func (bluechiStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (bluechiStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (bluechiStrategy) Canonicalize(obj runtime.Object) {
}

func (bluechiStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

// WarningsOnUpdate returns warnings for the given update.
func (bluechiStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}
