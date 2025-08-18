// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SettingSsr represents a SettingSsr struct.
type SettingSsr struct {
	// List of Conductor IP Addresses or Hosts to be used by the SSR Devices
	ConductorHosts []string `json:"conductor_hosts,omitempty"`
	// Token to be used by the SSR Devices to connect to the Conductor
	ConductorToken *string `json:"conductor_token,omitempty"`
	// Disable stats collection on SSR devices
	DisableStats         *bool                  `json:"disable_stats,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SettingSsr,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SettingSsr) String() string {
	return fmt.Sprintf(
		"SettingSsr[ConductorHosts=%v, ConductorToken=%v, DisableStats=%v, AdditionalProperties=%v]",
		s.ConductorHosts, s.ConductorToken, s.DisableStats, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SettingSsr.
// It customizes the JSON marshaling process for SettingSsr objects.
func (s SettingSsr) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"conductor_hosts", "conductor_token", "disable_stats"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SettingSsr object to a map representation for JSON marshaling.
func (s SettingSsr) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ConductorHosts != nil {
		structMap["conductor_hosts"] = s.ConductorHosts
	}
	if s.ConductorToken != nil {
		structMap["conductor_token"] = s.ConductorToken
	}
	if s.DisableStats != nil {
		structMap["disable_stats"] = s.DisableStats
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SettingSsr.
// It customizes the JSON unmarshaling process for SettingSsr objects.
func (s *SettingSsr) UnmarshalJSON(input []byte) error {
	var temp tempSettingSsr
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "conductor_hosts", "conductor_token", "disable_stats")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ConductorHosts = temp.ConductorHosts
	s.ConductorToken = temp.ConductorToken
	s.DisableStats = temp.DisableStats
	return nil
}

// tempSettingSsr is a temporary struct used for validating the fields of SettingSsr.
type tempSettingSsr struct {
	ConductorHosts []string `json:"conductor_hosts,omitempty"`
	ConductorToken *string  `json:"conductor_token,omitempty"`
	DisableStats   *bool    `json:"disable_stats,omitempty"`
}
