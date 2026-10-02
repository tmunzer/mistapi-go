// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SwitchRadius represents a SwitchRadius struct.
// RADIUS settings for switch management authentication. `auth_servers`, `auth_servers_retries`, and `auth_servers_timeout` are required when `enabled`==`true` and `use_different_radius`==`true`.
type SwitchRadius struct {
	// List of RADIUS authentication servers used for switch management authentication
	AuthServers []SwitchMgmtRadiusAuthServer `json:"auth_servers,omitempty"`
	// RADIUS auth session retries. Required when `enabled`==`true` and `use_different_radius`==`true`.
	AuthServersRetries *int `json:"auth_servers_retries,omitempty"`
	// RADIUS auth session timeout, in seconds. Required when `enabled`==`true` and `use_different_radius`==`true`.
	AuthServersTimeout *int `json:"auth_servers_timeout,omitempty"`
	// Whether RADIUS is enabled for switch management authentication
	Enabled *bool `json:"enabled,omitempty"`
	// Source network used for connectivity to the RADIUS servers
	Network *string `json:"network,omitempty"`
	// Whether to use alternate RADIUS settings instead of the default switch `radius_config`
	UseDifferentRadius   *bool                  `json:"use_different_radius,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchRadius,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchRadius) String() string {
	return fmt.Sprintf(
		"SwitchRadius[AuthServers=%v, AuthServersRetries=%v, AuthServersTimeout=%v, Enabled=%v, Network=%v, UseDifferentRadius=%v, AdditionalProperties=%v]",
		s.AuthServers, s.AuthServersRetries, s.AuthServersTimeout, s.Enabled, s.Network, s.UseDifferentRadius, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchRadius.
// It customizes the JSON marshaling process for SwitchRadius objects.
func (s SwitchRadius) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"auth_servers", "auth_servers_retries", "auth_servers_timeout", "enabled", "network", "use_different_radius"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchRadius object to a map representation for JSON marshaling.
func (s SwitchRadius) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AuthServers != nil {
		structMap["auth_servers"] = s.AuthServers
	}
	if s.AuthServersRetries != nil {
		structMap["auth_servers_retries"] = s.AuthServersRetries
	}
	if s.AuthServersTimeout != nil {
		structMap["auth_servers_timeout"] = s.AuthServersTimeout
	}
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.Network != nil {
		structMap["network"] = s.Network
	}
	if s.UseDifferentRadius != nil {
		structMap["use_different_radius"] = s.UseDifferentRadius
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchRadius.
// It customizes the JSON unmarshaling process for SwitchRadius objects.
func (s *SwitchRadius) UnmarshalJSON(input []byte) error {
	var temp tempSwitchRadius
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_servers", "auth_servers_retries", "auth_servers_timeout", "enabled", "network", "use_different_radius")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AuthServers = temp.AuthServers
	s.AuthServersRetries = temp.AuthServersRetries
	s.AuthServersTimeout = temp.AuthServersTimeout
	s.Enabled = temp.Enabled
	s.Network = temp.Network
	s.UseDifferentRadius = temp.UseDifferentRadius
	return nil
}

// tempSwitchRadius is a temporary struct used for validating the fields of SwitchRadius.
type tempSwitchRadius struct {
	AuthServers        []SwitchMgmtRadiusAuthServer `json:"auth_servers,omitempty"`
	AuthServersRetries *int                         `json:"auth_servers_retries,omitempty"`
	AuthServersTimeout *int                         `json:"auth_servers_timeout,omitempty"`
	Enabled            *bool                        `json:"enabled,omitempty"`
	Network            *string                      `json:"network,omitempty"`
	UseDifferentRadius *bool                        `json:"use_different_radius,omitempty"`
}
