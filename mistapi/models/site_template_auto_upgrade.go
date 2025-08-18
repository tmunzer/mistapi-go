// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SiteTemplateAutoUpgrade represents a SiteTemplateAutoUpgrade struct.
type SiteTemplateAutoUpgrade struct {
	// enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
	DayOfWeek            *DayOfWeekEnum         `json:"day_of_week,omitempty"`
	Enabled              *bool                  `json:"enabled,omitempty"`
	TimeOfDay            *string                `json:"time_of_day,omitempty"`
	Version              *string                `json:"version,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteTemplateAutoUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteTemplateAutoUpgrade) String() string {
	return fmt.Sprintf(
		"SiteTemplateAutoUpgrade[DayOfWeek=%v, Enabled=%v, TimeOfDay=%v, Version=%v, AdditionalProperties=%v]",
		s.DayOfWeek, s.Enabled, s.TimeOfDay, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteTemplateAutoUpgrade.
// It customizes the JSON marshaling process for SiteTemplateAutoUpgrade objects.
func (s SiteTemplateAutoUpgrade) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"day_of_week", "enabled", "time_of_day", "version"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteTemplateAutoUpgrade object to a map representation for JSON marshaling.
func (s SiteTemplateAutoUpgrade) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.DayOfWeek != nil {
		structMap["day_of_week"] = s.DayOfWeek
	}
	if s.Enabled != nil {
		structMap["enabled"] = s.Enabled
	}
	if s.TimeOfDay != nil {
		structMap["time_of_day"] = s.TimeOfDay
	}
	if s.Version != nil {
		structMap["version"] = s.Version
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteTemplateAutoUpgrade.
// It customizes the JSON unmarshaling process for SiteTemplateAutoUpgrade objects.
func (s *SiteTemplateAutoUpgrade) UnmarshalJSON(input []byte) error {
	var temp tempSiteTemplateAutoUpgrade
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "day_of_week", "enabled", "time_of_day", "version")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.DayOfWeek = temp.DayOfWeek
	s.Enabled = temp.Enabled
	s.TimeOfDay = temp.TimeOfDay
	s.Version = temp.Version
	return nil
}

// tempSiteTemplateAutoUpgrade is a temporary struct used for validating the fields of SiteTemplateAutoUpgrade.
type tempSiteTemplateAutoUpgrade struct {
	DayOfWeek *DayOfWeekEnum `json:"day_of_week,omitempty"`
	Enabled   *bool          `json:"enabled,omitempty"`
	TimeOfDay *string        `json:"time_of_day,omitempty"`
	Version   *string        `json:"version,omitempty"`
}
