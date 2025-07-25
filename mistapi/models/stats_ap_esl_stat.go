// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsApEslStat represents a StatsApEslStat struct.
type StatsApEslStat struct {
    Channel              Optional[int]          `json:"channel"`
    Connected            Optional[bool]         `json:"connected"`
    Type                 Optional[string]       `json:"type"`
    Up                   Optional[bool]         `json:"up"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApEslStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApEslStat) String() string {
    return fmt.Sprintf(
    	"StatsApEslStat[Channel=%v, Connected=%v, Type=%v, Up=%v, AdditionalProperties=%v]",
    	s.Channel, s.Connected, s.Type, s.Up, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApEslStat.
// It customizes the JSON marshaling process for StatsApEslStat objects.
func (s StatsApEslStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "channel", "connected", "type", "up"); err != nil {
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "connected", "type", "up")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Channel = temp.Channel
    s.Connected = temp.Connected
    s.Type = temp.Type
    s.Up = temp.Up
    return nil
}

// tempStatsApEslStat is a temporary struct used for validating the fields of StatsApEslStat.
type tempStatsApEslStat  struct {
    Channel   Optional[int]    `json:"channel"`
    Connected Optional[bool]   `json:"connected"`
    Type      Optional[string] `json:"type"`
    Up        Optional[bool]   `json:"up"`
}
