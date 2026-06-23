// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisClientLocation represents a MarvisClientLocation struct.
// Location collection settings for Marvis Client
type MarvisClientLocation struct {
	// Whether location collection is enabled for Marvis Client
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClientLocation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClientLocation) String() string {
	return fmt.Sprintf(
		"MarvisClientLocation[Enabled=%v, AdditionalProperties=%v]",
		m.Enabled, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClientLocation.
// It customizes the JSON marshaling process for MarvisClientLocation objects.
func (m MarvisClientLocation) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClientLocation object to a map representation for JSON marshaling.
func (m MarvisClientLocation) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Enabled != nil {
		structMap["enabled"] = m.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClientLocation.
// It customizes the JSON unmarshaling process for MarvisClientLocation objects.
func (m *MarvisClientLocation) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClientLocation
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Enabled = temp.Enabled
	return nil
}

// tempMarvisClientLocation is a temporary struct used for validating the fields of MarvisClientLocation.
type tempMarvisClientLocation struct {
	Enabled *bool `json:"enabled,omitempty"`
}
