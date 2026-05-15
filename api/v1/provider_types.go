package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Provider defines the way GEC should interact with the git provider
// +kubebuilder:object:root=true
type Provider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProviderSpec `json:"spec,omitempty"`
}

// ProviderSpec defines the desired state of the Provider
type ProviderSpec struct {
	// BaseURI is the base URI to join the repository. Trailing slash will be removed.
	// +kubebuilder:example=github.com/hermaeus-project
	BaseURI string `json:"baseURI"`

	// GitopsClass name managing the Provider
	GitopsClass string `json:"gitopsClass,omitempty"`

	// GitConnection defines the protocol and authentication method to the git repository. If omitted, the provider must be public
	// +optional
	Git GitConnection `json:"git,omitempty"`

	// Service defines what provider is used (Gitlab, GitHub, Forgejo, etc). If omitted, only bare git feature will be supported by the provider.
	// +optional
	Service ServiceConnection `json:"service,omitempty"`
}

// GitConnection defines the protocol and authentication method to the git repository
type GitConnection struct {
	// SSH enables and configures the credentials
	SSH GitConnectionSSH `json:"ssh,omitempty"`

	// HTTPS enables and configures the credentials
	HTTPS GitConnectionHTTPS `json:"https,omitempty"`
}

type GitConnectionSSH struct {
	// SecretName is the name of the secret containing the credentials. Must be in the same namespace as the GEC instances. Can be omitted in favor of unsecure username and key during exploration phase.
	SecretName string `json:"secretName,omitempty"`

	// KeyUsername is the key in the secret mapped to the username value
	KeyUsername string `json:"keyUsername,omitempty"`

	// KeyPrivateKey is the key in the secret mapped to the private key value
	KeyPrivateKey string `json:"keyPrivateKey,omitempty"`

	// UnsecureUsername provide an easy way to configure SSH. For exploration purposes only!
	UnsecureUsername string `json:"unsecureUsername,omitempty"`

	// UnsecurePrivateKey provide an easy way to configure SSH. For exploration purposes only!
	UnsecurePrivateKey string `json:"unsecurePrivateKey,omitempty"`
}

type GitConnectionHTTPS struct {
	// EnableTLS enforces HTTPS
	EnableTLS bool `json:"enableTLS,omitempty"`

	// BasicAuth enables and configure basic authentication to the git repository
	BasicAuth GitConnectionBasicAuth `json:"basicAuth,omitempty"`
}

type GitConnectionBasicAuth struct {
	// SecretName is the name of the secret containing the credentials. Must be in the same namespace as the GEC instances. Can be omitted in favor of unsecure username and password during exploration phase.
	SecretName string `json:"secretName,omitempty"`

	// KeyUsername is the key in the secret mapped to the username value
	KeyUsername string `json:"keyUsername,omitempty"`

	// KeyPassword is the key in the secret mapped to the password value
	KeyPassword string `json:"keyPassword,omitempty"`

	// UnsecureUsername provide an easy way to configure BasicAuth. For exploration purposes only!
	UnsecureUsername string `json:"unsecureUsername,omitempty"`

	// UnsecurePassword provide an easy way to configure BasicAuth. For exploration purposes only!
	UnsecurePassword string `json:"unsecurePassword,omitempty"`
}

// ServiceConnection defines what remote services is used by the provider and authentification.
type ServiceConnection struct {
	// Forgejo defines a Forgejo remote server
	// +optional
	Forgejo ProviderForgejo `json:"forgejo,omitempty"`

	// Gitlab defines a Gitlab remote server
	// +optional
	Gitlab ProviderGitlab `json:"gitlab,omitempty"`

	// GitHub defines a GitHub remote server
	// +optional
	Github ProviderGithub `json:"github,omitempty"`

	// TODO: add new supported providers here
}

type ProviderForgejo struct{}

type ProviderGitlab struct{}

type ProviderGithub struct{}
