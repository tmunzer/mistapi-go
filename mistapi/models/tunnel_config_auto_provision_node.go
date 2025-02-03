package models

import (
    "encoding/json"
    "fmt"
)

// TunnelConfigAutoProvisionNode represents a TunnelConfigAutoProvisionNode struct.
type TunnelConfigAutoProvisionNode struct {
    ProbeIps             []string               `json:"probe_ips,omitempty"`
    // Optional, only needed if `vars_only`==`false`
    WanNames             []string               `json:"wan_names,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelConfigAutoProvisionNode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelConfigAutoProvisionNode) String() string {
    return fmt.Sprintf(
    	"TunnelConfigAutoProvisionNode[ProbeIps=%v, WanNames=%v, AdditionalProperties=%v]",
    	t.ProbeIps, t.WanNames, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigAutoProvisionNode.
// It customizes the JSON marshaling process for TunnelConfigAutoProvisionNode objects.
func (t TunnelConfigAutoProvisionNode) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "probe_ips", "wan_names"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigAutoProvisionNode object to a map representation for JSON marshaling.
func (t TunnelConfigAutoProvisionNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.ProbeIps != nil {
        structMap["probe_ips"] = t.ProbeIps
    }
    if t.WanNames != nil {
        structMap["wan_names"] = t.WanNames
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigAutoProvisionNode.
// It customizes the JSON unmarshaling process for TunnelConfigAutoProvisionNode objects.
func (t *TunnelConfigAutoProvisionNode) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigAutoProvisionNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "probe_ips", "wan_names")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.ProbeIps = temp.ProbeIps
    t.WanNames = temp.WanNames
    return nil
}

// tempTunnelConfigAutoProvisionNode is a temporary struct used for validating the fields of TunnelConfigAutoProvisionNode.
type tempTunnelConfigAutoProvisionNode  struct {
    ProbeIps []string `json:"probe_ips,omitempty"`
    WanNames []string `json:"wan_names,omitempty"`
}
