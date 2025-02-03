package models

import (
    "encoding/json"
    "fmt"
)

// HaClusterConfigNode represents a HaClusterConfigNode struct.
type HaClusterConfigNode struct {
    // Node mac, should be unassigned
    Mac                  *string                `json:"mac,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for HaClusterConfigNode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (h HaClusterConfigNode) String() string {
    return fmt.Sprintf(
    	"HaClusterConfigNode[Mac=%v, AdditionalProperties=%v]",
    	h.Mac, h.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for HaClusterConfigNode.
// It customizes the JSON marshaling process for HaClusterConfigNode objects.
func (h HaClusterConfigNode) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(h.AdditionalProperties,
        "mac"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(h.toMap())
}

// toMap converts the HaClusterConfigNode object to a map representation for JSON marshaling.
func (h HaClusterConfigNode) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, h.AdditionalProperties)
    if h.Mac != nil {
        structMap["mac"] = h.Mac
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for HaClusterConfigNode.
// It customizes the JSON unmarshaling process for HaClusterConfigNode objects.
func (h *HaClusterConfigNode) UnmarshalJSON(input []byte) error {
    var temp tempHaClusterConfigNode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac")
    if err != nil {
    	return err
    }
    h.AdditionalProperties = additionalProperties
    
    h.Mac = temp.Mac
    return nil
}

// tempHaClusterConfigNode is a temporary struct used for validating the fields of HaClusterConfigNode.
type tempHaClusterConfigNode  struct {
    Mac *string `json:"mac,omitempty"`
}
