// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// RoutingPolicyTermAction represents a RoutingPolicyTermAction struct.
// When used as import policy
type RoutingPolicyTermAction struct {
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
	// Optional, for an import policy, local_preference can be changed
	LocalPreference *string `json:"local_preference,omitempty"`
	// When used as export policy, optional. By default, the local AS will be prepended, to change it
	PrependAsPath        []string               `json:"prepend_as_path,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RoutingPolicyTermAction,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RoutingPolicyTermAction) String() string {
	return fmt.Sprintf(
		"RoutingPolicyTermAction[Accept=%v, AddCommunity=%v, AddTargetVrfs=%v, Community=%v, ExcludeAsPath=%v, ExcludeCommunity=%v, ExportCommunities=%v, LocalPreference=%v, PrependAsPath=%v, AdditionalProperties=%v]",
		r.Accept, r.AddCommunity, r.AddTargetVrfs, r.Community, r.ExcludeAsPath, r.ExcludeCommunity, r.ExportCommunities, r.LocalPreference, r.PrependAsPath, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTermAction.
// It customizes the JSON marshaling process for RoutingPolicyTermAction objects.
func (r RoutingPolicyTermAction) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"accept", "add_community", "add_target_vrfs", "community", "exclude_as_path", "exclude_community", "export_communities", "local_preference", "prepend_as_path"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTermAction object to a map representation for JSON marshaling.
func (r RoutingPolicyTermAction) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Accept != nil {
		structMap["accept"] = r.Accept
	}
	if r.AddCommunity != nil {
		structMap["add_community"] = r.AddCommunity
	}
	if r.AddTargetVrfs != nil {
		structMap["add_target_vrfs"] = r.AddTargetVrfs
	}
	if r.Community != nil {
		structMap["community"] = r.Community
	}
	if r.ExcludeAsPath != nil {
		structMap["exclude_as_path"] = r.ExcludeAsPath
	}
	if r.ExcludeCommunity != nil {
		structMap["exclude_community"] = r.ExcludeCommunity
	}
	if r.ExportCommunities != nil {
		structMap["export_communities"] = r.ExportCommunities
	}
	if r.LocalPreference != nil {
		structMap["local_preference"] = r.LocalPreference
	}
	if r.PrependAsPath != nil {
		structMap["prepend_as_path"] = r.PrependAsPath
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RoutingPolicyTermAction.
// It customizes the JSON unmarshaling process for RoutingPolicyTermAction objects.
func (r *RoutingPolicyTermAction) UnmarshalJSON(input []byte) error {
	var temp tempRoutingPolicyTermAction
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accept", "add_community", "add_target_vrfs", "community", "exclude_as_path", "exclude_community", "export_communities", "local_preference", "prepend_as_path")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Accept = temp.Accept
	r.AddCommunity = temp.AddCommunity
	r.AddTargetVrfs = temp.AddTargetVrfs
	r.Community = temp.Community
	r.ExcludeAsPath = temp.ExcludeAsPath
	r.ExcludeCommunity = temp.ExcludeCommunity
	r.ExportCommunities = temp.ExportCommunities
	r.LocalPreference = temp.LocalPreference
	r.PrependAsPath = temp.PrependAsPath
	return nil
}

// tempRoutingPolicyTermAction is a temporary struct used for validating the fields of RoutingPolicyTermAction.
type tempRoutingPolicyTermAction struct {
	Accept            *bool    `json:"accept,omitempty"`
	AddCommunity      []string `json:"add_community,omitempty"`
	AddTargetVrfs     []string `json:"add_target_vrfs,omitempty"`
	Community         []string `json:"community,omitempty"`
	ExcludeAsPath     []string `json:"exclude_as_path,omitempty"`
	ExcludeCommunity  []string `json:"exclude_community,omitempty"`
	ExportCommunities []string `json:"export_communities,omitempty"`
	LocalPreference   *string  `json:"local_preference,omitempty"`
	PrependAsPath     []string `json:"prepend_as_path,omitempty"`
}
