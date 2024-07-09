package models

import (
    "encoding/json"
)

// ConstInsightMetricsPropertyParam represents a ConstInsightMetricsPropertyParam struct.
type ConstInsightMetricsPropertyParam struct {
    Required             *bool          `json:"required,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyParam.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyParam objects.
func (c ConstInsightMetricsPropertyParam) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyParam object to a map representation for JSON marshaling.
func (c ConstInsightMetricsPropertyParam) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Required != nil {
        structMap["required"] = c.Required
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyParam.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyParam objects.
func (c *ConstInsightMetricsPropertyParam) UnmarshalJSON(input []byte) error {
    var temp constInsightMetricsPropertyParam
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "required")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Required = temp.Required
    return nil
}

// constInsightMetricsPropertyParam is a temporary struct used for validating the fields of ConstInsightMetricsPropertyParam.
type constInsightMetricsPropertyParam  struct {
    Required *bool `json:"required,omitempty"`
}
