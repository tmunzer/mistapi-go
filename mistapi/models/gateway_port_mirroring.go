// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// GatewayPortMirroring represents a GatewayPortMirroring struct.
type GatewayPortMirroring struct {
    PortMirror           *GatewayPortMirroringPortMirror `json:"port_mirror,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPortMirroring,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortMirroring) String() string {
    return fmt.Sprintf(
    	"GatewayPortMirroring[PortMirror=%v, AdditionalProperties=%v]",
    	g.PortMirror, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortMirroring.
// It customizes the JSON marshaling process for GatewayPortMirroring objects.
func (g GatewayPortMirroring) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "port_mirror"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortMirroring object to a map representation for JSON marshaling.
func (g GatewayPortMirroring) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.PortMirror != nil {
        structMap["port_mirror"] = g.PortMirror.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortMirroring.
// It customizes the JSON unmarshaling process for GatewayPortMirroring objects.
func (g *GatewayPortMirroring) UnmarshalJSON(input []byte) error {
    var temp tempGatewayPortMirroring
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "port_mirror")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.PortMirror = temp.PortMirror
    return nil
}

// tempGatewayPortMirroring is a temporary struct used for validating the fields of GatewayPortMirroring.
type tempGatewayPortMirroring  struct {
    PortMirror *GatewayPortMirroringPortMirror `json:"port_mirror,omitempty"`
}
