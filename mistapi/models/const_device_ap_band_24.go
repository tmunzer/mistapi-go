package models

import (
    "encoding/json"
)

// ConstDeviceApBand24 represents a ConstDeviceApBand24 struct.
type ConstDeviceApBand24 struct {
    Band5ChannelsOp      *string                `json:"band5_channels_op,omitempty"`
    MaxClients           *int                   `json:"max_clients,omitempty"`
    MaxPower             *int                   `json:"max_power,omitempty"`
    MinPower             *int                   `json:"min_power,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceApBand24.
// It customizes the JSON marshaling process for ConstDeviceApBand24 objects.
func (c ConstDeviceApBand24) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "band5_channels_op", "max_clients", "max_power", "min_power"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceApBand24 object to a map representation for JSON marshaling.
func (c ConstDeviceApBand24) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Band5ChannelsOp != nil {
        structMap["band5_channels_op"] = c.Band5ChannelsOp
    }
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

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceApBand24.
// It customizes the JSON unmarshaling process for ConstDeviceApBand24 objects.
func (c *ConstDeviceApBand24) UnmarshalJSON(input []byte) error {
    var temp tempConstDeviceApBand24
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band5_channels_op", "max_clients", "max_power", "min_power")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Band5ChannelsOp = temp.Band5ChannelsOp
    c.MaxClients = temp.MaxClients
    c.MaxPower = temp.MaxPower
    c.MinPower = temp.MinPower
    return nil
}

// tempConstDeviceApBand24 is a temporary struct used for validating the fields of ConstDeviceApBand24.
type tempConstDeviceApBand24  struct {
    Band5ChannelsOp *string `json:"band5_channels_op,omitempty"`
    MaxClients      *int    `json:"max_clients,omitempty"`
    MaxPower        *int    `json:"max_power,omitempty"`
    MinPower        *int    `json:"min_power,omitempty"`
}
