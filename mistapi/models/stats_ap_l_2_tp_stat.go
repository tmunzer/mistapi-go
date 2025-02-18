package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsApL2tpStat represents a StatsApL2tpStat struct.
type StatsApL2tpStat struct {
    // List of sessions
    Sessions             []StatsApL2tpStatSession `json:"sessions,omitempty"`
    // enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
    State                *L2tpStateEnum           `json:"state,omitempty"`
    // Uptime
    Uptime               Optional[int]            `json:"uptime"`
    // WxlanTunnel ID
    WxtunnelId           Optional[uuid.UUID]      `json:"wxtunnel_id"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApL2tpStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApL2tpStat) String() string {
    return fmt.Sprintf(
    	"StatsApL2tpStat[Sessions=%v, State=%v, Uptime=%v, WxtunnelId=%v, AdditionalProperties=%v]",
    	s.Sessions, s.State, s.Uptime, s.WxtunnelId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApL2tpStat.
// It customizes the JSON marshaling process for StatsApL2tpStat objects.
func (s StatsApL2tpStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "sessions", "state", "uptime", "wxtunnel_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApL2tpStat object to a map representation for JSON marshaling.
func (s StatsApL2tpStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Sessions != nil {
        structMap["sessions"] = s.Sessions
    }
    if s.State != nil {
        structMap["state"] = s.State
    }
    if s.Uptime.IsValueSet() {
        if s.Uptime.Value() != nil {
            structMap["uptime"] = s.Uptime.Value()
        } else {
            structMap["uptime"] = nil
        }
    }
    if s.WxtunnelId.IsValueSet() {
        if s.WxtunnelId.Value() != nil {
            structMap["wxtunnel_id"] = s.WxtunnelId.Value()
        } else {
            structMap["wxtunnel_id"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApL2tpStat.
// It customizes the JSON unmarshaling process for StatsApL2tpStat objects.
func (s *StatsApL2tpStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsApL2tpStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "sessions", "state", "uptime", "wxtunnel_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Sessions = temp.Sessions
    s.State = temp.State
    s.Uptime = temp.Uptime
    s.WxtunnelId = temp.WxtunnelId
    return nil
}

// tempStatsApL2tpStat is a temporary struct used for validating the fields of StatsApL2tpStat.
type tempStatsApL2tpStat  struct {
    Sessions   []StatsApL2tpStatSession `json:"sessions,omitempty"`
    State      *L2tpStateEnum           `json:"state,omitempty"`
    Uptime     Optional[int]            `json:"uptime"`
    WxtunnelId Optional[uuid.UUID]      `json:"wxtunnel_id"`
}
