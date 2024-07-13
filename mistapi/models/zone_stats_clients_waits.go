package models

import (
    "encoding/json"
)

// ZoneStatsClientsWaits represents a ZoneStatsClientsWaits struct.
// client wait time right now
type ZoneStatsClientsWaits struct {
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

// MarshalJSON implements the json.Marshaler interface for ZoneStatsClientsWaits.
// It customizes the JSON marshaling process for ZoneStatsClientsWaits objects.
func (z ZoneStatsClientsWaits) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneStatsClientsWaits object to a map representation for JSON marshaling.
func (z ZoneStatsClientsWaits) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, z.AdditionalProperties)
    if z.Avg != nil {
        structMap["avg"] = z.Avg
    }
    if z.Max != nil {
        structMap["max"] = z.Max
    }
    if z.Min != nil {
        structMap["min"] = z.Min
    }
    if z.P95 != nil {
        structMap["p95"] = z.P95
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneStatsClientsWaits.
// It customizes the JSON unmarshaling process for ZoneStatsClientsWaits objects.
func (z *ZoneStatsClientsWaits) UnmarshalJSON(input []byte) error {
    var temp zoneStatsClientsWaits
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "avg", "max", "min", "p95")
    if err != nil {
    	return err
    }
    
    z.AdditionalProperties = additionalProperties
    z.Avg = temp.Avg
    z.Max = temp.Max
    z.Min = temp.Min
    z.P95 = temp.P95
    return nil
}

// zoneStatsClientsWaits is a temporary struct used for validating the fields of ZoneStatsClientsWaits.
type zoneStatsClientsWaits  struct {
    Avg *float64 `json:"avg,omitempty"`
    Max *float64 `json:"max,omitempty"`
    Min *float64 `json:"min,omitempty"`
    P95 *float64 `json:"p95,omitempty"`
}
