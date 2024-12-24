package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// OptionalStatVpnPeer represents a OptionalStatVpnPeer struct.
type OptionalStatVpnPeer struct {
    // Redundancy status of the associated interface
    IsActive             *bool                  `json:"is_active,omitempty"`
    LastSeen             *float64               `json:"last_seen,omitempty"`
    Latency              *float64               `json:"latency,omitempty"`
    Mos                  *float64               `json:"mos,omitempty"`
    Mtu                  *int                   `json:"mtu,omitempty"`
    // peer router mac address
    PeerMac              *string                `json:"peer_mac,omitempty"`
    // peer router device interface
    PeerPortId           *string                `json:"peer_port_id,omitempty"`
    PeerRouterName       *string                `json:"peer_router_name,omitempty"`
    PeerSiteId           *uuid.UUID             `json:"peer_site_id,omitempty"`
    // router device interface
    PortId               *string                `json:"port_id,omitempty"`
    RouterName           *string                `json:"router_name,omitempty"`
    // `ipsec`for SRX, `svr` for 128T
    Type                 *string                `json:"type,omitempty"`
    Up                   *bool                  `json:"up,omitempty"`
    Uptime               *int                   `json:"uptime,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OptionalStatVpnPeer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OptionalStatVpnPeer) String() string {
    return fmt.Sprintf(
    	"OptionalStatVpnPeer[IsActive=%v, LastSeen=%v, Latency=%v, Mos=%v, Mtu=%v, PeerMac=%v, PeerPortId=%v, PeerRouterName=%v, PeerSiteId=%v, PortId=%v, RouterName=%v, Type=%v, Up=%v, Uptime=%v, AdditionalProperties=%v]",
    	o.IsActive, o.LastSeen, o.Latency, o.Mos, o.Mtu, o.PeerMac, o.PeerPortId, o.PeerRouterName, o.PeerSiteId, o.PortId, o.RouterName, o.Type, o.Up, o.Uptime, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OptionalStatVpnPeer.
// It customizes the JSON marshaling process for OptionalStatVpnPeer objects.
func (o OptionalStatVpnPeer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "is_active", "last_seen", "latency", "mos", "mtu", "peer_mac", "peer_port_id", "peer_router_name", "peer_site_id", "port_id", "router_name", "type", "up", "uptime"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OptionalStatVpnPeer object to a map representation for JSON marshaling.
func (o OptionalStatVpnPeer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.IsActive != nil {
        structMap["is_active"] = o.IsActive
    }
    if o.LastSeen != nil {
        structMap["last_seen"] = o.LastSeen
    }
    if o.Latency != nil {
        structMap["latency"] = o.Latency
    }
    if o.Mos != nil {
        structMap["mos"] = o.Mos
    }
    if o.Mtu != nil {
        structMap["mtu"] = o.Mtu
    }
    if o.PeerMac != nil {
        structMap["peer_mac"] = o.PeerMac
    }
    if o.PeerPortId != nil {
        structMap["peer_port_id"] = o.PeerPortId
    }
    if o.PeerRouterName != nil {
        structMap["peer_router_name"] = o.PeerRouterName
    }
    if o.PeerSiteId != nil {
        structMap["peer_site_id"] = o.PeerSiteId
    }
    if o.PortId != nil {
        structMap["port_id"] = o.PortId
    }
    if o.RouterName != nil {
        structMap["router_name"] = o.RouterName
    }
    if o.Type != nil {
        structMap["type"] = o.Type
    }
    if o.Up != nil {
        structMap["up"] = o.Up
    }
    if o.Uptime != nil {
        structMap["uptime"] = o.Uptime
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OptionalStatVpnPeer.
// It customizes the JSON unmarshaling process for OptionalStatVpnPeer objects.
func (o *OptionalStatVpnPeer) UnmarshalJSON(input []byte) error {
    var temp tempOptionalStatVpnPeer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "is_active", "last_seen", "latency", "mos", "mtu", "peer_mac", "peer_port_id", "peer_router_name", "peer_site_id", "port_id", "router_name", "type", "up", "uptime")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.IsActive = temp.IsActive
    o.LastSeen = temp.LastSeen
    o.Latency = temp.Latency
    o.Mos = temp.Mos
    o.Mtu = temp.Mtu
    o.PeerMac = temp.PeerMac
    o.PeerPortId = temp.PeerPortId
    o.PeerRouterName = temp.PeerRouterName
    o.PeerSiteId = temp.PeerSiteId
    o.PortId = temp.PortId
    o.RouterName = temp.RouterName
    o.Type = temp.Type
    o.Up = temp.Up
    o.Uptime = temp.Uptime
    return nil
}

// tempOptionalStatVpnPeer is a temporary struct used for validating the fields of OptionalStatVpnPeer.
type tempOptionalStatVpnPeer  struct {
    IsActive       *bool      `json:"is_active,omitempty"`
    LastSeen       *float64   `json:"last_seen,omitempty"`
    Latency        *float64   `json:"latency,omitempty"`
    Mos            *float64   `json:"mos,omitempty"`
    Mtu            *int       `json:"mtu,omitempty"`
    PeerMac        *string    `json:"peer_mac,omitempty"`
    PeerPortId     *string    `json:"peer_port_id,omitempty"`
    PeerRouterName *string    `json:"peer_router_name,omitempty"`
    PeerSiteId     *uuid.UUID `json:"peer_site_id,omitempty"`
    PortId         *string    `json:"port_id,omitempty"`
    RouterName     *string    `json:"router_name,omitempty"`
    Type           *string    `json:"type,omitempty"`
    Up             *bool      `json:"up,omitempty"`
    Uptime         *int       `json:"uptime,omitempty"`
}
