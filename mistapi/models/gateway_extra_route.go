package models

import (
    "encoding/json"
)

// GatewayExtraRoute represents a GatewayExtraRoute struct.
// Property key is the destination CIDR (e.g. "10.0.0.0/8")
type GatewayExtraRoute struct {
    Via                  *string        `json:"via,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayExtraRoute.
// It customizes the JSON marshaling process for GatewayExtraRoute objects.
func (g GatewayExtraRoute) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayExtraRoute object to a map representation for JSON marshaling.
func (g GatewayExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Via != nil {
        structMap["via"] = g.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayExtraRoute.
// It customizes the JSON unmarshaling process for GatewayExtraRoute objects.
func (g *GatewayExtraRoute) UnmarshalJSON(input []byte) error {
    var temp gatewayExtraRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "via")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Via = temp.Via
    return nil
}

// gatewayExtraRoute is a temporary struct used for validating the fields of GatewayExtraRoute.
type gatewayExtraRoute  struct {
    Via *string `json:"via,omitempty"`
}
