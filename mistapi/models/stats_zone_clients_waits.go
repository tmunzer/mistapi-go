// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsZoneClientsWaits represents a StatsZoneClientsWaits struct.
// Client wait time right now
type StatsZoneClientsWaits struct {
    // Average wait time in seconds
    Avg                  *float64               `json:"avg,omitempty"`
    // Longest wait time in seconds
    Max                  *float64               `json:"max,omitempty"`
    // Shortest wait time in seconds
    Min                  *float64               `json:"min,omitempty"`
    // 95th percentile of all the wait time(s)
    P95                  *float64               `json:"p95,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsZoneClientsWaits,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsZoneClientsWaits) String() string {
    return fmt.Sprintf(
    	"StatsZoneClientsWaits[Avg=%v, Max=%v, Min=%v, P95=%v, AdditionalProperties=%v]",
    	s.Avg, s.Max, s.Min, s.P95, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsZoneClientsWaits.
// It customizes the JSON marshaling process for StatsZoneClientsWaits objects.
func (s StatsZoneClientsWaits) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "avg", "max", "min", "p95"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsZoneClientsWaits object to a map representation for JSON marshaling.
func (s StatsZoneClientsWaits) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Avg != nil {
        structMap["avg"] = s.Avg
    }
    if s.Max != nil {
        structMap["max"] = s.Max
    }
    if s.Min != nil {
        structMap["min"] = s.Min
    }
    if s.P95 != nil {
        structMap["p95"] = s.P95
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsZoneClientsWaits.
// It customizes the JSON unmarshaling process for StatsZoneClientsWaits objects.
func (s *StatsZoneClientsWaits) UnmarshalJSON(input []byte) error {
    var temp tempStatsZoneClientsWaits
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "avg", "max", "min", "p95")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Avg = temp.Avg
    s.Max = temp.Max
    s.Min = temp.Min
    s.P95 = temp.P95
    return nil
}

// tempStatsZoneClientsWaits is a temporary struct used for validating the fields of StatsZoneClientsWaits.
type tempStatsZoneClientsWaits  struct {
    Avg *float64 `json:"avg,omitempty"`
    Max *float64 `json:"max,omitempty"`
    Min *float64 `json:"min,omitempty"`
    P95 *float64 `json:"p95,omitempty"`
}
