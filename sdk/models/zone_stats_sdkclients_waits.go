package models

import (
    "encoding/json"
)

// ZoneStatsSdkclientsWaits represents a ZoneStatsSdkclientsWaits struct.
// sdkclient wait time right now
type ZoneStatsSdkclientsWaits struct {
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

// MarshalJSON implements the json.Marshaler interface for ZoneStatsSdkclientsWaits.
// It customizes the JSON marshaling process for ZoneStatsSdkclientsWaits objects.
func (z ZoneStatsSdkclientsWaits) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneStatsSdkclientsWaits object to a map representation for JSON marshaling.
func (z ZoneStatsSdkclientsWaits) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneStatsSdkclientsWaits.
// It customizes the JSON unmarshaling process for ZoneStatsSdkclientsWaits objects.
func (z *ZoneStatsSdkclientsWaits) UnmarshalJSON(input []byte) error {
    var temp zoneStatsSdkclientsWaits
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

// zoneStatsSdkclientsWaits is a temporary struct used for validating the fields of ZoneStatsSdkclientsWaits.
type zoneStatsSdkclientsWaits  struct {
    Avg *float64 `json:"avg,omitempty"`
    Max *float64 `json:"max,omitempty"`
    Min *float64 `json:"min,omitempty"`
    P95 *float64 `json:"p95,omitempty"`
}
