package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ApStatsL2TpStat represents a ApStatsL2TpStat struct.
type ApStatsL2TpStat struct {
    // list of sessions
    Sessions             []ApStatsL2TpStatSession `json:"sessions,omitempty"`
    State                *L2TpStateEnum           `json:"state,omitempty"`
    // uptime
    Uptime               *int                     `json:"uptime,omitempty"`
    // WxlanTunnel ID
    WxtunnelId           *uuid.UUID               `json:"wxtunnel_id,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsL2TpStat.
// It customizes the JSON marshaling process for ApStatsL2TpStat objects.
func (a ApStatsL2TpStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsL2TpStat object to a map representation for JSON marshaling.
func (a ApStatsL2TpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Sessions != nil {
        structMap["sessions"] = a.Sessions
    }
    if a.State != nil {
        structMap["state"] = a.State
    }
    if a.Uptime != nil {
        structMap["uptime"] = a.Uptime
    }
    if a.WxtunnelId != nil {
        structMap["wxtunnel_id"] = a.WxtunnelId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsL2TpStat.
// It customizes the JSON unmarshaling process for ApStatsL2TpStat objects.
func (a *ApStatsL2TpStat) UnmarshalJSON(input []byte) error {
    var temp apStatsL2TpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "sessions", "state", "uptime", "wxtunnel_id")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Sessions = temp.Sessions
    a.State = temp.State
    a.Uptime = temp.Uptime
    a.WxtunnelId = temp.WxtunnelId
    return nil
}

// apStatsL2TpStat is a temporary struct used for validating the fields of ApStatsL2TpStat.
type apStatsL2TpStat  struct {
    Sessions   []ApStatsL2TpStatSession `json:"sessions,omitempty"`
    State      *L2TpStateEnum           `json:"state,omitempty"`
    Uptime     *int                     `json:"uptime,omitempty"`
    WxtunnelId *uuid.UUID               `json:"wxtunnel_id,omitempty"`
}
