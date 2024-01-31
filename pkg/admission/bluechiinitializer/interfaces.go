/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bluechiinitializer

import (
	"k8s.io/apiserver/pkg/admission"
	informers "k8s.io/bluechi-server/pkg/generated/informers/externalversions"
)

// WantsInternalBlueChiInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsInternalBlueChiInformerFactory interface {
	SetInternalBlueChiInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
