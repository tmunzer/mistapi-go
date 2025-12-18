// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GwRoutingPolicyTermMatching represents a GwRoutingPolicyTermMatching struct.
// zero or more criteria/filter can be specified to match the term, all criteria have to be met
type GwRoutingPolicyTermMatching struct {
	AsPath    []BgpAs  `json:"as_path,omitempty"`
	Community []string `json:"community,omitempty"`
	Network   []string `json:"network,omitempty"`
	// zero or more criteria/filter can be specified to match the term, all criteria have to be met
	Prefix      []string                                  `json:"prefix,omitempty"`
	Protocol    []GwRoutingPolicyTermMatchingProtocolEnum `json:"protocol,omitempty"`
	RouteExists *GwRoutingPolicyTermMatchingRouteExists   `json:"route_exists,omitempty"`
	// overlay-facing criteria (used for bgp_config where via=vpn)
	VpnNeighborMac []string `json:"vpn_neighbor_mac,omitempty"`
	// overlay-facing criteria (used for bgp_config where via=vpn). ordered-
	VpnPath              []string                               `json:"vpn_path,omitempty"`
	VpnPathSla           *GwRoutingPolicyTermMatchingVpnPathSla `json:"vpn_path_sla,omitempty"`
	AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for GwRoutingPolicyTermMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GwRoutingPolicyTermMatching) String() string {
	return fmt.Sprintf(
		"GwRoutingPolicyTermMatching[AsPath=%v, Community=%v, Network=%v, Prefix=%v, Protocol=%v, RouteExists=%v, VpnNeighborMac=%v, VpnPath=%v, VpnPathSla=%v, AdditionalProperties=%v]",
		g.AsPath, g.Community, g.Network, g.Prefix, g.Protocol, g.RouteExists, g.VpnNeighborMac, g.VpnPath, g.VpnPathSla, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GwRoutingPolicyTermMatching.
// It customizes the JSON marshaling process for GwRoutingPolicyTermMatching objects.
func (g GwRoutingPolicyTermMatching) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"as_path", "community", "network", "prefix", "protocol", "route_exists", "vpn_neighbor_mac", "vpn_path", "vpn_path_sla"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GwRoutingPolicyTermMatching object to a map representation for JSON marshaling.
func (g GwRoutingPolicyTermMatching) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.AsPath != nil {
		structMap["as_path"] = g.AsPath
	}
	if g.Community != nil {
		structMap["community"] = g.Community
	}
	if g.Network != nil {
		structMap["network"] = g.Network
	}
	if g.Prefix != nil {
		structMap["prefix"] = g.Prefix
	}
	if g.Protocol != nil {
		structMap["protocol"] = g.Protocol
	}
	if g.RouteExists != nil {
		structMap["route_exists"] = g.RouteExists.toMap()
	}
	if g.VpnNeighborMac != nil {
		structMap["vpn_neighbor_mac"] = g.VpnNeighborMac
	}
	if g.VpnPath != nil {
		structMap["vpn_path"] = g.VpnPath
	}
	if g.VpnPathSla != nil {
		structMap["vpn_path_sla"] = g.VpnPathSla.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GwRoutingPolicyTermMatching.
// It customizes the JSON unmarshaling process for GwRoutingPolicyTermMatching objects.
func (g *GwRoutingPolicyTermMatching) UnmarshalJSON(input []byte) error {
	var temp tempGwRoutingPolicyTermMatching
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "as_path", "community", "network", "prefix", "protocol", "route_exists", "vpn_neighbor_mac", "vpn_path", "vpn_path_sla")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.AsPath = temp.AsPath
	g.Community = temp.Community
	g.Network = temp.Network
	g.Prefix = temp.Prefix
	g.Protocol = temp.Protocol
	g.RouteExists = temp.RouteExists
	g.VpnNeighborMac = temp.VpnNeighborMac
	g.VpnPath = temp.VpnPath
	g.VpnPathSla = temp.VpnPathSla
	return nil
}

// tempGwRoutingPolicyTermMatching is a temporary struct used for validating the fields of GwRoutingPolicyTermMatching.
type tempGwRoutingPolicyTermMatching struct {
	AsPath         []BgpAs                                   `json:"as_path,omitempty"`
	Community      []string                                  `json:"community,omitempty"`
	Network        []string                                  `json:"network,omitempty"`
	Prefix         []string                                  `json:"prefix,omitempty"`
	Protocol       []GwRoutingPolicyTermMatchingProtocolEnum `json:"protocol,omitempty"`
	RouteExists    *GwRoutingPolicyTermMatchingRouteExists   `json:"route_exists,omitempty"`
	VpnNeighborMac []string                                  `json:"vpn_neighbor_mac,omitempty"`
	VpnPath        []string                                  `json:"vpn_path,omitempty"`
	VpnPathSla     *GwRoutingPolicyTermMatchingVpnPathSla    `json:"vpn_path_sla,omitempty"`
}
