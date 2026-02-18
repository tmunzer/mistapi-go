// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ServicePolicySkyatpHttpInspection represents a ServicePolicySkyatpHttpInspection struct.
type ServicePolicySkyatpHttpInspection struct {
	Enabled *bool `json:"enabled,omitempty"`
	// enum: `standard`, `strict`
	Profile              *Profile2Enum          `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicySkyatpHttpInspection,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicySkyatpHttpInspection) String() string {
	return fmt.Sprintf(
		"ServicePolicySkyatpHttpInspection[Enabled=%v, Profile=%v, AdditionalProperties=%v]",
		s.Enabled, s.Profile, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicySkyatpHttpInspection.
// It customizes the JSON marshaling process for ServicePolicySkyatpHttpInspection objects.
func (s ServicePolicySkyatpHttpInspection) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"enabled", "profile"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicySkyatpHttpInspection object to a map representation for JSON marshaling.
func (s ServicePolicySkyatpHttpInspection) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.Profile != nil {
		structMap["profile"] = s.Profile
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicySkyatpHttpInspection.
// It customizes the JSON unmarshaling process for ServicePolicySkyatpHttpInspection objects.
func (s *ServicePolicySkyatpHttpInspection) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicySkyatpHttpInspection
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "profile")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Enabled = temp.Enabled
	s.Profile = temp.Profile
	return nil
}

// tempServicePolicySkyatpHttpInspection is a temporary struct used for validating the fields of ServicePolicySkyatpHttpInspection.
type tempServicePolicySkyatpHttpInspection struct {
	Enabled *bool         `json:"enabled,omitempty"`
	Profile *Profile2Enum `json:"profile,omitempty"`
}
