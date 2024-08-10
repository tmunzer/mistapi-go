package models

import (
    "encoding/json"
)

// VrrpConfigGroup represents a VrrpConfigGroup struct.
type VrrpConfigGroup struct {
    Priority             *int           `json:"priority,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VrrpConfigGroup.
// It customizes the JSON marshaling process for VrrpConfigGroup objects.
func (v VrrpConfigGroup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VrrpConfigGroup object to a map representation for JSON marshaling.
func (v VrrpConfigGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Priority != nil {
        structMap["priority"] = v.Priority
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrrpConfigGroup.
// It customizes the JSON unmarshaling process for VrrpConfigGroup objects.
func (v *VrrpConfigGroup) UnmarshalJSON(input []byte) error {
    var temp tempVrrpConfigGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "priority")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Priority = temp.Priority
    return nil
}

// tempVrrpConfigGroup is a temporary struct used for validating the fields of VrrpConfigGroup.
type tempVrrpConfigGroup  struct {
    Priority *int `json:"priority,omitempty"`
}
