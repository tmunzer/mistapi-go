// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// CaptureSwitchSwitches represents a CaptureSwitchSwitches struct.
type CaptureSwitchSwitches struct {
	// Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request
	Ports                map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{}                         `json:"_"`
}

// String implements the fmt.Stringer interface for CaptureSwitchSwitches,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CaptureSwitchSwitches) String() string {
	return fmt.Sprintf(
		"CaptureSwitchSwitches[Ports=%v, AdditionalProperties=%v]",
		c.Ports, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CaptureSwitchSwitches.
// It customizes the JSON marshaling process for CaptureSwitchSwitches objects.
func (c CaptureSwitchSwitches) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"ports"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureSwitchSwitches object to a map representation for JSON marshaling.
func (c CaptureSwitchSwitches) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Ports != nil {
		structMap["ports"] = c.Ports
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureSwitchSwitches.
// It customizes the JSON unmarshaling process for CaptureSwitchSwitches objects.
func (c *CaptureSwitchSwitches) UnmarshalJSON(input []byte) error {
	var temp tempCaptureSwitchSwitches
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ports")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Ports = temp.Ports
	return nil
}

// tempCaptureSwitchSwitches is a temporary struct used for validating the fields of CaptureSwitchSwitches.
type tempCaptureSwitchSwitches struct {
	Ports map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
}
