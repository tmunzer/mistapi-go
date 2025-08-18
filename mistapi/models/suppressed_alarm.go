// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SuppressedAlarm represents a SuppressedAlarm struct.
type SuppressedAlarm struct {
	// If `scope`==`site`. Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration`
	Applies *SuppressedAlarmApplies `json:"applies,omitempty"`
	// Duration, in seconds. Maximum duration is 86400 * 180 (180 days). 0 is to un-suppress alarms
	Duration *float64 `json:"duration,omitempty"`
	// Epoch_time in seconds, Default as now, accepted time range is from now to now + 7 days
	ScheduledTime *int `json:"scheduled_time,omitempty"`
	// level of scope. enum: `org`, `site`
	Scope                *SuppressedAlarmScopeEnum `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for SuppressedAlarm,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SuppressedAlarm) String() string {
	return fmt.Sprintf(
		"SuppressedAlarm[Applies=%v, Duration=%v, ScheduledTime=%v, Scope=%v, AdditionalProperties=%v]",
		s.Applies, s.Duration, s.ScheduledTime, s.Scope, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SuppressedAlarm.
// It customizes the JSON marshaling process for SuppressedAlarm objects.
func (s SuppressedAlarm) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"applies", "duration", "scheduled_time", "scope"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SuppressedAlarm object to a map representation for JSON marshaling.
func (s SuppressedAlarm) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Applies != nil {
		structMap["applies"] = s.Applies.toMap()
	}
	if s.Duration != nil {
		structMap["duration"] = s.Duration
	}
	if s.ScheduledTime != nil {
		structMap["scheduled_time"] = s.ScheduledTime
	}
	if s.Scope != nil {
		structMap["scope"] = s.Scope
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SuppressedAlarm.
// It customizes the JSON unmarshaling process for SuppressedAlarm objects.
func (s *SuppressedAlarm) UnmarshalJSON(input []byte) error {
	var temp tempSuppressedAlarm
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "applies", "duration", "scheduled_time", "scope")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Applies = temp.Applies
	s.Duration = temp.Duration
	s.ScheduledTime = temp.ScheduledTime
	s.Scope = temp.Scope
	return nil
}

// tempSuppressedAlarm is a temporary struct used for validating the fields of SuppressedAlarm.
type tempSuppressedAlarm struct {
	Applies       *SuppressedAlarmApplies   `json:"applies,omitempty"`
	Duration      *float64                  `json:"duration,omitempty"`
	ScheduledTime *int                      `json:"scheduled_time,omitempty"`
	Scope         *SuppressedAlarmScopeEnum `json:"scope,omitempty"`
}
