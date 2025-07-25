// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// GatewayPortWanSourceNat represents a GatewayPortWanSourceNat struct.
// Only if `usage`==`wan`, optional. By default, source-NAT is performed on all WAN Ports using the interface-ip
type GatewayPortWanSourceNat struct {
    // Or to disable the source-nat
    Disabled             *bool                  `json:"disabled,omitempty"`
    // If alternative nat_pool is desired
    NatPool              *string                `json:"nat_pool,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPortWanSourceNat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortWanSourceNat) String() string {
    return fmt.Sprintf(
    	"GatewayPortWanSourceNat[Disabled=%v, NatPool=%v, AdditionalProperties=%v]",
    	g.Disabled, g.NatPool, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortWanSourceNat.
// It customizes the JSON marshaling process for GatewayPortWanSourceNat objects.
func (g GatewayPortWanSourceNat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "disabled", "nat_pool"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortWanSourceNat object to a map representation for JSON marshaling.
func (g GatewayPortWanSourceNat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Disabled != nil {
        structMap["disabled"] = g.Disabled
    }
    if g.NatPool != nil {
        structMap["nat_pool"] = g.NatPool
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortWanSourceNat.
// It customizes the JSON unmarshaling process for GatewayPortWanSourceNat objects.
func (g *GatewayPortWanSourceNat) UnmarshalJSON(input []byte) error {
    var temp tempGatewayPortWanSourceNat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disabled", "nat_pool")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Disabled = temp.Disabled
    g.NatPool = temp.NatPool
    return nil
}

// tempGatewayPortWanSourceNat is a temporary struct used for validating the fields of GatewayPortWanSourceNat.
type tempGatewayPortWanSourceNat  struct {
    Disabled *bool   `json:"disabled,omitempty"`
    NatPool  *string `json:"nat_pool,omitempty"`
}
