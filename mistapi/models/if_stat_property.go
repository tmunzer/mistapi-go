package models

import (
    "encoding/json"
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
    RxBytes              *int                     `json:"rx_bytes,omitempty"`
    RxPkts               *int                     `json:"rx_pkts,omitempty"`
    ServpInfo            *IfStatPropertyServpInfo `json:"servp_info,omitempty"`
    TxBytes              *int                     `json:"tx_bytes,omitempty"`
    TxPkts               *int                     `json:"tx_pkts,omitempty"`
    Up                   *bool                    `json:"up,omitempty"`
    Vlan                 *int                     `json:"vlan,omitempty"`
    WanName              *string                  `json:"wan_name,omitempty"`
    WanType              *string                  `json:"wan_type,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IfStatProperty.
// It customizes the JSON marshaling process for IfStatProperty objects.
func (i IfStatProperty) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IfStatProperty object to a map representation for JSON marshaling.
func (i IfStatProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
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
    if i.RxBytes != nil {
        structMap["rx_bytes"] = i.RxBytes
    }
    if i.RxPkts != nil {
        structMap["rx_pkts"] = i.RxPkts
    }
    if i.ServpInfo != nil {
        structMap["servp_info"] = i.ServpInfo.toMap()
    }
    if i.TxBytes != nil {
        structMap["tx_bytes"] = i.TxBytes
    }
    if i.TxPkts != nil {
        structMap["tx_pkts"] = i.TxPkts
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "address_mode", "ips", "nat_addresses", "network_name", "port_id", "port_usage", "redundancy_state", "rx_bytes", "rx_pkts", "servp_info", "tx_bytes", "tx_pkts", "up", "vlan", "wan_name", "wan_type")
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
    RxBytes         *int                     `json:"rx_bytes,omitempty"`
    RxPkts          *int                     `json:"rx_pkts,omitempty"`
    ServpInfo       *IfStatPropertyServpInfo `json:"servp_info,omitempty"`
    TxBytes         *int                     `json:"tx_bytes,omitempty"`
    TxPkts          *int                     `json:"tx_pkts,omitempty"`
    Up              *bool                    `json:"up,omitempty"`
    Vlan            *int                     `json:"vlan,omitempty"`
    WanName         *string                  `json:"wan_name,omitempty"`
    WanType         *string                  `json:"wan_type,omitempty"`
}
