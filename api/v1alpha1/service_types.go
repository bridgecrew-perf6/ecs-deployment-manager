/*
Copyright 2022 br4sk4.

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
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Cluster        string              `json:"cluster,omitempty"`
	LaunchType     ecsTypes.LaunchType `json:"launchType,omitempty"`
	DesiredCount   int                 `json:"desiredCount,omitempty"`
	Subnets        []string            `json:"subnets,omitempty"`
	SecurityGroups []string            `json:"securityGroups,omitempty"`
	ContainerPort  int                 `json:"containerPort,omitempty"`
}

// ServiceStatus defines the observed state of Service
type ServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Synced     bool   `json:"synced,omitempty"`
	ServiceArn string `json:"serviceArn,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Service is the Schema for the services API
// +kubebuilder:printcolumn:name="Synced",type=string,JSONPath=`.status.synced`
// +kubebuilder:printcolumn:name="ARN",type=string,JSONPath=`.status.serviceArn`
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceSpec   `json:"spec,omitempty"`
	Status ServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServiceList contains a list of Service
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
