package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Worktree defines a revision to load and possibly manage in the HEGEC
// +kubebuilder:object:root=true
type Worktree struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorktreeSpec `json:"spec,omitempty"`
}

type WorktreeSpec struct {
	// RepositoryRef defines a reference to the repository the worktree belongs to
	RepositoryRef corev1.LocalObjectReference `json:"repositoryRef"`

	// WorktreeReference defines the reference to the revision
	Ref WorktreeReference `json:"ref"`
}

type WorktreeReference struct {
	// Branch explicitly defines the revision as a branch
	// +optional
	Branch string `json:"branch,omitempty"`

	// Tag explicitly defines the revision as a tag
	// +optional
	Tag string `json:"tag,omitempty"`

	// Commit explicitly defines the revision as a commit SHA
	// +optional
	Commit string `json:"commit,omitempty"`
}

// WorktreeList contains a list of Worktree
// +kubebuilder:object:root=true
type WorktreeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Worktree `json:"items"`
}
