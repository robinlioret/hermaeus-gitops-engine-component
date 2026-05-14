package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LeaderState int

const (
	LeaderStateMissing LeaderState = iota
	LeaderStateHealthy
)

// GitopsClass defines the scope for a set of GEC (often one tool)
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type GitopsClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitopsClassSpec   `json:"spec,omitempty"`
	Status GitopsClassStatus `json:"status,omitempty"`
}

// GitopsClassSpec defines the desired state of GitopsClass
type GitopsClassSpec struct {
	// Leader election policy
	// +optional
	Leader LeaderSpec `json:"leader,omitempty"`
}

// LeaderSpec defines the policy regarding the election of a new leader
type LeaderSpec struct {
	// Duration between two update of the leader status fields by the leader.
	// +optional
	ReportPeriod metav1.Duration `json:"reportPeriod,omitempty"`

	// Duration before triggering a new leader election.
	// +optional
	ReelectionTimeout metav1.Duration `json:"reelectionTimeout,omitempty"`
}

// GitopsClassStatus defines the observed state of GitopsClass
// +kubebuilder:object:generate=true
type GitopsClassStatus struct {
	// Leader current status
	// +optional
	Leader LeaderStatus `json:"leader,omitempty"`
}

// LeaderStatus defines the observed state of the Leader
// +kubebuilder:object:generate=true
type LeaderStatus struct {
	// Name of the leader. Namespace and Pod name
	// +optional
	Name string `json:"name"`

	// State of the leader.
	// +optional
	// +kubebuilder:default=Missing
	State LeaderState `json:"state"`

	// LastReportAt is the last time the leader reported healthy by editing the GitopsClass resource.
	// +optional
	LastReportAt metav1.Time `json:"lastReportAt,omitempty"`
}

// GitopsClassList contains a list of GitopsClass
// +kubebuilder:object:root=true
type GitopsClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitopsClass `json:"items"`
}
