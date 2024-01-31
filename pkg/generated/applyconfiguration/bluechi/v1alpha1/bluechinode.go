/* SPDX-License-Identifier: LGPL-2.1-or-later */

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "k8s.io/bluechi-server/pkg/apis/bluechi/v1alpha1"
)

// BlueChiNodeApplyConfiguration represents an declarative configuration of the BlueChiNode type for use
// with apply.
type BlueChiNodeApplyConfiguration struct {
	Name              *string              `json:"name,omitempty"`
	Status            *v1alpha1.NodeStatus `json:"status,omitempty"`
	LastSeenTimestamp *string              `json:"lastSeenTimestamp,omitempty"`
}

// BlueChiNodeApplyConfiguration constructs an declarative configuration of the BlueChiNode type for use with
// apply.
func BlueChiNode() *BlueChiNodeApplyConfiguration {
	return &BlueChiNodeApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *BlueChiNodeApplyConfiguration) WithName(value string) *BlueChiNodeApplyConfiguration {
	b.Name = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *BlueChiNodeApplyConfiguration) WithStatus(value v1alpha1.NodeStatus) *BlueChiNodeApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastSeenTimestamp sets the LastSeenTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastSeenTimestamp field is set to the value of the last call.
func (b *BlueChiNodeApplyConfiguration) WithLastSeenTimestamp(value string) *BlueChiNodeApplyConfiguration {
	b.LastSeenTimestamp = &value
	return b
}
