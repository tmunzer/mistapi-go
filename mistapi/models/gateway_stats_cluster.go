package models

import (
    "encoding/json"
)

// GatewayStatsCluster represents a GatewayStatsCluster struct.
type GatewayStatsCluster struct {
    State                Optional[string] `json:"state"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayStatsCluster.
// It customizes the JSON marshaling process for GatewayStatsCluster objects.
func (g GatewayStatsCluster) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayStatsCluster object to a map representation for JSON marshaling.
func (g GatewayStatsCluster) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.State.IsValueSet() {
        if g.State.Value() != nil {
            structMap["state"] = g.State.Value()
        } else {
            structMap["state"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayStatsCluster.
// It customizes the JSON unmarshaling process for GatewayStatsCluster objects.
func (g *GatewayStatsCluster) UnmarshalJSON(input []byte) error {
    var temp tempGatewayStatsCluster
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "state")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.State = temp.State
    return nil
}

// tempGatewayStatsCluster is a temporary struct used for validating the fields of GatewayStatsCluster.
type tempGatewayStatsCluster  struct {
    State Optional[string] `json:"state"`
}
