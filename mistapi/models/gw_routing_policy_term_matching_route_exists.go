// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GwRoutingPolicyTermMatchingRouteExists represents a GwRoutingPolicyTermMatchingRouteExists struct.
type GwRoutingPolicyTermMatchingRouteExists struct {
	Route *string `json:"route,omitempty"`
	// Name of the vrf instance, it can also be the name of the VPN or wan if they
	VrfName              *string                `json:"vrf_name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GwRoutingPolicyTermMatchingRouteExists,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GwRoutingPolicyTermMatchingRouteExists) String() string {
	return fmt.Sprintf(
		"GwRoutingPolicyTermMatchingRouteExists[Route=%v, VrfName=%v, AdditionalProperties=%v]",
		g.Route, g.VrfName, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GwRoutingPolicyTermMatchingRouteExists.
// It customizes the JSON marshaling process for GwRoutingPolicyTermMatchingRouteExists objects.
func (g GwRoutingPolicyTermMatchingRouteExists) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"route", "vrf_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GwRoutingPolicyTermMatchingRouteExists object to a map representation for JSON marshaling.
func (g GwRoutingPolicyTermMatchingRouteExists) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Route != nil {
		structMap["route"] = g.Route
	}
	if g.VrfName != nil {
		structMap["vrf_name"] = g.VrfName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GwRoutingPolicyTermMatchingRouteExists.
// It customizes the JSON unmarshaling process for GwRoutingPolicyTermMatchingRouteExists objects.
func (g *GwRoutingPolicyTermMatchingRouteExists) UnmarshalJSON(input []byte) error {
	var temp tempGwRoutingPolicyTermMatchingRouteExists
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "route", "vrf_name")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Route = temp.Route
	g.VrfName = temp.VrfName
	return nil
}

// tempGwRoutingPolicyTermMatchingRouteExists is a temporary struct used for validating the fields of GwRoutingPolicyTermMatchingRouteExists.
type tempGwRoutingPolicyTermMatchingRouteExists struct {
	Route   *string `json:"route,omitempty"`
	VrfName *string `json:"vrf_name,omitempty"`
}
