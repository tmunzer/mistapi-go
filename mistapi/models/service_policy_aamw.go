// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ServicePolicyAamw represents a ServicePolicyAamw struct.
// SRX only
type ServicePolicyAamw struct {
	// org-level Advanced Advance Anti Malware Profile (SkyAtp) Profile can be used, this takes precedence over 'profile'
	AamwprofileId *uuid.UUID `json:"aamwprofile_id,omitempty"`
	Enabled       *bool      `json:"enabled,omitempty"`
	// enum: `docsonly`, `executables`, `standard`
	Profile              *ServicePolicyAamwProfileEnum `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicyAamw,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicyAamw) String() string {
	return fmt.Sprintf(
		"ServicePolicyAamw[AamwprofileId=%v, Enabled=%v, Profile=%v, AdditionalProperties=%v]",
		s.AamwprofileId, s.Enabled, s.Profile, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicyAamw.
// It customizes the JSON marshaling process for ServicePolicyAamw objects.
func (s ServicePolicyAamw) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"aamwprofile_id", "enabled", "profile"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicyAamw object to a map representation for JSON marshaling.
func (s ServicePolicyAamw) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AamwprofileId != nil {
		structMap["aamwprofile_id"] = s.AamwprofileId
	}
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.Profile != nil {
		structMap["profile"] = s.Profile
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicyAamw.
// It customizes the JSON unmarshaling process for ServicePolicyAamw objects.
func (s *ServicePolicyAamw) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicyAamw
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aamwprofile_id", "enabled", "profile")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AamwprofileId = temp.AamwprofileId
	s.Enabled = temp.Enabled
	s.Profile = temp.Profile
	return nil
}

// tempServicePolicyAamw is a temporary struct used for validating the fields of ServicePolicyAamw.
type tempServicePolicyAamw struct {
	AamwprofileId *uuid.UUID                    `json:"aamwprofile_id,omitempty"`
	Enabled       *bool                         `json:"enabled,omitempty"`
	Profile       *ServicePolicyAamwProfileEnum `json:"profile,omitempty"`
}
