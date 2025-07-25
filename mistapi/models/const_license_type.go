// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ConstLicenseType represents a ConstLicenseType struct.
type ConstLicenseType struct {
    Description          *string                `json:"description,omitempty"`
    Includes             []string               `json:"includes,omitempty"`
    Key                  *string                `json:"key,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstLicenseType,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstLicenseType) String() string {
    return fmt.Sprintf(
    	"ConstLicenseType[Description=%v, Includes=%v, Key=%v, Name=%v, AdditionalProperties=%v]",
    	c.Description, c.Includes, c.Key, c.Name, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstLicenseType.
// It customizes the JSON marshaling process for ConstLicenseType objects.
func (c ConstLicenseType) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "description", "includes", "key", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstLicenseType object to a map representation for JSON marshaling.
func (c ConstLicenseType) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.Includes != nil {
        structMap["includes"] = c.Includes
    }
    if c.Key != nil {
        structMap["key"] = c.Key
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstLicenseType.
// It customizes the JSON unmarshaling process for ConstLicenseType objects.
func (c *ConstLicenseType) UnmarshalJSON(input []byte) error {
    var temp tempConstLicenseType
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "description", "includes", "key", "name")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Description = temp.Description
    c.Includes = temp.Includes
    c.Key = temp.Key
    c.Name = temp.Name
    return nil
}

// tempConstLicenseType is a temporary struct used for validating the fields of ConstLicenseType.
type tempConstLicenseType  struct {
    Description *string  `json:"description,omitempty"`
    Includes    []string `json:"includes,omitempty"`
    Key         *string  `json:"key,omitempty"`
    Name        *string  `json:"name,omitempty"`
}
