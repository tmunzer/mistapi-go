// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SleClassifierSummaryTrend represents a SleClassifierSummaryTrend struct.
// Time-series trend response for an SLE classifier
type SleClassifierSummaryTrend struct {
	// Time-series classifier trend detail for an SLE metric
	Classifier SleTrendClassifier `json:"classifier"`
	// Last timestamp in the classifier trend window
	End float64 `json:"end"`
	// SLE metric name represented by the trend
	Metric string `json:"metric"`
	// First timestamp in the classifier trend window
	Start                float64                `json:"start"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleClassifierSummaryTrend,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleClassifierSummaryTrend) String() string {
	return fmt.Sprintf(
		"SleClassifierSummaryTrend[Classifier=%v, End=%v, Metric=%v, Start=%v, AdditionalProperties=%v]",
		s.Classifier, s.End, s.Metric, s.Start, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleClassifierSummaryTrend.
// It customizes the JSON marshaling process for SleClassifierSummaryTrend objects.
func (s SleClassifierSummaryTrend) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"classifier", "end", "metric", "start"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleClassifierSummaryTrend object to a map representation for JSON marshaling.
func (s SleClassifierSummaryTrend) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["classifier"] = s.Classifier.toMap()
	structMap["end"] = s.End
	structMap["metric"] = s.Metric
	structMap["start"] = s.Start
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleClassifierSummaryTrend.
// It customizes the JSON unmarshaling process for SleClassifierSummaryTrend objects.
func (s *SleClassifierSummaryTrend) UnmarshalJSON(input []byte) error {
	var temp tempSleClassifierSummaryTrend
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "classifier", "end", "metric", "start")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Classifier = *temp.Classifier
	s.End = *temp.End
	s.Metric = *temp.Metric
	s.Start = *temp.Start
	return nil
}

// tempSleClassifierSummaryTrend is a temporary struct used for validating the fields of SleClassifierSummaryTrend.
type tempSleClassifierSummaryTrend struct {
	Classifier *SleTrendClassifier `json:"classifier"`
	End        *float64            `json:"end"`
	Metric     *string             `json:"metric"`
	Start      *float64            `json:"start"`
}

func (s *tempSleClassifierSummaryTrend) validate() error {
	var errs []string
	if s.Classifier == nil {
		errs = append(errs, "required field `classifier` is missing for type `sle_classifier_summary_trend`")
	}
	if s.End == nil {
		errs = append(errs, "required field `end` is missing for type `sle_classifier_summary_trend`")
	}
	if s.Metric == nil {
		errs = append(errs, "required field `metric` is missing for type `sle_classifier_summary_trend`")
	}
	if s.Start == nil {
		errs = append(errs, "required field `start` is missing for type `sle_classifier_summary_trend`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
