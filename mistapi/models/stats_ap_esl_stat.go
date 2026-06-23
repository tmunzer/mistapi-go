// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsApEslStat represents a StatsApEslStat struct.
// Electronic shelf label dongle status reported by an AP
type StatsApEslStat struct {
	// Radio channel used by the ESL dongle
	Channel Optional[int] `json:"channel"`
	// Whether the ESL dongle is connected
	Connected Optional[bool] `json:"connected"`
	// Network IP address of Hanshow and SoluM dongles
	Ip Optional[string] `json:"ip"`
	// Dongle MAC address for Hanshow and SoluM dongles
	Mac Optional[string] `json:"mac"`
	// Product ID of Hanshow and SoluM dongles
	ProductId Optional[string] `json:"product_id"`
	// ESL dongle type reported by the AP
	Type Optional[string] `json:"type"`
	// Whether the ESL dongle is operational
	Up Optional[bool] `json:"up"`
	// Vendor ID of Hanshow and SoluM dongles
	VendorId             Optional[string]       `json:"vendor_id"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApEslStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApEslStat) String() string {
	return fmt.Sprintf(
		"StatsApEslStat[Channel=%v, Connected=%v, Ip=%v, Mac=%v, ProductId=%v, Type=%v, Up=%v, VendorId=%v, AdditionalProperties=%v]",
		s.Channel, s.Connected, s.Ip, s.Mac, s.ProductId, s.Type, s.Up, s.VendorId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApEslStat.
// It customizes the JSON marshaling process for StatsApEslStat objects.
func (s StatsApEslStat) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"channel", "connected", "ip", "mac", "product_id", "type", "up", "vendor_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApEslStat object to a map representation for JSON marshaling.
func (s StatsApEslStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Channel.IsValueSet() {
		if s.Channel.Value() != nil {
			structMap["channel"] = s.Channel.Value()
		} else {
			structMap["channel"] = nil
		}
	}
	if s.Connected.IsValueSet() {
		if s.Connected.Value() != nil {
			structMap["connected"] = s.Connected.Value()
		} else {
			structMap["connected"] = nil
		}
	}
	if s.Ip.IsValueSet() {
		if s.Ip.Value() != nil {
			structMap["ip"] = s.Ip.Value()
		} else {
			structMap["ip"] = nil
		}
	}
	if s.Mac.IsValueSet() {
		if s.Mac.Value() != nil {
			structMap["mac"] = s.Mac.Value()
		} else {
			structMap["mac"] = nil
		}
	}
	if s.ProductId.IsValueSet() {
		if s.ProductId.Value() != nil {
			structMap["product_id"] = s.ProductId.Value()
		} else {
			structMap["product_id"] = nil
		}
	}
	if s.Type.IsValueSet() {
		if s.Type.Value() != nil {
			structMap["type"] = s.Type.Value()
		} else {
			structMap["type"] = nil
		}
	}
	if s.Up.IsValueSet() {
		if s.Up.Value() != nil {
			structMap["up"] = s.Up.Value()
		} else {
			structMap["up"] = nil
		}
	}
	if s.VendorId.IsValueSet() {
		if s.VendorId.Value() != nil {
			structMap["vendor_id"] = s.VendorId.Value()
		} else {
			structMap["vendor_id"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApEslStat.
// It customizes the JSON unmarshaling process for StatsApEslStat objects.
func (s *StatsApEslStat) UnmarshalJSON(input []byte) error {
	var temp tempStatsApEslStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "connected", "ip", "mac", "product_id", "type", "up", "vendor_id")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Channel = temp.Channel
	s.Connected = temp.Connected
	s.Ip = temp.Ip
	s.Mac = temp.Mac
	s.ProductId = temp.ProductId
	s.Type = temp.Type
	s.Up = temp.Up
	s.VendorId = temp.VendorId
	return nil
}

// tempStatsApEslStat is a temporary struct used for validating the fields of StatsApEslStat.
type tempStatsApEslStat struct {
	Channel   Optional[int]    `json:"channel"`
	Connected Optional[bool]   `json:"connected"`
	Ip        Optional[string] `json:"ip"`
	Mac       Optional[string] `json:"mac"`
	ProductId Optional[string] `json:"product_id"`
	Type      Optional[string] `json:"type"`
	Up        Optional[bool]   `json:"up"`
	VendorId  Optional[string] `json:"vendor_id"`
}
