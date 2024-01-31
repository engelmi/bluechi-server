/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bluechi

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
	Name              string
	Status            NodeStatus
	LastSeenTimestamp string
}

type BlueChiNodes []BlueChiNode

type BlueChiSystemSpec struct {
	Nodes BlueChiNodes
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BlueChiSystem struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   BlueChiSystemSpec
	Status BlueChiSystemStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BlueChiSystemList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []BlueChiSystem
}
