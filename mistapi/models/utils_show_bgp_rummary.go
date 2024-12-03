package models

import (
    "encoding/json"
)

// UtilsShowBgpRummary represents a UtilsShowBgpRummary struct.
type UtilsShowBgpRummary struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowBgpRummary.
// It customizes the JSON marshaling process for UtilsShowBgpRummary objects.
func (u UtilsShowBgpRummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "node"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowBgpRummary object to a map representation for JSON marshaling.
func (u UtilsShowBgpRummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowBgpRummary.
// It customizes the JSON unmarshaling process for UtilsShowBgpRummary objects.
func (u *UtilsShowBgpRummary) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowBgpRummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Node = temp.Node
    return nil
}

// tempUtilsShowBgpRummary is a temporary struct used for validating the fields of UtilsShowBgpRummary.
type tempUtilsShowBgpRummary  struct {
    Node *HaClusterNodeEnum `json:"node,omitempty"`
}
