// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ConstDeviceApBand6 represents a ConstDeviceApBand6 struct.
// 6 GHz radio capability limits for an AP model
type ConstDeviceApBand6 struct {
	// Maximum client count supported on the 6 GHz radio
	MaxClients *int `json:"max_clients,omitempty"`
	// Maximum transmit power for the 6 GHz radio, in dBm
	MaxPower *int `json:"max_power,omitempty"`
	// Minimum transmit power for the 6 GHz radio, in dBm
	MinPower             *int                   `json:"min_power,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstDeviceApBand6,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstDeviceApBand6) String() string {
	return fmt.Sprintf(
		"ConstDeviceApBand6[MaxClients=%v, MaxPower=%v, MinPower=%v, AdditionalProperties=%v]",
		c.MaxClients, c.MaxPower, c.MinPower, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceApBand6.
// It customizes the JSON marshaling process for ConstDeviceApBand6 objects.
func (c ConstDeviceApBand6) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"max_clients", "max_power", "min_power"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceApBand6 object to a map representation for JSON marshaling.
func (c ConstDeviceApBand6) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceApBand6.
// It customizes the JSON unmarshaling process for ConstDeviceApBand6 objects.
func (c *ConstDeviceApBand6) UnmarshalJSON(input []byte) error {
	var temp tempConstDeviceApBand6
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

// tempConstDeviceApBand6 is a temporary struct used for validating the fields of ConstDeviceApBand6.
type tempConstDeviceApBand6 struct {
	MaxClients *int `json:"max_clients,omitempty"`
	MaxPower   *int `json:"max_power,omitempty"`
	MinPower   *int `json:"min_power,omitempty"`
}
