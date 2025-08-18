// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApAutoUpgrade represents a StatsApAutoUpgrade struct.
type StatsApAutoUpgrade struct {
	Lastcheck            Optional[int64]        `json:"lastcheck"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApAutoUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApAutoUpgrade) String() string {
	return fmt.Sprintf(
		"StatsApAutoUpgrade[Lastcheck=%v, AdditionalProperties=%v]",
		s.Lastcheck, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApAutoUpgrade.
// It customizes the JSON marshaling process for StatsApAutoUpgrade objects.
func (s StatsApAutoUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"lastcheck"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApAutoUpgrade object to a map representation for JSON marshaling.
func (s StatsApAutoUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Lastcheck.IsValueSet() {
		if s.Lastcheck.Value() != nil {
			structMap["lastcheck"] = s.Lastcheck.Value()
		} else {
			structMap["lastcheck"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApAutoUpgrade.
// It customizes the JSON unmarshaling process for StatsApAutoUpgrade objects.
func (s *StatsApAutoUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempStatsApAutoUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "lastcheck")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Lastcheck = temp.Lastcheck
	return nil
}

// tempStatsApAutoUpgrade is a temporary struct used for validating the fields of StatsApAutoUpgrade.
type tempStatsApAutoUpgrade struct {
	Lastcheck Optional[int64] `json:"lastcheck"`
}
