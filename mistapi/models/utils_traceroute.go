// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UtilsTraceroute represents a UtilsTraceroute struct.
type UtilsTraceroute struct {
	// Host name
	Host *string `json:"host,omitempty"`
	// For SSR, optional, the source to initiate traceroute from
	Network *string `json:"network,omitempty"`
	// only for HA. enum: `node0`, `node1`
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// When `protocol`==`udp`, not supported in SSR. The udp port to use
	Port *int `json:"port,omitempty"`
	// enum: `icmp` (Only supported by AP/MxEdge), `udp`
	Protocol *UtilsTracerouteProtocolEnum `json:"protocol,omitempty"`
	// Not supported in SSR. Maximum time in seconds to wait for the response
	Timeout *int  `json:"timeout,omitempty"`
	UseIpv6 *bool `json:"use_ipv6,omitempty"`
	// For SRX, optional, the source to initiate traceroute from. by default, master VRF/RI is assumed
	Vrf                  *string                `json:"vrf,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsTraceroute,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsTraceroute) String() string {
	return fmt.Sprintf(
		"UtilsTraceroute[Host=%v, Network=%v, Node=%v, Port=%v, Protocol=%v, Timeout=%v, UseIpv6=%v, Vrf=%v, AdditionalProperties=%v]",
		u.Host, u.Network, u.Node, u.Port, u.Protocol, u.Timeout, u.UseIpv6, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsTraceroute.
// It customizes the JSON marshaling process for UtilsTraceroute objects.
func (u UtilsTraceroute) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"host", "network", "node", "port", "protocol", "timeout", "use_ipv6", "vrf"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsTraceroute object to a map representation for JSON marshaling.
func (u UtilsTraceroute) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Host != nil {
		structMap["host"] = u.Host
	}
	if u.Network != nil {
		structMap["network"] = u.Network
	}
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	if u.Port != nil {
		structMap["port"] = u.Port
	}
	if u.Protocol != nil {
		structMap["protocol"] = u.Protocol
	}
	if u.Timeout != nil {
		structMap["timeout"] = u.Timeout
	}
	if u.UseIpv6 != nil {
		structMap["use_ipv6"] = u.UseIpv6
	}
	if u.Vrf != nil {
		structMap["vrf"] = u.Vrf
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsTraceroute.
// It customizes the JSON unmarshaling process for UtilsTraceroute objects.
func (u *UtilsTraceroute) UnmarshalJSON(input []byte) error {
	var temp tempUtilsTraceroute
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "network", "node", "port", "protocol", "timeout", "use_ipv6", "vrf")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Host = temp.Host
	u.Network = temp.Network
	u.Node = temp.Node
	u.Port = temp.Port
	u.Protocol = temp.Protocol
	u.Timeout = temp.Timeout
	u.UseIpv6 = temp.UseIpv6
	u.Vrf = temp.Vrf
	return nil
}

// tempUtilsTraceroute is a temporary struct used for validating the fields of UtilsTraceroute.
type tempUtilsTraceroute struct {
	Host     *string                      `json:"host,omitempty"`
	Network  *string                      `json:"network,omitempty"`
	Node     *HaClusterNodeEnum           `json:"node,omitempty"`
	Port     *int                         `json:"port,omitempty"`
	Protocol *UtilsTracerouteProtocolEnum `json:"protocol,omitempty"`
	Timeout  *int                         `json:"timeout,omitempty"`
	UseIpv6  *bool                        `json:"use_ipv6,omitempty"`
	Vrf      *string                      `json:"vrf,omitempty"`
}
