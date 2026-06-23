// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ServicePolicyAntivirus represents a ServicePolicyAntivirus struct.
// SRX antivirus inspection settings for a service policy
type ServicePolicyAntivirus struct {
	// Organization-level antivirus profile ID; takes precedence over inline `profile` settings
	AvprofileId *uuid.UUID `json:"avprofile_id,omitempty"`
	// Whether antivirus inspection is enabled for the service policy
	Enabled *bool `json:"enabled,omitempty"`
	// Antivirus profile name to apply, such as `default`, `noftp`, `httponly`, or an AV profile key
	Profile              *string                `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicyAntivirus,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicyAntivirus) String() string {
	return fmt.Sprintf(
		"ServicePolicyAntivirus[AvprofileId=%v, Enabled=%v, Profile=%v, AdditionalProperties=%v]",
		s.AvprofileId, s.Enabled, s.Profile, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicyAntivirus.
// It customizes the JSON marshaling process for ServicePolicyAntivirus objects.
func (s ServicePolicyAntivirus) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"avprofile_id", "enabled", "profile"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicyAntivirus object to a map representation for JSON marshaling.
func (s ServicePolicyAntivirus) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AvprofileId != nil {
		structMap["avprofile_id"] = s.AvprofileId
	}
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.Profile != nil {
		structMap["profile"] = s.Profile
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicyAntivirus.
// It customizes the JSON unmarshaling process for ServicePolicyAntivirus objects.
func (s *ServicePolicyAntivirus) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicyAntivirus
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "avprofile_id", "enabled", "profile")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AvprofileId = temp.AvprofileId
	s.Enabled = temp.Enabled
	s.Profile = temp.Profile
	return nil
}

// tempServicePolicyAntivirus is a temporary struct used for validating the fields of ServicePolicyAntivirus.
type tempServicePolicyAntivirus struct {
	AvprofileId *uuid.UUID `json:"avprofile_id,omitempty"`
	Enabled     *bool      `json:"enabled,omitempty"`
	Profile     *string    `json:"profile,omitempty"`
}
