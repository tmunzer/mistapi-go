package models

import (
    "encoding/json"
)

// MxedgeStatsServiceStat represents a MxedgeStatsServiceStat struct.
type MxedgeStatsServiceStat struct {
    // external IP from ep-terminator’s point of view. valid only for service having its own cloud connection
    ExtIp                *string        `json:"ext_ip,omitempty"`
    // timestamp when the last stats is seen (cloud unix time, in second). valid only for service having its own stats or whole system (last among last_seen of all services)
    LastSeen             *float64       `json:"last_seen,omitempty"`
    // package/service installation state.
    PackageState         *string        `json:"package_state,omitempty"`
    // package/service installation state.
    PackageVersion       *string        `json:"package_version,omitempty"`
    // service running state.
    RunningState         *string        `json:"running_state,omitempty"`
    // service uptime.
    Uptime               *int           `json:"uptime,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsServiceStat.
// It customizes the JSON marshaling process for MxedgeStatsServiceStat objects.
func (m MxedgeStatsServiceStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsServiceStat object to a map representation for JSON marshaling.
func (m MxedgeStatsServiceStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.ExtIp != nil {
        structMap["ext_ip"] = m.ExtIp
    }
    if m.LastSeen != nil {
        structMap["last_seen"] = m.LastSeen
    }
    if m.PackageState != nil {
        structMap["package_state"] = m.PackageState
    }
    if m.PackageVersion != nil {
        structMap["package_version"] = m.PackageVersion
    }
    if m.RunningState != nil {
        structMap["running_state"] = m.RunningState
    }
    if m.Uptime != nil {
        structMap["uptime"] = m.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsServiceStat.
// It customizes the JSON unmarshaling process for MxedgeStatsServiceStat objects.
func (m *MxedgeStatsServiceStat) UnmarshalJSON(input []byte) error {
    var temp mxedgeStatsServiceStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ext_ip", "last_seen", "package_state", "package_version", "running_state", "uptime")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.ExtIp = temp.ExtIp
    m.LastSeen = temp.LastSeen
    m.PackageState = temp.PackageState
    m.PackageVersion = temp.PackageVersion
    m.RunningState = temp.RunningState
    m.Uptime = temp.Uptime
    return nil
}

// mxedgeStatsServiceStat is a temporary struct used for validating the fields of MxedgeStatsServiceStat.
type mxedgeStatsServiceStat  struct {
    ExtIp          *string  `json:"ext_ip,omitempty"`
    LastSeen       *float64 `json:"last_seen,omitempty"`
    PackageState   *string  `json:"package_state,omitempty"`
    PackageVersion *string  `json:"package_version,omitempty"`
    RunningState   *string  `json:"running_state,omitempty"`
    Uptime         *int     `json:"uptime,omitempty"`
}
