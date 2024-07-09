package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleSummarySle represents a SleSummarySle struct.
type SleSummarySle struct {
    Interval             float64              `json:"interval"`
    Name                 string               `json:"name"`
    Samples              SleSummarySleSamples `json:"samples"`
    XLabel               string               `json:"x_label"`
    YLabel               string               `json:"y_label"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleSummarySle.
// It customizes the JSON marshaling process for SleSummarySle objects.
func (s SleSummarySle) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleSummarySle object to a map representation for JSON marshaling.
func (s SleSummarySle) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["interval"] = s.Interval
    structMap["name"] = s.Name
    structMap["samples"] = s.Samples.toMap()
    structMap["x_label"] = s.XLabel
    structMap["y_label"] = s.YLabel
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleSummarySle.
// It customizes the JSON unmarshaling process for SleSummarySle objects.
func (s *SleSummarySle) UnmarshalJSON(input []byte) error {
    var temp sleSummarySle
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "interval", "name", "samples", "x_label", "y_label")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Interval = *temp.Interval
    s.Name = *temp.Name
    s.Samples = *temp.Samples
    s.XLabel = *temp.XLabel
    s.YLabel = *temp.YLabel
    return nil
}

// sleSummarySle is a temporary struct used for validating the fields of SleSummarySle.
type sleSummarySle  struct {
    Interval *float64              `json:"interval"`
    Name     *string               `json:"name"`
    Samples  *SleSummarySleSamples `json:"samples"`
    XLabel   *string               `json:"x_label"`
    YLabel   *string               `json:"y_label"`
}

func (s *sleSummarySle) validate() error {
    var errs []string
    if s.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `Sle_Summary_Sle`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sle_Summary_Sle`")
    }
    if s.Samples == nil {
        errs = append(errs, "required field `samples` is missing for type `Sle_Summary_Sle`")
    }
    if s.XLabel == nil {
        errs = append(errs, "required field `x_label` is missing for type `Sle_Summary_Sle`")
    }
    if s.YLabel == nil {
        errs = append(errs, "required field `y_label` is missing for type `Sle_Summary_Sle`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
