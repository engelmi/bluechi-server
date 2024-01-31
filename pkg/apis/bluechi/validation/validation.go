/* SPDX-License-Identifier: LGPL-2.1-or-later */

package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/bluechi-server/pkg/apis/bluechi"
)

// ValidateBlueChiSystem validates a BlueChiSystem.
func ValidateBlueChiSystem(f *bluechi.BlueChiSystem) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, ValidateBlueChiSystemSpec(&f.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidateBlueChiSystemSpec validates a BlueChiSystemSpec.
func ValidateBlueChiSystemSpec(s *bluechi.BlueChiSystemSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	return allErrs
}
