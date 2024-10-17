package models

import (
	"encoding/json"
)

// SuppressedAlarm represents a SuppressedAlarm struct.
type SuppressedAlarm struct {
	// if `scope`==`site`
	// Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration`
	Applies *SuppressedAlarmApplies `json:"applies,omitempty"`
	// duration, in seconds. Maximum duration is 86400 * 180 (180 days). 0 is to un-suppress alarms
	Duration *float64 `json:"duration,omitempty"`
	// poch_time in seconds, Default as now, accepted time range is from now to now + 7 days
	ScheduledTime *int `json:"scheduled_time,omitempty"`
	// level of scope. enum: `org`, `site`
	Scope                *SuppressedAlarmScopeEnum `json:"scope,omitempty"`
	AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SuppressedAlarm.
// It customizes the JSON marshaling process for SuppressedAlarm objects.
func (s SuppressedAlarm) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SuppressedAlarm object to a map representation for JSON marshaling.
func (s SuppressedAlarm) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "applies", "duration", "scheduled_time", "scope")
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
