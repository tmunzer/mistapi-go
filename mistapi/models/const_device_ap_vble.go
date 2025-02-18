package models

import (
    "encoding/json"
    "fmt"
)

// ConstDeviceApVble represents a ConstDeviceApVble struct.
type ConstDeviceApVble struct {
    BeaconRate           *int                   `json:"beacon_rate,omitempty"`
    Beams                *int                   `json:"beams,omitempty"`
    Power                *int                   `json:"power,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstDeviceApVble,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstDeviceApVble) String() string {
    return fmt.Sprintf(
    	"ConstDeviceApVble[BeaconRate=%v, Beams=%v, Power=%v, AdditionalProperties=%v]",
    	c.BeaconRate, c.Beams, c.Power, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceApVble.
// It customizes the JSON marshaling process for ConstDeviceApVble objects.
func (c ConstDeviceApVble) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "beacon_rate", "beams", "power"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceApVble object to a map representation for JSON marshaling.
func (c ConstDeviceApVble) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp tempConstDeviceApVble
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "beacon_rate", "beams", "power")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.BeaconRate = temp.BeaconRate
    c.Beams = temp.Beams
    c.Power = temp.Power
    return nil
}

// tempConstDeviceApVble is a temporary struct used for validating the fields of ConstDeviceApVble.
type tempConstDeviceApVble  struct {
    BeaconRate *int `json:"beacon_rate,omitempty"`
    Beams      *int `json:"beams,omitempty"`
    Power      *int `json:"power,omitempty"`
}
