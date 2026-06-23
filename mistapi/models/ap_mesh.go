// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApMesh represents a ApMesh struct.
// Wireless mesh settings for an access point
type ApMesh struct {
	// List of bands that mesh should use. For relay, the first viable one will be picked. enum: `24`, `5`, `6`
	Bands []Dot11BandEnum `json:"bands,omitempty"`
	// Whether mesh is enabled on this AP
	Enabled *bool `json:"enabled,omitempty"`
	// Mesh group, base AP(s) will only allow remote AP(s) in the same mesh group to join, 1-9, optional
	Group Optional[int] `json:"group"`
	// Mesh role for this AP, either base or remote. enum: `base`, `remote`
	Role *ApMeshRoleEnum `json:"role,omitempty"`
	// Whether to use WPA3 on the 5 GHz band for mesh links
	UseWpa3On5           *bool                  `json:"use_wpa3_on_5,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApMesh,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApMesh) String() string {
	return fmt.Sprintf(
		"ApMesh[Bands=%v, Enabled=%v, Group=%v, Role=%v, UseWpa3On5=%v, AdditionalProperties=%v]",
		a.Bands, a.Enabled, a.Group, a.Role, a.UseWpa3On5, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApMesh.
// It customizes the JSON marshaling process for ApMesh objects.
func (a ApMesh) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"bands", "enabled", "group", "role", "use_wpa3_on_5"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApMesh object to a map representation for JSON marshaling.
func (a ApMesh) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Bands != nil {
		structMap["bands"] = a.Bands
	}
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.Group.IsValueSet() {
		if a.Group.Value() != nil {
			structMap["group"] = a.Group.Value()
		} else {
			structMap["group"] = nil
		}
	}
	if a.Role != nil {
		structMap["role"] = a.Role
	}
	if a.UseWpa3On5 != nil {
		structMap["use_wpa3_on_5"] = a.UseWpa3On5
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApMesh.
// It customizes the JSON unmarshaling process for ApMesh objects.
func (a *ApMesh) UnmarshalJSON(input []byte) error {
	var temp tempApMesh
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bands", "enabled", "group", "role", "use_wpa3_on_5")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Bands = temp.Bands
	a.Enabled = temp.Enabled
	a.Group = temp.Group
	a.Role = temp.Role
	a.UseWpa3On5 = temp.UseWpa3On5
	return nil
}

// tempApMesh is a temporary struct used for validating the fields of ApMesh.
type tempApMesh struct {
	Bands      []Dot11BandEnum `json:"bands,omitempty"`
	Enabled    *bool           `json:"enabled,omitempty"`
	Group      Optional[int]   `json:"group"`
	Role       *ApMeshRoleEnum `json:"role,omitempty"`
	UseWpa3On5 *bool           `json:"use_wpa3_on_5,omitempty"`
}
