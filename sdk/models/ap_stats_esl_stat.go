package models

import (
    "encoding/json"
)

// ApStatsEslStat represents a ApStatsEslStat struct.
type ApStatsEslStat struct {
    Channel              *int           `json:"channel,omitempty"`
    Connected            *bool          `json:"connected,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    Up                   *bool          `json:"up,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsEslStat.
// It customizes the JSON marshaling process for ApStatsEslStat objects.
func (a ApStatsEslStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsEslStat object to a map representation for JSON marshaling.
func (a ApStatsEslStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.Connected != nil {
        structMap["connected"] = a.Connected
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    if a.Up != nil {
        structMap["up"] = a.Up
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsEslStat.
// It customizes the JSON unmarshaling process for ApStatsEslStat objects.
func (a *ApStatsEslStat) UnmarshalJSON(input []byte) error {
    var temp apStatsEslStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "connected", "type", "up")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Channel = temp.Channel
    a.Connected = temp.Connected
    a.Type = temp.Type
    a.Up = temp.Up
    return nil
}

// apStatsEslStat is a temporary struct used for validating the fields of ApStatsEslStat.
type apStatsEslStat  struct {
    Channel   *int    `json:"channel,omitempty"`
    Connected *bool   `json:"connected,omitempty"`
    Type      *string `json:"type,omitempty"`
    Up        *bool   `json:"up,omitempty"`
}
