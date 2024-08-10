package models

import (
    "encoding/json"
)

// GatewayPortVpnPath represents a GatewayPortVpnPath struct.
type GatewayPortVpnPath struct {
    // enum: `broadband`, `lte`
    BfdProfile           *GatewayPortVpnPathBfdProfileEnum `json:"bfd_profile,omitempty"`
    // whether to use tunnel mode. SSR only
    BfdUseTunnelMode     *bool                             `json:"bfd_use_tunnel_mode,omitempty"`
    // for a given VPN, when `path_selection.strategy`==`simple`, the preference for a path (lower is preferred)
    Preference           *int                              `json:"preference,omitempty"`
    // enum: `hub`, `spoke`
    Role                 *GatewayPortVpnPathRoleEnum       `json:"role,omitempty"`
    TrafficShaping       *GatewayTrafficShaping            `json:"traffic_shaping,omitempty"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayPortVpnPath.
// It customizes the JSON marshaling process for GatewayPortVpnPath objects.
func (g GatewayPortVpnPath) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayPortVpnPath object to a map representation for JSON marshaling.
func (g GatewayPortVpnPath) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bfd_profile", "bfd_use_tunnel_mode", "preference", "role", "traffic_shaping")
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
    TrafficShaping   *GatewayTrafficShaping            `json:"traffic_shaping,omitempty"`
}
