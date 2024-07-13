package models

import (
    "encoding/json"
)

// ApStatsUsbStat represents a ApStatsUsbStat struct.
type ApStatsUsbStat struct {
    Channel              Optional[int]    `json:"channel"`
    Connected            Optional[bool]   `json:"connected"`
    LastActivity         Optional[int]    `json:"last_activity"`
    Type                 Optional[string] `json:"type"`
    Up                   Optional[bool]   `json:"up"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsUsbStat.
// It customizes the JSON marshaling process for ApStatsUsbStat objects.
func (a ApStatsUsbStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsUsbStat object to a map representation for JSON marshaling.
func (a ApStatsUsbStat) toMap() map[string]any {
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
    if a.LastActivity.IsValueSet() {
        if a.LastActivity.Value() != nil {
            structMap["last_activity"] = a.LastActivity.Value()
        } else {
            structMap["last_activity"] = nil
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

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsUsbStat.
// It customizes the JSON unmarshaling process for ApStatsUsbStat objects.
func (a *ApStatsUsbStat) UnmarshalJSON(input []byte) error {
    var temp apStatsUsbStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "connected", "last_activity", "type", "up")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Channel = temp.Channel
    a.Connected = temp.Connected
    a.LastActivity = temp.LastActivity
    a.Type = temp.Type
    a.Up = temp.Up
    return nil
}

// apStatsUsbStat is a temporary struct used for validating the fields of ApStatsUsbStat.
type apStatsUsbStat  struct {
    Channel      Optional[int]    `json:"channel"`
    Connected    Optional[bool]   `json:"connected"`
    LastActivity Optional[int]    `json:"last_activity"`
    Type         Optional[string] `json:"type"`
    Up           Optional[bool]   `json:"up"`
}
