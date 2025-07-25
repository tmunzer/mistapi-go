// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleClassifier represents a SleClassifier struct.
type SleClassifier struct {
    Impact               SleClassifierImpact    `json:"impact"`
    Interval             float64                `json:"interval"`
    Name                 string                 `json:"name"`
    Samples              *SleClassifierSamples  `json:"samples,omitempty"`
    XLabel               string                 `json:"x_label"`
    YLabel               string                 `json:"y_label"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleClassifier,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleClassifier) String() string {
    return fmt.Sprintf(
    	"SleClassifier[Impact=%v, Interval=%v, Name=%v, Samples=%v, XLabel=%v, YLabel=%v, AdditionalProperties=%v]",
    	s.Impact, s.Interval, s.Name, s.Samples, s.XLabel, s.YLabel, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleClassifier.
// It customizes the JSON marshaling process for SleClassifier objects.
func (s SleClassifier) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "impact", "interval", "name", "samples", "x_label", "y_label"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleClassifier object to a map representation for JSON marshaling.
func (s SleClassifier) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["impact"] = s.Impact.toMap()
    structMap["interval"] = s.Interval
    structMap["name"] = s.Name
    if s.Samples != nil {
        structMap["samples"] = s.Samples.toMap()
    }
    structMap["x_label"] = s.XLabel
    structMap["y_label"] = s.YLabel
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleClassifier.
// It customizes the JSON unmarshaling process for SleClassifier objects.
func (s *SleClassifier) UnmarshalJSON(input []byte) error {
    var temp tempSleClassifier
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "impact", "interval", "name", "samples", "x_label", "y_label")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Impact = *temp.Impact
    s.Interval = *temp.Interval
    s.Name = *temp.Name
    s.Samples = temp.Samples
    s.XLabel = *temp.XLabel
    s.YLabel = *temp.YLabel
    return nil
}

// tempSleClassifier is a temporary struct used for validating the fields of SleClassifier.
type tempSleClassifier  struct {
    Impact   *SleClassifierImpact  `json:"impact"`
    Interval *float64              `json:"interval"`
    Name     *string               `json:"name"`
    Samples  *SleClassifierSamples `json:"samples,omitempty"`
    XLabel   *string               `json:"x_label"`
    YLabel   *string               `json:"y_label"`
}

func (s *tempSleClassifier) validate() error {
    var errs []string
    if s.Impact == nil {
        errs = append(errs, "required field `impact` is missing for type `sle_classifier`")
    }
    if s.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `sle_classifier`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sle_classifier`")
    }
    if s.XLabel == nil {
        errs = append(errs, "required field `x_label` is missing for type `sle_classifier`")
    }
    if s.YLabel == nil {
        errs = append(errs, "required field `y_label` is missing for type `sle_classifier`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
