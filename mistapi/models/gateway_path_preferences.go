package models

import (
    "encoding/json"
)

// GatewayPathPreferences represents a GatewayPathPreferences struct.
type GatewayPathPreferences struct {
    Paths                []GatewayPathPreferencesPath `json:"paths,omitempty"`
    // enum: `ecmp`, `ordered`, `weighted`
    Strategy             *GatewayPathStrategyEnum     `json:"strategy,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPathPreferences.
// It customizes the JSON marshaling process for GatewayPathPreferences objects.
func (g GatewayPathPreferences) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "paths", "strategy"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPathPreferences object to a map representation for JSON marshaling.
func (g GatewayPathPreferences) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
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
    var temp tempGatewayPathPreferences
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "paths", "strategy")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Paths = temp.Paths
    g.Strategy = temp.Strategy
    return nil
}

// tempGatewayPathPreferences is a temporary struct used for validating the fields of GatewayPathPreferences.
type tempGatewayPathPreferences  struct {
    Paths    []GatewayPathPreferencesPath `json:"paths,omitempty"`
    Strategy *GatewayPathStrategyEnum     `json:"strategy,omitempty"`
}
