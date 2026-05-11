// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisSelfDrivingDomain represents a MarvisSelfDrivingDomain struct.
type MarvisSelfDrivingDomain struct {
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisSelfDrivingDomain,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisSelfDrivingDomain) String() string {
	return fmt.Sprintf(
		"MarvisSelfDrivingDomain[Enabled=%v, AdditionalProperties=%v]",
		m.Enabled, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisSelfDrivingDomain.
// It customizes the JSON marshaling process for MarvisSelfDrivingDomain objects.
func (m MarvisSelfDrivingDomain) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisSelfDrivingDomain object to a map representation for JSON marshaling.
func (m MarvisSelfDrivingDomain) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Enabled != nil {
		structMap["enabled"] = m.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisSelfDrivingDomain.
// It customizes the JSON unmarshaling process for MarvisSelfDrivingDomain objects.
func (m *MarvisSelfDrivingDomain) UnmarshalJSON(input []byte) error {
	var temp tempMarvisSelfDrivingDomain
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

// tempMarvisSelfDrivingDomain is a temporary struct used for validating the fields of MarvisSelfDrivingDomain.
type tempMarvisSelfDrivingDomain struct {
	Enabled *bool `json:"enabled,omitempty"`
}
