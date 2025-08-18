// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApRadioConfigBand represents a StatsApRadioConfigBand struct.
type StatsApRadioConfigBand struct {
	AllowRrmDisable        Optional[bool]         `json:"allow_rrm_disable"`
	Bandwidth              Optional[float64]      `json:"bandwidth"`
	Channel                *int                   `json:"channel,omitempty"`
	Disabled               Optional[bool]         `json:"disabled"`
	DynamicChainingEnabled Optional[bool]         `json:"dynamic_chaining_enabled"`
	Power                  Optional[float64]      `json:"power"`
	PowerMax               Optional[float64]      `json:"power_max"`
	PowerMin               Optional[float64]      `json:"power_min"`
	RxChain                Optional[int]          `json:"rx_chain"`
	TxChain                Optional[int]          `json:"tx_chain"`
	AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApRadioConfigBand,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApRadioConfigBand) String() string {
	return fmt.Sprintf(
		"StatsApRadioConfigBand[AllowRrmDisable=%v, Bandwidth=%v, Channel=%v, Disabled=%v, DynamicChainingEnabled=%v, Power=%v, PowerMax=%v, PowerMin=%v, RxChain=%v, TxChain=%v, AdditionalProperties=%v]",
		s.AllowRrmDisable, s.Bandwidth, s.Channel, s.Disabled, s.DynamicChainingEnabled, s.Power, s.PowerMax, s.PowerMin, s.RxChain, s.TxChain, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApRadioConfigBand.
// It customizes the JSON marshaling process for StatsApRadioConfigBand objects.
func (s StatsApRadioConfigBand) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"allow_rrm_disable", "bandwidth", "channel", "disabled", "dynamic_chaining_enabled", "power", "power_max", "power_min", "rx_chain", "tx_chain"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApRadioConfigBand object to a map representation for JSON marshaling.
func (s StatsApRadioConfigBand) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AllowRrmDisable.IsValueSet() {
		if s.AllowRrmDisable.Value() != nil {
			structMap["allow_rrm_disable"] = s.AllowRrmDisable.Value()
		} else {
			structMap["allow_rrm_disable"] = nil
		}
	}
	if s.Bandwidth.IsValueSet() {
		if s.Bandwidth.Value() != nil {
			structMap["bandwidth"] = s.Bandwidth.Value()
		} else {
			structMap["bandwidth"] = nil
		}
	}
	if s.Channel != nil {
		structMap["channel"] = s.Channel
	}
	if s.Disabled.IsValueSet() {
		if s.Disabled.Value() != nil {
			structMap["disabled"] = s.Disabled.Value()
		} else {
			structMap["disabled"] = nil
		}
	}
	if s.DynamicChainingEnabled.IsValueSet() {
		if s.DynamicChainingEnabled.Value() != nil {
			structMap["dynamic_chaining_enabled"] = s.DynamicChainingEnabled.Value()
		} else {
			structMap["dynamic_chaining_enabled"] = nil
		}
	}
	if s.Power.IsValueSet() {
		if s.Power.Value() != nil {
			structMap["power"] = s.Power.Value()
		} else {
			structMap["power"] = nil
		}
	}
	if s.PowerMax.IsValueSet() {
		if s.PowerMax.Value() != nil {
			structMap["power_max"] = s.PowerMax.Value()
		} else {
			structMap["power_max"] = nil
		}
	}
	if s.PowerMin.IsValueSet() {
		if s.PowerMin.Value() != nil {
			structMap["power_min"] = s.PowerMin.Value()
		} else {
			structMap["power_min"] = nil
		}
	}
	if s.RxChain.IsValueSet() {
		if s.RxChain.Value() != nil {
			structMap["rx_chain"] = s.RxChain.Value()
		} else {
			structMap["rx_chain"] = nil
		}
	}
	if s.TxChain.IsValueSet() {
		if s.TxChain.Value() != nil {
			structMap["tx_chain"] = s.TxChain.Value()
		} else {
			structMap["tx_chain"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApRadioConfigBand.
// It customizes the JSON unmarshaling process for StatsApRadioConfigBand objects.
func (s *StatsApRadioConfigBand) UnmarshalJSON(input []byte) error {
	var temp tempStatsApRadioConfigBand
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allow_rrm_disable", "bandwidth", "channel", "disabled", "dynamic_chaining_enabled", "power", "power_max", "power_min", "rx_chain", "tx_chain")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AllowRrmDisable = temp.AllowRrmDisable
	s.Bandwidth = temp.Bandwidth
	s.Channel = temp.Channel
	s.Disabled = temp.Disabled
	s.DynamicChainingEnabled = temp.DynamicChainingEnabled
	s.Power = temp.Power
	s.PowerMax = temp.PowerMax
	s.PowerMin = temp.PowerMin
	s.RxChain = temp.RxChain
	s.TxChain = temp.TxChain
	return nil
}

// tempStatsApRadioConfigBand is a temporary struct used for validating the fields of StatsApRadioConfigBand.
type tempStatsApRadioConfigBand struct {
	AllowRrmDisable        Optional[bool]    `json:"allow_rrm_disable"`
	Bandwidth              Optional[float64] `json:"bandwidth"`
	Channel                *int              `json:"channel,omitempty"`
	Disabled               Optional[bool]    `json:"disabled"`
	DynamicChainingEnabled Optional[bool]    `json:"dynamic_chaining_enabled"`
	Power                  Optional[float64] `json:"power"`
	PowerMax               Optional[float64] `json:"power_max"`
	PowerMin               Optional[float64] `json:"power_min"`
	RxChain                Optional[int]     `json:"rx_chain"`
	TxChain                Optional[int]     `json:"tx_chain"`
}
