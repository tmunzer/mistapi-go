package models

import (
    "encoding/json"
)

// ZoneStatsAssetsWaits represents a ZoneStatsAssetsWaits struct.
// ble asset wait time right now
type ZoneStatsAssetsWaits struct {
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

// MarshalJSON implements the json.Marshaler interface for ZoneStatsAssetsWaits.
// It customizes the JSON marshaling process for ZoneStatsAssetsWaits objects.
func (z ZoneStatsAssetsWaits) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneStatsAssetsWaits object to a map representation for JSON marshaling.
func (z ZoneStatsAssetsWaits) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneStatsAssetsWaits.
// It customizes the JSON unmarshaling process for ZoneStatsAssetsWaits objects.
func (z *ZoneStatsAssetsWaits) UnmarshalJSON(input []byte) error {
    var temp tempZoneStatsAssetsWaits
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

// tempZoneStatsAssetsWaits is a temporary struct used for validating the fields of ZoneStatsAssetsWaits.
type tempZoneStatsAssetsWaits  struct {
    Avg *float64 `json:"avg,omitempty"`
    Max *float64 `json:"max,omitempty"`
    Min *float64 `json:"min,omitempty"`
    P95 *float64 `json:"p95,omitempty"`
}
