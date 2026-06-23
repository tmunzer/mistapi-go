// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisClientTelemetry represents a MarvisClientTelemetry struct.
// Note: some stats are not collected when it's not connected to Mist infrastructure
type MarvisClientTelemetry struct {
	// Whether telemetry collection is enabled for Marvis Client
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClientTelemetry,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClientTelemetry) String() string {
	return fmt.Sprintf(
		"MarvisClientTelemetry[Enabled=%v, AdditionalProperties=%v]",
		m.Enabled, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClientTelemetry.
// It customizes the JSON marshaling process for MarvisClientTelemetry objects.
func (m MarvisClientTelemetry) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClientTelemetry object to a map representation for JSON marshaling.
func (m MarvisClientTelemetry) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Enabled != nil {
		structMap["enabled"] = m.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClientTelemetry.
// It customizes the JSON unmarshaling process for MarvisClientTelemetry objects.
func (m *MarvisClientTelemetry) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClientTelemetry
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

// tempMarvisClientTelemetry is a temporary struct used for validating the fields of MarvisClientTelemetry.
type tempMarvisClientTelemetry struct {
	Enabled *bool `json:"enabled,omitempty"`
}
