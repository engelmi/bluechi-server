/* SPDX-License-Identifier: LGPL-2.1-or-later */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	bluechiv1alpha1 "k8s.io/bluechi-server/pkg/apis/bluechi/v1alpha1"
	versioned "k8s.io/bluechi-server/pkg/generated/clientset/versioned"
	internalinterfaces "k8s.io/bluechi-server/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/bluechi-server/pkg/generated/listers/bluechi/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
)

// BlueChiSystemInformer provides access to a shared informer and lister for
// BlueChiSystems.
type BlueChiSystemInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BlueChiSystemLister
}

type blueChiSystemInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBlueChiSystemInformer constructs a new informer for BlueChiSystem type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBlueChiSystemInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBlueChiSystemInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBlueChiSystemInformer constructs a new informer for BlueChiSystem type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBlueChiSystemInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OrgV1alpha1().BlueChiSystems(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OrgV1alpha1().BlueChiSystems(namespace).Watch(context.TODO(), options)
			},
		},
		&bluechiv1alpha1.BlueChiSystem{},
		resyncPeriod,
		indexers,
	)
}

func (f *blueChiSystemInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBlueChiSystemInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *blueChiSystemInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bluechiv1alpha1.BlueChiSystem{}, f.defaultInformer)
}

func (f *blueChiSystemInformer) Lister() v1alpha1.BlueChiSystemLister {
	return v1alpha1.NewBlueChiSystemLister(f.Informer().GetIndexer())
}