package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleClassifierSamples represents a SleClassifierSamples struct.
type SleClassifierSamples struct {
    Degraded             []float64      `json:"degraded"`
    Duration             []float64      `json:"duration"`
    Total                []float64      `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleClassifierSamples.
// It customizes the JSON marshaling process for SleClassifierSamples objects.
func (s SleClassifierSamples) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleClassifierSamples object to a map representation for JSON marshaling.
func (s SleClassifierSamples) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["degraded"] = s.Degraded
    structMap["duration"] = s.Duration
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleClassifierSamples.
// It customizes the JSON unmarshaling process for SleClassifierSamples objects.
func (s *SleClassifierSamples) UnmarshalJSON(input []byte) error {
    var temp sleClassifierSamples
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "degraded", "duration", "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Degraded = *temp.Degraded
    s.Duration = *temp.Duration
    s.Total = *temp.Total
    return nil
}

// sleClassifierSamples is a temporary struct used for validating the fields of SleClassifierSamples.
type sleClassifierSamples  struct {
    Degraded *[]float64 `json:"degraded"`
    Duration *[]float64 `json:"duration"`
    Total    *[]float64 `json:"total"`
}

func (s *sleClassifierSamples) validate() error {
    var errs []string
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `Sle_Classifier_Samples`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `Sle_Classifier_Samples`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Sle_Classifier_Samples`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}