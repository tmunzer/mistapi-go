// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SynthetictestConfig represents a SynthetictestConfig struct.
type SynthetictestConfig struct {
	// enum: `auto`, `high`, `low`
	Aggressiveness *SynthetictestConfigAggressivenessEnum `json:"aggressiveness,omitempty"`
	// Custom probes to be used for synthetic tests
	CustomProbes map[string]SynthetictestConfigCustomProbe `json:"custom_probes,omitempty"`
	Disabled     *bool                                     `json:"disabled,omitempty"`
	// List of networks to be used for synthetic tests
	LanNetworks          []SynthetictestConfigLanNetwork  `json:"lan_networks,omitempty"`
	Vlans                []SynthetictestConfigVlan        `json:"vlans,omitempty"` // Deprecated
	WanSpeedtest         *SynthetictestConfigWanSpeedtest `json:"wan_speedtest,omitempty"`
	AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestConfig) String() string {
	return fmt.Sprintf(
		"SynthetictestConfig[Aggressiveness=%v, CustomProbes=%v, Disabled=%v, LanNetworks=%v, Vlans=%v, WanSpeedtest=%v, AdditionalProperties=%v]",
		s.Aggressiveness, s.CustomProbes, s.Disabled, s.LanNetworks, s.Vlans, s.WanSpeedtest, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfig.
// It customizes the JSON marshaling process for SynthetictestConfig objects.
func (s SynthetictestConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"aggressiveness", "custom_probes", "disabled", "lan_networks", "vlans", "wan_speedtest"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestConfig object to a map representation for JSON marshaling.
func (s SynthetictestConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Aggressiveness != nil {
		structMap["aggressiveness"] = s.Aggressiveness
	}
	if s.CustomProbes != nil {
		structMap["custom_probes"] = s.CustomProbes
	}
	if s.Disabled != nil {
		structMap["disabled"] = s.Disabled
	}
	if s.LanNetworks != nil {
		structMap["lan_networks"] = s.LanNetworks
	}
	if s.Vlans != nil {
		structMap["vlans"] = s.Vlans
	}
	if s.WanSpeedtest != nil {
		structMap["wan_speedtest"] = s.WanSpeedtest.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestConfig.
// It customizes the JSON unmarshaling process for SynthetictestConfig objects.
func (s *SynthetictestConfig) UnmarshalJSON(input []byte) error {
	var temp tempSynthetictestConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aggressiveness", "custom_probes", "disabled", "lan_networks", "vlans", "wan_speedtest")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Aggressiveness = temp.Aggressiveness
	s.CustomProbes = temp.CustomProbes
	s.Disabled = temp.Disabled
	s.LanNetworks = temp.LanNetworks
	s.Vlans = temp.Vlans
	s.WanSpeedtest = temp.WanSpeedtest
	return nil
}

// tempSynthetictestConfig is a temporary struct used for validating the fields of SynthetictestConfig.
type tempSynthetictestConfig struct {
	Aggressiveness *SynthetictestConfigAggressivenessEnum    `json:"aggressiveness,omitempty"`
	CustomProbes   map[string]SynthetictestConfigCustomProbe `json:"custom_probes,omitempty"`
	Disabled       *bool                                     `json:"disabled,omitempty"`
	LanNetworks    []SynthetictestConfigLanNetwork           `json:"lan_networks,omitempty"`
	Vlans          []SynthetictestConfigVlan                 `json:"vlans,omitempty"`
	WanSpeedtest   *SynthetictestConfigWanSpeedtest          `json:"wan_speedtest,omitempty"`
}
