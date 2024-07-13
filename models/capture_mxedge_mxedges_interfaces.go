package models

import (
    "encoding/json"
)

// CaptureMxedgeMxedgesInterfaces represents a CaptureMxedgeMxedgesInterfaces struct.
// Property key is the Port name (e.g. "port1")
type CaptureMxedgeMxedgesInterfaces struct {
    // tcpdump expression common for wired,radiotap
    TcpdumpExpression    *string        `json:"tcpdump_expression,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureMxedgeMxedgesInterfaces.
// It customizes the JSON marshaling process for CaptureMxedgeMxedgesInterfaces objects.
func (c CaptureMxedgeMxedgesInterfaces) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureMxedgeMxedgesInterfaces object to a map representation for JSON marshaling.
func (c CaptureMxedgeMxedgesInterfaces) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.TcpdumpExpression != nil {
        structMap["tcpdump_expression"] = c.TcpdumpExpression
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureMxedgeMxedgesInterfaces.
// It customizes the JSON unmarshaling process for CaptureMxedgeMxedgesInterfaces objects.
func (c *CaptureMxedgeMxedgesInterfaces) UnmarshalJSON(input []byte) error {
    var temp captureMxedgeMxedgesInterfaces
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

// captureMxedgeMxedgesInterfaces is a temporary struct used for validating the fields of CaptureMxedgeMxedgesInterfaces.
type captureMxedgeMxedgesInterfaces  struct {
    TcpdumpExpression *string `json:"tcpdump_expression,omitempty"`
}
