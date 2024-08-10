package models

import (
    "encoding/json"
)

// CaptureSwitchPortsTcpdumpExpression represents a CaptureSwitchPortsTcpdumpExpression struct.
type CaptureSwitchPortsTcpdumpExpression struct {
    // tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
    TcpdumpExpression    *string        `json:"tcpdump_expression,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureSwitchPortsTcpdumpExpression.
// It customizes the JSON marshaling process for CaptureSwitchPortsTcpdumpExpression objects.
func (c CaptureSwitchPortsTcpdumpExpression) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureSwitchPortsTcpdumpExpression object to a map representation for JSON marshaling.
func (c CaptureSwitchPortsTcpdumpExpression) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.TcpdumpExpression != nil {
        structMap["tcpdump_expression"] = c.TcpdumpExpression
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureSwitchPortsTcpdumpExpression.
// It customizes the JSON unmarshaling process for CaptureSwitchPortsTcpdumpExpression objects.
func (c *CaptureSwitchPortsTcpdumpExpression) UnmarshalJSON(input []byte) error {
    var temp tempCaptureSwitchPortsTcpdumpExpression
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "tcpdump_expression")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.TcpdumpExpression = temp.TcpdumpExpression
    return nil
}

// tempCaptureSwitchPortsTcpdumpExpression is a temporary struct used for validating the fields of CaptureSwitchPortsTcpdumpExpression.
type tempCaptureSwitchPortsTcpdumpExpression  struct {
    TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
}
