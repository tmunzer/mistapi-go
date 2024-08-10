package models

import (
    "encoding/json"
)

// ApStatsEslStat represents a ApStatsEslStat struct.
type ApStatsEslStat struct {
    Channel              Optional[int]    `json:"channel"`
    Connected            Optional[bool]   `json:"connected"`
    Type                 Optional[string] `json:"type"`
    Up                   Optional[bool]   `json:"up"`
    AdditionalProperties map[string]any   `json:"_"`
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
    if a.Channel.IsValueSet() {
        if a.Channel.Value() != nil {
            structMap["channel"] = a.Channel.Value()
        } else {
            structMap["channel"] = nil
        }
    }
    if a.Connected.IsValueSet() {
        if a.Connected.Value() != nil {
            structMap["connected"] = a.Connected.Value()
        } else {
            structMap["connected"] = nil
        }
    }
    if a.Type.IsValueSet() {
        if a.Type.Value() != nil {
            structMap["type"] = a.Type.Value()
        } else {
            structMap["type"] = nil
        }
    }
    if a.Up.IsValueSet() {
        if a.Up.Value() != nil {
            structMap["up"] = a.Up.Value()
        } else {
            structMap["up"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsEslStat.
// It customizes the JSON unmarshaling process for ApStatsEslStat objects.
func (a *ApStatsEslStat) UnmarshalJSON(input []byte) error {
    var temp tempApStatsEslStat
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

// tempApStatsEslStat is a temporary struct used for validating the fields of ApStatsEslStat.
type tempApStatsEslStat  struct {
    Channel   Optional[int]    `json:"channel"`
    Connected Optional[bool]   `json:"connected"`
    Type      Optional[string] `json:"type"`
    Up        Optional[bool]   `json:"up"`
}
