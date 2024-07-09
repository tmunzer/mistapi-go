package models

import (
    "encoding/json"
)

// UtilsShowBgpRummary represents a UtilsShowBgpRummary struct.
type UtilsShowBgpRummary struct {
    // only for HA
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowBgpRummary.
// It customizes the JSON marshaling process for UtilsShowBgpRummary objects.
func (u UtilsShowBgpRummary) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowBgpRummary object to a map representation for JSON marshaling.
func (u UtilsShowBgpRummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowBgpRummary.
// It customizes the JSON unmarshaling process for UtilsShowBgpRummary objects.
func (u *UtilsShowBgpRummary) UnmarshalJSON(input []byte) error {
    var temp utilsShowBgpRummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "node")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Node = temp.Node
    return nil
}

// utilsShowBgpRummary is a temporary struct used for validating the fields of UtilsShowBgpRummary.
type utilsShowBgpRummary  struct {
    Node *HaClusterNodeEnum `json:"node,omitempty"`
}
