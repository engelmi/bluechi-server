/* SPDX-License-Identifier: LGPL-2.1-or-later */

package install

import (
	"github.com/engelmi/bluechi-server/pkg/apis/bluechi"
	"github.com/engelmi/bluechi-server/pkg/apis/bluechi/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(bluechi.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}
