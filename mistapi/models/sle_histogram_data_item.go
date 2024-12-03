package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleHistogramDataItem represents a SleHistogramDataItem struct.
type SleHistogramDataItem struct {
    Range                []float64              `json:"range,omitempty"`
    Value                float64                `json:"value"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleHistogramDataItem.
// It customizes the JSON marshaling process for SleHistogramDataItem objects.
func (s SleHistogramDataItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "range", "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleHistogramDataItem object to a map representation for JSON marshaling.
func (s SleHistogramDataItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Range != nil {
        structMap["range"] = s.Range
    }
    structMap["value"] = s.Value
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleHistogramDataItem.
// It customizes the JSON unmarshaling process for SleHistogramDataItem objects.
func (s *SleHistogramDataItem) UnmarshalJSON(input []byte) error {
    var temp tempSleHistogramDataItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "range", "value")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Range = temp.Range
    s.Value = *temp.Value
    return nil
}

// tempSleHistogramDataItem is a temporary struct used for validating the fields of SleHistogramDataItem.
type tempSleHistogramDataItem  struct {
    Range []float64 `json:"range,omitempty"`
    Value *float64  `json:"value"`
}

func (s *tempSleHistogramDataItem) validate() error {
    var errs []string
    if s.Value == nil {
        errs = append(errs, "required field `value` is missing for type `sle_histogram_data_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
