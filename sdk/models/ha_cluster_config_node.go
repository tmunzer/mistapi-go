package models

import (
    "encoding/json"
)

// HaClusterConfigNode represents a HaClusterConfigNode struct.
type HaClusterConfigNode struct {
    // node mac, should be unassigned
    Mac                  *string        `json:"mac,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for HaClusterConfigNode.
// It customizes the JSON marshaling process for HaClusterConfigNode objects.
func (h HaClusterConfigNode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(h.toMap())
}

// toMap converts the HaClusterConfigNode object to a map representation for JSON marshaling.
func (h HaClusterConfigNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, h.AdditionalProperties)
    if h.Mac != nil {
        structMap["mac"] = h.Mac
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for HaClusterConfigNode.
// It customizes the JSON unmarshaling process for HaClusterConfigNode objects.
func (h *HaClusterConfigNode) UnmarshalJSON(input []byte) error {
    var temp haClusterConfigNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac")
    if err != nil {
    	return err
    }
    
    h.AdditionalProperties = additionalProperties
    h.Mac = temp.Mac
    return nil
}

// haClusterConfigNode is a temporary struct used for validating the fields of HaClusterConfigNode.
type haClusterConfigNode  struct {
    Mac *string `json:"mac,omitempty"`
}
