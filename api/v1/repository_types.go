package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Repository defines a Repository accessible to HEGEC
// +kubebuilder:object:root=true
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RepositorySpec `json:"spec,omitempty"`
}

// RepositorySpec defines the spec of a Repository
type RepositorySpec struct {
	// ProviderRef is a reference to the Provider for this repository.
	ProviderRef corev1.LocalObjectReference `json:"providerName"`

	// Path is append to the provider baseURI to get the full path to the repo
	// +kubebuilder:example=hermaeus-gec
	Path string `json:"path"`
}

// RepositoryList contains a list of Repository
// +kubebuilder:object:root=true
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}
