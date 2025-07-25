// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// IfStatProperty represents a IfStatProperty struct.
type IfStatProperty struct {
    AddressMode          *string                  `json:"address_mode,omitempty"`
    Ips                  []string                 `json:"ips,omitempty"`
    NatAddresses         []string                 `json:"nat_addresses,omitempty"`
    NetworkName          *string                  `json:"network_name,omitempty"`
    PortId               *string                  `json:"port_id,omitempty"`
    PortUsage            *string                  `json:"port_usage,omitempty"`
    RedundancyState      *string                  `json:"redundancy_state,omitempty"`
    // Amount of traffic received since connection
    RxBytes              Optional[int64]          `json:"rx_bytes"`
    // Amount of packets received since connection
    RxPkts               Optional[int64]          `json:"rx_pkts"`
    ServpInfo            *IfStatPropertyServpInfo `json:"servp_info,omitempty"`
    // Amount of traffic sent since connection
    TxBytes              Optional[int64]          `json:"tx_bytes"`
    // Amount of packets sent since connection
    TxPkts               Optional[int64]          `json:"tx_pkts"`
    Up                   *bool                    `json:"up,omitempty"`
    Vlan                 *int                     `json:"vlan,omitempty"`
    WanName              *string                  `json:"wan_name,omitempty"`
    WanType              *string                  `json:"wan_type,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for IfStatProperty,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IfStatProperty) String() string {
    return fmt.Sprintf(
    	"IfStatProperty[AddressMode=%v, Ips=%v, NatAddresses=%v, NetworkName=%v, PortId=%v, PortUsage=%v, RedundancyState=%v, RxBytes=%v, RxPkts=%v, ServpInfo=%v, TxBytes=%v, TxPkts=%v, Up=%v, Vlan=%v, WanName=%v, WanType=%v, AdditionalProperties=%v]",
    	i.AddressMode, i.Ips, i.NatAddresses, i.NetworkName, i.PortId, i.PortUsage, i.RedundancyState, i.RxBytes, i.RxPkts, i.ServpInfo, i.TxBytes, i.TxPkts, i.Up, i.Vlan, i.WanName, i.WanType, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IfStatProperty.
// It customizes the JSON marshaling process for IfStatProperty objects.
func (i IfStatProperty) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "address_mode", "ips", "nat_addresses", "network_name", "port_id", "port_usage", "redundancy_state", "rx_bytes", "rx_pkts", "servp_info", "tx_bytes", "tx_pkts", "up", "vlan", "wan_name", "wan_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IfStatProperty object to a map representation for JSON marshaling.
func (i IfStatProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.AddressMode != nil {
        structMap["address_mode"] = i.AddressMode
    }
    if i.Ips != nil {
        structMap["ips"] = i.Ips
    }
    if i.NatAddresses != nil {
        structMap["nat_addresses"] = i.NatAddresses
    }
    if i.NetworkName != nil {
        structMap["network_name"] = i.NetworkName
    }
    if i.PortId != nil {
        structMap["port_id"] = i.PortId
    }
    if i.PortUsage != nil {
        structMap["port_usage"] = i.PortUsage
    }
    if i.RedundancyState != nil {
        structMap["redundancy_state"] = i.RedundancyState
    }
    if i.RxBytes.IsValueSet() {
        if i.RxBytes.Value() != nil {
            structMap["rx_bytes"] = i.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if i.RxPkts.IsValueSet() {
        if i.RxPkts.Value() != nil {
            structMap["rx_pkts"] = i.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
        }
    }
    if i.ServpInfo != nil {
        structMap["servp_info"] = i.ServpInfo.toMap()
    }
    if i.TxBytes.IsValueSet() {
        if i.TxBytes.Value() != nil {
            structMap["tx_bytes"] = i.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if i.TxPkts.IsValueSet() {
        if i.TxPkts.Value() != nil {
            structMap["tx_pkts"] = i.TxPkts.Value()
        } else {
            structMap["tx_pkts"] = nil
        }
    }
    if i.Up != nil {
        structMap["up"] = i.Up
    }
    if i.Vlan != nil {
        structMap["vlan"] = i.Vlan
    }
    if i.WanName != nil {
        structMap["wan_name"] = i.WanName
    }
    if i.WanType != nil {
        structMap["wan_type"] = i.WanType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IfStatProperty.
// It customizes the JSON unmarshaling process for IfStatProperty objects.
func (i *IfStatProperty) UnmarshalJSON(input []byte) error {
    var temp tempIfStatProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "address_mode", "ips", "nat_addresses", "network_name", "port_id", "port_usage", "redundancy_state", "rx_bytes", "rx_pkts", "servp_info", "tx_bytes", "tx_pkts", "up", "vlan", "wan_name", "wan_type")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.AddressMode = temp.AddressMode
    i.Ips = temp.Ips
    i.NatAddresses = temp.NatAddresses
    i.NetworkName = temp.NetworkName
    i.PortId = temp.PortId
    i.PortUsage = temp.PortUsage
    i.RedundancyState = temp.RedundancyState
    i.RxBytes = temp.RxBytes
    i.RxPkts = temp.RxPkts
    i.ServpInfo = temp.ServpInfo
    i.TxBytes = temp.TxBytes
    i.TxPkts = temp.TxPkts
    i.Up = temp.Up
    i.Vlan = temp.Vlan
    i.WanName = temp.WanName
    i.WanType = temp.WanType
    return nil
}

// tempIfStatProperty is a temporary struct used for validating the fields of IfStatProperty.
type tempIfStatProperty  struct {
    AddressMode     *string                  `json:"address_mode,omitempty"`
    Ips             []string                 `json:"ips,omitempty"`
    NatAddresses    []string                 `json:"nat_addresses,omitempty"`
    NetworkName     *string                  `json:"network_name,omitempty"`
    PortId          *string                  `json:"port_id,omitempty"`
    PortUsage       *string                  `json:"port_usage,omitempty"`
    RedundancyState *string                  `json:"redundancy_state,omitempty"`
    RxBytes         Optional[int64]          `json:"rx_bytes"`
    RxPkts          Optional[int64]          `json:"rx_pkts"`
    ServpInfo       *IfStatPropertyServpInfo `json:"servp_info,omitempty"`
    TxBytes         Optional[int64]          `json:"tx_bytes"`
    TxPkts          Optional[int64]          `json:"tx_pkts"`
    Up              *bool                    `json:"up,omitempty"`
    Vlan            *int                     `json:"vlan,omitempty"`
    WanName         *string                  `json:"wan_name,omitempty"`
    WanType         *string                  `json:"wan_type,omitempty"`
}
