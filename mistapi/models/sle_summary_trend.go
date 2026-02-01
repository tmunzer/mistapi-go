// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SleSummaryTrend represents a SleSummaryTrend struct.
type SleSummaryTrend struct {
	Classifiers          []SleTrendClassifier   `json:"classifiers"`
	End                  float64                `json:"end"`
	Sle                  SleSummarySle          `json:"sle"`
	Start                float64                `json:"start"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleSummaryTrend,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleSummaryTrend) String() string {
	return fmt.Sprintf(
		"SleSummaryTrend[Classifiers=%v, End=%v, Sle=%v, Start=%v, AdditionalProperties=%v]",
		s.Classifiers, s.End, s.Sle, s.Start, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleSummaryTrend.
// It customizes the JSON marshaling process for SleSummaryTrend objects.
func (s SleSummaryTrend) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"classifiers", "end", "sle", "start"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleSummaryTrend object to a map representation for JSON marshaling.
func (s SleSummaryTrend) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["classifiers"] = s.Classifiers
	structMap["end"] = s.End
	structMap["sle"] = s.Sle.toMap()
	structMap["start"] = s.Start
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleSummaryTrend.
// It customizes the JSON unmarshaling process for SleSummaryTrend objects.
func (s *SleSummaryTrend) UnmarshalJSON(input []byte) error {
	var temp tempSleSummaryTrend
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "classifiers", "end", "sle", "start")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Classifiers = *temp.Classifiers
	s.End = *temp.End
	s.Sle = *temp.Sle
	s.Start = *temp.Start
	return nil
}

// tempSleSummaryTrend is a temporary struct used for validating the fields of SleSummaryTrend.
type tempSleSummaryTrend struct {
	Classifiers *[]SleTrendClassifier `json:"classifiers"`
	End         *float64              `json:"end"`
	Sle         *SleSummarySle        `json:"sle"`
	Start       *float64              `json:"start"`
}

func (s *tempSleSummaryTrend) validate() error {
	var errs []string
	if s.Classifiers == nil {
		errs = append(errs, "required field `classifiers` is missing for type `sle_summary_trend`")
	}
	if s.End == nil {
		errs = append(errs, "required field `end` is missing for type `sle_summary_trend`")
	}
	if s.Sle == nil {
		errs = append(errs, "required field `sle` is missing for type `sle_summary_trend`")
	}
	if s.Start == nil {
		errs = append(errs, "required field `start` is missing for type `sle_summary_trend`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
