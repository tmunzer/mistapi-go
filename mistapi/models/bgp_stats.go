package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// BgpStats represents a BgpStats struct.
type BgpStats struct {
    // if this is created for evpn overlay
    EvpnOverlay          *bool              `json:"evpn_overlay,omitempty"`
    // if this is created for overlay
    ForOverlay           *bool              `json:"for_overlay,omitempty"`
    // AS
    LocalAs              *int               `json:"local_as,omitempty"`
    // router mac address
    Mac                  *string            `json:"mac,omitempty"`
    Neighbor             *string            `json:"neighbor,omitempty"`
    NeighborAs           *int               `json:"neighbor_as,omitempty"`
    // if it's another device in the same org
    NeighborMac          *string            `json:"neighbor_mac,omitempty"`
    // node0/node1
    Node                 *string            `json:"node,omitempty"`
    OrgId                *uuid.UUID         `json:"org_id,omitempty"`
    RxPkts               *int               `json:"rx_pkts,omitempty"`
    // number of received routes
    RxRoutes             *int               `json:"rx_routes,omitempty"`
    SiteId               *uuid.UUID         `json:"site_id,omitempty"`
    // enum: `active`, `connect`, `established`, `idle`, `open_config`, `open_sent`
    State                *BgpStatsStateEnum `json:"state,omitempty"`
    Timestamp            *float64           `json:"timestamp,omitempty"`
    TxPkts               *int               `json:"tx_pkts,omitempty"`
    TxRoutes             *int               `json:"tx_routes,omitempty"`
    Up                   *bool              `json:"up,omitempty"`
    Uptime               *int               `json:"uptime,omitempty"`
    VrfName              *string            `json:"vrf_name,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BgpStats.
// It customizes the JSON marshaling process for BgpStats objects.
func (b BgpStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BgpStats object to a map representation for JSON marshaling.
func (b BgpStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.EvpnOverlay != nil {
        structMap["evpn_overlay"] = b.EvpnOverlay
    }
    if b.ForOverlay != nil {
        structMap["for_overlay"] = b.ForOverlay
    }
    if b.LocalAs != nil {
        structMap["local_as"] = b.LocalAs
    }
    if b.Mac != nil {
        structMap["mac"] = b.Mac
    }
    if b.Neighbor != nil {
        structMap["neighbor"] = b.Neighbor
    }
    if b.NeighborAs != nil {
        structMap["neighbor_as"] = b.NeighborAs
    }
    if b.NeighborMac != nil {
        structMap["neighbor_mac"] = b.NeighborMac
    }
    if b.Node != nil {
        structMap["node"] = b.Node
    }
    if b.OrgId != nil {
        structMap["org_id"] = b.OrgId
    }
    if b.RxPkts != nil {
        structMap["rx_pkts"] = b.RxPkts
    }
    if b.RxRoutes != nil {
        structMap["rx_routes"] = b.RxRoutes
    }
    if b.SiteId != nil {
        structMap["site_id"] = b.SiteId
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

// UnmarshalJSON implements the json.Unmarshaler interface for BgpStats.
// It customizes the JSON unmarshaling process for BgpStats objects.
func (b *BgpStats) UnmarshalJSON(input []byte) error {
    var temp tempBgpStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "evpn_overlay", "for_overlay", "local_as", "mac", "neighbor", "neighbor_as", "neighbor_mac", "node", "org_id", "rx_pkts", "rx_routes", "site_id", "state", "timestamp", "tx_pkts", "tx_routes", "up", "uptime", "vrf_name")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.EvpnOverlay = temp.EvpnOverlay
    b.ForOverlay = temp.ForOverlay
    b.LocalAs = temp.LocalAs
    b.Mac = temp.Mac
    b.Neighbor = temp.Neighbor
    b.NeighborAs = temp.NeighborAs
    b.NeighborMac = temp.NeighborMac
    b.Node = temp.Node
    b.OrgId = temp.OrgId
    b.RxPkts = temp.RxPkts
    b.RxRoutes = temp.RxRoutes
    b.SiteId = temp.SiteId
    b.State = temp.State
    b.Timestamp = temp.Timestamp
    b.TxPkts = temp.TxPkts
    b.TxRoutes = temp.TxRoutes
    b.Up = temp.Up
    b.Uptime = temp.Uptime
    b.VrfName = temp.VrfName
    return nil
}

// tempBgpStats is a temporary struct used for validating the fields of BgpStats.
type tempBgpStats  struct {
    EvpnOverlay *bool              `json:"evpn_overlay,omitempty"`
    ForOverlay  *bool              `json:"for_overlay,omitempty"`
    LocalAs     *int               `json:"local_as,omitempty"`
    Mac         *string            `json:"mac,omitempty"`
    Neighbor    *string            `json:"neighbor,omitempty"`
    NeighborAs  *int               `json:"neighbor_as,omitempty"`
    NeighborMac *string            `json:"neighbor_mac,omitempty"`
    Node        *string            `json:"node,omitempty"`
    OrgId       *uuid.UUID         `json:"org_id,omitempty"`
    RxPkts      *int               `json:"rx_pkts,omitempty"`
    RxRoutes    *int               `json:"rx_routes,omitempty"`
    SiteId      *uuid.UUID         `json:"site_id,omitempty"`
    State       *BgpStatsStateEnum `json:"state,omitempty"`
    Timestamp   *float64           `json:"timestamp,omitempty"`
    TxPkts      *int               `json:"tx_pkts,omitempty"`
    TxRoutes    *int               `json:"tx_routes,omitempty"`
    Up          *bool              `json:"up,omitempty"`
    Uptime      *int               `json:"uptime,omitempty"`
    VrfName     *string            `json:"vrf_name,omitempty"`
}
