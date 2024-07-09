package models

import (
    "encoding/json"
)

// GatewayStatsDhcpdStatLan represents a GatewayStatsDhcpdStatLan struct.
type GatewayStatsDhcpdStatLan struct {
    NumIps               *int           `json:"num_ips,omitempty"`
    NumLeased            *int           `json:"num_leased,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayStatsDhcpdStatLan.
// It customizes the JSON marshaling process for GatewayStatsDhcpdStatLan objects.
func (g GatewayStatsDhcpdStatLan) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayStatsDhcpdStatLan object to a map representation for JSON marshaling.
func (g GatewayStatsDhcpdStatLan) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.NumIps != nil {
        structMap["num_ips"] = g.NumIps
    }
    if g.NumLeased != nil {
        structMap["num_leased"] = g.NumLeased
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayStatsDhcpdStatLan.
// It customizes the JSON unmarshaling process for GatewayStatsDhcpdStatLan objects.
func (g *GatewayStatsDhcpdStatLan) UnmarshalJSON(input []byte) error {
    var temp gatewayStatsDhcpdStatLan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_ips", "num_leased")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.NumIps = temp.NumIps
    g.NumLeased = temp.NumLeased
    return nil
}

// gatewayStatsDhcpdStatLan is a temporary struct used for validating the fields of GatewayStatsDhcpdStatLan.
type gatewayStatsDhcpdStatLan  struct {
    NumIps    *int `json:"num_ips,omitempty"`
    NumLeased *int `json:"num_leased,omitempty"`
}
