package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleClassifierSamples represents a SleClassifierSamples struct.
type SleClassifierSamples struct {
    Degraded             []float64              `json:"degraded"`
    Duration             []float64              `json:"duration"`
    Total                []float64              `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleClassifierSamples,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleClassifierSamples) String() string {
    return fmt.Sprintf(
    	"SleClassifierSamples[Degraded=%v, Duration=%v, Total=%v, AdditionalProperties=%v]",
    	s.Degraded, s.Duration, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleClassifierSamples.
// It customizes the JSON marshaling process for SleClassifierSamples objects.
func (s SleClassifierSamples) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "degraded", "duration", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleClassifierSamples object to a map representation for JSON marshaling.
func (s SleClassifierSamples) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["degraded"] = s.Degraded
    structMap["duration"] = s.Duration
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleClassifierSamples.
// It customizes the JSON unmarshaling process for SleClassifierSamples objects.
func (s *SleClassifierSamples) UnmarshalJSON(input []byte) error {
    var temp tempSleClassifierSamples
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "degraded", "duration", "total")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Degraded = *temp.Degraded
    s.Duration = *temp.Duration
    s.Total = *temp.Total
    return nil
}

// tempSleClassifierSamples is a temporary struct used for validating the fields of SleClassifierSamples.
type tempSleClassifierSamples  struct {
    Degraded *[]float64 `json:"degraded"`
    Duration *[]float64 `json:"duration"`
    Total    *[]float64 `json:"total"`
}

func (s *tempSleClassifierSamples) validate() error {
    var errs []string
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `sle_classifier_samples`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `sle_classifier_samples`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `sle_classifier_samples`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
