package models

import (
    "encoding/json"
)

// GatewayStatsClusterStat represents a GatewayStatsClusterStat struct.
type GatewayStatsClusterStat struct {
    Status               *string        `json:"status,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayStatsClusterStat.
// It customizes the JSON marshaling process for GatewayStatsClusterStat objects.
func (g GatewayStatsClusterStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayStatsClusterStat object to a map representation for JSON marshaling.
func (g GatewayStatsClusterStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Status != nil {
        structMap["status"] = g.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayStatsClusterStat.
// It customizes the JSON unmarshaling process for GatewayStatsClusterStat objects.
func (g *GatewayStatsClusterStat) UnmarshalJSON(input []byte) error {
    var temp gatewayStatsClusterStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "status")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Status = temp.Status
    return nil
}

// gatewayStatsClusterStat is a temporary struct used for validating the fields of GatewayStatsClusterStat.
type gatewayStatsClusterStat  struct {
    Status *string `json:"status,omitempty"`
}
