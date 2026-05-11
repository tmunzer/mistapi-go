// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsZigbeeJoin represents a UtilsZigbeeJoin struct.
type UtilsZigbeeJoin struct {
	// Duration in seconds for which new Zigbee end device joins are permitted. Range is 30–3600
	Duration             *int                   `json:"duration,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsZigbeeJoin,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsZigbeeJoin) String() string {
	return fmt.Sprintf(
		"UtilsZigbeeJoin[Duration=%v, AdditionalProperties=%v]",
		u.Duration, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsZigbeeJoin.
// It customizes the JSON marshaling process for UtilsZigbeeJoin objects.
func (u UtilsZigbeeJoin) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"duration"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsZigbeeJoin object to a map representation for JSON marshaling.
func (u UtilsZigbeeJoin) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Duration != nil {
		structMap["duration"] = u.Duration
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsZigbeeJoin.
// It customizes the JSON unmarshaling process for UtilsZigbeeJoin objects.
func (u *UtilsZigbeeJoin) UnmarshalJSON(input []byte) error {
	var temp tempUtilsZigbeeJoin
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Duration = temp.Duration
	return nil
}

// tempUtilsZigbeeJoin is a temporary struct used for validating the fields of UtilsZigbeeJoin.
type tempUtilsZigbeeJoin struct {
	Duration *int `json:"duration,omitempty"`
}
