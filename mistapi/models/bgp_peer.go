package models

import (
    "encoding/json"
    "fmt"
)

// BgpPeer represents a BgpPeer struct.
// Only present when `bgp_peers` in `fields` query parameter
type BgpPeer struct {
    // If this is created for evpn overlay
    EvpnOverlay          *bool                  `json:"evpn_overlay,omitempty"`
    // If this is created for overlay
    ForOverlay           *bool                  `json:"for_overlay,omitempty"`
    // BGP AS, value in range 1-4294967295
    LocalAs              *BgpAs                 `json:"local_as,omitempty"`
    Neighbor             *string                `json:"neighbor,omitempty"`
    // BGP AS, value in range 1-4294967295
    NeighborAs           *BgpAs                 `json:"neighbor_as,omitempty"`
    // If it's another device in the same org
    NeighborMac          *string                `json:"neighbor_mac,omitempty"`
    // Node0/node1
    Node                 *string                `json:"node,omitempty"`
    RxPkts               *int                   `json:"rx_pkts,omitempty"`
    // Number of received routes
    RxRoutes             *int                   `json:"rx_routes,omitempty"`
    // enum: `active`, `connect`, `established`, `idle`, `open_config`, `open_sent`
    State                *BgpStatsStateEnum     `json:"state,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    TxPkts               *int                   `json:"tx_pkts,omitempty"`
    TxRoutes             *int                   `json:"tx_routes,omitempty"`
    Up                   *bool                  `json:"up,omitempty"`
    Uptime               *int                   `json:"uptime,omitempty"`
    VrfName              *string                `json:"vrf_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BgpPeer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BgpPeer) String() string {
    return fmt.Sprintf(
    	"BgpPeer[EvpnOverlay=%v, ForOverlay=%v, LocalAs=%v, Neighbor=%v, NeighborAs=%v, NeighborMac=%v, Node=%v, RxPkts=%v, RxRoutes=%v, State=%v, Timestamp=%v, TxPkts=%v, TxRoutes=%v, Up=%v, Uptime=%v, VrfName=%v, AdditionalProperties=%v]",
    	b.EvpnOverlay, b.ForOverlay, b.LocalAs, b.Neighbor, b.NeighborAs, b.NeighborMac, b.Node, b.RxPkts, b.RxRoutes, b.State, b.Timestamp, b.TxPkts, b.TxRoutes, b.Up, b.Uptime, b.VrfName, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BgpPeer.
// It customizes the JSON marshaling process for BgpPeer objects.
func (b BgpPeer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "evpn_overlay", "for_overlay", "local_as", "neighbor", "neighbor_as", "neighbor_mac", "node", "rx_pkts", "rx_routes", "state", "timestamp", "tx_pkts", "tx_routes", "up", "uptime", "vrf_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BgpPeer object to a map representation for JSON marshaling.
func (b BgpPeer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.EvpnOverlay != nil {
        structMap["evpn_overlay"] = b.EvpnOverlay
    }
    if b.ForOverlay != nil {
        structMap["for_overlay"] = b.ForOverlay
    }
    if b.LocalAs != nil {
        structMap["local_as"] = b.LocalAs.toMap()
    }
    if b.Neighbor != nil {
        structMap["neighbor"] = b.Neighbor
    }
    if b.NeighborAs != nil {
        structMap["neighbor_as"] = b.NeighborAs.toMap()
    }
    if b.NeighborMac != nil {
        structMap["neighbor_mac"] = b.NeighborMac
    }
    if b.Node != nil {
        structMap["node"] = b.Node
    }
    if b.RxPkts != nil {
        structMap["rx_pkts"] = b.RxPkts
    }
    if b.RxRoutes != nil {
        structMap["rx_routes"] = b.RxRoutes
    }
    if b.State != nil {
        structMap["state"] = b.State
    }
    if b.Timestamp != nil {
        structMap["timestamp"] = b.Timestamp
    }
    if b.TxPkts != nil {
        structMap["tx_pkts"] = b.TxPkts
    }
    if b.TxRoutes != nil {
        structMap["tx_routes"] = b.TxRoutes
    }
    if b.Up != nil {
        structMap["up"] = b.Up
    }
    if b.Uptime != nil {
        structMap["uptime"] = b.Uptime
    }
    if b.VrfName != nil {
        structMap["vrf_name"] = b.VrfName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BgpPeer.
// It customizes the JSON unmarshaling process for BgpPeer objects.
func (b *BgpPeer) UnmarshalJSON(input []byte) error {
    var temp tempBgpPeer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "evpn_overlay", "for_overlay", "local_as", "neighbor", "neighbor_as", "neighbor_mac", "node", "rx_pkts", "rx_routes", "state", "timestamp", "tx_pkts", "tx_routes", "up", "uptime", "vrf_name")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.EvpnOverlay = temp.EvpnOverlay
    b.ForOverlay = temp.ForOverlay
    b.LocalAs = temp.LocalAs
    b.Neighbor = temp.Neighbor
    b.NeighborAs = temp.NeighborAs
    b.NeighborMac = temp.NeighborMac
    b.Node = temp.Node
    b.RxPkts = temp.RxPkts
    b.RxRoutes = temp.RxRoutes
    b.State = temp.State
    b.Timestamp = temp.Timestamp
    b.TxPkts = temp.TxPkts
    b.TxRoutes = temp.TxRoutes
    b.Up = temp.Up
    b.Uptime = temp.Uptime
    b.VrfName = temp.VrfName
    return nil
}

// tempBgpPeer is a temporary struct used for validating the fields of BgpPeer.
type tempBgpPeer  struct {
    EvpnOverlay *bool              `json:"evpn_overlay,omitempty"`
    ForOverlay  *bool              `json:"for_overlay,omitempty"`
    LocalAs     *BgpAs             `json:"local_as,omitempty"`
    Neighbor    *string            `json:"neighbor,omitempty"`
    NeighborAs  *BgpAs             `json:"neighbor_as,omitempty"`
    NeighborMac *string            `json:"neighbor_mac,omitempty"`
    Node        *string            `json:"node,omitempty"`
    RxPkts      *int               `json:"rx_pkts,omitempty"`
    RxRoutes    *int               `json:"rx_routes,omitempty"`
    State       *BgpStatsStateEnum `json:"state,omitempty"`
    Timestamp   *float64           `json:"timestamp,omitempty"`
    TxPkts      *int               `json:"tx_pkts,omitempty"`
    TxRoutes    *int               `json:"tx_routes,omitempty"`
    Up          *bool              `json:"up,omitempty"`
    Uptime      *int               `json:"uptime,omitempty"`
    VrfName     *string            `json:"vrf_name,omitempty"`
}
