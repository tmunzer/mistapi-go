package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsWanTunnel represents a StatsWanTunnel struct.
type StatsWanTunnel struct {
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
    PeerIp               string                      `json:"peer_ip"`
    // enum: `primary`, `secondary`
    Priority             *StatsWanTunnelPriorityEnum `json:"priority,omitempty"`
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
    AdditionalProperties map[string]interface{}      `json:"_"`
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
    if s.RxBytes != nil {
        structMap["rx_bytes"] = s.RxBytes
    }
    if s.RxPkts != nil {
        structMap["rx_pkts"] = s.RxPkts
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.TunnelName != nil {
        structMap["tunnel_name"] = s.TunnelName
    }
    if s.TxBytes != nil {
        structMap["tx_bytes"] = s.TxBytes
    }
    if s.TxPkts != nil {
        structMap["tx_pkts"] = s.TxPkts
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
type tempStatsWanTunnel  struct {
    AuthAlgo    *string                     `json:"auth_algo,omitempty"`
    EncryptAlgo *string                     `json:"encrypt_algo,omitempty"`
    IkeVersion  *string                     `json:"ike_version,omitempty"`
    Ip          *string                     `json:"ip,omitempty"`
    LastEvent   *string                     `json:"last_event,omitempty"`
    Mac         *string                     `json:"mac,omitempty"`
    Node        *string                     `json:"node,omitempty"`
    OrgId       *uuid.UUID                  `json:"org_id,omitempty"`
    PeerHost    *string                     `json:"peer_host,omitempty"`
    PeerIp      *string                     `json:"peer_ip"`
    Priority    *StatsWanTunnelPriorityEnum `json:"priority,omitempty"`
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

func (s *tempStatsWanTunnel) validate() error {
    var errs []string
    if s.PeerIp == nil {
        errs = append(errs, "required field `peer_ip` is missing for type `stats_wan_tunnel`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
