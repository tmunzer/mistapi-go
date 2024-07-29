package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// WanTunnelStats represents a WanTunnelStats struct.
type WanTunnelStats struct {
    // authentication algorithm
    AuthAlgo             *string                     `json:"auth_algo,omitempty"`
    // encryption algorithm
    EncryptAlgo          *string                     `json:"encrypt_algo,omitempty"`
    // ike version
    IkeVersion           *string                     `json:"ike_version,omitempty"`
    // ip address
    Ip                   *string                     `json:"ip,omitempty"`
    // reason of why the tunnel is down
    LastEvent            *string                     `json:"last_event,omitempty"`
    // router mac address
    Mac                  *string                     `json:"mac,omitempty"`
    // node0/node1
    Node                 *string                     `json:"node,omitempty"`
    OrgId                *uuid.UUID                  `json:"org_id,omitempty"`
    // peer host
    PeerHost             *string                     `json:"peer_host,omitempty"`
    // peer ip address
    PeerIp               *string                     `json:"peer_ip,omitempty"`
    // enum: `primary`, `secondary`
    Priority             *WanTunnelStatsPriorityEnum `json:"priority,omitempty"`
    // enum: `gre`, `ipsec`
    Protocol             *WanTunnelProtocolEnum      `json:"protocol,omitempty"`
    RxBytes              *int                        `json:"rx_bytes,omitempty"`
    RxPkts               *int                        `json:"rx_pkts,omitempty"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    // Mist Tunnel Name
    TunnelName           *string                     `json:"tunnel_name,omitempty"`
    TxBytes              *int                        `json:"tx_bytes,omitempty"`
    TxPkts               *int                        `json:"tx_pkts,omitempty"`
    Up                   *bool                       `json:"up,omitempty"`
    // duration from first (or last) SA was established
    Uptime               *int                        `json:"uptime,omitempty"`
    // wan interface name
    WanName              *string                     `json:"wan_name,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WanTunnelStats.
// It customizes the JSON marshaling process for WanTunnelStats objects.
func (w WanTunnelStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WanTunnelStats object to a map representation for JSON marshaling.
func (w WanTunnelStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AuthAlgo != nil {
        structMap["auth_algo"] = w.AuthAlgo
    }
    if w.EncryptAlgo != nil {
        structMap["encrypt_algo"] = w.EncryptAlgo
    }
    if w.IkeVersion != nil {
        structMap["ike_version"] = w.IkeVersion
    }
    if w.Ip != nil {
        structMap["ip"] = w.Ip
    }
    if w.LastEvent != nil {
        structMap["last_event"] = w.LastEvent
    }
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    if w.Node != nil {
        structMap["node"] = w.Node
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.PeerHost != nil {
        structMap["peer_host"] = w.PeerHost
    }
    if w.PeerIp != nil {
        structMap["peer_ip"] = w.PeerIp
    }
    if w.Priority != nil {
        structMap["priority"] = w.Priority
    }
    if w.Protocol != nil {
        structMap["protocol"] = w.Protocol
    }
    if w.RxBytes != nil {
        structMap["rx_bytes"] = w.RxBytes
    }
    if w.RxPkts != nil {
        structMap["rx_pkts"] = w.RxPkts
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.TunnelName != nil {
        structMap["tunnel_name"] = w.TunnelName
    }
    if w.TxBytes != nil {
        structMap["tx_bytes"] = w.TxBytes
    }
    if w.TxPkts != nil {
        structMap["tx_pkts"] = w.TxPkts
    }
    if w.Up != nil {
        structMap["up"] = w.Up
    }
    if w.Uptime != nil {
        structMap["uptime"] = w.Uptime
    }
    if w.WanName != nil {
        structMap["wan_name"] = w.WanName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WanTunnelStats.
// It customizes the JSON unmarshaling process for WanTunnelStats objects.
func (w *WanTunnelStats) UnmarshalJSON(input []byte) error {
    var temp wanTunnelStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "mac", "node", "org_id", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "site_id", "tunnel_name", "tx_bytes", "tx_pkts", "up", "uptime", "wan_name")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.AuthAlgo = temp.AuthAlgo
    w.EncryptAlgo = temp.EncryptAlgo
    w.IkeVersion = temp.IkeVersion
    w.Ip = temp.Ip
    w.LastEvent = temp.LastEvent
    w.Mac = temp.Mac
    w.Node = temp.Node
    w.OrgId = temp.OrgId
    w.PeerHost = temp.PeerHost
    w.PeerIp = temp.PeerIp
    w.Priority = temp.Priority
    w.Protocol = temp.Protocol
    w.RxBytes = temp.RxBytes
    w.RxPkts = temp.RxPkts
    w.SiteId = temp.SiteId
    w.TunnelName = temp.TunnelName
    w.TxBytes = temp.TxBytes
    w.TxPkts = temp.TxPkts
    w.Up = temp.Up
    w.Uptime = temp.Uptime
    w.WanName = temp.WanName
    return nil
}

// wanTunnelStats is a temporary struct used for validating the fields of WanTunnelStats.
type wanTunnelStats  struct {
    AuthAlgo    *string                     `json:"auth_algo,omitempty"`
    EncryptAlgo *string                     `json:"encrypt_algo,omitempty"`
    IkeVersion  *string                     `json:"ike_version,omitempty"`
    Ip          *string                     `json:"ip,omitempty"`
    LastEvent   *string                     `json:"last_event,omitempty"`
    Mac         *string                     `json:"mac,omitempty"`
    Node        *string                     `json:"node,omitempty"`
    OrgId       *uuid.UUID                  `json:"org_id,omitempty"`
    PeerHost    *string                     `json:"peer_host,omitempty"`
    PeerIp      *string                     `json:"peer_ip,omitempty"`
    Priority    *WanTunnelStatsPriorityEnum `json:"priority,omitempty"`
    Protocol    *WanTunnelProtocolEnum      `json:"protocol,omitempty"`
    RxBytes     *int                        `json:"rx_bytes,omitempty"`
    RxPkts      *int                        `json:"rx_pkts,omitempty"`
    SiteId      *uuid.UUID                  `json:"site_id,omitempty"`
    TunnelName  *string                     `json:"tunnel_name,omitempty"`
    TxBytes     *int                        `json:"tx_bytes,omitempty"`
    TxPkts      *int                        `json:"tx_pkts,omitempty"`
    Up          *bool                       `json:"up,omitempty"`
    Uptime      *int                        `json:"uptime,omitempty"`
    WanName     *string                     `json:"wan_name,omitempty"`
}
