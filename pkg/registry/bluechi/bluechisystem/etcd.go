/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bluechisystem

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/bluechi-server/pkg/apis/bluechi"
	"k8s.io/bluechi-server/pkg/registry"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		NewFunc:                   func() runtime.Object { return &bluechi.BlueChiSystem{} },
		NewListFunc:               func() runtime.Object { return &bluechi.BlueChiSystemList{} },
		PredicateFunc:             MatchBlueChiSystem,
		DefaultQualifiedResource:  bluechi.Resource("bluechisystems"),
		SingularQualifiedResource: bluechi.Resource("bluechisystem"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		// TODO: define table converter that exposes more than name/creation timestamp
		TableConvertor: rest.NewDefaultTableConvertor(bluechi.Resource("bluechisystems")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{Store: store}, nil
}
