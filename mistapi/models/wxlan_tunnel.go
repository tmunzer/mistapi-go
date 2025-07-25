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

// WxlanTunnel represents a WxlanTunnel struct.
// WxLAn Tunnel
type WxlanTunnel struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // Dynamic Multipoint VPN configurations
    Dmvpn                *WxlanTunnelDmvpn      `json:"dmvpn,omitempty"`
    // Determined during creation time and cannot be toggled. A management tunnel cannot be used by wxlan rule or by wlan
    ForMgmt              *bool                  `json:"for_mgmt,omitempty"`
    ForSite              *bool                  `json:"for_site,omitempty"`
    // In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries.
    HelloInterval        *int                   `json:"hello_interval,omitempty"`
    HelloRetries         *int                   `json:"hello_retries,omitempty"`
    // Optional, overwrite the hostname in SCCRQ control message, default is  or null, %H and %M can be used, which will be replace with corresponding values:
    // * %H: name of the ap if provided (and will be stripped so it can be used for hostname) and fallbacks to MAC
    // * %M: MAC (e.g. 5c5b350e0060)
    Hostname             *string                `json:"hostname,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // IPSec-related configurations; requires DMVPN be enabled
    Ipsec                *WxlanTunnelIpsec      `json:"ipsec,omitempty"`
    // Whether it’s static/unmanaged (i.e. no control session). As the session configurations are not compatible, cannot be toggled.
    IsStatic             *bool                  `json:"is_static,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU
    Mtu                  *int                   `json:"mtu,omitempty"`
    // The name of the tunnel
    Name                 string                 `json:"name"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // List of remote peers’ IP or hostname
    Peers                []string               `json:"peers,omitempty"`
    // Optional, overwrite the router-id in SCCRQ control message, default is "" or null, can also be an IPv4 address
    RouterId             *string                `json:"router_id,omitempty"`
    // Secret, ‘’ if no auth is used
    Secret               *string                `json:"secret,omitempty"`
    // Sessions to be established with the tunnel. Has to be >= 1 in order for this tunnel to be useful. For management tunnel, it can only have 1
    Sessions             []WxlanTunnelSession   `json:"sessions,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // UDP port if `use_udp`==`true`
    UdpPort              *int                   `json:"udp_port,omitempty"`
    // Whether to use UDP instead of IP (proto=115, which is default of L2TPv3)
    UseUdp               *bool                  `json:"use_udp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WxlanTunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WxlanTunnel) String() string {
    return fmt.Sprintf(
    	"WxlanTunnel[CreatedTime=%v, Dmvpn=%v, ForMgmt=%v, ForSite=%v, HelloInterval=%v, HelloRetries=%v, Hostname=%v, Id=%v, Ipsec=%v, IsStatic=%v, ModifiedTime=%v, Mtu=%v, Name=%v, OrgId=%v, Peers=%v, RouterId=%v, Secret=%v, Sessions=%v, SiteId=%v, UdpPort=%v, UseUdp=%v, AdditionalProperties=%v]",
    	w.CreatedTime, w.Dmvpn, w.ForMgmt, w.ForSite, w.HelloInterval, w.HelloRetries, w.Hostname, w.Id, w.Ipsec, w.IsStatic, w.ModifiedTime, w.Mtu, w.Name, w.OrgId, w.Peers, w.RouterId, w.Secret, w.Sessions, w.SiteId, w.UdpPort, w.UseUdp, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WxlanTunnel.
// It customizes the JSON marshaling process for WxlanTunnel objects.
func (w WxlanTunnel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "created_time", "dmvpn", "for_mgmt", "for_site", "hello_interval", "hello_retries", "hostname", "id", "ipsec", "is_static", "modified_time", "mtu", "name", "org_id", "peers", "router_id", "secret", "sessions", "site_id", "udp_port", "use_udp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WxlanTunnel object to a map representation for JSON marshaling.
func (w WxlanTunnel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.CreatedTime != nil {
        structMap["created_time"] = w.CreatedTime
    }
    if w.Dmvpn != nil {
        structMap["dmvpn"] = w.Dmvpn.toMap()
    }
    if w.ForMgmt != nil {
        structMap["for_mgmt"] = w.ForMgmt
    }
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    if w.HelloInterval != nil {
        structMap["hello_interval"] = w.HelloInterval
    }
    if w.HelloRetries != nil {
        structMap["hello_retries"] = w.HelloRetries
    }
    if w.Hostname != nil {
        structMap["hostname"] = w.Hostname
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.Ipsec != nil {
        structMap["ipsec"] = w.Ipsec.toMap()
    }
    if w.IsStatic != nil {
        structMap["is_static"] = w.IsStatic
    }
    if w.ModifiedTime != nil {
        structMap["modified_time"] = w.ModifiedTime
    }
    if w.Mtu != nil {
        structMap["mtu"] = w.Mtu
    }
    structMap["name"] = w.Name
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.Peers != nil {
        structMap["peers"] = w.Peers
    }
    if w.RouterId != nil {
        structMap["router_id"] = w.RouterId
    }
    if w.Secret != nil {
        structMap["secret"] = w.Secret
    }
    if w.Sessions != nil {
        structMap["sessions"] = w.Sessions
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.UdpPort != nil {
        structMap["udp_port"] = w.UdpPort
    }
    if w.UseUdp != nil {
        structMap["use_udp"] = w.UseUdp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxlanTunnel.
// It customizes the JSON unmarshaling process for WxlanTunnel objects.
func (w *WxlanTunnel) UnmarshalJSON(input []byte) error {
    var temp tempWxlanTunnel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "dmvpn", "for_mgmt", "for_site", "hello_interval", "hello_retries", "hostname", "id", "ipsec", "is_static", "modified_time", "mtu", "name", "org_id", "peers", "router_id", "secret", "sessions", "site_id", "udp_port", "use_udp")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.CreatedTime = temp.CreatedTime
    w.Dmvpn = temp.Dmvpn
    w.ForMgmt = temp.ForMgmt
    w.ForSite = temp.ForSite
    w.HelloInterval = temp.HelloInterval
    w.HelloRetries = temp.HelloRetries
    w.Hostname = temp.Hostname
    w.Id = temp.Id
    w.Ipsec = temp.Ipsec
    w.IsStatic = temp.IsStatic
    w.ModifiedTime = temp.ModifiedTime
    w.Mtu = temp.Mtu
    w.Name = *temp.Name
    w.OrgId = temp.OrgId
    w.Peers = temp.Peers
    w.RouterId = temp.RouterId
    w.Secret = temp.Secret
    w.Sessions = temp.Sessions
    w.SiteId = temp.SiteId
    w.UdpPort = temp.UdpPort
    w.UseUdp = temp.UseUdp
    return nil
}

// tempWxlanTunnel is a temporary struct used for validating the fields of WxlanTunnel.
type tempWxlanTunnel  struct {
    CreatedTime   *float64             `json:"created_time,omitempty"`
    Dmvpn         *WxlanTunnelDmvpn    `json:"dmvpn,omitempty"`
    ForMgmt       *bool                `json:"for_mgmt,omitempty"`
    ForSite       *bool                `json:"for_site,omitempty"`
    HelloInterval *int                 `json:"hello_interval,omitempty"`
    HelloRetries  *int                 `json:"hello_retries,omitempty"`
    Hostname      *string              `json:"hostname,omitempty"`
    Id            *uuid.UUID           `json:"id,omitempty"`
    Ipsec         *WxlanTunnelIpsec    `json:"ipsec,omitempty"`
    IsStatic      *bool                `json:"is_static,omitempty"`
    ModifiedTime  *float64             `json:"modified_time,omitempty"`
    Mtu           *int                 `json:"mtu,omitempty"`
    Name          *string              `json:"name"`
    OrgId         *uuid.UUID           `json:"org_id,omitempty"`
    Peers         []string             `json:"peers,omitempty"`
    RouterId      *string              `json:"router_id,omitempty"`
    Secret        *string              `json:"secret,omitempty"`
    Sessions      []WxlanTunnelSession `json:"sessions,omitempty"`
    SiteId        *uuid.UUID           `json:"site_id,omitempty"`
    UdpPort       *int                 `json:"udp_port,omitempty"`
    UseUdp        *bool                `json:"use_udp,omitempty"`
}

func (w *tempWxlanTunnel) validate() error {
    var errs []string
    if w.Name == nil {
        errs = append(errs, "required field `name` is missing for type `wxlan_tunnel`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
