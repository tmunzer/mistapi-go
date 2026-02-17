// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ServicePolicySkyatpDnsTunnelDetection represents a ServicePolicySkyatpDnsTunnelDetection struct.
type ServicePolicySkyatpDnsTunnelDetection struct {
	Enabled *bool `json:"enabled,omitempty"`
	// enum: `default`, `standard`, `strict`
	Profile              *SecintelProfileProfileActionEnum `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicySkyatpDnsTunnelDetection,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicySkyatpDnsTunnelDetection) String() string {
	return fmt.Sprintf(
		"ServicePolicySkyatpDnsTunnelDetection[Enabled=%v, Profile=%v, AdditionalProperties=%v]",
		s.Enabled, s.Profile, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicySkyatpDnsTunnelDetection.
// It customizes the JSON marshaling process for ServicePolicySkyatpDnsTunnelDetection objects.
func (s ServicePolicySkyatpDnsTunnelDetection) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"enabled", "profile"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicySkyatpDnsTunnelDetection object to a map representation for JSON marshaling.
func (s ServicePolicySkyatpDnsTunnelDetection) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicySkyatpDnsTunnelDetection.
// It customizes the JSON unmarshaling process for ServicePolicySkyatpDnsTunnelDetection objects.
func (s *ServicePolicySkyatpDnsTunnelDetection) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicySkyatpDnsTunnelDetection
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

// tempServicePolicySkyatpDnsTunnelDetection is a temporary struct used for validating the fields of ServicePolicySkyatpDnsTunnelDetection.
type tempServicePolicySkyatpDnsTunnelDetection struct {
	Enabled *bool                             `json:"enabled,omitempty"`
	Profile *SecintelProfileProfileActionEnum `json:"profile,omitempty"`
}
