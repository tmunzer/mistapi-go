package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleClassifierSummary represents a SleClassifierSummary struct.
type SleClassifierSummary struct {
    Classifier           SleClassifier              `json:"classifier"`
    End                  float64                    `json:"end"`
    Failures             []interface{}              `json:"failures"`
    Impact               SleClassifierSummaryImpact `json:"impact"`
    Metric               string                     `json:"metric"`
    Start                float64                    `json:"start"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleClassifierSummary.
// It customizes the JSON marshaling process for SleClassifierSummary objects.
func (s SleClassifierSummary) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleClassifierSummary object to a map representation for JSON marshaling.
func (s SleClassifierSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["classifier"] = s.Classifier.toMap()
    structMap["end"] = s.End
    structMap["failures"] = s.Failures
    structMap["impact"] = s.Impact.toMap()
    structMap["metric"] = s.Metric
    structMap["start"] = s.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleClassifierSummary.
// It customizes the JSON unmarshaling process for SleClassifierSummary objects.
func (s *SleClassifierSummary) UnmarshalJSON(input []byte) error {
    var temp tempSleClassifierSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "classifier", "end", "failures", "impact", "metric", "start")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Classifier = *temp.Classifier
    s.End = *temp.End
    s.Failures = *temp.Failures
    s.Impact = *temp.Impact
    s.Metric = *temp.Metric
    s.Start = *temp.Start
    return nil
}

// tempSleClassifierSummary is a temporary struct used for validating the fields of SleClassifierSummary.
type tempSleClassifierSummary  struct {
    Classifier *SleClassifier              `json:"classifier"`
    End        *float64                    `json:"end"`
    Failures   *[]interface{}              `json:"failures"`
    Impact     *SleClassifierSummaryImpact `json:"impact"`
    Metric     *string                     `json:"metric"`
    Start      *float64                    `json:"start"`
}

func (s *tempSleClassifierSummary) validate() error {
    var errs []string
    if s.Classifier == nil {
        errs = append(errs, "required field `classifier` is missing for type `sle_classifier_summary`")
    }
    if s.End == nil {
        errs = append(errs, "required field `end` is missing for type `sle_classifier_summary`")
    }
    if s.Failures == nil {
        errs = append(errs, "required field `failures` is missing for type `sle_classifier_summary`")
    }
    if s.Impact == nil {
        errs = append(errs, "required field `impact` is missing for type `sle_classifier_summary`")
    }
    if s.Metric == nil {
        errs = append(errs, "required field `metric` is missing for type `sle_classifier_summary`")
    }
    if s.Start == nil {
        errs = append(errs, "required field `start` is missing for type `sle_classifier_summary`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
