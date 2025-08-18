// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApCentrak represents a ApCentrak struct.
type ApCentrak struct {
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApCentrak,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApCentrak) String() string {
	return fmt.Sprintf(
		"ApCentrak[Enabled=%v, AdditionalProperties=%v]",
		a.Enabled, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApCentrak.
// It customizes the JSON marshaling process for ApCentrak objects.
func (a ApCentrak) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApCentrak object to a map representation for JSON marshaling.
func (a ApCentrak) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApCentrak.
// It customizes the JSON unmarshaling process for ApCentrak objects.
func (a *ApCentrak) UnmarshalJSON(input []byte) error {
	var temp tempApCentrak
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Enabled = temp.Enabled
	return nil
}

// tempApCentrak is a temporary struct used for validating the fields of ApCentrak.
type tempApCentrak struct {
	Enabled *bool `json:"enabled,omitempty"`
}
