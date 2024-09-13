package models

import (
    "encoding/json"
)

// UtilsShowDhcpLeases represents a UtilsShowDhcpLeases struct.
type UtilsShowDhcpLeases struct {
    // DHCP network for the leases, returns full table if not specified
    Network              *string            `json:"network,omitempty"`
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowDhcpLeases.
// It customizes the JSON marshaling process for UtilsShowDhcpLeases objects.
func (u UtilsShowDhcpLeases) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowDhcpLeases object to a map representation for JSON marshaling.
func (u UtilsShowDhcpLeases) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Network != nil {
        structMap["network"] = u.Network
    }
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowDhcpLeases.
// It customizes the JSON unmarshaling process for UtilsShowDhcpLeases objects.
func (u *UtilsShowDhcpLeases) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowDhcpLeases
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "network", "node")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Network = temp.Network
    u.Node = temp.Node
    return nil
}

// tempUtilsShowDhcpLeases is a temporary struct used for validating the fields of UtilsShowDhcpLeases.
type tempUtilsShowDhcpLeases  struct {
    Network *string            `json:"network,omitempty"`
    Node    *HaClusterNodeEnum `json:"node,omitempty"`
}
