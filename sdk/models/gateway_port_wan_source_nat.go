package models

import (
    "encoding/json"
)

// GatewayPortWanSourceNat represents a GatewayPortWanSourceNat struct.
// optional, by default, source-NAT is performed on all WAN Ports using the interface-ip
type GatewayPortWanSourceNat struct {
    // or to disable the source-nat
    Disabled             *bool          `json:"disabled,omitempty"`
    // if alternative nat_pool is desired
    NatPool              *string        `json:"nat_pool,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortWanSourceNat.
// It customizes the JSON marshaling process for GatewayPortWanSourceNat objects.
func (g GatewayPortWanSourceNat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortWanSourceNat object to a map representation for JSON marshaling.
func (g GatewayPortWanSourceNat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
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
    var temp gatewayPortWanSourceNat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disabled", "nat_pool")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Disabled = temp.Disabled
    g.NatPool = temp.NatPool
    return nil
}

// gatewayPortWanSourceNat is a temporary struct used for validating the fields of GatewayPortWanSourceNat.
type gatewayPortWanSourceNat  struct {
    Disabled *bool   `json:"disabled,omitempty"`
    NatPool  *string `json:"nat_pool,omitempty"`
}