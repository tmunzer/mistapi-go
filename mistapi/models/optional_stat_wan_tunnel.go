package models

import (
    "encoding/json"
    "fmt"
)

// OptionalStatWanTunnel represents a OptionalStatWanTunnel struct.
type OptionalStatWanTunnel struct {
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

// String implements the fmt.Stringer interface for OptionalStatWanTunnel,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OptionalStatWanTunnel) String() string {
    return fmt.Sprintf(
    	"OptionalStatWanTunnel[AuthAlgo=%v, EncryptAlgo=%v, IkeVersion=%v, Ip=%v, LastEvent=%v, LastFlapped=%v, Node=%v, PeerHost=%v, PeerIp=%v, Priority=%v, Protocol=%v, RxBytes=%v, RxPkts=%v, TunnelName=%v, TxBytes=%v, TxPkts=%v, Up=%v, Uptime=%v, WanName=%v, AdditionalProperties=%v]",
    	o.AuthAlgo, o.EncryptAlgo, o.IkeVersion, o.Ip, o.LastEvent, o.LastFlapped, o.Node, o.PeerHost, o.PeerIp, o.Priority, o.Protocol, o.RxBytes, o.RxPkts, o.TunnelName, o.TxBytes, o.TxPkts, o.Up, o.Uptime, o.WanName, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OptionalStatWanTunnel.
// It customizes the JSON marshaling process for OptionalStatWanTunnel objects.
func (o OptionalStatWanTunnel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "last_flapped", "node", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "tunnel_name", "tx_bytes", "tx_pkts", "up", "uptime", "wan_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OptionalStatWanTunnel object to a map representation for JSON marshaling.
func (o OptionalStatWanTunnel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AuthAlgo != nil {
        structMap["auth_algo"] = o.AuthAlgo
    }
    if o.EncryptAlgo != nil {
        structMap["encrypt_algo"] = o.EncryptAlgo
    }
    if o.IkeVersion != nil {
        structMap["ike_version"] = o.IkeVersion
    }
    if o.Ip != nil {
        structMap["ip"] = o.Ip
    }
    if o.LastEvent != nil {
        structMap["last_event"] = o.LastEvent
    }
    if o.LastFlapped != nil {
        structMap["last_flapped"] = o.LastFlapped
    }
    if o.Node != nil {
        structMap["node"] = o.Node
    }
    if o.PeerHost != nil {
        structMap["peer_host"] = o.PeerHost
    }
    if o.PeerIp != nil {
        structMap["peer_ip"] = o.PeerIp
    }
    if o.Priority != nil {
        structMap["priority"] = o.Priority
    }
    if o.Protocol != nil {
        structMap["protocol"] = o.Protocol
    }
    if o.RxBytes != nil {
        structMap["rx_bytes"] = o.RxBytes
    }
    if o.RxPkts != nil {
        structMap["rx_pkts"] = o.RxPkts
    }
    if o.TunnelName != nil {
        structMap["tunnel_name"] = o.TunnelName
    }
    if o.TxBytes != nil {
        structMap["tx_bytes"] = o.TxBytes
    }
    if o.TxPkts != nil {
        structMap["tx_pkts"] = o.TxPkts
    }
    if o.Up != nil {
        structMap["up"] = o.Up
    }
    if o.Uptime != nil {
        structMap["uptime"] = o.Uptime
    }
    if o.WanName != nil {
        structMap["wan_name"] = o.WanName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OptionalStatWanTunnel.
// It customizes the JSON unmarshaling process for OptionalStatWanTunnel objects.
func (o *OptionalStatWanTunnel) UnmarshalJSON(input []byte) error {
    var temp tempOptionalStatWanTunnel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_algo", "encrypt_algo", "ike_version", "ip", "last_event", "last_flapped", "node", "peer_host", "peer_ip", "priority", "protocol", "rx_bytes", "rx_pkts", "tunnel_name", "tx_bytes", "tx_pkts", "up", "uptime", "wan_name")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.AuthAlgo = temp.AuthAlgo
    o.EncryptAlgo = temp.EncryptAlgo
    o.IkeVersion = temp.IkeVersion
    o.Ip = temp.Ip
    o.LastEvent = temp.LastEvent
    o.LastFlapped = temp.LastFlapped
    o.Node = temp.Node
    o.PeerHost = temp.PeerHost
    o.PeerIp = temp.PeerIp
    o.Priority = temp.Priority
    o.Protocol = temp.Protocol
    o.RxBytes = temp.RxBytes
    o.RxPkts = temp.RxPkts
    o.TunnelName = temp.TunnelName
    o.TxBytes = temp.TxBytes
    o.TxPkts = temp.TxPkts
    o.Up = temp.Up
    o.Uptime = temp.Uptime
    o.WanName = temp.WanName
    return nil
}

// tempOptionalStatWanTunnel is a temporary struct used for validating the fields of OptionalStatWanTunnel.
type tempOptionalStatWanTunnel  struct {
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
