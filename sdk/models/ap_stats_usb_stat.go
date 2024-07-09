package models

import (
    "encoding/json"
)

// ApStatsUsbStat represents a ApStatsUsbStat struct.
type ApStatsUsbStat struct {
    Channel              *int           `json:"channel,omitempty"`
    Connected            *bool          `json:"connected,omitempty"`
    LastActivity         *int           `json:"last_activity,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    Up                   *bool          `json:"up,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
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
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.Connected != nil {
        structMap["connected"] = a.Connected
    }
    if a.LastActivity != nil {
        structMap["last_activity"] = a.LastActivity
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    if a.Up != nil {
        structMap["up"] = a.Up
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
    Channel      *int    `json:"channel,omitempty"`
    Connected    *bool   `json:"connected,omitempty"`
    LastActivity *int    `json:"last_activity,omitempty"`
    Type         *string `json:"type,omitempty"`
    Up           *bool   `json:"up,omitempty"`
}
