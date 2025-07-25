// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ConstState represents a ConstState struct.
type ConstState struct {
    IsoCode              *string                `json:"iso_code,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstState,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstState) String() string {
    return fmt.Sprintf(
    	"ConstState[IsoCode=%v, Name=%v, AdditionalProperties=%v]",
    	c.IsoCode, c.Name, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstState.
// It customizes the JSON marshaling process for ConstState objects.
func (c ConstState) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "iso_code", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstState object to a map representation for JSON marshaling.
func (c ConstState) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "iso_code", "name")
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
