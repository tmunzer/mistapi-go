// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsShowDot1x represents a UtilsShowDot1x struct.
// 802.1X table lookup request for device command output
type UtilsShowDot1x struct {
	// Refresh duration in seconds; set only when `interval` is nonzero
	Duration *int `json:"duration,omitempty"`
	// Refresh interval in seconds for repeated command output
	Interval *int `json:"interval,omitempty"`
	// Device port identifier filter for the 802.1X table lookup
	PortId               *string                `json:"port_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowDot1x,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowDot1x) String() string {
	return fmt.Sprintf(
		"UtilsShowDot1x[Duration=%v, Interval=%v, PortId=%v, AdditionalProperties=%v]",
		u.Duration, u.Interval, u.PortId, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowDot1x.
// It customizes the JSON marshaling process for UtilsShowDot1x objects.
func (u UtilsShowDot1x) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"duration", "interval", "port_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowDot1x object to a map representation for JSON marshaling.
func (u UtilsShowDot1x) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Duration != nil {
		structMap["duration"] = u.Duration
	}
	if u.Interval != nil {
		structMap["interval"] = u.Interval
	}
	if u.PortId != nil {
		structMap["port_id"] = u.PortId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowDot1x.
// It customizes the JSON unmarshaling process for UtilsShowDot1x objects.
func (u *UtilsShowDot1x) UnmarshalJSON(input []byte) error {
	var temp tempUtilsShowDot1x
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "interval", "port_id")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Duration = temp.Duration
	u.Interval = temp.Interval
	u.PortId = temp.PortId
	return nil
}

// tempUtilsShowDot1x is a temporary struct used for validating the fields of UtilsShowDot1x.
type tempUtilsShowDot1x struct {
	Duration *int    `json:"duration,omitempty"`
	Interval *int    `json:"interval,omitempty"`
	PortId   *string `json:"port_id,omitempty"`
}
