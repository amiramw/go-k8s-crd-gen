package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExtensionsConfig is a top-level type. A client is created for it.
type ExtensionsConfig struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Status TestTypeStatus `json:"status,omitempty"`
	// +optional
	//string author `json:author`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExtensionsConfigList is a top-level list type. The client methods for lists are automatically created.
// You are not supposed to create a separated client for this one.
type ExtensionsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ExtensionsConfig `json:"items"`
}

type TestTypeStatus struct {
	Blah string
}
