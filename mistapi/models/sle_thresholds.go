// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SleThresholds represents a SleThresholds struct.
type SleThresholds struct {
	// Capacity, in %
	Capacity *int `json:"capacity,omitempty"`
	// Coverage, in dBm
	Coverage *int `json:"coverage,omitempty"`
	// Throughput, in Mbps
	Throughput *int `json:"throughput,omitempty"`
	// Time to connect, in seconds
	TimeToConnect        *int                   `json:"time-to-connect,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleThresholds,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleThresholds) String() string {
	return fmt.Sprintf(
		"SleThresholds[Capacity=%v, Coverage=%v, Throughput=%v, TimeToConnect=%v, AdditionalProperties=%v]",
		s.Capacity, s.Coverage, s.Throughput, s.TimeToConnect, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleThresholds.
// It customizes the JSON marshaling process for SleThresholds objects.
func (s SleThresholds) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"capacity", "coverage", "throughput", "time-to-connect"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleThresholds object to a map representation for JSON marshaling.
func (s SleThresholds) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Capacity != nil {
		structMap["capacity"] = s.Capacity
	}
	if s.Coverage != nil {
		structMap["coverage"] = s.Coverage
	}
	if s.Throughput != nil {
		structMap["throughput"] = s.Throughput
	}
	if s.TimeToConnect != nil {
		structMap["time-to-connect"] = s.TimeToConnect
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleThresholds.
// It customizes the JSON unmarshaling process for SleThresholds objects.
func (s *SleThresholds) UnmarshalJSON(input []byte) error {
	var temp tempSleThresholds
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "capacity", "coverage", "throughput", "time-to-connect")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Capacity = temp.Capacity
	s.Coverage = temp.Coverage
	s.Throughput = temp.Throughput
	s.TimeToConnect = temp.TimeToConnect
	return nil
}

// tempSleThresholds is a temporary struct used for validating the fields of SleThresholds.
type tempSleThresholds struct {
	Capacity      *int `json:"capacity,omitempty"`
	Coverage      *int `json:"coverage,omitempty"`
	Throughput    *int `json:"throughput,omitempty"`
	TimeToConnect *int `json:"time-to-connect,omitempty"`
}
