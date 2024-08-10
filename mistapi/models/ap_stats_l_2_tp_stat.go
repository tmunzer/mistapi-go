package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ApStatsL2tpStat represents a ApStatsL2tpStat struct.
type ApStatsL2tpStat struct {
    // list of sessions
    Sessions             []ApStatsL2tpStatSession `json:"sessions,omitempty"`
    // enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
    State                *L2tpStateEnum           `json:"state,omitempty"`
    // uptime
    Uptime               Optional[int]            `json:"uptime"`
    // WxlanTunnel ID
    WxtunnelId           Optional[uuid.UUID]      `json:"wxtunnel_id"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsL2tpStat.
// It customizes the JSON marshaling process for ApStatsL2tpStat objects.
func (a ApStatsL2tpStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsL2tpStat object to a map representation for JSON marshaling.
func (a ApStatsL2tpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Sessions != nil {
        structMap["sessions"] = a.Sessions
    }
    if a.State != nil {
        structMap["state"] = a.State
    }
    if a.Uptime.IsValueSet() {
        if a.Uptime.Value() != nil {
            structMap["uptime"] = a.Uptime.Value()
        } else {
            structMap["uptime"] = nil
        }
    }
    if a.WxtunnelId.IsValueSet() {
        if a.WxtunnelId.Value() != nil {
            structMap["wxtunnel_id"] = a.WxtunnelId.Value()
        } else {
            structMap["wxtunnel_id"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsL2tpStat.
// It customizes the JSON unmarshaling process for ApStatsL2tpStat objects.
func (a *ApStatsL2tpStat) UnmarshalJSON(input []byte) error {
    var temp tempApStatsL2tpStat
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

// tempApStatsL2tpStat is a temporary struct used for validating the fields of ApStatsL2tpStat.
type tempApStatsL2tpStat  struct {
    Sessions   []ApStatsL2tpStatSession `json:"sessions,omitempty"`
    State      *L2tpStateEnum           `json:"state,omitempty"`
    Uptime     Optional[int]            `json:"uptime"`
    WxtunnelId Optional[uuid.UUID]      `json:"wxtunnel_id"`
}
