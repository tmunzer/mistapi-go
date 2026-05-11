// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsAssetServicePacket represents a StatsAssetServicePacket struct.
// Service data advertisement from a BLE asset
type StatsAssetServicePacket struct {
	// Service data payload (hex encoded)
	Data *string `json:"data,omitempty"`
	// Unix timestamp when this service data was last received
	LastRxTime *int `json:"last_rx_time,omitempty"`
	// Total number of times this service data was received
	RxCnt *int `json:"rx_cnt,omitempty"`
	// Service UUID
	Uuid                 *string                `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsAssetServicePacket,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsAssetServicePacket) String() string {
	return fmt.Sprintf(
		"StatsAssetServicePacket[Data=%v, LastRxTime=%v, RxCnt=%v, Uuid=%v, AdditionalProperties=%v]",
		s.Data, s.LastRxTime, s.RxCnt, s.Uuid, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsAssetServicePacket.
// It customizes the JSON marshaling process for StatsAssetServicePacket objects.
func (s StatsAssetServicePacket) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"data", "last_rx_time", "rx_cnt", "uuid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsAssetServicePacket object to a map representation for JSON marshaling.
func (s StatsAssetServicePacket) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Data != nil {
		structMap["data"] = s.Data
	}
	if s.LastRxTime != nil {
		structMap["last_rx_time"] = s.LastRxTime
	}
	if s.RxCnt != nil {
		structMap["rx_cnt"] = s.RxCnt
	}
	if s.Uuid != nil {
		structMap["uuid"] = s.Uuid
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsAssetServicePacket.
// It customizes the JSON unmarshaling process for StatsAssetServicePacket objects.
func (s *StatsAssetServicePacket) UnmarshalJSON(input []byte) error {
	var temp tempStatsAssetServicePacket
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "data", "last_rx_time", "rx_cnt", "uuid")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Data = temp.Data
	s.LastRxTime = temp.LastRxTime
	s.RxCnt = temp.RxCnt
	s.Uuid = temp.Uuid
	return nil
}

// tempStatsAssetServicePacket is a temporary struct used for validating the fields of StatsAssetServicePacket.
type tempStatsAssetServicePacket struct {
	Data       *string `json:"data,omitempty"`
	LastRxTime *int    `json:"last_rx_time,omitempty"`
	RxCnt      *int    `json:"rx_cnt,omitempty"`
	Uuid       *string `json:"uuid,omitempty"`
}
