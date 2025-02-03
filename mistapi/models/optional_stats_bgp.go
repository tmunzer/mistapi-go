package models

import (
    "encoding/json"
    "fmt"
)

// OptionalStatsBgp represents a OptionalStatsBgp struct.
// Only present when `bgp_peers` in `fields` query parameter
type OptionalStatsBgp struct {
    // If this is created for evpn overlay
    EvpnOverlay          *bool                  `json:"evpn_overlay,omitempty"`
    // If this is created for overlay
    ForOverlay           *bool                  `json:"for_overlay,omitempty"`
    // AS
    LocalAs              *int                   `json:"local_as,omitempty"`
    Neighbor             *string                `json:"neighbor,omitempty"`
    NeighborAs           *int                   `json:"neighbor_as,omitempty"`
    // If it's another device in the same org
    NeighborMac          *string                `json:"neighbor_mac,omitempty"`
    // Node0/node1
    Node                 *string                `json:"node,omitempty"`
    RxPkts               *int                   `json:"rx_pkts,omitempty"`
    // Number of received routes
    RxRoutes             *int                   `json:"rx_routes,omitempty"`
    // enum: `active`, `connect`, `established`, `idle`, `open_config`, `open_sent`
    State                *BgpStatsStateEnum     `json:"state,omitempty"`
    Timestamp            *float64               `json:"timestamp,omitempty"`
    TxPkts               *int                   `json:"tx_pkts,omitempty"`
    TxRoutes             *int                   `json:"tx_routes,omitempty"`
    Up                   *bool                  `json:"up,omitempty"`
    Uptime               *int                   `json:"uptime,omitempty"`
    VrfName              *string                `json:"vrf_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OptionalStatsBgp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OptionalStatsBgp) String() string {
    return fmt.Sprintf(
    	"OptionalStatsBgp[EvpnOverlay=%v, ForOverlay=%v, LocalAs=%v, Neighbor=%v, NeighborAs=%v, NeighborMac=%v, Node=%v, RxPkts=%v, RxRoutes=%v, State=%v, Timestamp=%v, TxPkts=%v, TxRoutes=%v, Up=%v, Uptime=%v, VrfName=%v, AdditionalProperties=%v]",
    	o.EvpnOverlay, o.ForOverlay, o.LocalAs, o.Neighbor, o.NeighborAs, o.NeighborMac, o.Node, o.RxPkts, o.RxRoutes, o.State, o.Timestamp, o.TxPkts, o.TxRoutes, o.Up, o.Uptime, o.VrfName, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OptionalStatsBgp.
// It customizes the JSON marshaling process for OptionalStatsBgp objects.
func (o OptionalStatsBgp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "evpn_overlay", "for_overlay", "local_as", "neighbor", "neighbor_as", "neighbor_mac", "node", "rx_pkts", "rx_routes", "state", "timestamp", "tx_pkts", "tx_routes", "up", "uptime", "vrf_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OptionalStatsBgp object to a map representation for JSON marshaling.
func (o OptionalStatsBgp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.EvpnOverlay != nil {
        structMap["evpn_overlay"] = o.EvpnOverlay
    }
    if o.ForOverlay != nil {
        structMap["for_overlay"] = o.ForOverlay
    }
    if o.LocalAs != nil {
        structMap["local_as"] = o.LocalAs
    }
    if o.Neighbor != nil {
        structMap["neighbor"] = o.Neighbor
    }
    if o.NeighborAs != nil {
        structMap["neighbor_as"] = o.NeighborAs
    }
    if o.NeighborMac != nil {
        structMap["neighbor_mac"] = o.NeighborMac
    }
    if o.Node != nil {
        structMap["node"] = o.Node
    }
    if o.RxPkts != nil {
        structMap["rx_pkts"] = o.RxPkts
    }
    if o.RxRoutes != nil {
        structMap["rx_routes"] = o.RxRoutes
    }
    if o.State != nil {
        structMap["state"] = o.State
    }
    if o.Timestamp != nil {
        structMap["timestamp"] = o.Timestamp
    }
    if o.TxPkts != nil {
        structMap["tx_pkts"] = o.TxPkts
    }
    if o.TxRoutes != nil {
        structMap["tx_routes"] = o.TxRoutes
    }
    if o.Up != nil {
        structMap["up"] = o.Up
    }
    if o.Uptime != nil {
        structMap["uptime"] = o.Uptime
    }
    if o.VrfName != nil {
        structMap["vrf_name"] = o.VrfName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OptionalStatsBgp.
// It customizes the JSON unmarshaling process for OptionalStatsBgp objects.
func (o *OptionalStatsBgp) UnmarshalJSON(input []byte) error {
    var temp tempOptionalStatsBgp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "evpn_overlay", "for_overlay", "local_as", "neighbor", "neighbor_as", "neighbor_mac", "node", "rx_pkts", "rx_routes", "state", "timestamp", "tx_pkts", "tx_routes", "up", "uptime", "vrf_name")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.EvpnOverlay = temp.EvpnOverlay
    o.ForOverlay = temp.ForOverlay
    o.LocalAs = temp.LocalAs
    o.Neighbor = temp.Neighbor
    o.NeighborAs = temp.NeighborAs
    o.NeighborMac = temp.NeighborMac
    o.Node = temp.Node
    o.RxPkts = temp.RxPkts
    o.RxRoutes = temp.RxRoutes
    o.State = temp.State
    o.Timestamp = temp.Timestamp
    o.TxPkts = temp.TxPkts
    o.TxRoutes = temp.TxRoutes
    o.Up = temp.Up
    o.Uptime = temp.Uptime
    o.VrfName = temp.VrfName
    return nil
}

// tempOptionalStatsBgp is a temporary struct used for validating the fields of OptionalStatsBgp.
type tempOptionalStatsBgp  struct {
    EvpnOverlay *bool              `json:"evpn_overlay,omitempty"`
    ForOverlay  *bool              `json:"for_overlay,omitempty"`
    LocalAs     *int               `json:"local_as,omitempty"`
    Neighbor    *string            `json:"neighbor,omitempty"`
    NeighborAs  *int               `json:"neighbor_as,omitempty"`
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
