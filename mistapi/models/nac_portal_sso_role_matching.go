// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// NacPortalSsoRoleMatching represents a NacPortalSsoRoleMatching struct.
type NacPortalSsoRoleMatching struct {
	Assigned             *string                `json:"assigned,omitempty"`
	Match                *string                `json:"match,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NacPortalSsoRoleMatching,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacPortalSsoRoleMatching) String() string {
	return fmt.Sprintf(
		"NacPortalSsoRoleMatching[Assigned=%v, Match=%v, AdditionalProperties=%v]",
		n.Assigned, n.Match, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacPortalSsoRoleMatching.
// It customizes the JSON marshaling process for NacPortalSsoRoleMatching objects.
func (n NacPortalSsoRoleMatching) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"assigned", "match"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NacPortalSsoRoleMatching object to a map representation for JSON marshaling.
func (n NacPortalSsoRoleMatching) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Assigned != nil {
		structMap["assigned"] = n.Assigned
	}
	if n.Match != nil {
		structMap["match"] = n.Match
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacPortalSsoRoleMatching.
// It customizes the JSON unmarshaling process for NacPortalSsoRoleMatching objects.
func (n *NacPortalSsoRoleMatching) UnmarshalJSON(input []byte) error {
	var temp tempNacPortalSsoRoleMatching
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "assigned", "match")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.Assigned = temp.Assigned
	n.Match = temp.Match
	return nil
}

// tempNacPortalSsoRoleMatching is a temporary struct used for validating the fields of NacPortalSsoRoleMatching.
type tempNacPortalSsoRoleMatching struct {
	Assigned *string `json:"assigned,omitempty"`
	Match    *string `json:"match,omitempty"`
}
