package models

import (
    "encoding/json"
)

// ConstDeviceSwitchDefault represents a ConstDeviceSwitchDefault struct.
type ConstDeviceSwitchDefault struct {
    Ports                *string                `json:"_ports,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceSwitchDefault.
// It customizes the JSON marshaling process for ConstDeviceSwitchDefault objects.
func (c ConstDeviceSwitchDefault) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "_ports"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceSwitchDefault object to a map representation for JSON marshaling.
func (c ConstDeviceSwitchDefault) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Ports != nil {
        structMap["_ports"] = c.Ports
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceSwitchDefault.
// It customizes the JSON unmarshaling process for ConstDeviceSwitchDefault objects.
func (c *ConstDeviceSwitchDefault) UnmarshalJSON(input []byte) error {
    var temp tempConstDeviceSwitchDefault
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "_ports")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Ports = temp.Ports
    return nil
}

// tempConstDeviceSwitchDefault is a temporary struct used for validating the fields of ConstDeviceSwitchDefault.
type tempConstDeviceSwitchDefault  struct {
    Ports *string `json:"_ports,omitempty"`
}
