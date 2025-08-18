// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// GatewayWanProbeOverride represents a GatewayWanProbeOverride struct.
// Only if `usage`==`wan`
type GatewayWanProbeOverride struct {
    Ip6s                 []string                                 `json:"ip6s,omitempty"`
    Ips                  []string                                 `json:"ips,omitempty"`
    // enum: `broadband`, `lte`
    ProbeProfile         *GatewayWanProbeOverrideProbeProfileEnum `json:"probe_profile,omitempty"`
    AdditionalProperties map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayWanProbeOverride,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayWanProbeOverride) String() string {
    return fmt.Sprintf(
    	"GatewayWanProbeOverride[Ip6s=%v, Ips=%v, ProbeProfile=%v, AdditionalProperties=%v]",
    	g.Ip6s, g.Ips, g.ProbeProfile, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayWanProbeOverride.
// It customizes the JSON marshaling process for GatewayWanProbeOverride objects.
func (g GatewayWanProbeOverride) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "ip6s", "ips", "probe_profile"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayWanProbeOverride object to a map representation for JSON marshaling.
func (g GatewayWanProbeOverride) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Ip6s != nil {
        structMap["ip6s"] = g.Ip6s
    }
    if g.Ips != nil {
        structMap["ips"] = g.Ips
    }
    if g.ProbeProfile != nil {
        structMap["probe_profile"] = g.ProbeProfile
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayWanProbeOverride.
// It customizes the JSON unmarshaling process for GatewayWanProbeOverride objects.
func (g *GatewayWanProbeOverride) UnmarshalJSON(input []byte) error {
    var temp tempGatewayWanProbeOverride
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip6s", "ips", "probe_profile")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Ip6s = temp.Ip6s
    g.Ips = temp.Ips
    g.ProbeProfile = temp.ProbeProfile
    return nil
}

// tempGatewayWanProbeOverride is a temporary struct used for validating the fields of GatewayWanProbeOverride.
type tempGatewayWanProbeOverride  struct {
    Ip6s         []string                                 `json:"ip6s,omitempty"`
    Ips          []string                                 `json:"ips,omitempty"`
    ProbeProfile *GatewayWanProbeOverrideProbeProfileEnum `json:"probe_profile,omitempty"`
}
