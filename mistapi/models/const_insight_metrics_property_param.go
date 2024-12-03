package models

import (
    "encoding/json"
)

// ConstInsightMetricsPropertyParam represents a ConstInsightMetricsPropertyParam struct.
type ConstInsightMetricsPropertyParam struct {
    Required             *bool                  `json:"required,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstInsightMetricsPropertyParam.
// It customizes the JSON marshaling process for ConstInsightMetricsPropertyParam objects.
func (c ConstInsightMetricsPropertyParam) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "required"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstInsightMetricsPropertyParam object to a map representation for JSON marshaling.
func (c ConstInsightMetricsPropertyParam) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Required != nil {
        structMap["required"] = c.Required
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstInsightMetricsPropertyParam.
// It customizes the JSON unmarshaling process for ConstInsightMetricsPropertyParam objects.
func (c *ConstInsightMetricsPropertyParam) UnmarshalJSON(input []byte) error {
    var temp tempConstInsightMetricsPropertyParam
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "required")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Required = temp.Required
    return nil
}

// tempConstInsightMetricsPropertyParam is a temporary struct used for validating the fields of ConstInsightMetricsPropertyParam.
type tempConstInsightMetricsPropertyParam  struct {
    Required *bool `json:"required,omitempty"`
}
