// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// GwRoutingPolicyTermAction represents a GwRoutingPolicyTermAction struct.
// When used as import policy
type GwRoutingPolicyTermAction struct {
	Accept       *bool    `json:"accept,omitempty"`
	AddCommunity []string `json:"add_community,omitempty"`
	// For SSR, hub decides how VRF routes are leaked on spoke
	AddTargetVrfs []string `json:"add_target_vrfs,omitempty"`
	// When used as export policy, optional
	Community []string `json:"community,omitempty"`
	// When used as export policy, optional. To exclude certain AS
	ExcludeAsPath    []string `json:"exclude_as_path,omitempty"`
	ExcludeCommunity []string `json:"exclude_community,omitempty"`
	// When used as export policy, optional
	ExportCommunities []string `json:"export_communities,omitempty"`
	// Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`)
	LocalPreference *RoutingPolicyLocalPreference `json:"local_preference,omitempty"`
	// When used as export policy, optional. By default, the local AS will be prepended, to change it. Can be a Variable (e.g. `{{as_path}}`)
	PrependAsPath        []string               `json:"prepend_as_path,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GwRoutingPolicyTermAction,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GwRoutingPolicyTermAction) String() string {
	return fmt.Sprintf(
		"GwRoutingPolicyTermAction[Accept=%v, AddCommunity=%v, AddTargetVrfs=%v, Community=%v, ExcludeAsPath=%v, ExcludeCommunity=%v, ExportCommunities=%v, LocalPreference=%v, PrependAsPath=%v, AdditionalProperties=%v]",
		g.Accept, g.AddCommunity, g.AddTargetVrfs, g.Community, g.ExcludeAsPath, g.ExcludeCommunity, g.ExportCommunities, g.LocalPreference, g.PrependAsPath, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GwRoutingPolicyTermAction.
// It customizes the JSON marshaling process for GwRoutingPolicyTermAction objects.
func (g GwRoutingPolicyTermAction) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(g.AdditionalProperties,
		"accept", "add_community", "add_target_vrfs", "community", "exclude_as_path", "exclude_community", "export_communities", "local_preference", "prepend_as_path"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(g.toMap())
}

// toMap converts the GwRoutingPolicyTermAction object to a map representation for JSON marshaling.
func (g GwRoutingPolicyTermAction) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, g.AdditionalProperties)
	if g.Accept != nil {
		structMap["accept"] = g.Accept
	}
	if g.AddCommunity != nil {
		structMap["add_community"] = g.AddCommunity
	}
	if g.AddTargetVrfs != nil {
		structMap["add_target_vrfs"] = g.AddTargetVrfs
	}
	if g.Community != nil {
		structMap["community"] = g.Community
	}
	if g.ExcludeAsPath != nil {
		structMap["exclude_as_path"] = g.ExcludeAsPath
	}
	if g.ExcludeCommunity != nil {
		structMap["exclude_community"] = g.ExcludeCommunity
	}
	if g.ExportCommunities != nil {
		structMap["export_communities"] = g.ExportCommunities
	}
	if g.LocalPreference != nil {
		structMap["local_preference"] = g.LocalPreference.toMap()
	}
	if g.PrependAsPath != nil {
		structMap["prepend_as_path"] = g.PrependAsPath
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GwRoutingPolicyTermAction.
// It customizes the JSON unmarshaling process for GwRoutingPolicyTermAction objects.
func (g *GwRoutingPolicyTermAction) UnmarshalJSON(input []byte) error {
	var temp tempGwRoutingPolicyTermAction
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accept", "add_community", "add_target_vrfs", "community", "exclude_as_path", "exclude_community", "export_communities", "local_preference", "prepend_as_path")
	if err != nil {
		return err
	}
	g.AdditionalProperties = additionalProperties

	g.Accept = temp.Accept
	g.AddCommunity = temp.AddCommunity
	g.AddTargetVrfs = temp.AddTargetVrfs
	g.Community = temp.Community
	g.ExcludeAsPath = temp.ExcludeAsPath
	g.ExcludeCommunity = temp.ExcludeCommunity
	g.ExportCommunities = temp.ExportCommunities
	g.LocalPreference = temp.LocalPreference
	g.PrependAsPath = temp.PrependAsPath
	return nil
}

// tempGwRoutingPolicyTermAction is a temporary struct used for validating the fields of GwRoutingPolicyTermAction.
type tempGwRoutingPolicyTermAction struct {
	Accept            *bool                         `json:"accept,omitempty"`
	AddCommunity      []string                      `json:"add_community,omitempty"`
	AddTargetVrfs     []string                      `json:"add_target_vrfs,omitempty"`
	Community         []string                      `json:"community,omitempty"`
	ExcludeAsPath     []string                      `json:"exclude_as_path,omitempty"`
	ExcludeCommunity  []string                      `json:"exclude_community,omitempty"`
	ExportCommunities []string                      `json:"export_communities,omitempty"`
	LocalPreference   *RoutingPolicyLocalPreference `json:"local_preference,omitempty"`
	PrependAsPath     []string                      `json:"prepend_as_path,omitempty"`
}
