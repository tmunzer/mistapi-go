// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// VpnPeerStat represents a VpnPeerStat struct.
type VpnPeerStat struct {
    // Redundancy status of the associated interface
    IsActive             *bool                  `json:"is_active,omitempty"`
    // Last seen timestamp
    LastSeen             Optional[float64]      `json:"last_seen"`
    Latency              *float64               `json:"latency,omitempty"`
    // Router mac address
    Mac                  *string                `json:"mac,omitempty"`
    Mos                  *float64               `json:"mos,omitempty"`
    Mtu                  *int                   `json:"mtu,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // Peer router mac address
    PeerMac              *string                `json:"peer_mac,omitempty"`
    // Peer router device interface
    PeerPortId           *string                `json:"peer_port_id,omitempty"`
    PeerRouterName       *string                `json:"peer_router_name,omitempty"`
    PeerSiteId           *uuid.UUID             `json:"peer_site_id,omitempty"`
    // Router device interface
    PortId               *string                `json:"port_id,omitempty"`
    RouterName           *string                `json:"router_name,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // `ipsec`for SRX, `svr` for 128T
    Type                 *string                `json:"type,omitempty"`
    Up                   *bool                  `json:"up,omitempty"`
    Uptime               *int                   `json:"uptime,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VpnPeerStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VpnPeerStat) String() string {
    return fmt.Sprintf(
    	"VpnPeerStat[IsActive=%v, LastSeen=%v, Latency=%v, Mac=%v, Mos=%v, Mtu=%v, OrgId=%v, PeerMac=%v, PeerPortId=%v, PeerRouterName=%v, PeerSiteId=%v, PortId=%v, RouterName=%v, SiteId=%v, Type=%v, Up=%v, Uptime=%v, AdditionalProperties=%v]",
    	v.IsActive, v.LastSeen, v.Latency, v.Mac, v.Mos, v.Mtu, v.OrgId, v.PeerMac, v.PeerPortId, v.PeerRouterName, v.PeerSiteId, v.PortId, v.RouterName, v.SiteId, v.Type, v.Up, v.Uptime, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VpnPeerStat.
// It customizes the JSON marshaling process for VpnPeerStat objects.
func (v VpnPeerStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "is_active", "last_seen", "latency", "mac", "mos", "mtu", "org_id", "peer_mac", "peer_port_id", "peer_router_name", "peer_site_id", "port_id", "router_name", "site_id", "type", "up", "uptime"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VpnPeerStat object to a map representation for JSON marshaling.
func (v VpnPeerStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.IsActive != nil {
        structMap["is_active"] = v.IsActive
    }
    if v.LastSeen.IsValueSet() {
        if v.LastSeen.Value() != nil {
            structMap["last_seen"] = v.LastSeen.Value()
        } else {
            structMap["last_seen"] = nil
        }
    }
    if v.Latency != nil {
        structMap["latency"] = v.Latency
    }
    if v.Mac != nil {
        structMap["mac"] = v.Mac
    }
    if v.Mos != nil {
        structMap["mos"] = v.Mos
    }
    if v.Mtu != nil {
        structMap["mtu"] = v.Mtu
    }
    if v.OrgId != nil {
        structMap["org_id"] = v.OrgId
    }
    if v.PeerMac != nil {
        structMap["peer_mac"] = v.PeerMac
    }
    if v.PeerPortId != nil {
        structMap["peer_port_id"] = v.PeerPortId
    }
    if v.PeerRouterName != nil {
        structMap["peer_router_name"] = v.PeerRouterName
    }
    if v.PeerSiteId != nil {
        structMap["peer_site_id"] = v.PeerSiteId
    }
    if v.PortId != nil {
        structMap["port_id"] = v.PortId
    }
    if v.RouterName != nil {
        structMap["router_name"] = v.RouterName
    }
    if v.SiteId != nil {
        structMap["site_id"] = v.SiteId
    }
    if v.Type != nil {
        structMap["type"] = v.Type
    }
    if v.Up != nil {
        structMap["up"] = v.Up
    }
    if v.Uptime != nil {
        structMap["uptime"] = v.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VpnPeerStat.
// It customizes the JSON unmarshaling process for VpnPeerStat objects.
func (v *VpnPeerStat) UnmarshalJSON(input []byte) error {
    var temp tempVpnPeerStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "is_active", "last_seen", "latency", "mac", "mos", "mtu", "org_id", "peer_mac", "peer_port_id", "peer_router_name", "peer_site_id", "port_id", "router_name", "site_id", "type", "up", "uptime")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.IsActive = temp.IsActive
    v.LastSeen = temp.LastSeen
    v.Latency = temp.Latency
    v.Mac = temp.Mac
    v.Mos = temp.Mos
    v.Mtu = temp.Mtu
    v.OrgId = temp.OrgId
    v.PeerMac = temp.PeerMac
    v.PeerPortId = temp.PeerPortId
    v.PeerRouterName = temp.PeerRouterName
    v.PeerSiteId = temp.PeerSiteId
    v.PortId = temp.PortId
    v.RouterName = temp.RouterName
    v.SiteId = temp.SiteId
    v.Type = temp.Type
    v.Up = temp.Up
    v.Uptime = temp.Uptime
    return nil
}

// tempVpnPeerStat is a temporary struct used for validating the fields of VpnPeerStat.
type tempVpnPeerStat  struct {
    IsActive       *bool             `json:"is_active,omitempty"`
    LastSeen       Optional[float64] `json:"last_seen"`
    Latency        *float64          `json:"latency,omitempty"`
    Mac            *string           `json:"mac,omitempty"`
    Mos            *float64          `json:"mos,omitempty"`
    Mtu            *int              `json:"mtu,omitempty"`
    OrgId          *uuid.UUID        `json:"org_id,omitempty"`
    PeerMac        *string           `json:"peer_mac,omitempty"`
    PeerPortId     *string           `json:"peer_port_id,omitempty"`
    PeerRouterName *string           `json:"peer_router_name,omitempty"`
    PeerSiteId     *uuid.UUID        `json:"peer_site_id,omitempty"`
    PortId         *string           `json:"port_id,omitempty"`
    RouterName     *string           `json:"router_name,omitempty"`
    SiteId         *uuid.UUID        `json:"site_id,omitempty"`
    Type           *string           `json:"type,omitempty"`
    Up             *bool             `json:"up,omitempty"`
    Uptime         *int              `json:"uptime,omitempty"`
}
