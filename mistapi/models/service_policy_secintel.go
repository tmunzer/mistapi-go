// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ServicePolicySecintel represents a ServicePolicySecintel struct.
// For SRX Only
type ServicePolicySecintel struct {
    Enabled              *bool                             `json:"enabled,omitempty"`
    // enum: `default`, `standard`, `strict`
    Profile              *ServicePolicySecintelProfileEnum `json:"profile,omitempty"`
    // org-level secintel Profile can be used, this takes precedence over 'profile'
    SecintelprofileId    *string                           `json:"secintelprofile_id,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicySecintel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicySecintel) String() string {
    return fmt.Sprintf(
    	"ServicePolicySecintel[Enabled=%v, Profile=%v, SecintelprofileId=%v, AdditionalProperties=%v]",
    	s.Enabled, s.Profile, s.SecintelprofileId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicySecintel.
// It customizes the JSON marshaling process for ServicePolicySecintel objects.
func (s ServicePolicySecintel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled", "profile", "secintelprofile_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicySecintel object to a map representation for JSON marshaling.
func (s ServicePolicySecintel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    if s.Profile != nil {
        structMap["profile"] = s.Profile
    }
    if s.SecintelprofileId != nil {
        structMap["secintelprofile_id"] = s.SecintelprofileId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicySecintel.
// It customizes the JSON unmarshaling process for ServicePolicySecintel objects.
func (s *ServicePolicySecintel) UnmarshalJSON(input []byte) error {
    var temp tempServicePolicySecintel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled", "profile", "secintelprofile_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    s.Profile = temp.Profile
    s.SecintelprofileId = temp.SecintelprofileId
    return nil
}

// tempServicePolicySecintel is a temporary struct used for validating the fields of ServicePolicySecintel.
type tempServicePolicySecintel  struct {
    Enabled           *bool                             `json:"enabled,omitempty"`
    Profile           *ServicePolicySecintelProfileEnum `json:"profile,omitempty"`
    SecintelprofileId *string                           `json:"secintelprofile_id,omitempty"`
}
