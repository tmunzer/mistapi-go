// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsShowArp represents a UtilsShowArp struct.
// ARP table lookup request for device command output
type UtilsShowArp struct {
	// Refresh duration in seconds; set only when `interval` is nonzero
	Duration *int `json:"duration,omitempty"`
	// Refresh interval in seconds for repeated command output
	Interval *int `json:"interval,omitempty"`
	// Address filter for the ARP table lookup
	Ip *string `json:"ip,omitempty"`
	// HA cluster node selector. enum: `node0`, `node1`
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// Device port identifier filter for the ARP table lookup
	PortId *string `json:"port_id,omitempty"`
	// Routing instance or VRF filter for the ARP table lookup
	Vrf                  *string                `json:"vrf,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowArp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowArp) String() string {
	return fmt.Sprintf(
		"UtilsShowArp[Duration=%v, Interval=%v, Ip=%v, Node=%v, PortId=%v, Vrf=%v, AdditionalProperties=%v]",
		u.Duration, u.Interval, u.Ip, u.Node, u.PortId, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowArp.
// It customizes the JSON marshaling process for UtilsShowArp objects.
func (u UtilsShowArp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"duration", "interval", "ip", "node", "port_id", "vrf"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowArp object to a map representation for JSON marshaling.
func (u UtilsShowArp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Duration != nil {
		structMap["duration"] = u.Duration
	}
	if u.Interval != nil {
		structMap["interval"] = u.Interval
	}
	if u.Ip != nil {
		structMap["ip"] = u.Ip
	}
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	if u.PortId != nil {
		structMap["port_id"] = u.PortId
	}
	if u.Vrf != nil {
		structMap["vrf"] = u.Vrf
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowArp.
// It customizes the JSON unmarshaling process for UtilsShowArp objects.
func (u *UtilsShowArp) UnmarshalJSON(input []byte) error {
	var temp tempUtilsShowArp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "interval", "ip", "node", "port_id", "vrf")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Duration = temp.Duration
	u.Interval = temp.Interval
	u.Ip = temp.Ip
	u.Node = temp.Node
	u.PortId = temp.PortId
	u.Vrf = temp.Vrf
	return nil
}

// tempUtilsShowArp is a temporary struct used for validating the fields of UtilsShowArp.
type tempUtilsShowArp struct {
	Duration *int               `json:"duration,omitempty"`
	Interval *int               `json:"interval,omitempty"`
	Ip       *string            `json:"ip,omitempty"`
	Node     *HaClusterNodeEnum `json:"node,omitempty"`
	PortId   *string            `json:"port_id,omitempty"`
	Vrf      *string            `json:"vrf,omitempty"`
}
