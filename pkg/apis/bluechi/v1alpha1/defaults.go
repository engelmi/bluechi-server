/* SPDX-License-Identifier: LGPL-2.1-or-later */

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_BlueChiSystemSpec sets defaults for BlueChiSystem spec
func SetDefaults_BlueChiSystemSpec(obj *BlueChiSystemSpec) {

}
