/* SPDX-License-Identifier: LGPL-2.1-or-later */

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type BlueChiSystemStatus string

const (
	SystemStatusUp       = BlueChiSystemStatus("up")
	SystemStatusDegraded = BlueChiSystemStatus("degraded")
	SystemStatusDown     = BlueChiSystemStatus("down")
)

type NodeStatus string

const (
	NodeStatusOnline  = NodeStatus("online")
	NodeStatusOffline = NodeStatus("offline")
)

type BlueChiNode struct {
	Name              string     `json:"name,omitempty"`
	Status            NodeStatus `json:"status,omitempty"`
	LastSeenTimestamp string     `json:"lastSeenTimestamp,omitempty"`
}

type BlueChiNodes []BlueChiNode

type BlueChiSystemSpec struct {
	Nodes BlueChiNodes `json:"nodes,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BlueChiSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlueChiSystemSpec   `json:"spec,omitempty"`
	Status BlueChiSystemStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BlueChiSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []BlueChiSystem `json:"items"`
}
