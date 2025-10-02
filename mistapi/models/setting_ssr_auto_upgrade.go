// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SettingSsrAutoUpgrade represents a SettingSsrAutoUpgrade struct.
// auto_upgrade device first time it is onboarded
type SettingSsrAutoUpgrade struct {
	// upgrade channel to follow. enum: `alpha`, `beta`, `stable`
	Channel *SsrUpgradeChannelEnum `json:"channel,omitempty"`
	// Property key is the SSR model (e.g. "SSR130").
	CustomVersions       map[string]string      `json:"custom_versions,omitempty"`
	Enabled              *bool                  `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SettingSsrAutoUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SettingSsrAutoUpgrade) String() string {
	return fmt.Sprintf(
		"SettingSsrAutoUpgrade[Channel=%v, CustomVersions=%v, Enabled=%v, AdditionalProperties=%v]",
		s.Channel, s.CustomVersions, s.Enabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SettingSsrAutoUpgrade.
// It customizes the JSON marshaling process for SettingSsrAutoUpgrade objects.
func (s SettingSsrAutoUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"channel", "custom_versions", "enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SettingSsrAutoUpgrade object to a map representation for JSON marshaling.
func (s SettingSsrAutoUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Channel != nil {
		structMap["channel"] = s.Channel
	}
	if s.CustomVersions != nil {
		structMap["custom_versions"] = s.CustomVersions
	}
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SettingSsrAutoUpgrade.
// It customizes the JSON unmarshaling process for SettingSsrAutoUpgrade objects.
func (s *SettingSsrAutoUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempSettingSsrAutoUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "custom_versions", "enabled")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Channel = temp.Channel
	s.CustomVersions = temp.CustomVersions
	s.Enabled = temp.Enabled
	return nil
}

// tempSettingSsrAutoUpgrade is a temporary struct used for validating the fields of SettingSsrAutoUpgrade.
type tempSettingSsrAutoUpgrade struct {
	Channel        *SsrUpgradeChannelEnum `json:"channel,omitempty"`
	CustomVersions map[string]string      `json:"custom_versions,omitempty"`
	Enabled        *bool                  `json:"enabled,omitempty"`
}
