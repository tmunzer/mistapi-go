// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisClientSyntheticTest represents a MarvisClientSyntheticTest struct.
// Synthetic test settings for Marvis Client
type MarvisClientSyntheticTest struct {
	// Whether synthetic testing is enabled for Marvis Client
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClientSyntheticTest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClientSyntheticTest) String() string {
	return fmt.Sprintf(
		"MarvisClientSyntheticTest[Enabled=%v, AdditionalProperties=%v]",
		m.Enabled, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClientSyntheticTest.
// It customizes the JSON marshaling process for MarvisClientSyntheticTest objects.
func (m MarvisClientSyntheticTest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClientSyntheticTest object to a map representation for JSON marshaling.
func (m MarvisClientSyntheticTest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Enabled != nil {
		structMap["enabled"] = m.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClientSyntheticTest.
// It customizes the JSON unmarshaling process for MarvisClientSyntheticTest objects.
func (m *MarvisClientSyntheticTest) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClientSyntheticTest
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

// tempMarvisClientSyntheticTest is a temporary struct used for validating the fields of MarvisClientSyntheticTest.
type tempMarvisClientSyntheticTest struct {
	Enabled *bool `json:"enabled,omitempty"`
}
