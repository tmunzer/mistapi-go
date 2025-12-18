// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwRoutingPolicyTermAction represents a SwRoutingPolicyTermAction struct.
// When used as import policy
type SwRoutingPolicyTermAction struct {
	Accept *bool `json:"accept,omitempty"`
	// When used as export policy, optional
	Community []string `json:"community,omitempty"`
	// Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`)
	LocalPreference *RoutingPolicyLocalPreference `json:"local_preference,omitempty"`
	// When used as export policy, optional. By default, the local AS will be prepended, to change it. Can be a Variable (e.g. `{{as_path}}`)
	PrependAsPath        []string               `json:"prepend_as_path,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwRoutingPolicyTermAction,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwRoutingPolicyTermAction) String() string {
	return fmt.Sprintf(
		"SwRoutingPolicyTermAction[Accept=%v, Community=%v, LocalPreference=%v, PrependAsPath=%v, AdditionalProperties=%v]",
		s.Accept, s.Community, s.LocalPreference, s.PrependAsPath, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwRoutingPolicyTermAction.
// It customizes the JSON marshaling process for SwRoutingPolicyTermAction objects.
func (s SwRoutingPolicyTermAction) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"accept", "community", "local_preference", "prepend_as_path"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwRoutingPolicyTermAction object to a map representation for JSON marshaling.
func (s SwRoutingPolicyTermAction) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Accept != nil {
		structMap["accept"] = s.Accept
	}
	if s.Community != nil {
		structMap["community"] = s.Community
	}
	if s.LocalPreference != nil {
		structMap["local_preference"] = s.LocalPreference.toMap()
	}
	if s.PrependAsPath != nil {
		structMap["prepend_as_path"] = s.PrependAsPath
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwRoutingPolicyTermAction.
// It customizes the JSON unmarshaling process for SwRoutingPolicyTermAction objects.
func (s *SwRoutingPolicyTermAction) UnmarshalJSON(input []byte) error {
	var temp tempSwRoutingPolicyTermAction
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accept", "community", "local_preference", "prepend_as_path")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Accept = temp.Accept
	s.Community = temp.Community
	s.LocalPreference = temp.LocalPreference
	s.PrependAsPath = temp.PrependAsPath
	return nil
}

// tempSwRoutingPolicyTermAction is a temporary struct used for validating the fields of SwRoutingPolicyTermAction.
type tempSwRoutingPolicyTermAction struct {
	Accept          *bool                         `json:"accept,omitempty"`
	Community       []string                      `json:"community,omitempty"`
	LocalPreference *RoutingPolicyLocalPreference `json:"local_preference,omitempty"`
	PrependAsPath   []string                      `json:"prepend_as_path,omitempty"`
}
