// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MarvisConfigAction represents a MarvisConfigAction struct.
// A Marvis-injected config action record
type MarvisConfigAction struct {
	// Admin UUID associated with the config action
	AdminId *uuid.UUID `json:"admin_id,omitempty"`
	// UUID of the config action
	Id *uuid.UUID `json:"id,omitempty"`
	// Device MAC address
	Mac *string `json:"mac,omitempty"`
	// Operation type (e.g. disable_port, enable_port, update_mtu, add_vlans_to_port)
	Op *string `json:"op,omitempty"`
	// Organization UUID
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Port identifier (e.g. ge-0/0/13)
	PortId *string `json:"port_id,omitempty"`
	// Reason for the config action (e.g. rogue_dhcp_server_detected)
	Reason *string `json:"reason,omitempty"`
	// Site UUID
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Source of the config action (e.g. marvis)
	Src *string `json:"src,omitempty"`
	// Timestamp when the config action was recorded, in epoch seconds
	Timestamp *float64 `json:"timestamp,omitempty"`
	// Config type (e.g. wired)
	Type *string `json:"type,omitempty"`
	// List of VLAN IDs involved in the config action
	VlanIds              []int                  `json:"vlan_ids,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisConfigAction,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisConfigAction) String() string {
	return fmt.Sprintf(
		"MarvisConfigAction[AdminId=%v, Id=%v, Mac=%v, Op=%v, OrgId=%v, PortId=%v, Reason=%v, SiteId=%v, Src=%v, Timestamp=%v, Type=%v, VlanIds=%v, AdditionalProperties=%v]",
		m.AdminId, m.Id, m.Mac, m.Op, m.OrgId, m.PortId, m.Reason, m.SiteId, m.Src, m.Timestamp, m.Type, m.VlanIds, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisConfigAction.
// It customizes the JSON marshaling process for MarvisConfigAction objects.
func (m MarvisConfigAction) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"admin_id", "id", "mac", "op", "org_id", "port_id", "reason", "site_id", "src", "timestamp", "type", "vlan_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisConfigAction object to a map representation for JSON marshaling.
func (m MarvisConfigAction) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.AdminId != nil {
		structMap["admin_id"] = m.AdminId
	}
	if m.Id != nil {
		structMap["id"] = m.Id
	}
	if m.Mac != nil {
		structMap["mac"] = m.Mac
	}
	if m.Op != nil {
		structMap["op"] = m.Op
	}
	if m.OrgId != nil {
		structMap["org_id"] = m.OrgId
	}
	if m.PortId != nil {
		structMap["port_id"] = m.PortId
	}
	if m.Reason != nil {
		structMap["reason"] = m.Reason
	}
	if m.SiteId != nil {
		structMap["site_id"] = m.SiteId
	}
	if m.Src != nil {
		structMap["src"] = m.Src
	}
	if m.Timestamp != nil {
		structMap["timestamp"] = m.Timestamp
	}
	if m.Type != nil {
		structMap["type"] = m.Type
	}
	if m.VlanIds != nil {
		structMap["vlan_ids"] = m.VlanIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisConfigAction.
// It customizes the JSON unmarshaling process for MarvisConfigAction objects.
func (m *MarvisConfigAction) UnmarshalJSON(input []byte) error {
	var temp tempMarvisConfigAction
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "admin_id", "id", "mac", "op", "org_id", "port_id", "reason", "site_id", "src", "timestamp", "type", "vlan_ids")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.AdminId = temp.AdminId
	m.Id = temp.Id
	m.Mac = temp.Mac
	m.Op = temp.Op
	m.OrgId = temp.OrgId
	m.PortId = temp.PortId
	m.Reason = temp.Reason
	m.SiteId = temp.SiteId
	m.Src = temp.Src
	m.Timestamp = temp.Timestamp
	m.Type = temp.Type
	m.VlanIds = temp.VlanIds
	return nil
}

// tempMarvisConfigAction is a temporary struct used for validating the fields of MarvisConfigAction.
type tempMarvisConfigAction struct {
	AdminId   *uuid.UUID `json:"admin_id,omitempty"`
	Id        *uuid.UUID `json:"id,omitempty"`
	Mac       *string    `json:"mac,omitempty"`
	Op        *string    `json:"op,omitempty"`
	OrgId     *uuid.UUID `json:"org_id,omitempty"`
	PortId    *string    `json:"port_id,omitempty"`
	Reason    *string    `json:"reason,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	Src       *string    `json:"src,omitempty"`
	Timestamp *float64   `json:"timestamp,omitempty"`
	Type      *string    `json:"type,omitempty"`
	VlanIds   []int      `json:"vlan_ids,omitempty"`
}
