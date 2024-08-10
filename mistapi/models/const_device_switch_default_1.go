package models

import (
    "encoding/json"
)

// ConstDeviceSwitchDefault1 represents a ConstDeviceSwitchDefault1 struct.
// Object Key is the interface type name (e.g. "lan_ports", "wan_ports", ...)
type ConstDeviceSwitchDefault1 struct {
    Ports                *string        `json:"_ports,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceSwitchDefault1.
// It customizes the JSON marshaling process for ConstDeviceSwitchDefault1 objects.
func (c ConstDeviceSwitchDefault1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceSwitchDefault1 object to a map representation for JSON marshaling.
func (c ConstDeviceSwitchDefault1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Ports != nil {
        structMap["_ports"] = c.Ports
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceSwitchDefault1.
// It customizes the JSON unmarshaling process for ConstDeviceSwitchDefault1 objects.
func (c *ConstDeviceSwitchDefault1) UnmarshalJSON(input []byte) error {
    var temp tempConstDeviceSwitchDefault1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "_ports")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Ports = temp.Ports
    return nil
}

// tempConstDeviceSwitchDefault1 is a temporary struct used for validating the fields of ConstDeviceSwitchDefault1.
type tempConstDeviceSwitchDefault1  struct {
    Ports *string `json:"_ports,omitempty"`
}
