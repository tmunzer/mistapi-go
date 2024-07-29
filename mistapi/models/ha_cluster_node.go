package models

import (
    "encoding/json"
)

// HaClusterNode represents a HaClusterNode struct.
type HaClusterNode struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for HaClusterNode.
// It customizes the JSON marshaling process for HaClusterNode objects.
func (h HaClusterNode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(h.toMap())
}

// toMap converts the HaClusterNode object to a map representation for JSON marshaling.
func (h HaClusterNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, h.AdditionalProperties)
    if h.Node != nil {
        structMap["node"] = h.Node
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for HaClusterNode.
// It customizes the JSON unmarshaling process for HaClusterNode objects.
func (h *HaClusterNode) UnmarshalJSON(input []byte) error {
    var temp haClusterNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "node")
    if err != nil {
    	return err
    }
    
    h.AdditionalProperties = additionalProperties
    h.Node = temp.Node
    return nil
}

// haClusterNode is a temporary struct used for validating the fields of HaClusterNode.
type haClusterNode  struct {
    Node *HaClusterNodeEnum `json:"node,omitempty"`
}
