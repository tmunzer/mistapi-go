package models

import (
    "encoding/json"
)

// WxlanTunnelDmvpn represents a WxlanTunnelDmvpn struct.
// Dynamic Multipoint VPN configurations
type WxlanTunnelDmvpn struct {
    // whether DMVPN is enabled
    Enabled              *bool          `json:"enabled,omitempty"`
    // optional; the holding time for NHRP ‘registration requests’  and ‘resolution replies’ sent from the Mist AP (in seconds); default 600
    HoldingTime          *int           `json:"holding_time,omitempty"`
    // optional; list of IPv4 DMVPN peer host ip-addresses to which traffic is forwarded
    HostRoutes           []string       `json:"host_routes,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WxlanTunnelDmvpn.
// It customizes the JSON marshaling process for WxlanTunnelDmvpn objects.
func (w WxlanTunnelDmvpn) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WxlanTunnelDmvpn object to a map representation for JSON marshaling.
func (w WxlanTunnelDmvpn) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.HoldingTime != nil {
        structMap["holding_time"] = w.HoldingTime
    }
    if w.HostRoutes != nil {
        structMap["host_routes"] = w.HostRoutes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanTunnelDmvpn.
// It customizes the JSON unmarshaling process for WxlanTunnelDmvpn objects.
func (w *WxlanTunnelDmvpn) UnmarshalJSON(input []byte) error {
    var temp wxlanTunnelDmvpn
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "holding_time", "host_routes")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Enabled = temp.Enabled
    w.HoldingTime = temp.HoldingTime
    w.HostRoutes = temp.HostRoutes
    return nil
}

// wxlanTunnelDmvpn is a temporary struct used for validating the fields of WxlanTunnelDmvpn.
type wxlanTunnelDmvpn  struct {
    Enabled     *bool    `json:"enabled,omitempty"`
    HoldingTime *int     `json:"holding_time,omitempty"`
    HostRoutes  []string `json:"host_routes,omitempty"`
}
