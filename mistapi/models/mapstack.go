// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// Mapstack represents a Mapstack struct.
// Map Stack filter or creation payload
type Mapstack struct {
	// The name of the map stack
	Name                 *string                `json:"name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Mapstack,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Mapstack) String() string {
	return fmt.Sprintf(
		"Mapstack[Name=%v, AdditionalProperties=%v]",
		m.Name, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Mapstack.
// It customizes the JSON marshaling process for Mapstack objects.
func (m Mapstack) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the Mapstack object to a map representation for JSON marshaling.
func (m Mapstack) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Mapstack.
// It customizes the JSON unmarshaling process for Mapstack objects.
func (m *Mapstack) UnmarshalJSON(input []byte) error {
	var temp tempMapstack
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Name = temp.Name
	return nil
}

// tempMapstack is a temporary struct used for validating the fields of Mapstack.
type tempMapstack struct {
	Name *string `json:"name,omitempty"`
}
