/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	rufio "github.com/tinkerbell/rufio/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InventoryWorkflowStatus string

type TinkWorkflowStatus string

type TaskWorkflowStatus string

type ConditionType string

const (
	InventoryFinalizer = "finalizer.inventory.harvesterhci.io"
)

const (
	BMCObjectCreated InventoryWorkflowStatus = "bmcObjectCreated"

	TinkHardwareSubmitted InventoryWorkflowStatus = "tinkerbellHardwareCreated"
	InventoryReady        InventoryWorkflowStatus = "inventoryNodeReady"
	InventoryRunning      InventoryWorkflowStatus = "inventoryRunning"
)

const (
	BMCTaskRequest      ConditionType = "bmcTaskRequested"
	BMCTaskSubmitted    ConditionType = "bmcTaskSubmitted"
	BMCTaskComplete     ConditionType = "bmcTaskCompleted"
	BMCTaskError        ConditionType = "bmcTaskError"
	TinkWorkflowCreated ConditionType = "tinkWorkflowCreated"
	TinkWorkflowError   ConditionType = "tinkWorkflowError"
)

// InventorySpec defines the desired state of Inventory
type InventorySpec struct {
	PrimaryDisk                   string `json:"primaryDisk"`
	PXEBootInterface              `json:"managementInterface"`
	rufio.BaseboardManagementSpec `json:"baseboardSpec"`
}

type BMCSecretReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type PXEBootInterface struct {
	Address     string   `json:"address,omitempty"`
	Gateway     string   `json:"gateway,omitempty"`
	Netmask     string   `json:"netmask,omitempty"`
	MacAddress  string   `json:"macAddress"`
	NameServers []string `json:"nameServers,omitempty"`
}

// InventoryStatus defines the observed state of Inventory
type InventoryStatus struct {
	Status            InventoryWorkflowStatus `json:"status,omitempty"`
	GeneratedPassword string                  `json:"generatedPassword,omitempty"`
	HardwareID        string                  `json:"hardwareID,omitempty"`
	Conditions        []Conditions            `json:"conditions,omitempty"`
}

type Conditions struct {
	Type           ConditionType `json:"type"`
	StartTime      metav1.Time   `json:"startTime"`
	LastUpdateTime metav1.Time   `json:"lastUpdateTime,omitempty"`
	Message        string        `json:"message,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Inventory is the Schema for the inventories API
type Inventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InventorySpec   `json:"spec,omitempty"`
	Status InventoryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InventoryList contains a list of Inventory
type InventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Inventory `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Inventory{}, &InventoryList{})
}