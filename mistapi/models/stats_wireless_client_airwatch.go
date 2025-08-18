// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// StatsWirelessClientAirwatch represents a StatsWirelessClientAirwatch struct.
// Information if airwatch enabled
type StatsWirelessClientAirwatch struct {
	Authorized           bool                   `json:"authorized"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWirelessClientAirwatch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWirelessClientAirwatch) String() string {
	return fmt.Sprintf(
		"StatsWirelessClientAirwatch[Authorized=%v, AdditionalProperties=%v]",
		s.Authorized, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClientAirwatch.
// It customizes the JSON marshaling process for StatsWirelessClientAirwatch objects.
func (s StatsWirelessClientAirwatch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"authorized"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClientAirwatch object to a map representation for JSON marshaling.
func (s StatsWirelessClientAirwatch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["authorized"] = s.Authorized
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWirelessClientAirwatch.
// It customizes the JSON unmarshaling process for StatsWirelessClientAirwatch objects.
func (s *StatsWirelessClientAirwatch) UnmarshalJSON(input []byte) error {
	var temp tempStatsWirelessClientAirwatch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "authorized")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Authorized = *temp.Authorized
	return nil
}

// tempStatsWirelessClientAirwatch is a temporary struct used for validating the fields of StatsWirelessClientAirwatch.
type tempStatsWirelessClientAirwatch struct {
	Authorized *bool `json:"authorized"`
}

func (s *tempStatsWirelessClientAirwatch) validate() error {
	var errs []string
	if s.Authorized == nil {
		errs = append(errs, "required field `authorized` is missing for type `stats_wireless_client_airwatch`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
