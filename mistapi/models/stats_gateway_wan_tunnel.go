package models

import (
    "encoding/json"
    "fmt"
)

// StatsGatewayWanTunnel represents a StatsGatewayWanTunnel struct.
type StatsGatewayWanTunnel struct {
    // Authentication algorithm
    AuthAlgo             *string                     `json:"auth_algo,omitempty"`
    // Encryption algorithm
    EncryptAlgo          *string                     `json:"encrypt_algo,omitempty"`
    // IKE version
    IkeVersion           *string                     `json:"ike_version,omitempty"`
    // IPaddress
    Ip                   *string                     `json:"ip,omitempty"`
    // Reason of why the tunnel is down
    LastEvent            *string                     `json:"last_event,omitempty"`
    // Indicates when the port was last flapped
    LastFlapped          *float64                    `json:"last_flapped,omitempty"`
    // Node0/node1
    Node                 *string                     `json:"node,omitempty"`
    // Peer host
    PeerHost             *string                     `json:"peer_host,omitempty"`
    // Peer ip address
    PeerIp               *string                     `json:"peer_ip,omitempty"`
    // enum: `primary`, `secondary`
    Priority             *StatsWanTunnelPriorityEnum `json:"priority,omitempty"`
    // enum: `gre`, `ipsec`
    Protocol             *WanTunnelProtocolEnum      `json:"protocol,omitempty"`
    RxBytes              *int                        `json:"rx_bytes,omitempty"`
    RxPkts               *int                        `json:"rx_pkts,omitempty"`
    // Mist Tunnel Name
    TunnelName           *string                     `json:"tunnel_name,omitempty"`
    TxBytes              *int                        `json:"tx_bytes,omitempty"`
    TxPkts               *int                        `json:"tx_pkts,omitempty"`
    Up                   *bool                       `json:"up,omitempty"`
    // Duration from first (or last) SA was established
    Uptime               *int                        `json:"uptime,omitempty"`
    // WAN interface name
    WanName              *string                     `json:"wan_name,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for StatsGatewayWanTunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsGatewayWanTunnel) String() string {
    return fmt.Sprintf(
    	"StatsGatewayWanTunnel[AuthAlgo=%v, EncryptAlgo=%v, IkeVersion=%v, Ip=%v, LastEvent=%v, LastFlapped=%v, Node=%v, PeerHost=%v, PeerIp=%v, Priority=%v, Protocol=%v, RxBytes=%v, RxPkts=%v, TunnelName=%v, TxBytes=%v, TxPkts=%v, Up=%v, Uptime=%v, WanName=%v, AdditionalProperties=%v]",
    	s.AuthAlgo, s.EncryptAlgo, s.IkeVersion, s.Ip, s.LastEvent, s.LastFlapped, s.Node, s.PeerHost, s.PeerIp, s.Priority, s.Protocol, s.RxBytes, s.RxPkts, s.TunnelName, s.TxBytes, s.TxPkts, s.Up, s.Uptime, s.WanName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsGatewayWanTunnel.
// It customizes the JSON marshaling process for StatsGatewayWanTunnel objects.
func (s StatsGatewayWanTunnel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "last_flapped", "node", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "tunnel_name", "tx_bytes", "tx_pkts", "up", "uptime", "wan_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGatewayWanTunnel object to a map representation for JSON marshaling.
func (s StatsGatewayWanTunnel) toMap() map[string]any {
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
    if s.LastFlapped != nil {
        structMap["last_flapped"] = s.LastFlapped
    }
    if s.Node != nil {
        structMap["node"] = s.Node
    }
    if s.PeerHost != nil {
        structMap["peer_host"] = s.PeerHost
    }
    if s.PeerIp != nil {
        structMap["peer_ip"] = s.PeerIp
    }
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

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGatewayWanTunnel.
// It customizes the JSON unmarshaling process for StatsGatewayWanTunnel objects.
func (s *StatsGatewayWanTunnel) UnmarshalJSON(input []byte) error {
    var temp tempStatsGatewayWanTunnel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "last_flapped", "node", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "tunnel_name", "tx_bytes", "tx_pkts", "up", "uptime", "wan_name")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AuthAlgo = temp.AuthAlgo
    s.EncryptAlgo = temp.EncryptAlgo
    s.IkeVersion = temp.IkeVersion
    s.Ip = temp.Ip
    s.LastEvent = temp.LastEvent
    s.LastFlapped = temp.LastFlapped
    s.Node = temp.Node
    s.PeerHost = temp.PeerHost
    s.PeerIp = temp.PeerIp
    s.Priority = temp.Priority
    s.Protocol = temp.Protocol
    s.RxBytes = temp.RxBytes
    s.RxPkts = temp.RxPkts
    s.TunnelName = temp.TunnelName
    s.TxBytes = temp.TxBytes
    s.TxPkts = temp.TxPkts
    s.Up = temp.Up
    s.Uptime = temp.Uptime
    s.WanName = temp.WanName
    return nil
}

// tempStatsGatewayWanTunnel is a temporary struct used for validating the fields of StatsGatewayWanTunnel.
type tempStatsGatewayWanTunnel  struct {
    AuthAlgo    *string                     `json:"auth_algo,omitempty"`
    EncryptAlgo *string                     `json:"encrypt_algo,omitempty"`
    IkeVersion  *string                     `json:"ike_version,omitempty"`
    Ip          *string                     `json:"ip,omitempty"`
    LastEvent   *string                     `json:"last_event,omitempty"`
    LastFlapped *float64                    `json:"last_flapped,omitempty"`
    Node        *string                     `json:"node,omitempty"`
    PeerHost    *string                     `json:"peer_host,omitempty"`
    PeerIp      *string                     `json:"peer_ip,omitempty"`
    Priority    *StatsWanTunnelPriorityEnum `json:"priority,omitempty"`
    Protocol    *WanTunnelProtocolEnum      `json:"protocol,omitempty"`
    RxBytes     *int                        `json:"rx_bytes,omitempty"`
    RxPkts      *int                        `json:"rx_pkts,omitempty"`
    TunnelName  *string                     `json:"tunnel_name,omitempty"`
    TxBytes     *int                        `json:"tx_bytes,omitempty"`
    TxPkts      *int                        `json:"tx_pkts,omitempty"`
    Up          *bool                       `json:"up,omitempty"`
    Uptime      *int                        `json:"uptime,omitempty"`
    WanName     *string                     `json:"wan_name,omitempty"`
}
