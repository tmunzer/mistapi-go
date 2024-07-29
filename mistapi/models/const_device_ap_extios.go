package models

import (
    "encoding/json"
)

// ConstDeviceApExtios represents a ConstDeviceApExtios struct.
type ConstDeviceApExtios struct {
    // enum: `IN`, `OUT`
    DefaultDir           *ConstDeviceApExtiosDefaultDirEnum `json:"default_dir,omitempty"`
    Input                *bool                              `json:"input,omitempty"`
    Output               *bool                              `json:"output,omitempty"`
    AdditionalProperties map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceApExtios.
// It customizes the JSON marshaling process for ConstDeviceApExtios objects.
func (c ConstDeviceApExtios) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceApExtios object to a map representation for JSON marshaling.
func (c ConstDeviceApExtios) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.DefaultDir != nil {
        structMap["default_dir"] = c.DefaultDir
    }
    if c.Input != nil {
        structMap["input"] = c.Input
    }
    if c.Output != nil {
        structMap["output"] = c.Output
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceApExtios.
// It customizes the JSON unmarshaling process for ConstDeviceApExtios objects.
func (c *ConstDeviceApExtios) UnmarshalJSON(input []byte) error {
    var temp constDeviceApExtios
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "default_dir", "input", "output")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.DefaultDir = temp.DefaultDir
    c.Input = temp.Input
    c.Output = temp.Output
    return nil
}

// constDeviceApExtios is a temporary struct used for validating the fields of ConstDeviceApExtios.
type constDeviceApExtios  struct {
    DefaultDir *ConstDeviceApExtiosDefaultDirEnum `json:"default_dir,omitempty"`
    Input      *bool                              `json:"input,omitempty"`
    Output     *bool                              `json:"output,omitempty"`
}
