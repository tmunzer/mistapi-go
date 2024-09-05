package models

import (
    "encoding/json"
)

// UtilsShowRoute represents a UtilsShowRoute struct.
type UtilsShowRoute struct {
    // duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
    Duration             *int                        `json:"duration,omitempty"`
    // rate at which output will refresh
    Interval             *int                        `json:"interval,omitempty"`
    // IP of the neighbor
    Neighbor             *string                     `json:"neighbor,omitempty"`
    Node                 *HaClusterNode              `json:"node,omitempty"`
    // route prefix
    Prefix               *string                     `json:"prefix,omitempty"`
    // enum: `any`, `bgp`, `direct`, `evpn`, `ospf`, `static`
    Protocol             *UtilsShowRouteProtocolEnum `json:"protocol,omitempty"`
    // if specified, dump bot received and advertised, if not specified, both will be shown
    // * for SSR, show bgp neighbors 10.250.18.202 received-routes/advertised-routes
    // * for SRX and Switches, show route receive_protocol/advertise_protocol bgp 192.168.255.12'
    Route                *string                     `json:"route,omitempty"`
    // VRF name
    Vrf                  *string                     `json:"vrf,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowRoute.
// It customizes the JSON marshaling process for UtilsShowRoute objects.
func (u UtilsShowRoute) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowRoute object to a map representation for JSON marshaling.
func (u UtilsShowRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Duration != nil {
        structMap["duration"] = u.Duration
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.Neighbor != nil {
        structMap["neighbor"] = u.Neighbor
    }
    if u.Node != nil {
        structMap["node"] = u.Node.toMap()
    }
    if u.Prefix != nil {
        structMap["prefix"] = u.Prefix
    }
    if u.Protocol != nil {
        structMap["protocol"] = u.Protocol
    }
    if u.Route != nil {
        structMap["route"] = u.Route
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowRoute.
// It customizes the JSON unmarshaling process for UtilsShowRoute objects.
func (u *UtilsShowRoute) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "interval", "neighbor", "node", "prefix", "protocol", "route", "vrf")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Duration = temp.Duration
    u.Interval = temp.Interval
    u.Neighbor = temp.Neighbor
    u.Node = temp.Node
    u.Prefix = temp.Prefix
    u.Protocol = temp.Protocol
    u.Route = temp.Route
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsShowRoute is a temporary struct used for validating the fields of UtilsShowRoute.
type tempUtilsShowRoute  struct {
    Duration *int                        `json:"duration,omitempty"`
    Interval *int                        `json:"interval,omitempty"`
    Neighbor *string                     `json:"neighbor,omitempty"`
    Node     *HaClusterNode              `json:"node,omitempty"`
    Prefix   *string                     `json:"prefix,omitempty"`
    Protocol *UtilsShowRouteProtocolEnum `json:"protocol,omitempty"`
    Route    *string                     `json:"route,omitempty"`
    Vrf      *string                     `json:"vrf,omitempty"`
}
