// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// StatsWanTunnel represents a StatsWanTunnel struct.
// WAN tunnel statistics record returned by tunnel search
type StatsWanTunnel struct {
	// Authentication algorithm negotiated for the tunnel
	AuthAlgo *string `json:"auth_algo,omitempty"`
	// Encryption algorithm negotiated for the tunnel
	EncryptAlgo *string `json:"encrypt_algo,omitempty"`
	// IKE version used to establish the tunnel
	IkeVersion *string `json:"ike_version,omitempty"`
	// Local IP address used by the tunnel
	Ip *string `json:"ip,omitempty"`
	// Most recent reason the tunnel went down
	LastEvent *string `json:"last_event,omitempty"`
	// Router MAC address reporting the tunnel statistics
	Mac *string `json:"mac,omitempty"`
	// HA node handling the tunnel, such as node0 or node1
	Node *string `json:"node,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Hostname or configured peer host of the remote tunnel endpoint
	PeerHost *string `json:"peer_host,omitempty"`
	// IP address of the remote tunnel endpoint
	PeerIp string `json:"peer_ip"`
	// Relative preference assigned to the tunnel. enum: `primary`, `secondary`
	Priority *TunnelPriorityEnum `json:"priority,omitempty"`
	// Tunnel protocol used for the connection. enum: `gre`, `ipsec`
	Protocol *WanTunnelProtocolEnum `json:"protocol,omitempty"`
	// Amount of traffic received since connection
	RxBytes Optional[int64] `json:"rx_bytes"`
	// Amount of packets received since connection
	RxPkts Optional[int64] `json:"rx_pkts"`
	// Unique identifier of a Mist site
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Name of the Mist-managed WAN tunnel
	TunnelName *string `json:"tunnel_name,omitempty"`
	// Amount of traffic sent since connection
	TxBytes Optional[int64] `json:"tx_bytes"`
	// Amount of packets sent since connection
	TxPkts Optional[int64] `json:"tx_pkts"`
	// Indicates whether the tunnel is currently up
	Up *bool `json:"up,omitempty"`
	// Duration since the tunnel security association was established
	Uptime *int `json:"uptime,omitempty"`
	// Name of the WAN interface carrying the tunnel
	WanName              *string                `json:"wan_name,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWanTunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWanTunnel) String() string {
	return fmt.Sprintf(
		"StatsWanTunnel[AuthAlgo=%v, EncryptAlgo=%v, IkeVersion=%v, Ip=%v, LastEvent=%v, Mac=%v, Node=%v, OrgId=%v, PeerHost=%v, PeerIp=%v, Priority=%v, Protocol=%v, RxBytes=%v, RxPkts=%v, SiteId=%v, TunnelName=%v, TxBytes=%v, TxPkts=%v, Up=%v, Uptime=%v, WanName=%v, AdditionalProperties=%v]",
		s.AuthAlgo, s.EncryptAlgo, s.IkeVersion, s.Ip, s.LastEvent, s.Mac, s.Node, s.OrgId, s.PeerHost, s.PeerIp, s.Priority, s.Protocol, s.RxBytes, s.RxPkts, s.SiteId, s.TunnelName, s.TxBytes, s.TxPkts, s.Up, s.Uptime, s.WanName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsWanTunnel.
// It customizes the JSON marshaling process for StatsWanTunnel objects.
func (s StatsWanTunnel) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "mac", "node", "org_id", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "site_id", "tunnel_name", "tx_bytes", "tx_pkts", "up", "uptime", "wan_name"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsWanTunnel object to a map representation for JSON marshaling.
func (s StatsWanTunnel) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AuthAlgo != nil {
		structMap["auth_algo"] = s.AuthAlgo
	}
	if s.EncryptAlgo != nil {
		structMap["encrypt_algo"] = s.EncryptAlgo
	}
	if s.IkeVersion != nil {
		structMap["ike_version"] = s.IkeVersion
	}
	if s.Ip != nil {
		structMap["ip"] = s.Ip
	}
	if s.LastEvent != nil {
		structMap["last_event"] = s.LastEvent
	}
	if s.Mac != nil {
		structMap["mac"] = s.Mac
	}
	if s.Node != nil {
		structMap["node"] = s.Node
	}
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	if s.PeerHost != nil {
		structMap["peer_host"] = s.PeerHost
	}
	structMap["peer_ip"] = s.PeerIp
	if s.Priority != nil {
		structMap["priority"] = s.Priority
	}
	if s.Protocol != nil {
		structMap["protocol"] = s.Protocol
	}
	if s.RxBytes.IsValueSet() {
		if s.RxBytes.Value() != nil {
			structMap["rx_bytes"] = s.RxBytes.Value()
		} else {
			structMap["rx_bytes"] = nil
		}
	}
	if s.RxPkts.IsValueSet() {
		if s.RxPkts.Value() != nil {
			structMap["rx_pkts"] = s.RxPkts.Value()
		} else {
			structMap["rx_pkts"] = nil
		}
	}
	if s.SiteId != nil {
		structMap["site_id"] = s.SiteId
	}
	if s.TunnelName != nil {
		structMap["tunnel_name"] = s.TunnelName
	}
	if s.TxBytes.IsValueSet() {
		if s.TxBytes.Value() != nil {
			structMap["tx_bytes"] = s.TxBytes.Value()
		} else {
			structMap["tx_bytes"] = nil
		}
	}
	if s.TxPkts.IsValueSet() {
		if s.TxPkts.Value() != nil {
			structMap["tx_pkts"] = s.TxPkts.Value()
		} else {
			structMap["tx_pkts"] = nil
		}
	}
	if s.Up != nil {
		structMap["up"] = s.Up
	}
	if s.Uptime != nil {
		structMap["uptime"] = s.Uptime
	}
	if s.WanName != nil {
		structMap["wan_name"] = s.WanName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWanTunnel.
// It customizes the JSON unmarshaling process for StatsWanTunnel objects.
func (s *StatsWanTunnel) UnmarshalJSON(input []byte) error {
	var temp tempStatsWanTunnel
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "mac", "node", "org_id", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "site_id", "tunnel_name", "tx_bytes", "tx_pkts", "up", "uptime", "wan_name")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AuthAlgo = temp.AuthAlgo
	s.EncryptAlgo = temp.EncryptAlgo
	s.IkeVersion = temp.IkeVersion
	s.Ip = temp.Ip
	s.LastEvent = temp.LastEvent
	s.Mac = temp.Mac
	s.Node = temp.Node
	s.OrgId = temp.OrgId
	s.PeerHost = temp.PeerHost
	s.PeerIp = *temp.PeerIp
	s.Priority = temp.Priority
	s.Protocol = temp.Protocol
	s.RxBytes = temp.RxBytes
	s.RxPkts = temp.RxPkts
	s.SiteId = temp.SiteId
	s.TunnelName = temp.TunnelName
	s.TxBytes = temp.TxBytes
	s.TxPkts = temp.TxPkts
	s.Up = temp.Up
	s.Uptime = temp.Uptime
	s.WanName = temp.WanName
	return nil
}

// tempStatsWanTunnel is a temporary struct used for validating the fields of StatsWanTunnel.
type tempStatsWanTunnel struct {
	AuthAlgo    *string                `json:"auth_algo,omitempty"`
	EncryptAlgo *string                `json:"encrypt_algo,omitempty"`
	IkeVersion  *string                `json:"ike_version,omitempty"`
	Ip          *string                `json:"ip,omitempty"`
	LastEvent   *string                `json:"last_event,omitempty"`
	Mac         *string                `json:"mac,omitempty"`
	Node        *string                `json:"node,omitempty"`
	OrgId       *uuid.UUID             `json:"org_id,omitempty"`
	PeerHost    *string                `json:"peer_host,omitempty"`
	PeerIp      *string                `json:"peer_ip"`
	Priority    *TunnelPriorityEnum    `json:"priority,omitempty"`
	Protocol    *WanTunnelProtocolEnum `json:"protocol,omitempty"`
	RxBytes     Optional[int64]        `json:"rx_bytes"`
	RxPkts      Optional[int64]        `json:"rx_pkts"`
	SiteId      *uuid.UUID             `json:"site_id,omitempty"`
	TunnelName  *string                `json:"tunnel_name,omitempty"`
	TxBytes     Optional[int64]        `json:"tx_bytes"`
	TxPkts      Optional[int64]        `json:"tx_pkts"`
	Up          *bool                  `json:"up,omitempty"`
	Uptime      *int                   `json:"uptime,omitempty"`
	WanName     *string                `json:"wan_name,omitempty"`
}

func (s *tempStatsWanTunnel) validate() error {
	var errs []string
	if s.PeerIp == nil {
		errs = append(errs, "required field `peer_ip` is missing for type `stats_wan_tunnel`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
