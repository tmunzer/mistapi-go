// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ServicePolicySkyatpIotDevicePolicy represents a ServicePolicySkyatpIotDevicePolicy struct.
type ServicePolicySkyatpIotDevicePolicy struct {
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicySkyatpIotDevicePolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicySkyatpIotDevicePolicy) String() string {
	return fmt.Sprintf(
		"ServicePolicySkyatpIotDevicePolicy[Enabled=%v, AdditionalProperties=%v]",
		s.Enabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicySkyatpIotDevicePolicy.
// It customizes the JSON marshaling process for ServicePolicySkyatpIotDevicePolicy objects.
func (s ServicePolicySkyatpIotDevicePolicy) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicySkyatpIotDevicePolicy object to a map representation for JSON marshaling.
func (s ServicePolicySkyatpIotDevicePolicy) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicySkyatpIotDevicePolicy.
// It customizes the JSON unmarshaling process for ServicePolicySkyatpIotDevicePolicy objects.
func (s *ServicePolicySkyatpIotDevicePolicy) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicySkyatpIotDevicePolicy
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

// tempServicePolicySkyatpIotDevicePolicy is a temporary struct used for validating the fields of ServicePolicySkyatpIotDevicePolicy.
type tempServicePolicySkyatpIotDevicePolicy struct {
	Enabled *bool `json:"enabled,omitempty"`
}
