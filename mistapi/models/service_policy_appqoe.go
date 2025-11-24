// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ServicePolicyAppqoe represents a ServicePolicyAppqoe struct.
// SRX only
type ServicePolicyAppqoe struct {
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicyAppqoe,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicyAppqoe) String() string {
	return fmt.Sprintf(
		"ServicePolicyAppqoe[Enabled=%v, AdditionalProperties=%v]",
		s.Enabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicyAppqoe.
// It customizes the JSON marshaling process for ServicePolicyAppqoe objects.
func (s ServicePolicyAppqoe) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicyAppqoe object to a map representation for JSON marshaling.
func (s ServicePolicyAppqoe) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicyAppqoe.
// It customizes the JSON unmarshaling process for ServicePolicyAppqoe objects.
func (s *ServicePolicyAppqoe) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicyAppqoe
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Enabled = temp.Enabled
	return nil
}

// tempServicePolicyAppqoe is a temporary struct used for validating the fields of ServicePolicyAppqoe.
type tempServicePolicyAppqoe struct {
	Enabled *bool `json:"enabled,omitempty"`
}
