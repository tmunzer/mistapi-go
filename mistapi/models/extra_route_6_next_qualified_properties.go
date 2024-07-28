package models

import (
    "encoding/json"
)

// ExtraRoute6NextQualifiedProperties represents a ExtraRoute6NextQualifiedProperties struct.
type ExtraRoute6NextQualifiedProperties struct {
    Metric               Optional[int]  `json:"metric"`
    Preference           Optional[int]  `json:"preference"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ExtraRoute6NextQualifiedProperties.
// It customizes the JSON marshaling process for ExtraRoute6NextQualifiedProperties objects.
func (e ExtraRoute6NextQualifiedProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the ExtraRoute6NextQualifiedProperties object to a map representation for JSON marshaling.
func (e ExtraRoute6NextQualifiedProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Metric.IsValueSet() {
        if e.Metric.Value() != nil {
            structMap["metric"] = e.Metric.Value()
        } else {
            structMap["metric"] = nil
        }
    }
    if e.Preference.IsValueSet() {
        if e.Preference.Value() != nil {
            structMap["preference"] = e.Preference.Value()
        } else {
            structMap["preference"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExtraRoute6NextQualifiedProperties.
// It customizes the JSON unmarshaling process for ExtraRoute6NextQualifiedProperties objects.
func (e *ExtraRoute6NextQualifiedProperties) UnmarshalJSON(input []byte) error {
    var temp extraRoute6NextQualifiedProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "metric", "preference")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Metric = temp.Metric
    e.Preference = temp.Preference
    return nil
}

// extraRoute6NextQualifiedProperties is a temporary struct used for validating the fields of ExtraRoute6NextQualifiedProperties.
type extraRoute6NextQualifiedProperties  struct {
    Metric     Optional[int] `json:"metric"`
    Preference Optional[int] `json:"preference"`
}