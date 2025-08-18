// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SleHistogram represents a SleHistogram struct.
type SleHistogram struct {
	Data                 []SleHistogramDataItem `json:"data"`
	End                  float64                `json:"end"`
	Metric               string                 `json:"metric"`
	Start                float64                `json:"start"`
	XLabel               string                 `json:"x_label"`
	YLabel               string                 `json:"y_label"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleHistogram,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleHistogram) String() string {
	return fmt.Sprintf(
		"SleHistogram[Data=%v, End=%v, Metric=%v, Start=%v, XLabel=%v, YLabel=%v, AdditionalProperties=%v]",
		s.Data, s.End, s.Metric, s.Start, s.XLabel, s.YLabel, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleHistogram.
// It customizes the JSON marshaling process for SleHistogram objects.
func (s SleHistogram) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"data", "end", "metric", "start", "x_label", "y_label"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleHistogram object to a map representation for JSON marshaling.
func (s SleHistogram) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["data"] = s.Data
	structMap["end"] = s.End
	structMap["metric"] = s.Metric
	structMap["start"] = s.Start
	structMap["x_label"] = s.XLabel
	structMap["y_label"] = s.YLabel
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleHistogram.
// It customizes the JSON unmarshaling process for SleHistogram objects.
func (s *SleHistogram) UnmarshalJSON(input []byte) error {
	var temp tempSleHistogram
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "data", "end", "metric", "start", "x_label", "y_label")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Data = *temp.Data
	s.End = *temp.End
	s.Metric = *temp.Metric
	s.Start = *temp.Start
	s.XLabel = *temp.XLabel
	s.YLabel = *temp.YLabel
	return nil
}

// tempSleHistogram is a temporary struct used for validating the fields of SleHistogram.
type tempSleHistogram struct {
	Data   *[]SleHistogramDataItem `json:"data"`
	End    *float64                `json:"end"`
	Metric *string                 `json:"metric"`
	Start  *float64                `json:"start"`
	XLabel *string                 `json:"x_label"`
	YLabel *string                 `json:"y_label"`
}

func (s *tempSleHistogram) validate() error {
	var errs []string
	if s.Data == nil {
		errs = append(errs, "required field `data` is missing for type `sle_histogram`")
	}
	if s.End == nil {
		errs = append(errs, "required field `end` is missing for type `sle_histogram`")
	}
	if s.Metric == nil {
		errs = append(errs, "required field `metric` is missing for type `sle_histogram`")
	}
	if s.Start == nil {
		errs = append(errs, "required field `start` is missing for type `sle_histogram`")
	}
	if s.XLabel == nil {
		errs = append(errs, "required field `x_label` is missing for type `sle_histogram`")
	}
	if s.YLabel == nil {
		errs = append(errs, "required field `y_label` is missing for type `sle_histogram`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
