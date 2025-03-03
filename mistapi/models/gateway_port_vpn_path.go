package models

import (
    "encoding/json"
    "fmt"
)

// GatewayPortVpnPath represents a GatewayPortVpnPath struct.
type GatewayPortVpnPath struct {
    // Only if the VPN `type`==`hub_spoke`. enum: `broadband`, `lte`
    BfdProfile           *GatewayPortVpnPathBfdProfileEnum `json:"bfd_profile,omitempty"`
    // Only if the VPN `type`==`hub_spoke`. Whether to use tunnel mode. SSR only
    BfdUseTunnelMode     *bool                             `json:"bfd_use_tunnel_mode,omitempty"`
    // Only if the VPN `type`==`hub_spoke`. For a given VPN, when `path_selection.strategy`==`simple`, the preference for a path (lower is preferred)
    Preference           *int                              `json:"preference,omitempty"`
    // If the VPN `type`==`hub_spoke`, enum: `hub`, `spoke`. If the VPN `type`==`mesh`, enum: `mesh`
    Role                 *GatewayPortVpnPathRoleEnum       `json:"role,omitempty"`
    // Only if the VPN `type`==`hub_spoke`
    TrafficShaping       *GatewayPortVpnPathTrafficShaping `json:"traffic_shaping,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for GatewayPortVpnPath,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GatewayPortVpnPath) String() string {
    return fmt.Sprintf(
    	"GatewayPortVpnPath[BfdProfile=%v, BfdUseTunnelMode=%v, Preference=%v, Role=%v, TrafficShaping=%v, AdditionalProperties=%v]",
    	g.BfdProfile, g.BfdUseTunnelMode, g.Preference, g.Role, g.TrafficShaping, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortVpnPath.
// It customizes the JSON marshaling process for GatewayPortVpnPath objects.
func (g GatewayPortVpnPath) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "bfd_profile", "bfd_use_tunnel_mode", "preference", "role", "traffic_shaping"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortVpnPath object to a map representation for JSON marshaling.
func (g GatewayPortVpnPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.BfdProfile != nil {
        structMap["bfd_profile"] = g.BfdProfile
    }
    if g.BfdUseTunnelMode != nil {
        structMap["bfd_use_tunnel_mode"] = g.BfdUseTunnelMode
    }
    if g.Preference != nil {
        structMap["preference"] = g.Preference
    }
    if g.Role != nil {
        structMap["role"] = g.Role
    }
    if g.TrafficShaping != nil {
        structMap["traffic_shaping"] = g.TrafficShaping.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayPortVpnPath.
// It customizes the JSON unmarshaling process for GatewayPortVpnPath objects.
func (g *GatewayPortVpnPath) UnmarshalJSON(input []byte) error {
    var temp tempGatewayPortVpnPath
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bfd_profile", "bfd_use_tunnel_mode", "preference", "role", "traffic_shaping")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.BfdProfile = temp.BfdProfile
    g.BfdUseTunnelMode = temp.BfdUseTunnelMode
    g.Preference = temp.Preference
    g.Role = temp.Role
    g.TrafficShaping = temp.TrafficShaping
    return nil
}

// tempGatewayPortVpnPath is a temporary struct used for validating the fields of GatewayPortVpnPath.
type tempGatewayPortVpnPath  struct {
    BfdProfile       *GatewayPortVpnPathBfdProfileEnum `json:"bfd_profile,omitempty"`
    BfdUseTunnelMode *bool                             `json:"bfd_use_tunnel_mode,omitempty"`
    Preference       *int                              `json:"preference,omitempty"`
    Role             *GatewayPortVpnPathRoleEnum       `json:"role,omitempty"`
    TrafficShaping   *GatewayPortVpnPathTrafficShaping `json:"traffic_shaping,omitempty"`
}
