package models

import (
    "encoding/json"
)

// TunnelConfigsAutoProvisionNode represents a TunnelConfigsAutoProvisionNode struct.
type TunnelConfigsAutoProvisionNode struct {
    NumHosts             *string        `json:"num_hosts,omitempty"`
    // optional, only needed if `vars_only`==`false`
    WanNames             []string       `json:"wan_names,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TunnelConfigsAutoProvisionNode.
// It customizes the JSON marshaling process for TunnelConfigsAutoProvisionNode objects.
func (t TunnelConfigsAutoProvisionNode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(t.toMap())
}

// toMap converts the TunnelConfigsAutoProvisionNode object to a map representation for JSON marshaling.
func (t TunnelConfigsAutoProvisionNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, t.AdditionalProperties)
    if t.NumHosts != nil {
        structMap["num_hosts"] = t.NumHosts
    }
    if t.WanNames != nil {
        structMap["wan_names"] = t.WanNames
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelConfigsAutoProvisionNode.
// It customizes the JSON unmarshaling process for TunnelConfigsAutoProvisionNode objects.
func (t *TunnelConfigsAutoProvisionNode) UnmarshalJSON(input []byte) error {
    var temp tempTunnelConfigsAutoProvisionNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_hosts", "wan_names")
    if err != nil {
    	return err
    }
    
    t.AdditionalProperties = additionalProperties
    t.NumHosts = temp.NumHosts
    t.WanNames = temp.WanNames
    return nil
}

// tempTunnelConfigsAutoProvisionNode is a temporary struct used for validating the fields of TunnelConfigsAutoProvisionNode.
type tempTunnelConfigsAutoProvisionNode  struct {
    NumHosts *string  `json:"num_hosts,omitempty"`
    WanNames []string `json:"wan_names,omitempty"`
}
