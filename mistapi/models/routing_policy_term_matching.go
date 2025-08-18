// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// RoutingPolicyTermMatching represents a RoutingPolicyTermMatching struct.
// zero or more criteria/filter can be specified to match the term, all criteria have to be met
type RoutingPolicyTermMatching struct {
	// takes regular expression
	AsPath    []string `json:"as_path,omitempty"`
	Community []string `json:"community,omitempty"`
	Network   []string `json:"network,omitempty"`
	// zero or more criteria/filter can be specified to match the term, all criteria have to be met
	Prefix []string `json:"prefix,omitempty"`
	// `direct`, `bgp`, `osp`, `static`, `aggregate`...
	Protocol    []string                              `json:"protocol,omitempty"`
	RouteExists *RoutingPolicyTermMatchingRouteExists `json:"route_exists,omitempty"`
	// overlay-facing criteria (used for bgp_config where via=vpn)
	VpnNeighborMac []string `json:"vpn_neighbor_mac,omitempty"`
	// overlay-facing criteria (used for bgp_config where via=vpn). ordered-
	VpnPath              []string                             `json:"vpn_path,omitempty"`
	VpnPathSla           *RoutingPolicyTermMatchingVpnPathSla `json:"vpn_path_sla,omitempty"`
	AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for RoutingPolicyTermMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RoutingPolicyTermMatching) String() string {
	return fmt.Sprintf(
		"RoutingPolicyTermMatching[AsPath=%v, Community=%v, Network=%v, Prefix=%v, Protocol=%v, RouteExists=%v, VpnNeighborMac=%v, VpnPath=%v, VpnPathSla=%v, AdditionalProperties=%v]",
		r.AsPath, r.Community, r.Network, r.Prefix, r.Protocol, r.RouteExists, r.VpnNeighborMac, r.VpnPath, r.VpnPathSla, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTermMatching.
// It customizes the JSON marshaling process for RoutingPolicyTermMatching objects.
func (r RoutingPolicyTermMatching) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"as_path", "community", "network", "prefix", "protocol", "route_exists", "vpn_neighbor_mac", "vpn_path", "vpn_path_sla"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTermMatching object to a map representation for JSON marshaling.
func (r RoutingPolicyTermMatching) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
	var temp tempRoutingPolicyTermMatching
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "as_path", "community", "network", "prefix", "protocol", "route_exists", "vpn_neighbor_mac", "vpn_path", "vpn_path_sla")
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

// tempRoutingPolicyTermMatching is a temporary struct used for validating the fields of RoutingPolicyTermMatching.
type tempRoutingPolicyTermMatching struct {
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
