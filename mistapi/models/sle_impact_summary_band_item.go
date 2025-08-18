// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SleImpactSummaryBandItem represents a SleImpactSummaryBandItem struct.
type SleImpactSummaryBandItem struct {
	Band                 string                 `json:"band"`
	Degraded             float64                `json:"degraded"`
	Duration             float64                `json:"duration"`
	Name                 string                 `json:"name"`
	Total                float64                `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactSummaryBandItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactSummaryBandItem) String() string {
	return fmt.Sprintf(
		"SleImpactSummaryBandItem[Band=%v, Degraded=%v, Duration=%v, Name=%v, Total=%v, AdditionalProperties=%v]",
		s.Band, s.Degraded, s.Duration, s.Name, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryBandItem.
// It customizes the JSON marshaling process for SleImpactSummaryBandItem objects.
func (s SleImpactSummaryBandItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"band", "degraded", "duration", "name", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryBandItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryBandItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["band"] = s.Band
	structMap["degraded"] = s.Degraded
	structMap["duration"] = s.Duration
	structMap["name"] = s.Name
	structMap["total"] = s.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactSummaryBandItem.
// It customizes the JSON unmarshaling process for SleImpactSummaryBandItem objects.
func (s *SleImpactSummaryBandItem) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactSummaryBandItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "degraded", "duration", "name", "total")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Band = *temp.Band
	s.Degraded = *temp.Degraded
	s.Duration = *temp.Duration
	s.Name = *temp.Name
	s.Total = *temp.Total
	return nil
}

// tempSleImpactSummaryBandItem is a temporary struct used for validating the fields of SleImpactSummaryBandItem.
type tempSleImpactSummaryBandItem struct {
	Band     *string  `json:"band"`
	Degraded *float64 `json:"degraded"`
	Duration *float64 `json:"duration"`
	Name     *string  `json:"name"`
	Total    *float64 `json:"total"`
}

func (s *tempSleImpactSummaryBandItem) validate() error {
	var errs []string
	if s.Band == nil {
		errs = append(errs, "required field `band` is missing for type `sle_impact_summary_band_item`")
	}
	if s.Degraded == nil {
		errs = append(errs, "required field `degraded` is missing for type `sle_impact_summary_band_item`")
	}
	if s.Duration == nil {
		errs = append(errs, "required field `duration` is missing for type `sle_impact_summary_band_item`")
	}
	if s.Name == nil {
		errs = append(errs, "required field `name` is missing for type `sle_impact_summary_band_item`")
	}
	if s.Total == nil {
		errs = append(errs, "required field `total` is missing for type `sle_impact_summary_band_item`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
