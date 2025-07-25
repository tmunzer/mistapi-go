// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleImpactedAps represents a SleImpactedAps struct.
type SleImpactedAps struct {
    Aps                  []SleImpactedApsAp     `json:"aps"`
    Classifier           string                 `json:"classifier"`
    End                  float64                `json:"end"`
    Failure              string                 `json:"failure"`
    Limit                float64                `json:"limit"`
    Metric               string                 `json:"metric"`
    Page                 float64                `json:"page"`
    Start                float64                `json:"start"`
    TotalCount           float64                `json:"total_count"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedAps,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedAps) String() string {
    return fmt.Sprintf(
    	"SleImpactedAps[Aps=%v, Classifier=%v, End=%v, Failure=%v, Limit=%v, Metric=%v, Page=%v, Start=%v, TotalCount=%v, AdditionalProperties=%v]",
    	s.Aps, s.Classifier, s.End, s.Failure, s.Limit, s.Metric, s.Page, s.Start, s.TotalCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedAps.
// It customizes the JSON marshaling process for SleImpactedAps objects.
func (s SleImpactedAps) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "aps", "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedAps object to a map representation for JSON marshaling.
func (s SleImpactedAps) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["aps"] = s.Aps
    structMap["classifier"] = s.Classifier
    structMap["end"] = s.End
    structMap["failure"] = s.Failure
    structMap["limit"] = s.Limit
    structMap["metric"] = s.Metric
    structMap["page"] = s.Page
    structMap["start"] = s.Start
    structMap["total_count"] = s.TotalCount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedAps.
// It customizes the JSON unmarshaling process for SleImpactedAps objects.
func (s *SleImpactedAps) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedAps
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aps", "classifier", "end", "failure", "limit", "metric", "page", "start", "total_count")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Aps = *temp.Aps
    s.Classifier = *temp.Classifier
    s.End = *temp.End
    s.Failure = *temp.Failure
    s.Limit = *temp.Limit
    s.Metric = *temp.Metric
    s.Page = *temp.Page
    s.Start = *temp.Start
    s.TotalCount = *temp.TotalCount
    return nil
}

// tempSleImpactedAps is a temporary struct used for validating the fields of SleImpactedAps.
type tempSleImpactedAps  struct {
    Aps        *[]SleImpactedApsAp `json:"aps"`
    Classifier *string             `json:"classifier"`
    End        *float64            `json:"end"`
    Failure    *string             `json:"failure"`
    Limit      *float64            `json:"limit"`
    Metric     *string             `json:"metric"`
    Page       *float64            `json:"page"`
    Start      *float64            `json:"start"`
    TotalCount *float64            `json:"total_count"`
}

func (s *tempSleImpactedAps) validate() error {
    var errs []string
    if s.Aps == nil {
        errs = append(errs, "required field `aps` is missing for type `sle_impacted_aps`")
    }
    if s.Classifier == nil {
        errs = append(errs, "required field `classifier` is missing for type `sle_impacted_aps`")
    }
    if s.End == nil {
        errs = append(errs, "required field `end` is missing for type `sle_impacted_aps`")
    }
    if s.Failure == nil {
        errs = append(errs, "required field `failure` is missing for type `sle_impacted_aps`")
    }
    if s.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `sle_impacted_aps`")
    }
    if s.Metric == nil {
        errs = append(errs, "required field `metric` is missing for type `sle_impacted_aps`")
    }
    if s.Page == nil {
        errs = append(errs, "required field `page` is missing for type `sle_impacted_aps`")
    }
    if s.Start == nil {
        errs = append(errs, "required field `start` is missing for type `sle_impacted_aps`")
    }
    if s.TotalCount == nil {
        errs = append(errs, "required field `total_count` is missing for type `sle_impacted_aps`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
