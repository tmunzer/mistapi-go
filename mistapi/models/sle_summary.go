// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleSummary represents a SleSummary struct.
type SleSummary struct {
    Classifiers          []SleClassifier        `json:"classifiers"`
    End                  float64                `json:"end"`
    Events               []interface{}          `json:"events"`
    Impact               SleSummaryImpact       `json:"impact"`
    Sle                  SleSummarySle          `json:"sle"`
    Start                float64                `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleSummary) String() string {
    return fmt.Sprintf(
    	"SleSummary[Classifiers=%v, End=%v, Events=%v, Impact=%v, Sle=%v, Start=%v, AdditionalProperties=%v]",
    	s.Classifiers, s.End, s.Events, s.Impact, s.Sle, s.Start, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleSummary.
// It customizes the JSON marshaling process for SleSummary objects.
func (s SleSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "classifiers", "end", "events", "impact", "sle", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleSummary object to a map representation for JSON marshaling.
func (s SleSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["classifiers"] = s.Classifiers
    structMap["end"] = s.End
    structMap["events"] = s.Events
    structMap["impact"] = s.Impact.toMap()
    structMap["sle"] = s.Sle.toMap()
    structMap["start"] = s.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleSummary.
// It customizes the JSON unmarshaling process for SleSummary objects.
func (s *SleSummary) UnmarshalJSON(input []byte) error {
    var temp tempSleSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "classifiers", "end", "events", "impact", "sle", "start")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Classifiers = *temp.Classifiers
    s.End = *temp.End
    s.Events = *temp.Events
    s.Impact = *temp.Impact
    s.Sle = *temp.Sle
    s.Start = *temp.Start
    return nil
}

// tempSleSummary is a temporary struct used for validating the fields of SleSummary.
type tempSleSummary  struct {
    Classifiers *[]SleClassifier  `json:"classifiers"`
    End         *float64          `json:"end"`
    Events      *[]interface{}    `json:"events"`
    Impact      *SleSummaryImpact `json:"impact"`
    Sle         *SleSummarySle    `json:"sle"`
    Start       *float64          `json:"start"`
}

func (s *tempSleSummary) validate() error {
    var errs []string
    if s.Classifiers == nil {
        errs = append(errs, "required field `classifiers` is missing for type `sle_summary`")
    }
    if s.End == nil {
        errs = append(errs, "required field `end` is missing for type `sle_summary`")
    }
    if s.Events == nil {
        errs = append(errs, "required field `events` is missing for type `sle_summary`")
    }
    if s.Impact == nil {
        errs = append(errs, "required field `impact` is missing for type `sle_summary`")
    }
    if s.Sle == nil {
        errs = append(errs, "required field `sle` is missing for type `sle_summary`")
    }
    if s.Start == nil {
        errs = append(errs, "required field `start` is missing for type `sle_summary`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
