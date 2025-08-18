// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SleImpactedSwitches represents a SleImpactedSwitches struct.
type SleImpactedSwitches struct {
	Classifier           *string                     `json:"classifier,omitempty"`
	End                  *int                        `json:"end,omitempty"`
	Failure              *string                     `json:"failure,omitempty"`
	Limit                *int                        `json:"limit,omitempty"`
	Metric               *string                     `json:"metric,omitempty"`
	Page                 *int                        `json:"page,omitempty"`
	Start                *int                        `json:"start,omitempty"`
	Switches             []SleImpactedSwitchesSwitch `json:"switches,omitempty"`
	TotalCount           *int                        `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedSwitches,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedSwitches) String() string {
	return fmt.Sprintf(
		"SleImpactedSwitches[Classifier=%v, End=%v, Failure=%v, Limit=%v, Metric=%v, Page=%v, Start=%v, Switches=%v, TotalCount=%v, AdditionalProperties=%v]",
		s.Classifier, s.End, s.Failure, s.Limit, s.Metric, s.Page, s.Start, s.Switches, s.TotalCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedSwitches.
// It customizes the JSON marshaling process for SleImpactedSwitches objects.
func (s SleImpactedSwitches) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"classifier", "end", "failure", "limit", "metric", "page", "start", "switches", "total_count"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedSwitches object to a map representation for JSON marshaling.
func (s SleImpactedSwitches) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Classifier != nil {
		structMap["classifier"] = s.Classifier
	}
	if s.End != nil {
		structMap["end"] = s.End
	}
	if s.Failure != nil {
		structMap["failure"] = s.Failure
	}
	if s.Limit != nil {
		structMap["limit"] = s.Limit
	}
	if s.Metric != nil {
		structMap["metric"] = s.Metric
	}
	if s.Page != nil {
		structMap["page"] = s.Page
	}
	if s.Start != nil {
		structMap["start"] = s.Start
	}
	if s.Switches != nil {
		structMap["switches"] = s.Switches
	}
	if s.TotalCount != nil {
		structMap["total_count"] = s.TotalCount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedSwitches.
// It customizes the JSON unmarshaling process for SleImpactedSwitches objects.
func (s *SleImpactedSwitches) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactedSwitches
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "classifier", "end", "failure", "limit", "metric", "page", "start", "switches", "total_count")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Classifier = temp.Classifier
	s.End = temp.End
	s.Failure = temp.Failure
	s.Limit = temp.Limit
	s.Metric = temp.Metric
	s.Page = temp.Page
	s.Start = temp.Start
	s.Switches = temp.Switches
	s.TotalCount = temp.TotalCount
	return nil
}

// tempSleImpactedSwitches is a temporary struct used for validating the fields of SleImpactedSwitches.
type tempSleImpactedSwitches struct {
	Classifier *string                     `json:"classifier,omitempty"`
	End        *int                        `json:"end,omitempty"`
	Failure    *string                     `json:"failure,omitempty"`
	Limit      *int                        `json:"limit,omitempty"`
	Metric     *string                     `json:"metric,omitempty"`
	Page       *int                        `json:"page,omitempty"`
	Start      *int                        `json:"start,omitempty"`
	Switches   []SleImpactedSwitchesSwitch `json:"switches,omitempty"`
	TotalCount *int                        `json:"total_count,omitempty"`
}
