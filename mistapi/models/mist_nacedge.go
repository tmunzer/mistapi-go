// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MistNacedge represents a MistNacedge struct.
// Mist NAC Site Survivability settings for the site
type MistNacedge struct {
	// Cache of last auth result; in seconds
	AuthTtl *int `json:"auth_ttl,omitempty"`
	// Site UUIDs whose auth requests are cached by NAC Edges in the cluster
	CachingSiteIds []uuid.UUID `json:"caching_site_ids,omitempty"`
	// Default vlan for all dot1x devices, if different from default_vlan
	DefaultDot1xVlan *string `json:"default_dot1x_vlan,omitempty"`
	// Default vlan to assign for devices not in the cache
	DefaultVlan *string `json:"default_vlan,omitempty"`
	// Whether Mist Site Survivability is enabled for the site
	Enabled *bool `json:"enabled,omitempty"`
	// List of NAC Edges used for the Site Survivability feature
	MxedgeHosts          []string               `json:"mxedge_hosts,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MistNacedge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MistNacedge) String() string {
	return fmt.Sprintf(
		"MistNacedge[AuthTtl=%v, CachingSiteIds=%v, DefaultDot1xVlan=%v, DefaultVlan=%v, Enabled=%v, MxedgeHosts=%v, AdditionalProperties=%v]",
		m.AuthTtl, m.CachingSiteIds, m.DefaultDot1xVlan, m.DefaultVlan, m.Enabled, m.MxedgeHosts, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MistNacedge.
// It customizes the JSON marshaling process for MistNacedge objects.
func (m MistNacedge) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"auth_ttl", "caching_site_ids", "default_dot1x_vlan", "default_vlan", "enabled", "mxedge_hosts"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MistNacedge object to a map representation for JSON marshaling.
func (m MistNacedge) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.AuthTtl != nil {
		structMap["auth_ttl"] = m.AuthTtl
	}
	if m.CachingSiteIds != nil {
		structMap["caching_site_ids"] = m.CachingSiteIds
	}
	if m.DefaultDot1xVlan != nil {
		structMap["default_dot1x_vlan"] = m.DefaultDot1xVlan
	}
	if m.DefaultVlan != nil {
		structMap["default_vlan"] = m.DefaultVlan
	}
	if m.Enabled != nil {
		structMap["enabled"] = m.Enabled
	}
	if m.MxedgeHosts != nil {
		structMap["mxedge_hosts"] = m.MxedgeHosts
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MistNacedge.
// It customizes the JSON unmarshaling process for MistNacedge objects.
func (m *MistNacedge) UnmarshalJSON(input []byte) error {
	var temp tempMistNacedge
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_ttl", "caching_site_ids", "default_dot1x_vlan", "default_vlan", "enabled", "mxedge_hosts")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.AuthTtl = temp.AuthTtl
	m.CachingSiteIds = temp.CachingSiteIds
	m.DefaultDot1xVlan = temp.DefaultDot1xVlan
	m.DefaultVlan = temp.DefaultVlan
	m.Enabled = temp.Enabled
	m.MxedgeHosts = temp.MxedgeHosts
	return nil
}

// tempMistNacedge is a temporary struct used for validating the fields of MistNacedge.
type tempMistNacedge struct {
	AuthTtl          *int        `json:"auth_ttl,omitempty"`
	CachingSiteIds   []uuid.UUID `json:"caching_site_ids,omitempty"`
	DefaultDot1xVlan *string     `json:"default_dot1x_vlan,omitempty"`
	DefaultVlan      *string     `json:"default_vlan,omitempty"`
	Enabled          *bool       `json:"enabled,omitempty"`
	MxedgeHosts      []string    `json:"mxedge_hosts,omitempty"`
}
