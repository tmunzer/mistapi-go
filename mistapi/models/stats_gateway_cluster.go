package models

import (
    "encoding/json"
)

// StatsGatewayCluster represents a StatsGatewayCluster struct.
type StatsGatewayCluster struct {
    State                Optional[string] `json:"state"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayCluster.
// It customizes the JSON marshaling process for StatsGatewayCluster objects.
func (s StatsGatewayCluster) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayCluster object to a map representation for JSON marshaling.
func (s StatsGatewayCluster) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.State.IsValueSet() {
        if s.State.Value() != nil {
            structMap["state"] = s.State.Value()
        } else {
            structMap["state"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayCluster.
// It customizes the JSON unmarshaling process for StatsGatewayCluster objects.
func (s *StatsGatewayCluster) UnmarshalJSON(input []byte) error {
    var temp tempStatsGatewayCluster
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "state")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.State = temp.State
    return nil
}

// tempStatsGatewayCluster is a temporary struct used for validating the fields of StatsGatewayCluster.
type tempStatsGatewayCluster  struct {
    State Optional[string] `json:"state"`
}
