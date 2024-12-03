package models

import (
    "encoding/json"
)

// ConstDeviceApBand5 represents a ConstDeviceApBand5 struct.
type ConstDeviceApBand5 struct {
    MaxClients           *int                   `json:"max_clients,omitempty"`
    MaxPower             *int                   `json:"max_power,omitempty"`
    MinPower             *int                   `json:"min_power,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceApBand5.
// It customizes the JSON marshaling process for ConstDeviceApBand5 objects.
func (c ConstDeviceApBand5) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "max_clients", "max_power", "min_power"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceApBand5 object to a map representation for JSON marshaling.
func (c ConstDeviceApBand5) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.MaxClients != nil {
        structMap["max_clients"] = c.MaxClients
    }
    if c.MaxPower != nil {
        structMap["max_power"] = c.MaxPower
    }
    if c.MinPower != nil {
        structMap["min_power"] = c.MinPower
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceApBand5.
// It customizes the JSON unmarshaling process for ConstDeviceApBand5 objects.
func (c *ConstDeviceApBand5) UnmarshalJSON(input []byte) error {
    var temp tempConstDeviceApBand5
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "max_clients", "max_power", "min_power")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.MaxClients = temp.MaxClients
    c.MaxPower = temp.MaxPower
    c.MinPower = temp.MinPower
    return nil
}

// tempConstDeviceApBand5 is a temporary struct used for validating the fields of ConstDeviceApBand5.
type tempConstDeviceApBand5  struct {
    MaxClients *int `json:"max_clients,omitempty"`
    MaxPower   *int `json:"max_power,omitempty"`
    MinPower   *int `json:"min_power,omitempty"`
}
