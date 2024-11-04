package models

import (
    "encoding/json"
)

// StatsZoneClientsWaits represents a StatsZoneClientsWaits struct.
// client wait time right now
type StatsZoneClientsWaits struct {
    // average wait time in seconds
    Avg                  *float64       `json:"avg,omitempty"`
    // longest wait time in seconds
    Max                  *float64       `json:"max,omitempty"`
    // shortest wait time in seconds
    Min                  *float64       `json:"min,omitempty"`
    // 95th percentile of all the wait time(s)
    P95                  *float64       `json:"p95,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsZoneClientsWaits.
// It customizes the JSON marshaling process for StatsZoneClientsWaits objects.
func (s StatsZoneClientsWaits) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsZoneClientsWaits object to a map representation for JSON marshaling.
func (s StatsZoneClientsWaits) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "avg", "max", "min", "p95")
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
