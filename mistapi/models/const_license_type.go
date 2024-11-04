package models

import (
    "encoding/json"
)

// ConstLicenseType represents a ConstLicenseType struct.
type ConstLicenseType struct {
    Description          *string        `json:"description,omitempty"`
    Includes             []string       `json:"includes,omitempty"`
    Key                  *string        `json:"key,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstLicenseType.
// It customizes the JSON marshaling process for ConstLicenseType objects.
func (c ConstLicenseType) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstLicenseType object to a map representation for JSON marshaling.
func (c ConstLicenseType) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "description", "includes", "key", "name")
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
