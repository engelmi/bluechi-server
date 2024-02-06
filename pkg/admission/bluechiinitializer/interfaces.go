/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bluechiinitializer

import (
	informers "github.com/engelmi/bluechi-server/pkg/generated/informers/externalversions"
	"k8s.io/apiserver/pkg/admission"
)

// WantsInternalBlueChiInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsInternalBlueChiInformerFactory interface {
	SetInternalBlueChiInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
