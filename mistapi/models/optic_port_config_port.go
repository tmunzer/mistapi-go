// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OpticPortConfigPort represents a OpticPortConfigPort struct.
type OpticPortConfigPort struct {
	// Enable channelization
	Channelized *bool `json:"channelized,omitempty"`
	// Interface speed (e.g. `25g`, `50g`), use the chassis speed by default
	Speed                *string                `json:"speed,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OpticPortConfigPort,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OpticPortConfigPort) String() string {
	return fmt.Sprintf(
		"OpticPortConfigPort[Channelized=%v, Speed=%v, AdditionalProperties=%v]",
		o.Channelized, o.Speed, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OpticPortConfigPort.
// It customizes the JSON marshaling process for OpticPortConfigPort objects.
func (o OpticPortConfigPort) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"channelized", "speed"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OpticPortConfigPort object to a map representation for JSON marshaling.
func (o OpticPortConfigPort) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Channelized != nil {
		structMap["channelized"] = o.Channelized
	}
	if o.Speed != nil {
		structMap["speed"] = o.Speed
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OpticPortConfigPort.
// It customizes the JSON unmarshaling process for OpticPortConfigPort objects.
func (o *OpticPortConfigPort) UnmarshalJSON(input []byte) error {
	var temp tempOpticPortConfigPort
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channelized", "speed")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Channelized = temp.Channelized
	o.Speed = temp.Speed
	return nil
}

// tempOpticPortConfigPort is a temporary struct used for validating the fields of OpticPortConfigPort.
type tempOpticPortConfigPort struct {
	Channelized *bool   `json:"channelized,omitempty"`
	Speed       *string `json:"speed,omitempty"`
}
