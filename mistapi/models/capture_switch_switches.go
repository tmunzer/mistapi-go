package models

import (
    "encoding/json"
)

// CaptureSwitchSwitches represents a CaptureSwitchSwitches struct.
type CaptureSwitchSwitches struct {
    // Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request
    Ports                map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
    AdditionalProperties map[string]any                                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureSwitchSwitches.
// It customizes the JSON marshaling process for CaptureSwitchSwitches objects.
func (c CaptureSwitchSwitches) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CaptureSwitchSwitches object to a map representation for JSON marshaling.
func (c CaptureSwitchSwitches) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Ports != nil {
        structMap["ports"] = c.Ports
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureSwitchSwitches.
// It customizes the JSON unmarshaling process for CaptureSwitchSwitches objects.
func (c *CaptureSwitchSwitches) UnmarshalJSON(input []byte) error {
    var temp tempCaptureSwitchSwitches
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ports")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Ports = temp.Ports
    return nil
}

// tempCaptureSwitchSwitches is a temporary struct used for validating the fields of CaptureSwitchSwitches.
type tempCaptureSwitchSwitches  struct {
    Ports map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
}
