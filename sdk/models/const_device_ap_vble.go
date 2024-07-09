package models

import (
    "encoding/json"
)

// ConstDeviceApVble represents a ConstDeviceApVble struct.
type ConstDeviceApVble struct {
    BeaconRate           *int           `json:"beacon_rate,omitempty"`
    Beams                *int           `json:"beams,omitempty"`
    Power                *int           `json:"power,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceApVble.
// It customizes the JSON marshaling process for ConstDeviceApVble objects.
func (c ConstDeviceApVble) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceApVble object to a map representation for JSON marshaling.
func (c ConstDeviceApVble) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.BeaconRate != nil {
        structMap["beacon_rate"] = c.BeaconRate
    }
    if c.Beams != nil {
        structMap["beams"] = c.Beams
    }
    if c.Power != nil {
        structMap["power"] = c.Power
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceApVble.
// It customizes the JSON unmarshaling process for ConstDeviceApVble objects.
func (c *ConstDeviceApVble) UnmarshalJSON(input []byte) error {
    var temp constDeviceApVble
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "beacon_rate", "beams", "power")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.BeaconRate = temp.BeaconRate
    c.Beams = temp.Beams
    c.Power = temp.Power
    return nil
}

// constDeviceApVble is a temporary struct used for validating the fields of ConstDeviceApVble.
type constDeviceApVble  struct {
    BeaconRate *int `json:"beacon_rate,omitempty"`
    Beams      *int `json:"beams,omitempty"`
    Power      *int `json:"power,omitempty"`
}
