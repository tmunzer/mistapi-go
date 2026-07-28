// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApZigbeeStat represents a StatsApZigbeeStat struct.
// ZigBee statistics reported by an AP, present only when ZigBee is enabled on the AP
type StatsApZigbeeStat struct {
	// Connection status of the IoT proxy
	IotproxyStatus Optional[string] `json:"iotproxy_status"`
	// Number of IoT endpoints connected through the AP
	NumIotendpoints      Optional[int]          `json:"num_iotendpoints"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApZigbeeStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApZigbeeStat) String() string {
	return fmt.Sprintf(
		"StatsApZigbeeStat[IotproxyStatus=%v, NumIotendpoints=%v, AdditionalProperties=%v]",
		s.IotproxyStatus, s.NumIotendpoints, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApZigbeeStat.
// It customizes the JSON marshaling process for StatsApZigbeeStat objects.
func (s StatsApZigbeeStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"iotproxy_status", "num_iotendpoints"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApZigbeeStat object to a map representation for JSON marshaling.
func (s StatsApZigbeeStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.IotproxyStatus.IsValueSet() {
		if s.IotproxyStatus.Value() != nil {
			structMap["iotproxy_status"] = s.IotproxyStatus.Value()
		} else {
			structMap["iotproxy_status"] = nil
		}
	}
	if s.NumIotendpoints.IsValueSet() {
		if s.NumIotendpoints.Value() != nil {
			structMap["num_iotendpoints"] = s.NumIotendpoints.Value()
		} else {
			structMap["num_iotendpoints"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApZigbeeStat.
// It customizes the JSON unmarshaling process for StatsApZigbeeStat objects.
func (s *StatsApZigbeeStat) UnmarshalJSON(input []byte) error {
	var temp tempStatsApZigbeeStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "iotproxy_status", "num_iotendpoints")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.IotproxyStatus = temp.IotproxyStatus
	s.NumIotendpoints = temp.NumIotendpoints
	return nil
}

// tempStatsApZigbeeStat is a temporary struct used for validating the fields of StatsApZigbeeStat.
type tempStatsApZigbeeStat struct {
	IotproxyStatus  Optional[string] `json:"iotproxy_status"`
	NumIotendpoints Optional[int]    `json:"num_iotendpoints"`
}
