package models

import (
    "encoding/json"
)

// GatewayStatsSpuStat represents a GatewayStatsSpuStat struct.
type GatewayStatsSpuStat struct {
    SpuCpu               *int           `json:"spu_cpu,omitempty"`
    SpuCurrentSession    *float64       `json:"spu_current_session,omitempty"`
    SpuMaxSession        *float64       `json:"spu_max_session,omitempty"`
    SpuMemory            *int           `json:"spu_memory,omitempty"`
    SpuPendingSession    *float64       `json:"spu_pending_session,omitempty"`
    SpuValidSession      *float64       `json:"spu_valid_session,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayStatsSpuStat.
// It customizes the JSON marshaling process for GatewayStatsSpuStat objects.
func (g GatewayStatsSpuStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayStatsSpuStat object to a map representation for JSON marshaling.
func (g GatewayStatsSpuStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.SpuCpu != nil {
        structMap["spu_cpu"] = g.SpuCpu
    }
    if g.SpuCurrentSession != nil {
        structMap["spu_current_session"] = g.SpuCurrentSession
    }
    if g.SpuMaxSession != nil {
        structMap["spu_max_session"] = g.SpuMaxSession
    }
    if g.SpuMemory != nil {
        structMap["spu_memory"] = g.SpuMemory
    }
    if g.SpuPendingSession != nil {
        structMap["spu_pending_session"] = g.SpuPendingSession
    }
    if g.SpuValidSession != nil {
        structMap["spu_valid_session"] = g.SpuValidSession
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayStatsSpuStat.
// It customizes the JSON unmarshaling process for GatewayStatsSpuStat objects.
func (g *GatewayStatsSpuStat) UnmarshalJSON(input []byte) error {
    var temp gatewayStatsSpuStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "spu_cpu", "spu_current_session", "spu_max_session", "spu_memory", "spu_pending_session", "spu_valid_session")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.SpuCpu = temp.SpuCpu
    g.SpuCurrentSession = temp.SpuCurrentSession
    g.SpuMaxSession = temp.SpuMaxSession
    g.SpuMemory = temp.SpuMemory
    g.SpuPendingSession = temp.SpuPendingSession
    g.SpuValidSession = temp.SpuValidSession
    return nil
}

// gatewayStatsSpuStat is a temporary struct used for validating the fields of GatewayStatsSpuStat.
type gatewayStatsSpuStat  struct {
    SpuCpu            *int     `json:"spu_cpu,omitempty"`
    SpuCurrentSession *float64 `json:"spu_current_session,omitempty"`
    SpuMaxSession     *float64 `json:"spu_max_session,omitempty"`
    SpuMemory         *int     `json:"spu_memory,omitempty"`
    SpuPendingSession *float64 `json:"spu_pending_session,omitempty"`
    SpuValidSession   *float64 `json:"spu_valid_session,omitempty"`
}
