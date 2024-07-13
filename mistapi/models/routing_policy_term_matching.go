package models

import (
    "encoding/json"
)

// RoutingPolicyTermMatching represents a RoutingPolicyTermMatching struct.
// zero or more criteria/filter can be specified to match the term, all criteria have to be met
type RoutingPolicyTermMatching struct {
    // takes regular expression
    AsPath               []string                              `json:"as_path,omitempty"`
    Community            []string                              `json:"community,omitempty"`
    Network              []string                              `json:"network,omitempty"`
    // zero or more criteria/filter can be specified to match the term, all criteria have to be met
    Prefix               []string                              `json:"prefix,omitempty"`
    // `direct`, `bgp`, `osp`, ...
    Protocol             []string                              `json:"protocol,omitempty"`
    RouteExists          *RoutingPolicyTermMatchingRouteExists `json:"route_exists,omitempty"`
    // overlay-facing criteria (used for bgp_config where via=vpn)
    VpnNeighborMac       []string                              `json:"vpn_neighbor_mac,omitempty"`
    // overlay-facing criteria (used for bgp_config where via=vpn)
    // ordered-
    VpnPath              []string                              `json:"vpn_path,omitempty"`
    VpnPathSla           *RoutingPolicyTermMatchingVpnPathSla  `json:"vpn_path_sla,omitempty"`
    AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTermMatching.
// It customizes the JSON marshaling process for RoutingPolicyTermMatching objects.
func (r RoutingPolicyTermMatching) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTermMatching object to a map representation for JSON marshaling.
func (r RoutingPolicyTermMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.AsPath != nil {
        structMap["as_path"] = r.AsPath
    }
    if r.Community != nil {
        structMap["community"] = r.Community
    }
    if r.Network != nil {
        structMap["network"] = r.Network
    }
    if r.Prefix != nil {
        structMap["prefix"] = r.Prefix
    }
    if r.Protocol != nil {
        structMap["protocol"] = r.Protocol
    }
    if r.RouteExists != nil {
        structMap["route_exists"] = r.RouteExists.toMap()
    }
    if r.VpnNeighborMac != nil {
        structMap["vpn_neighbor_mac"] = r.VpnNeighborMac
    }
    if r.VpnPath != nil {
        structMap["vpn_path"] = r.VpnPath
    }
    if r.VpnPathSla != nil {
        structMap["vpn_path_sla"] = r.VpnPathSla.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RoutingPolicyTermMatching.
// It customizes the JSON unmarshaling process for RoutingPolicyTermMatching objects.
func (r *RoutingPolicyTermMatching) UnmarshalJSON(input []byte) error {
    var temp routingPolicyTermMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "as_path", "community", "network", "prefix", "protocol", "route_exists", "vpn_neighbor_mac", "vpn_path", "vpn_path_sla")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.AsPath = temp.AsPath
    r.Community = temp.Community
    r.Network = temp.Network
    r.Prefix = temp.Prefix
    r.Protocol = temp.Protocol
    r.RouteExists = temp.RouteExists
    r.VpnNeighborMac = temp.VpnNeighborMac
    r.VpnPath = temp.VpnPath
    r.VpnPathSla = temp.VpnPathSla
    return nil
}

// routingPolicyTermMatching is a temporary struct used for validating the fields of RoutingPolicyTermMatching.
type routingPolicyTermMatching  struct {
    AsPath         []string                              `json:"as_path,omitempty"`
    Community      []string                              `json:"community,omitempty"`
    Network        []string                              `json:"network,omitempty"`
    Prefix         []string                              `json:"prefix,omitempty"`
    Protocol       []string                              `json:"protocol,omitempty"`
    RouteExists    *RoutingPolicyTermMatchingRouteExists `json:"route_exists,omitempty"`
    VpnNeighborMac []string                              `json:"vpn_neighbor_mac,omitempty"`
    VpnPath        []string                              `json:"vpn_path,omitempty"`
    VpnPathSla     *RoutingPolicyTermMatchingVpnPathSla  `json:"vpn_path_sla,omitempty"`
}