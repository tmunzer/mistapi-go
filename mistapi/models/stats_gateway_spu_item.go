package models

import (
    "encoding/json"
)

// StatsGatewaySpuItem represents a StatsGatewaySpuItem struct.
type StatsGatewaySpuItem struct {
    SpuCpu               *int           `json:"spu_cpu,omitempty"`
    SpuCurrentSession    *int           `json:"spu_current_session,omitempty"`
    SpuMaxSession        *int           `json:"spu_max_session,omitempty"`
    SpuMemory            *int           `json:"spu_memory,omitempty"`
    SpuPendingSession    *int           `json:"spu_pending_session,omitempty"`
    SpuValidSession      *int           `json:"spu_valid_session,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewaySpuItem.
// It customizes the JSON marshaling process for StatsGatewaySpuItem objects.
func (s StatsGatewaySpuItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewaySpuItem object to a map representation for JSON marshaling.
func (s StatsGatewaySpuItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.SpuCpu != nil {
        structMap["spu_cpu"] = s.SpuCpu
    }
    if s.SpuCurrentSession != nil {
        structMap["spu_current_session"] = s.SpuCurrentSession
    }
    if s.SpuMaxSession != nil {
        structMap["spu_max_session"] = s.SpuMaxSession
    }
    if s.SpuMemory != nil {
        structMap["spu_memory"] = s.SpuMemory
    }
    if s.SpuPendingSession != nil {
        structMap["spu_pending_session"] = s.SpuPendingSession
    }
    if s.SpuValidSession != nil {
        structMap["spu_valid_session"] = s.SpuValidSession
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewaySpuItem.
// It customizes the JSON unmarshaling process for StatsGatewaySpuItem objects.
func (s *StatsGatewaySpuItem) UnmarshalJSON(input []byte) error {
    var temp tempStatsGatewaySpuItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "spu_cpu", "spu_current_session", "spu_max_session", "spu_memory", "spu_pending_session", "spu_valid_session")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.SpuCpu = temp.SpuCpu
    s.SpuCurrentSession = temp.SpuCurrentSession
    s.SpuMaxSession = temp.SpuMaxSession
    s.SpuMemory = temp.SpuMemory
    s.SpuPendingSession = temp.SpuPendingSession
    s.SpuValidSession = temp.SpuValidSession
    return nil
}

// tempStatsGatewaySpuItem is a temporary struct used for validating the fields of StatsGatewaySpuItem.
type tempStatsGatewaySpuItem  struct {
    SpuCpu            *int `json:"spu_cpu,omitempty"`
    SpuCurrentSession *int `json:"spu_current_session,omitempty"`
    SpuMaxSession     *int `json:"spu_max_session,omitempty"`
    SpuMemory         *int `json:"spu_memory,omitempty"`
    SpuPendingSession *int `json:"spu_pending_session,omitempty"`
    SpuValidSession   *int `json:"spu_valid_session,omitempty"`
}