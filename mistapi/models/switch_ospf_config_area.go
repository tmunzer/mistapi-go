// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchOspfConfigArea represents a SwitchOspfConfigArea struct.
type SwitchOspfConfigArea struct {
	// Disable OSPF summary routes for this area
	NoSummary            *bool                  `json:"no_summary,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchOspfConfigArea,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchOspfConfigArea) String() string {
	return fmt.Sprintf(
		"SwitchOspfConfigArea[NoSummary=%v, AdditionalProperties=%v]",
		s.NoSummary, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchOspfConfigArea.
// It customizes the JSON marshaling process for SwitchOspfConfigArea objects.
func (s SwitchOspfConfigArea) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"no_summary"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchOspfConfigArea object to a map representation for JSON marshaling.
func (s SwitchOspfConfigArea) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.NoSummary != nil {
		structMap["no_summary"] = s.NoSummary
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchOspfConfigArea.
// It customizes the JSON unmarshaling process for SwitchOspfConfigArea objects.
func (s *SwitchOspfConfigArea) UnmarshalJSON(input []byte) error {
	var temp tempSwitchOspfConfigArea
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "no_summary")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.NoSummary = temp.NoSummary
	return nil
}

// tempSwitchOspfConfigArea is a temporary struct used for validating the fields of SwitchOspfConfigArea.
type tempSwitchOspfConfigArea struct {
	NoSummary *bool `json:"no_summary,omitempty"`
}
