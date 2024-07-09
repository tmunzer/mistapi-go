package models

import (
    "encoding/json"
)

// GatewayPortMirroring represents a GatewayPortMirroring struct.
type GatewayPortMirroring struct {
    PortMirror           *GatewayPortMirroringPortMirror `json:"port_mirror,omitempty"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortMirroring.
// It customizes the JSON marshaling process for GatewayPortMirroring objects.
func (g GatewayPortMirroring) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortMirroring object to a map representation for JSON marshaling.
func (g GatewayPortMirroring) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.PortMirror != nil {
        structMap["port_mirror"] = g.PortMirror.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortMirroring.
// It customizes the JSON unmarshaling process for GatewayPortMirroring objects.
func (g *GatewayPortMirroring) UnmarshalJSON(input []byte) error {
    var temp gatewayPortMirroring
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "port_mirror")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.PortMirror = temp.PortMirror
    return nil
}

// gatewayPortMirroring is a temporary struct used for validating the fields of GatewayPortMirroring.
type gatewayPortMirroring  struct {
    PortMirror *GatewayPortMirroringPortMirror `json:"port_mirror,omitempty"`
}
