// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApRadioStat represents a StatsApRadioStat struct.
type StatsApRadioStat struct {
	// Radio stat
	Band24 *ApRadioStat `json:"band_24,omitempty"`
	// Radio stat
	Band5 *ApRadioStat `json:"band_5,omitempty"`
	// Radio stat
	Band6                *ApRadioStat           `json:"band_6,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApRadioStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApRadioStat) String() string {
	return fmt.Sprintf(
		"StatsApRadioStat[Band24=%v, Band5=%v, Band6=%v, AdditionalProperties=%v]",
		s.Band24, s.Band5, s.Band6, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApRadioStat.
// It customizes the JSON marshaling process for StatsApRadioStat objects.
func (s StatsApRadioStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"band_24", "band_5", "band_6"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApRadioStat object to a map representation for JSON marshaling.
func (s StatsApRadioStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Band24 != nil {
		structMap["band_24"] = s.Band24.toMap()
	}
	if s.Band5 != nil {
		structMap["band_5"] = s.Band5.toMap()
	}
	if s.Band6 != nil {
		structMap["band_6"] = s.Band6.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApRadioStat.
// It customizes the JSON unmarshaling process for StatsApRadioStat objects.
func (s *StatsApRadioStat) UnmarshalJSON(input []byte) error {
	var temp tempStatsApRadioStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band_24", "band_5", "band_6")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Band24 = temp.Band24
	s.Band5 = temp.Band5
	s.Band6 = temp.Band6
	return nil
}

// tempStatsApRadioStat is a temporary struct used for validating the fields of StatsApRadioStat.
type tempStatsApRadioStat struct {
	Band24 *ApRadioStat `json:"band_24,omitempty"`
	Band5  *ApRadioStat `json:"band_5,omitempty"`
	Band6  *ApRadioStat `json:"band_6,omitempty"`
}
