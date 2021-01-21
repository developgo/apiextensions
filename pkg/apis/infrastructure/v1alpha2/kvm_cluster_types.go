package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api/api/v1alpha3"
)

const (
	KindKVMCluster = "KVMCluster"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=giantswarm;kvm
// +k8s:openapi-gen=true

// KVMCluster is the infrastructure provider referenced in upstream Cluster API Cluster CRs for the Giant Swarm KVM
// platform.
type KVMCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KVMClusterSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	Status KVMClusterStatus `json:"status"`
}

// KVMClusterSpec is the spec part for the KVMCluster resource.
// +k8s:openapi-gen=true
type KVMClusterSpec struct {
	// Endpoint used to connect to the target cluster's Kubernetes API server.
	ControlPlaneEndpoint v1alpha3.APIEndpoint `json:"controlPlaneEndpoint"`
	// Cluster specification details.
	Cluster KVMClusterSpecCluster `json:"cluster"`
	// Provider-specific configuration details.
	Provider KVMClusterSpecProvider `json:"provider"`
}

// KVMClusterSpecCluster provides general cluster specification details.
// +k8s:openapi-gen=true
type KVMClusterSpecCluster struct {
	// User-friendly description that should explain the purpose of the cluster to humans.
	Description string `json:"description"`
	// IP range of the cluster pods.
	ClusterIPRange string `json:"clusterIPRange"`
	// DNS configuration details.
	DNS KVMClusterSpecClusterDNS `json:"dns"`
	// Configuration for OpenID Connect (OIDC) authentication.
	OIDC KVMClusterSpecClusterOIDC `json:"oidc,omitempty"`
	// Calico configuration.
	Network KVMClusterSpecClusterNetwork `json:"network"`
	// Nodes is a slice representing the nodes of the cluster.
	Nodes []KVMClusterSpecClusterNode `json:"nodes"`
}

// KVMClusterSpecClusterNetwork describes the Calico CNI network configuration.
// +k8s:openapi-gen=true
type KVMClusterSpecClusterNetwork struct {
	// MTU of overlay network.
	MTU int `json:"role"`
	// Subnet IP of overlay network.
	Subnet string `json:"subnet"`
	// Subnet mask bits of overlay network.
	CIDR int `json:"cidr"`
}

// KVMClusterSpecClusterNode describes a single node in a cluster.
// +k8s:openapi-gen=true
type KVMClusterSpecClusterNode struct {
	// User-friendly description that should explain the purpose of the cluster to humans.
	ID string `json:"id"`
	// DNS configuration details.
	Role string `json:"role"`
	// Machine size details.
	Size KVMMachineSpecSize `json:"size"`
	// Reference to a provider-specific resource.
	InfrastructureRef corev1.TypedLocalObjectReference `json:"infrastructureRef"`
}

// KVMClusterSpecClusterOIDC holds configuration for OpenID Connect (OIDC) authentication.
// +k8s:openapi-gen=true
type KVMClusterSpecClusterOIDC struct {
	Claims    KVMClusterSpecClusterOIDCClaims `json:"claims,omitempty"`
	ClientID  string                          `json:"clientID,omitempty"`
	IssuerURL string                          `json:"issuerURL,omitempty"`
}

// KVMClusterSpecClusterOIDCClaims defines OIDC claims.
// +k8s:openapi-gen=true
type KVMClusterSpecClusterOIDCClaims struct {
	Username string `json:"username,omitempty"`
	Groups   string `json:"groups,omitempty"`
}

// KVMClusterSpecClusterDNS holds DNS configuration details.
// +k8s:openapi-gen=true
type KVMClusterSpecClusterDNS struct {
	Domain string `json:"domain"`
}

// KVMClusterSpecProvider provides provider-specific cluster specification.
// +k8s:openapi-gen=true
type KVMClusterSpecProvider struct {
	EndpointUpdaterImage string                               `json:"endpointUpdaterImage"`
	MachineImage         string                               `json:"machineImage"`
	MachineStorageType   string                               `json:"machineStorageType"`
	FlannelVNI           int                                  `json:"flannelVNI"`
	PortMappings         []KVMClusterSpecProviderPortMappings `json:"portMappings"`
}

// +k8s:openapi-gen=true
type KVMClusterSpecProviderPortMappings struct {
	Name       string `json:"name"`
	NodePort   int    `json:"nodePort"`
	TargetPort int    `json:"targetPort"`
}

// AWSClusterStatus holds status information about the cluster, populated once the cluster is in creation or created.
// +k8s:openapi-gen=true
type KVMClusterStatus struct {
	// +kubebuilder:validation:Optional
	// True when the infrastructure is ready to be used.
	Ready bool `json:"ready"`
	// +kubebuilder:validation:Optional
	// Cluster-specific status details, including conditions and versions.
	Cluster CommonClusterStatus `json:"cluster,omitempty"`
	// +kubebuilder:validation:Optional
	// Provider-specific status details.
	Provider KVMClusterStatusProvider `json:"provider,omitempty"`
}

// AWSClusterStatusProvider holds provider-specific status details.
// +k8s:openapi-gen=true
type KVMClusterStatusProvider struct {
	// +kubebuilder:validation:Optional
	// +nullable
	// NodeIndexes is a map from nodeID -> nodeIndex. This is used to create deterministic iSCSI initiator names.
	NodeIndexes map[string]int `json:"nodeIndexes"`
}

// KVMClusterList is the type returned when listing KVMCLuster resources.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KVMClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KVMCluster `json:"items"`
}
