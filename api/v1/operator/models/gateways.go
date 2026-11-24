package models

// GateWays Status
//
// +k8s:deepcopy-gen=true
//
// swagger:model GateWays

// GateWaysConfig Cluster configuration exposed by the remote cluster
//
// +k8s:deepcopy-gen=true
//
// swagger:model GateWaysConfig
type GateWaysConfig struct {

	// The NameSpace
	NameSpace string `json:"nameSpace,omitempty"`

	// Whether the remote cluster information is locally cached by kvstoremesh
	Listeners int64 `json:"Listeners,omitempty"`

	// Whether the configuration is required to be present
	TLS bool `json:"TLS,omitempty"`
}
