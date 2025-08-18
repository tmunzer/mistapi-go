// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// JseDevice represents a JseDevice struct.
type JseDevice struct {
	// When available
	ExtIp *string `json:"ext_ip,omitempty"`
	// Last seen timestamp
	LastSeen             Optional[float64]      `json:"last_seen"`
	Mac                  *string                `json:"mac,omitempty"`
	Model                *string                `json:"model,omitempty"`
	Serial               *string                `json:"serial,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JseDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JseDevice) String() string {
	return fmt.Sprintf(
		"JseDevice[ExtIp=%v, LastSeen=%v, Mac=%v, Model=%v, Serial=%v, AdditionalProperties=%v]",
		j.ExtIp, j.LastSeen, j.Mac, j.Model, j.Serial, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JseDevice.
// It customizes the JSON marshaling process for JseDevice objects.
func (j JseDevice) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"ext_ip", "last_seen", "mac", "model", "serial"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JseDevice object to a map representation for JSON marshaling.
func (j JseDevice) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.ExtIp != nil {
		structMap["ext_ip"] = j.ExtIp
	}
	if j.LastSeen.IsValueSet() {
		if j.LastSeen.Value() != nil {
			structMap["last_seen"] = j.LastSeen.Value()
		} else {
			structMap["last_seen"] = nil
		}
	}
	if j.Mac != nil {
		structMap["mac"] = j.Mac
	}
	if j.Model != nil {
		structMap["model"] = j.Model
	}
	if j.Serial != nil {
		structMap["serial"] = j.Serial
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JseDevice.
// It customizes the JSON unmarshaling process for JseDevice objects.
func (j *JseDevice) UnmarshalJSON(input []byte) error {
	var temp tempJseDevice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ext_ip", "last_seen", "mac", "model", "serial")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.ExtIp = temp.ExtIp
	j.LastSeen = temp.LastSeen
	j.Mac = temp.Mac
	j.Model = temp.Model
	j.Serial = temp.Serial
	return nil
}

// tempJseDevice is a temporary struct used for validating the fields of JseDevice.
type tempJseDevice struct {
	ExtIp    *string           `json:"ext_ip,omitempty"`
	LastSeen Optional[float64] `json:"last_seen"`
	Mac      *string           `json:"mac,omitempty"`
	Model    *string           `json:"model,omitempty"`
	Serial   *string           `json:"serial,omitempty"`
}
