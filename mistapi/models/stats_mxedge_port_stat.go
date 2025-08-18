// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsMxedgePortStat represents a StatsMxedgePortStat struct.
type StatsMxedgePortStat struct {
	FullDuplex *bool   `json:"full_duplex,omitempty"`
	Mac        *string `json:"mac,omitempty"`
	// Amount of traffic received since connection
	RxBytes  Optional[int64] `json:"rx_bytes"`
	RxErrors *int            `json:"rx_errors,omitempty"`
	// Amount of packets received since connection
	RxPkts Optional[int64] `json:"rx_pkts"`
	Speed  *int            `json:"speed,omitempty"`
	State  *string         `json:"state,omitempty"`
	// Amount of traffic sent since connection
	TxBytes  Optional[int64] `json:"tx_bytes"`
	TxErrors *int            `json:"tx_errors,omitempty"`
	// Amount of packets sent since connection
	TxPkts               Optional[int64]        `json:"tx_pkts"`
	Up                   *bool                  `json:"up,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgePortStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgePortStat) String() string {
	return fmt.Sprintf(
		"StatsMxedgePortStat[FullDuplex=%v, Mac=%v, RxBytes=%v, RxErrors=%v, RxPkts=%v, Speed=%v, State=%v, TxBytes=%v, TxErrors=%v, TxPkts=%v, Up=%v, AdditionalProperties=%v]",
		s.FullDuplex, s.Mac, s.RxBytes, s.RxErrors, s.RxPkts, s.Speed, s.State, s.TxBytes, s.TxErrors, s.TxPkts, s.Up, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgePortStat.
// It customizes the JSON marshaling process for StatsMxedgePortStat objects.
func (s StatsMxedgePortStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"full_duplex", "mac", "rx_bytes", "rx_errors", "rx_pkts", "speed", "state", "tx_bytes", "tx_errors", "tx_pkts", "up"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgePortStat object to a map representation for JSON marshaling.
func (s StatsMxedgePortStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.FullDuplex != nil {
		structMap["full_duplex"] = s.FullDuplex
	}
	if s.Mac != nil {
		structMap["mac"] = s.Mac
	}
	if s.RxBytes.IsValueSet() {
		if s.RxBytes.Value() != nil {
			structMap["rx_bytes"] = s.RxBytes.Value()
		} else {
			structMap["rx_bytes"] = nil
		}
	}
	if s.RxErrors != nil {
		structMap["rx_errors"] = s.RxErrors
	}
	if s.RxPkts.IsValueSet() {
		if s.RxPkts.Value() != nil {
			structMap["rx_pkts"] = s.RxPkts.Value()
		} else {
			structMap["rx_pkts"] = nil
		}
	}
	if s.Speed != nil {
		structMap["speed"] = s.Speed
	}
	if s.State != nil {
		structMap["state"] = s.State
	}
	if s.TxBytes.IsValueSet() {
		if s.TxBytes.Value() != nil {
			structMap["tx_bytes"] = s.TxBytes.Value()
		} else {
			structMap["tx_bytes"] = nil
		}
	}
	if s.TxErrors != nil {
		structMap["tx_errors"] = s.TxErrors
	}
	if s.TxPkts.IsValueSet() {
		if s.TxPkts.Value() != nil {
			structMap["tx_pkts"] = s.TxPkts.Value()
		} else {
			structMap["tx_pkts"] = nil
		}
	}
	if s.Up != nil {
		structMap["up"] = s.Up
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgePortStat.
// It customizes the JSON unmarshaling process for StatsMxedgePortStat objects.
func (s *StatsMxedgePortStat) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgePortStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "full_duplex", "mac", "rx_bytes", "rx_errors", "rx_pkts", "speed", "state", "tx_bytes", "tx_errors", "tx_pkts", "up")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.FullDuplex = temp.FullDuplex
	s.Mac = temp.Mac
	s.RxBytes = temp.RxBytes
	s.RxErrors = temp.RxErrors
	s.RxPkts = temp.RxPkts
	s.Speed = temp.Speed
	s.State = temp.State
	s.TxBytes = temp.TxBytes
	s.TxErrors = temp.TxErrors
	s.TxPkts = temp.TxPkts
	s.Up = temp.Up
	return nil
}

// tempStatsMxedgePortStat is a temporary struct used for validating the fields of StatsMxedgePortStat.
type tempStatsMxedgePortStat struct {
	FullDuplex *bool           `json:"full_duplex,omitempty"`
	Mac        *string         `json:"mac,omitempty"`
	RxBytes    Optional[int64] `json:"rx_bytes"`
	RxErrors   *int            `json:"rx_errors,omitempty"`
	RxPkts     Optional[int64] `json:"rx_pkts"`
	Speed      *int            `json:"speed,omitempty"`
	State      *string         `json:"state,omitempty"`
	TxBytes    Optional[int64] `json:"tx_bytes"`
	TxErrors   *int            `json:"tx_errors,omitempty"`
	TxPkts     Optional[int64] `json:"tx_pkts"`
	Up         *bool           `json:"up,omitempty"`
}
