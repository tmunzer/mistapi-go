package models

import (
    "encoding/json"
)

// ConstState represents a ConstState struct.
type ConstState struct {
    IsoCode              *string        `json:"iso_code,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstState.
// It customizes the JSON marshaling process for ConstState objects.
func (c ConstState) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstState object to a map representation for JSON marshaling.
func (c ConstState) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.IsoCode != nil {
        structMap["iso_code"] = c.IsoCode
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstState.
// It customizes the JSON unmarshaling process for ConstState objects.
func (c *ConstState) UnmarshalJSON(input []byte) error {
    var temp tempConstState
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "iso_code", "name")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.IsoCode = temp.IsoCode
    c.Name = temp.Name
    return nil
}

// tempConstState is a temporary struct used for validating the fields of ConstState.
type tempConstState  struct {
    IsoCode *string `json:"iso_code,omitempty"`
    Name    *string `json:"name,omitempty"`
}
