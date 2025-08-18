// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsShowDot1x represents a UtilsShowDot1x struct.
type UtilsShowDot1x struct {
	// Duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
	Duration *int `json:"duration,omitempty"`
	// Rate at which output will refresh
	Interval *int `json:"interval,omitempty"`
	// Device Port ID
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
