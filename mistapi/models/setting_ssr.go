// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SettingSsr represents a SettingSsr struct.
// SSR management settings for device onboarding and connectivity
type SettingSsr struct {
	// Automatic firmware upgrade settings applied when an SSR device is first onboarded
	AutoUpgrade *SettingSsrAutoUpgrade `json:"auto_upgrade,omitempty"`
	// Conductor IP addresses or hostnames used by SSR devices
	ConductorHosts []string `json:"conductor_hosts,omitempty"`
	// Registration token used by SSR devices to connect to the conductor
	ConductorToken *string `json:"conductor_token,omitempty"`
	// Whether stats collection is disabled on SSR devices
	DisableStats *bool `json:"disable_stats,omitempty"`
	// SSR proxy configuration to talk to Mist
	Proxy                *SsrProxy              `json:"proxy,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SettingSsr,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SettingSsr) String() string {
	return fmt.Sprintf(
		"SettingSsr[AutoUpgrade=%v, ConductorHosts=%v, ConductorToken=%v, DisableStats=%v, Proxy=%v, AdditionalProperties=%v]",
		s.AutoUpgrade, s.ConductorHosts, s.ConductorToken, s.DisableStats, s.Proxy, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SettingSsr.
// It customizes the JSON marshaling process for SettingSsr objects.
func (s SettingSsr) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"auto_upgrade", "conductor_hosts", "conductor_token", "disable_stats", "proxy"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SettingSsr object to a map representation for JSON marshaling.
func (s SettingSsr) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AutoUpgrade != nil {
		structMap["auto_upgrade"] = s.AutoUpgrade.toMap()
	}
	if s.ConductorHosts != nil {
		structMap["conductor_hosts"] = s.ConductorHosts
	}
	if s.ConductorToken != nil {
		structMap["conductor_token"] = s.ConductorToken
	}
	if s.DisableStats != nil {
		structMap["disable_stats"] = s.DisableStats
	}
	if s.Proxy != nil {
		structMap["proxy"] = s.Proxy.toMap()
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_upgrade", "conductor_hosts", "conductor_token", "disable_stats", "proxy")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AutoUpgrade = temp.AutoUpgrade
	s.ConductorHosts = temp.ConductorHosts
	s.ConductorToken = temp.ConductorToken
	s.DisableStats = temp.DisableStats
	s.Proxy = temp.Proxy
	return nil
}

// tempSettingSsr is a temporary struct used for validating the fields of SettingSsr.
type tempSettingSsr struct {
	AutoUpgrade    *SettingSsrAutoUpgrade `json:"auto_upgrade,omitempty"`
	ConductorHosts []string               `json:"conductor_hosts,omitempty"`
	ConductorToken *string                `json:"conductor_token,omitempty"`
	DisableStats   *bool                  `json:"disable_stats,omitempty"`
	Proxy          *SsrProxy              `json:"proxy,omitempty"`
}
