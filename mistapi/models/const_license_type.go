// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ConstLicenseType represents a ConstLicenseType struct.
// License type definition returned by the constants API
type ConstLicenseType struct {
	// Human-readable description of the license type
	Description *string `json:"description,omitempty"`
	// Level at which the license is enforced. enum: `org`, `site`.
	EnforcementLevel *ConstLicenseTypeEnforcementLevelEnum `json:"enforcement_level,omitempty"`
	// License type keys this license type entitles
	EntitledLicenses []string `json:"entitled_licenses,omitempty"`
	// License group this license type belongs to
	Group *string `json:"group,omitempty"`
	// License SKU components included by a license type
	Includes []string `json:"includes,omitempty"`
	// Machine-readable license type key
	Key *string `json:"key,omitempty"`
	// Display name of the license type
	Name *string `json:"name,omitempty"`
	// License type identifier (SKU)
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstLicenseType,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstLicenseType) String() string {
	return fmt.Sprintf(
		"ConstLicenseType[Description=%v, EnforcementLevel=%v, EntitledLicenses=%v, Group=%v, Includes=%v, Key=%v, Name=%v, Type=%v, AdditionalProperties=%v]",
		c.Description, c.EnforcementLevel, c.EntitledLicenses, c.Group, c.Includes, c.Key, c.Name, c.Type, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstLicenseType.
// It customizes the JSON marshaling process for ConstLicenseType objects.
func (c ConstLicenseType) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"description", "enforcement_level", "entitled_licenses", "group", "includes", "key", "name", "type"); err != nil {
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
	if c.EnforcementLevel != nil {
		structMap["enforcement_level"] = c.EnforcementLevel
	}
	if c.EntitledLicenses != nil {
		structMap["entitled_licenses"] = c.EntitledLicenses
	}
	if c.Group != nil {
		structMap["group"] = c.Group
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
	if c.Type != nil {
		structMap["type"] = c.Type
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "description", "enforcement_level", "entitled_licenses", "group", "includes", "key", "name", "type")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Description = temp.Description
	c.EnforcementLevel = temp.EnforcementLevel
	c.EntitledLicenses = temp.EntitledLicenses
	c.Group = temp.Group
	c.Includes = temp.Includes
	c.Key = temp.Key
	c.Name = temp.Name
	c.Type = temp.Type
	return nil
}

// tempConstLicenseType is a temporary struct used for validating the fields of ConstLicenseType.
type tempConstLicenseType struct {
	Description      *string                               `json:"description,omitempty"`
	EnforcementLevel *ConstLicenseTypeEnforcementLevelEnum `json:"enforcement_level,omitempty"`
	EntitledLicenses []string                              `json:"entitled_licenses,omitempty"`
	Group            *string                               `json:"group,omitempty"`
	Includes         []string                              `json:"includes,omitempty"`
	Key              *string                               `json:"key,omitempty"`
	Name             *string                               `json:"name,omitempty"`
	Type             *string                               `json:"type,omitempty"`
}
