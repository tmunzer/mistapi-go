// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// StatsGatewayMacTableStats represents a StatsGatewayMacTableStats struct.
type StatsGatewayMacTableStats struct {
	MacTableCount          *int                   `json:"mac_table_count,omitempty"`
	MaxMacEntriesSupported *int                   `json:"max_mac_entries_supported,omitempty"`
	AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsGatewayMacTableStats,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsGatewayMacTableStats) String() string {
	return fmt.Sprintf(
		"StatsGatewayMacTableStats[MacTableCount=%v, MaxMacEntriesSupported=%v, AdditionalProperties=%v]",
		s.MacTableCount, s.MaxMacEntriesSupported, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayMacTableStats.
// It customizes the JSON marshaling process for StatsGatewayMacTableStats objects.
func (s StatsGatewayMacTableStats) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"mac_table_count", "max_mac_entries_supported"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayMacTableStats object to a map representation for JSON marshaling.
func (s StatsGatewayMacTableStats) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.MacTableCount != nil {
		structMap["mac_table_count"] = s.MacTableCount
	}
	if s.MaxMacEntriesSupported != nil {
		structMap["max_mac_entries_supported"] = s.MaxMacEntriesSupported
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayMacTableStats.
// It customizes the JSON unmarshaling process for StatsGatewayMacTableStats objects.
func (s *StatsGatewayMacTableStats) UnmarshalJSON(input []byte) error {
	var temp tempStatsGatewayMacTableStats
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac_table_count", "max_mac_entries_supported")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.MacTableCount = temp.MacTableCount
	s.MaxMacEntriesSupported = temp.MaxMacEntriesSupported
	return nil
}

// tempStatsGatewayMacTableStats is a temporary struct used for validating the fields of StatsGatewayMacTableStats.
type tempStatsGatewayMacTableStats struct {
	MacTableCount          *int `json:"mac_table_count,omitempty"`
	MaxMacEntriesSupported *int `json:"max_mac_entries_supported,omitempty"`
}
