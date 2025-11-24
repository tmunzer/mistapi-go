// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// OspfPeerStatsSearchResultsItems represents a OspfPeerStatsSearchResultsItems struct.
type OspfPeerStatsSearchResultsItems struct {
	// Activity timer
	DeadTime *int `json:"dead_time,omitempty"`
	// Router MAC address
	Mac *string `json:"mac,omitempty"`
	// Router org ID
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Neighbor address (IP)
	PeerIp *string `json:"peer_ip,omitempty"`
	// Interface name
	PortId *string `json:"port_id,omitempty"`
	// Neighbor priority, 0-255
	Priority *int `json:"priority,omitempty"`
	// Router site ID
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Eg. full, down, 2way, init, exstart, exchange, loading
	State *string `json:"state,omitempty"`
	// Sampling time (in epoch seconds)
	Timestamp *float64 `json:"timestamp,omitempty"`
	// True if state is full
	Up *bool `json:"up,omitempty"`
	// Instance name, e.g. master
	VrfName              *string                `json:"vrf_name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OspfPeerStatsSearchResultsItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OspfPeerStatsSearchResultsItems) String() string {
	return fmt.Sprintf(
		"OspfPeerStatsSearchResultsItems[DeadTime=%v, Mac=%v, OrgId=%v, PeerIp=%v, PortId=%v, Priority=%v, SiteId=%v, State=%v, Timestamp=%v, Up=%v, VrfName=%v, AdditionalProperties=%v]",
		o.DeadTime, o.Mac, o.OrgId, o.PeerIp, o.PortId, o.Priority, o.SiteId, o.State, o.Timestamp, o.Up, o.VrfName, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OspfPeerStatsSearchResultsItems.
// It customizes the JSON marshaling process for OspfPeerStatsSearchResultsItems objects.
func (o OspfPeerStatsSearchResultsItems) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"dead_time", "mac", "org_id", "peer_ip", "port_id", "priority", "site_id", "state", "timestamp", "up", "vrf_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OspfPeerStatsSearchResultsItems object to a map representation for JSON marshaling.
func (o OspfPeerStatsSearchResultsItems) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.DeadTime != nil {
		structMap["dead_time"] = o.DeadTime
	}
	if o.Mac != nil {
		structMap["mac"] = o.Mac
	}
	if o.OrgId != nil {
		structMap["org_id"] = o.OrgId
	}
	if o.PeerIp != nil {
		structMap["peer_ip"] = o.PeerIp
	}
	if o.PortId != nil {
		structMap["port_id"] = o.PortId
	}
	if o.Priority != nil {
		structMap["priority"] = o.Priority
	}
	if o.SiteId != nil {
		structMap["site_id"] = o.SiteId
	}
	if o.State != nil {
		structMap["state"] = o.State
	}
	if o.Timestamp != nil {
		structMap["timestamp"] = o.Timestamp
	}
	if o.Up != nil {
		structMap["up"] = o.Up
	}
	if o.VrfName != nil {
		structMap["vrf_name"] = o.VrfName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OspfPeerStatsSearchResultsItems.
// It customizes the JSON unmarshaling process for OspfPeerStatsSearchResultsItems objects.
func (o *OspfPeerStatsSearchResultsItems) UnmarshalJSON(input []byte) error {
	var temp tempOspfPeerStatsSearchResultsItems
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dead_time", "mac", "org_id", "peer_ip", "port_id", "priority", "site_id", "state", "timestamp", "up", "vrf_name")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.DeadTime = temp.DeadTime
	o.Mac = temp.Mac
	o.OrgId = temp.OrgId
	o.PeerIp = temp.PeerIp
	o.PortId = temp.PortId
	o.Priority = temp.Priority
	o.SiteId = temp.SiteId
	o.State = temp.State
	o.Timestamp = temp.Timestamp
	o.Up = temp.Up
	o.VrfName = temp.VrfName
	return nil
}

// tempOspfPeerStatsSearchResultsItems is a temporary struct used for validating the fields of OspfPeerStatsSearchResultsItems.
type tempOspfPeerStatsSearchResultsItems struct {
	DeadTime  *int       `json:"dead_time,omitempty"`
	Mac       *string    `json:"mac,omitempty"`
	OrgId     *uuid.UUID `json:"org_id,omitempty"`
	PeerIp    *string    `json:"peer_ip,omitempty"`
	PortId    *string    `json:"port_id,omitempty"`
	Priority  *int       `json:"priority,omitempty"`
	SiteId    *uuid.UUID `json:"site_id,omitempty"`
	State     *string    `json:"state,omitempty"`
	Timestamp *float64   `json:"timestamp,omitempty"`
	Up        *bool      `json:"up,omitempty"`
	VrfName   *string    `json:"vrf_name,omitempty"`
}
