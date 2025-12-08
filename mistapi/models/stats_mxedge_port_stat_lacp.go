// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsMxedgePortStatLacp represents a StatsMxedgePortStatLacp struct.
type StatsMxedgePortStatLacp struct {
	MuxState             *string                `json:"mux_state,omitempty"`
	RxLacpdu             *int                   `json:"rx_lacpdu,omitempty"`
	RxState              *string                `json:"rx_state,omitempty"`
	TxLacpdu             *int                   `json:"tx_lacpdu,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedgePortStatLacp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedgePortStatLacp) String() string {
	return fmt.Sprintf(
		"StatsMxedgePortStatLacp[MuxState=%v, RxLacpdu=%v, RxState=%v, TxLacpdu=%v, AdditionalProperties=%v]",
		s.MuxState, s.RxLacpdu, s.RxState, s.TxLacpdu, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedgePortStatLacp.
// It customizes the JSON marshaling process for StatsMxedgePortStatLacp objects.
func (s StatsMxedgePortStatLacp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"mux_state", "rx_lacpdu", "rx_state", "tx_lacpdu"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedgePortStatLacp object to a map representation for JSON marshaling.
func (s StatsMxedgePortStatLacp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.MuxState != nil {
		structMap["mux_state"] = s.MuxState
	}
	if s.RxLacpdu != nil {
		structMap["rx_lacpdu"] = s.RxLacpdu
	}
	if s.RxState != nil {
		structMap["rx_state"] = s.RxState
	}
	if s.TxLacpdu != nil {
		structMap["tx_lacpdu"] = s.TxLacpdu
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedgePortStatLacp.
// It customizes the JSON unmarshaling process for StatsMxedgePortStatLacp objects.
func (s *StatsMxedgePortStatLacp) UnmarshalJSON(input []byte) error {
	var temp tempStatsMxedgePortStatLacp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mux_state", "rx_lacpdu", "rx_state", "tx_lacpdu")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.MuxState = temp.MuxState
	s.RxLacpdu = temp.RxLacpdu
	s.RxState = temp.RxState
	s.TxLacpdu = temp.TxLacpdu
	return nil
}

// tempStatsMxedgePortStatLacp is a temporary struct used for validating the fields of StatsMxedgePortStatLacp.
type tempStatsMxedgePortStatLacp struct {
	MuxState *string `json:"mux_state,omitempty"`
	RxLacpdu *int    `json:"rx_lacpdu,omitempty"`
	RxState  *string `json:"rx_state,omitempty"`
	TxLacpdu *int    `json:"tx_lacpdu,omitempty"`
}
