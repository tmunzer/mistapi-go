// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// StatsGatewayVpnPeer represents a StatsGatewayVpnPeer struct.
// VPN peer path statistics reported by a gateway
type StatsGatewayVpnPeer struct {
	// Whether this VPN peer path is the active redundant path
	IsActive *bool `json:"is_active,omitempty"`
	// Last sampled VPN peer jitter, in milliseconds
	Jitter *float64 `json:"jitter,omitempty"`
	// Timestamp indicating when the entity was last seen
	LastSeen Optional[float64] `json:"last_seen"`
	// Last sampled VPN peer latency, in milliseconds
	Latency *float64 `json:"latency,omitempty"`
	// Packet loss observed for the VPN peer, as a percentage
	Loss *float64 `json:"loss,omitempty"`
	// Mean Opinion Score for VPN link quality, from 0 to 5
	Mos *float64 `json:"mos,omitempty"`
	// Maximum transmission unit for the VPN peer path, in bytes
	Mtu *int `json:"mtu,omitempty"`
	// Peer router MAC address for the VPN link
	PeerMac *string `json:"peer_mac,omitempty"`
	// Peer router interface identifier for the VPN link
	PeerPortId *string `json:"peer_port_id,omitempty"`
	// Peer router name reported for the VPN link
	PeerRouterName *string `json:"peer_router_name,omitempty"`
	// Unique identifier of a Mist site
	PeerSiteId *uuid.UUID `json:"peer_site_id,omitempty"`
	// Local router interface identifier for the VPN link
	PortId *string `json:"port_id,omitempty"`
	// Local router name reported for the VPN link
	RouterName *string `json:"router_name,omitempty"`
	// VPN implementation type for the peer, such as `ipsec` for SRX or `svr` for SSR
	Type *string `json:"type,omitempty"`
	// Whether the VPN peer is currently up
	Up *bool `json:"up,omitempty"`
	// Gateway-reported VPN peer uptime value, in seconds
	Uptime               *int                   `json:"uptime,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsGatewayVpnPeer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsGatewayVpnPeer) String() string {
	return fmt.Sprintf(
		"StatsGatewayVpnPeer[IsActive=%v, Jitter=%v, LastSeen=%v, Latency=%v, Loss=%v, Mos=%v, Mtu=%v, PeerMac=%v, PeerPortId=%v, PeerRouterName=%v, PeerSiteId=%v, PortId=%v, RouterName=%v, Type=%v, Up=%v, Uptime=%v, AdditionalProperties=%v]",
		s.IsActive, s.Jitter, s.LastSeen, s.Latency, s.Loss, s.Mos, s.Mtu, s.PeerMac, s.PeerPortId, s.PeerRouterName, s.PeerSiteId, s.PortId, s.RouterName, s.Type, s.Up, s.Uptime, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayVpnPeer.
// It customizes the JSON marshaling process for StatsGatewayVpnPeer objects.
func (s StatsGatewayVpnPeer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"is_active", "jitter", "last_seen", "latency", "loss", "mos", "mtu", "peer_mac", "peer_port_id", "peer_router_name", "peer_site_id", "port_id", "router_name", "type", "up", "uptime"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayVpnPeer object to a map representation for JSON marshaling.
func (s StatsGatewayVpnPeer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.IsActive != nil {
		structMap["is_active"] = s.IsActive
	}
	if s.Jitter != nil {
		structMap["jitter"] = s.Jitter
	}
	if s.LastSeen.IsValueSet() {
		if s.LastSeen.Value() != nil {
			structMap["last_seen"] = s.LastSeen.Value()
		} else {
			structMap["last_seen"] = nil
		}
	}
	if s.Latency != nil {
		structMap["latency"] = s.Latency
	}
	if s.Loss != nil {
		structMap["loss"] = s.Loss
	}
	if s.Mos != nil {
		structMap["mos"] = s.Mos
	}
	if s.Mtu != nil {
		structMap["mtu"] = s.Mtu
	}
	if s.PeerMac != nil {
		structMap["peer_mac"] = s.PeerMac
	}
	if s.PeerPortId != nil {
		structMap["peer_port_id"] = s.PeerPortId
	}
	if s.PeerRouterName != nil {
		structMap["peer_router_name"] = s.PeerRouterName
	}
	if s.PeerSiteId != nil {
		structMap["peer_site_id"] = s.PeerSiteId
	}
	if s.PortId != nil {
		structMap["port_id"] = s.PortId
	}
	if s.RouterName != nil {
		structMap["router_name"] = s.RouterName
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	if s.Up != nil {
		structMap["up"] = s.Up
	}
	if s.Uptime != nil {
		structMap["uptime"] = s.Uptime
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayVpnPeer.
// It customizes the JSON unmarshaling process for StatsGatewayVpnPeer objects.
func (s *StatsGatewayVpnPeer) UnmarshalJSON(input []byte) error {
	var temp tempStatsGatewayVpnPeer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "is_active", "jitter", "last_seen", "latency", "loss", "mos", "mtu", "peer_mac", "peer_port_id", "peer_router_name", "peer_site_id", "port_id", "router_name", "type", "up", "uptime")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.IsActive = temp.IsActive
	s.Jitter = temp.Jitter
	s.LastSeen = temp.LastSeen
	s.Latency = temp.Latency
	s.Loss = temp.Loss
	s.Mos = temp.Mos
	s.Mtu = temp.Mtu
	s.PeerMac = temp.PeerMac
	s.PeerPortId = temp.PeerPortId
	s.PeerRouterName = temp.PeerRouterName
	s.PeerSiteId = temp.PeerSiteId
	s.PortId = temp.PortId
	s.RouterName = temp.RouterName
	s.Type = temp.Type
	s.Up = temp.Up
	s.Uptime = temp.Uptime
	return nil
}

// tempStatsGatewayVpnPeer is a temporary struct used for validating the fields of StatsGatewayVpnPeer.
type tempStatsGatewayVpnPeer struct {
	IsActive       *bool             `json:"is_active,omitempty"`
	Jitter         *float64          `json:"jitter,omitempty"`
	LastSeen       Optional[float64] `json:"last_seen"`
	Latency        *float64          `json:"latency,omitempty"`
	Loss           *float64          `json:"loss,omitempty"`
	Mos            *float64          `json:"mos,omitempty"`
	Mtu            *int              `json:"mtu,omitempty"`
	PeerMac        *string           `json:"peer_mac,omitempty"`
	PeerPortId     *string           `json:"peer_port_id,omitempty"`
	PeerRouterName *string           `json:"peer_router_name,omitempty"`
	PeerSiteId     *uuid.UUID        `json:"peer_site_id,omitempty"`
	PortId         *string           `json:"port_id,omitempty"`
	RouterName     *string           `json:"router_name,omitempty"`
	Type           *string           `json:"type,omitempty"`
	Up             *bool             `json:"up,omitempty"`
	Uptime         *int              `json:"uptime,omitempty"`
}
