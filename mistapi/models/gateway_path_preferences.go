package models

import (
    "encoding/json"
)

// GatewayPathPreferences represents a GatewayPathPreferences struct.
type GatewayPathPreferences struct {
    Paths                []GatewayPathPreferencesPath `json:"paths,omitempty"`
    // enum: `ecmp`, `ordered`, `weighted`
    Strategy             *GatewayPathStrategyEnum     `json:"strategy,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPathPreferences.
// It customizes the JSON marshaling process for GatewayPathPreferences objects.
func (g GatewayPathPreferences) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPathPreferences object to a map representation for JSON marshaling.
func (g GatewayPathPreferences) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Paths != nil {
        structMap["paths"] = g.Paths
    }
    if g.Strategy != nil {
        structMap["strategy"] = g.Strategy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPathPreferences.
// It customizes the JSON unmarshaling process for GatewayPathPreferences objects.
func (g *GatewayPathPreferences) UnmarshalJSON(input []byte) error {
    var temp gatewayPathPreferences
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "paths", "strategy")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Paths = temp.Paths
    g.Strategy = temp.Strategy
    return nil
}

// gatewayPathPreferences is a temporary struct used for validating the fields of GatewayPathPreferences.
type gatewayPathPreferences  struct {
    Paths    []GatewayPathPreferencesPath `json:"paths,omitempty"`
    Strategy *GatewayPathStrategyEnum     `json:"strategy,omitempty"`
}
