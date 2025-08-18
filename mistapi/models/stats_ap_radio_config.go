// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApRadioConfig represents a StatsApRadioConfig struct.
type StatsApRadioConfig struct {
	Band24               *StatsApRadioConfigBand `json:"band_24,omitempty"`
	Band24Usage          Optional[string]        `json:"band_24_usage"`
	Band5                *StatsApRadioConfigBand `json:"band_5,omitempty"`
	Band6                *StatsApRadioConfigBand `json:"band_6,omitempty"`
	ScanningEnabled      *bool                   `json:"scanning_enabled,omitempty"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApRadioConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApRadioConfig) String() string {
	return fmt.Sprintf(
		"StatsApRadioConfig[Band24=%v, Band24Usage=%v, Band5=%v, Band6=%v, ScanningEnabled=%v, AdditionalProperties=%v]",
		s.Band24, s.Band24Usage, s.Band5, s.Band6, s.ScanningEnabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApRadioConfig.
// It customizes the JSON marshaling process for StatsApRadioConfig objects.
func (s StatsApRadioConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"band_24", "band_24_usage", "band_5", "band_6", "scanning_enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApRadioConfig object to a map representation for JSON marshaling.
func (s StatsApRadioConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Band24 != nil {
		structMap["band_24"] = s.Band24.toMap()
	}
	if s.Band24Usage.IsValueSet() {
		if s.Band24Usage.Value() != nil {
			structMap["band_24_usage"] = s.Band24Usage.Value()
		} else {
			structMap["band_24_usage"] = nil
		}
	}
	if s.Band5 != nil {
		structMap["band_5"] = s.Band5.toMap()
	}
	if s.Band6 != nil {
		structMap["band_6"] = s.Band6.toMap()
	}
	if s.ScanningEnabled != nil {
		structMap["scanning_enabled"] = s.ScanningEnabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApRadioConfig.
// It customizes the JSON unmarshaling process for StatsApRadioConfig objects.
func (s *StatsApRadioConfig) UnmarshalJSON(input []byte) error {
	var temp tempStatsApRadioConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band_24", "band_24_usage", "band_5", "band_6", "scanning_enabled")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Band24 = temp.Band24
	s.Band24Usage = temp.Band24Usage
	s.Band5 = temp.Band5
	s.Band6 = temp.Band6
	s.ScanningEnabled = temp.ScanningEnabled
	return nil
}

// tempStatsApRadioConfig is a temporary struct used for validating the fields of StatsApRadioConfig.
type tempStatsApRadioConfig struct {
	Band24          *StatsApRadioConfigBand `json:"band_24,omitempty"`
	Band24Usage     Optional[string]        `json:"band_24_usage"`
	Band5           *StatsApRadioConfigBand `json:"band_5,omitempty"`
	Band6           *StatsApRadioConfigBand `json:"band_6,omitempty"`
	ScanningEnabled *bool                   `json:"scanning_enabled,omitempty"`
}
