package models

import (
	"encoding/json"
)

// RoutingPolicyTermAction represents a RoutingPolicyTermAction struct.
// when used as import policy
type RoutingPolicyTermAction struct {
	Accept       *bool    `json:"accept,omitempty"`
	AddCommunity []string `json:"add_community,omitempty"`
	// for SSR, hub decides how VRF routes are leaked on spoke
	AddTargetVrfs []string `json:"add_target_vrfs,omitempty"`
	// when used as export policy, optional
	Community []string `json:"community,omitempty"`
	// when used as export policy, optional. To exclude certain AS
	ExcludeAsPath    []string `json:"exclude_as_path,omitempty"`
	ExcludeCommunity []string `json:"exclude_community,omitempty"`
	// when used as export policy, optional
	ExportCommunitites []string `json:"export_communitites,omitempty"`
	// optional, for an import policy, local_preference can be changed
	LocalPreference *string `json:"local_preference,omitempty"`
	// when used as export policy, optional. By default, the local AS will be prepended, to change it
	PrependAsPath        []string       `json:"prepend_as_path,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RoutingPolicyTermAction.
// It customizes the JSON marshaling process for RoutingPolicyTermAction objects.
func (r RoutingPolicyTermAction) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RoutingPolicyTermAction object to a map representation for JSON marshaling.
func (r RoutingPolicyTermAction) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
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
	if r.ExportCommunitites != nil {
		structMap["export_communitites"] = r.ExportCommunitites
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "accept", "add_community", "add_target_vrfs", "community", "exclude_as_path", "exclude_community", "export_communitites", "local_preference", "prepend_as_path")
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
	r.ExportCommunitites = temp.ExportCommunitites
	r.LocalPreference = temp.LocalPreference
	r.PrependAsPath = temp.PrependAsPath
	return nil
}

// tempRoutingPolicyTermAction is a temporary struct used for validating the fields of RoutingPolicyTermAction.
type tempRoutingPolicyTermAction struct {
	Accept             *bool    `json:"accept,omitempty"`
	AddCommunity       []string `json:"add_community,omitempty"`
	AddTargetVrfs      []string `json:"add_target_vrfs,omitempty"`
	Community          []string `json:"community,omitempty"`
	ExcludeAsPath      []string `json:"exclude_as_path,omitempty"`
	ExcludeCommunity   []string `json:"exclude_community,omitempty"`
	ExportCommunitites []string `json:"export_communitites,omitempty"`
	LocalPreference    *string  `json:"local_preference,omitempty"`
	PrependAsPath      []string `json:"prepend_as_path,omitempty"`
}
