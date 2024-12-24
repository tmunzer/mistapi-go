package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleSummarySleSamples represents a SleSummarySleSamples struct.
type SleSummarySleSamples struct {
    Degraded             []float64              `json:"degraded"`
    Total                []float64              `json:"total"`
    Value                []float64              `json:"value"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleSummarySleSamples,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleSummarySleSamples) String() string {
    return fmt.Sprintf(
    	"SleSummarySleSamples[Degraded=%v, Total=%v, Value=%v, AdditionalProperties=%v]",
    	s.Degraded, s.Total, s.Value, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleSummarySleSamples.
// It customizes the JSON marshaling process for SleSummarySleSamples objects.
func (s SleSummarySleSamples) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "degraded", "total", "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleSummarySleSamples object to a map representation for JSON marshaling.
func (s SleSummarySleSamples) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["degraded"] = s.Degraded
    structMap["total"] = s.Total
    structMap["value"] = s.Value
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleSummarySleSamples.
// It customizes the JSON unmarshaling process for SleSummarySleSamples objects.
func (s *SleSummarySleSamples) UnmarshalJSON(input []byte) error {
    var temp tempSleSummarySleSamples
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "degraded", "total", "value")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Degraded = *temp.Degraded
    s.Total = *temp.Total
    s.Value = *temp.Value
    return nil
}

// tempSleSummarySleSamples is a temporary struct used for validating the fields of SleSummarySleSamples.
type tempSleSummarySleSamples  struct {
    Degraded *[]float64 `json:"degraded"`
    Total    *[]float64 `json:"total"`
    Value    *[]float64 `json:"value"`
}

func (s *tempSleSummarySleSamples) validate() error {
    var errs []string
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `sle_summary_sle_samples`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `sle_summary_sle_samples`")
    }
    if s.Value == nil {
        errs = append(errs, "required field `value` is missing for type `sle_summary_sle_samples`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
