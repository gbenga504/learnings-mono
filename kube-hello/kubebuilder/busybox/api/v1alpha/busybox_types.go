/*
Copyright 2024.

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

package v1alpha

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BusyboxSpec defines the desired state of Busybox.
type BusyboxSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size is the number of replicas of the busybox image that should be run
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:ExclusiveMaximum=false
	Size int32 `json:"size,omitempty"`
}

// BusyboxStatus defines the observed state of Busybox.
type BusyboxStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// These markers "+kubebuilder:object:root=true" tells our controller-gen tool that this is the root object of our API
// and these objects represents Kinds in Kubernetes API. The "+kubebuilder:subresource:status" marker tells the controller-gen tool to generate the status subresource for our API.
// Then the object genetator generates an implementation of the runtime.Object interface for us, which is the standard interface that all types representing Kinds must implement.
// check the zz_generated.deepcopy.go file for the implementation of DeepCopyInto, DeepCopy, DeepCopyObject functions for the Busybox and BusyboxList types which are methods defined on runtime.Object interface.

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Busybox is the Schema for the busyboxes API.
type Busybox struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BusyboxSpec   `json:"spec,omitempty"`
	Status BusyboxStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BusyboxList contains a list of Busybox.
type BusyboxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Busybox `json:"items"`
}

// Finally, we add the Go types to the API group. This allows us to add the types in this API group to any Scheme.
func init() {
	SchemeBuilder.Register(&Busybox{}, &BusyboxList{})
}
