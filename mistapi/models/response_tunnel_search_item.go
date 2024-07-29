package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ResponseTunnelSearchItem represents a ResponseTunnelSearchItem struct.
type ResponseTunnelSearchItem struct {
    Ap                   *string                     `json:"ap,omitempty"`
    ForSite              *bool                       `json:"for_site,omitempty"`
    LastSeen             *float64                    `json:"last_seen,omitempty"`
    MxclusterId          *uuid.UUID                  `json:"mxcluster_id,omitempty"`
    MxedgeId             *uuid.UUID                  `json:"mxedge_id,omitempty"`
    MxtunnelId           *uuid.UUID                  `json:"mxtunnel_id,omitempty"`
    OrgId                *uuid.UUID                  `json:"org_id,omitempty"`
    // MxEdge ID of the peer(mist edge to mist edge tunnel)
    PeerMxedgeId         *uuid.UUID                  `json:"peer_mxedge_id,omitempty"`
    RemoteIp             *string                     `json:"remote_ip,omitempty"`
    RemotePort           *int                        `json:"remote_port,omitempty"`
    RxControlPkts        *int                        `json:"rx_control_pkts,omitempty"`
    // list of sessions
    Sessions             []MxtunnelStatsSession      `json:"sessions,omitempty"`
    SiteId               *uuid.UUID                  `json:"site_id,omitempty"`
    // enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
    State                *MxtunnelStatsStateEnum     `json:"state,omitempty"`
    TxControlPkts        *int                        `json:"tx_control_pkts,omitempty"`
    // duration from first (or last) SA was established
    Uptime               *int                        `json:"uptime,omitempty"`
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
    // Mist Tunnel Name
    TunnelName           *string                     `json:"tunnel_name,omitempty"`
    TxBytes              *int                        `json:"tx_bytes,omitempty"`
    TxPkts               *int                        `json:"tx_pkts,omitempty"`
    Up                   *bool                       `json:"up,omitempty"`
    // wan interface name
    WanName              *string                     `json:"wan_name,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseTunnelSearchItem.
// It customizes the JSON marshaling process for ResponseTunnelSearchItem objects.
func (r ResponseTunnelSearchItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseTunnelSearchItem object to a map representation for JSON marshaling.
func (r ResponseTunnelSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Ap != nil {
        structMap["ap"] = r.Ap
    }
    if r.ForSite != nil {
        structMap["for_site"] = r.ForSite
    }
    if r.LastSeen != nil {
        structMap["last_seen"] = r.LastSeen
    }
    if r.MxclusterId != nil {
        structMap["mxcluster_id"] = r.MxclusterId
    }
    if r.MxedgeId != nil {
        structMap["mxedge_id"] = r.MxedgeId
    }
    if r.MxtunnelId != nil {
        structMap["mxtunnel_id"] = r.MxtunnelId
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.PeerMxedgeId != nil {
        structMap["peer_mxedge_id"] = r.PeerMxedgeId
    }
    if r.RemoteIp != nil {
        structMap["remote_ip"] = r.RemoteIp
    }
    if r.RemotePort != nil {
        structMap["remote_port"] = r.RemotePort
    }
    if r.RxControlPkts != nil {
        structMap["rx_control_pkts"] = r.RxControlPkts
    }
    if r.Sessions != nil {
        structMap["sessions"] = r.Sessions
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.State != nil {
        structMap["state"] = r.State
    }
    if r.TxControlPkts != nil {
        structMap["tx_control_pkts"] = r.TxControlPkts
    }
    if r.Uptime != nil {
        structMap["uptime"] = r.Uptime
    }
    if r.AuthAlgo != nil {
        structMap["auth_algo"] = r.AuthAlgo
    }
    if r.EncryptAlgo != nil {
        structMap["encrypt_algo"] = r.EncryptAlgo
    }
    if r.IkeVersion != nil {
        structMap["ike_version"] = r.IkeVersion
    }
    if r.Ip != nil {
        structMap["ip"] = r.Ip
    }
    if r.LastEvent != nil {
        structMap["last_event"] = r.LastEvent
    }
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    if r.Node != nil {
        structMap["node"] = r.Node
    }
    if r.PeerHost != nil {
        structMap["peer_host"] = r.PeerHost
    }
    if r.PeerIp != nil {
        structMap["peer_ip"] = r.PeerIp
    }
    if r.Priority != nil {
        structMap["priority"] = r.Priority
    }
    if r.Protocol != nil {
        structMap["protocol"] = r.Protocol
    }
    if r.RxBytes != nil {
        structMap["rx_bytes"] = r.RxBytes
    }
    if r.RxPkts != nil {
        structMap["rx_pkts"] = r.RxPkts
    }
    if r.TunnelName != nil {
        structMap["tunnel_name"] = r.TunnelName
    }
    if r.TxBytes != nil {
        structMap["tx_bytes"] = r.TxBytes
    }
    if r.TxPkts != nil {
        structMap["tx_pkts"] = r.TxPkts
    }
    if r.Up != nil {
        structMap["up"] = r.Up
    }
    if r.WanName != nil {
        structMap["wan_name"] = r.WanName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseTunnelSearchItem.
// It customizes the JSON unmarshaling process for ResponseTunnelSearchItem objects.
func (r *ResponseTunnelSearchItem) UnmarshalJSON(input []byte) error {
    var temp responseTunnelSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "for_site", "last_seen", "mxcluster_id", "mxedge_id", "mxtunnel_id", "org_id", "peer_mxedge_id", "remote_ip", "remote_port", "rx_control_pkts", "sessions", "site_id", "state", "tx_control_pkts", "uptime", "auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "mac", "node", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "tunnel_name", "tx_bytes", "tx_pkts", "up", "wan_name")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Ap = temp.Ap
    r.ForSite = temp.ForSite
    r.LastSeen = temp.LastSeen
    r.MxclusterId = temp.MxclusterId
    r.MxedgeId = temp.MxedgeId
    r.MxtunnelId = temp.MxtunnelId
    r.OrgId = temp.OrgId
    r.PeerMxedgeId = temp.PeerMxedgeId
    r.RemoteIp = temp.RemoteIp
    r.RemotePort = temp.RemotePort
    r.RxControlPkts = temp.RxControlPkts
    r.Sessions = temp.Sessions
    r.SiteId = temp.SiteId
    r.State = temp.State
    r.TxControlPkts = temp.TxControlPkts
    r.Uptime = temp.Uptime
    r.AuthAlgo = temp.AuthAlgo
    r.EncryptAlgo = temp.EncryptAlgo
    r.IkeVersion = temp.IkeVersion
    r.Ip = temp.Ip
    r.LastEvent = temp.LastEvent
    r.Mac = temp.Mac
    r.Node = temp.Node
    r.PeerHost = temp.PeerHost
    r.PeerIp = temp.PeerIp
    r.Priority = temp.Priority
    r.Protocol = temp.Protocol
    r.RxBytes = temp.RxBytes
    r.RxPkts = temp.RxPkts
    r.TunnelName = temp.TunnelName
    r.TxBytes = temp.TxBytes
    r.TxPkts = temp.TxPkts
    r.Up = temp.Up
    r.WanName = temp.WanName
    return nil
}

// responseTunnelSearchItem is a temporary struct used for validating the fields of ResponseTunnelSearchItem.
type responseTunnelSearchItem  struct {
    Ap            *string                     `json:"ap,omitempty"`
    ForSite       *bool                       `json:"for_site,omitempty"`
    LastSeen      *float64                    `json:"last_seen,omitempty"`
    MxclusterId   *uuid.UUID                  `json:"mxcluster_id,omitempty"`
    MxedgeId      *uuid.UUID                  `json:"mxedge_id,omitempty"`
    MxtunnelId    *uuid.UUID                  `json:"mxtunnel_id,omitempty"`
    OrgId         *uuid.UUID                  `json:"org_id,omitempty"`
    PeerMxedgeId  *uuid.UUID                  `json:"peer_mxedge_id,omitempty"`
    RemoteIp      *string                     `json:"remote_ip,omitempty"`
    RemotePort    *int                        `json:"remote_port,omitempty"`
    RxControlPkts *int                        `json:"rx_control_pkts,omitempty"`
    Sessions      []MxtunnelStatsSession      `json:"sessions,omitempty"`
    SiteId        *uuid.UUID                  `json:"site_id,omitempty"`
    State         *MxtunnelStatsStateEnum     `json:"state,omitempty"`
    TxControlPkts *int                        `json:"tx_control_pkts,omitempty"`
    Uptime        *int                        `json:"uptime,omitempty"`
    AuthAlgo      *string                     `json:"auth_algo,omitempty"`
    EncryptAlgo   *string                     `json:"encrypt_algo,omitempty"`
    IkeVersion    *string                     `json:"ike_version,omitempty"`
    Ip            *string                     `json:"ip,omitempty"`
    LastEvent     *string                     `json:"last_event,omitempty"`
    Mac           *string                     `json:"mac,omitempty"`
    Node          *string                     `json:"node,omitempty"`
    PeerHost      *string                     `json:"peer_host,omitempty"`
    PeerIp        *string                     `json:"peer_ip,omitempty"`
    Priority      *WanTunnelStatsPriorityEnum `json:"priority,omitempty"`
    Protocol      *WanTunnelProtocolEnum      `json:"protocol,omitempty"`
    RxBytes       *int                        `json:"rx_bytes,omitempty"`
    RxPkts        *int                        `json:"rx_pkts,omitempty"`
    TunnelName    *string                     `json:"tunnel_name,omitempty"`
    TxBytes       *int                        `json:"tx_bytes,omitempty"`
    TxPkts        *int                        `json:"tx_pkts,omitempty"`
    Up            *bool                       `json:"up,omitempty"`
    WanName       *string                     `json:"wan_name,omitempty"`
}
