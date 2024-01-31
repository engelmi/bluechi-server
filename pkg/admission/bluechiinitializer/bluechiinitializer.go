/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bluechiinitializer

import (
	"k8s.io/apiserver/pkg/admission"
	informers "k8s.io/bluechi-server/pkg/generated/informers/externalversions"
)

type pluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = pluginInitializer{}

// New creates an instance of bluechi admission plugins initializer.
func New(informers informers.SharedInformerFactory) pluginInitializer {
	return pluginInitializer{
		informers: informers,
	}
}

// Initialize checks the initialization interfaces implemented by a plugin
// and provide the appropriate initialization data
func (i pluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsInternalBlueChiInformerFactory); ok {
		wants.SetInternalBlueChiInformerFactory(i.informers)
	}
}
