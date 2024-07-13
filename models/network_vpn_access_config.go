package models

import (
    "encoding/json"
)

// NetworkVpnAccessConfig represents a NetworkVpnAccessConfig struct.
type NetworkVpnAccessConfig struct {
    // if `routed`==`true`, whether to advertise an aggregated subnet toward HUB this is useful when there are multiple networks on SPOKE's side
    AdvertisedSubnet          *string                                  `json:"advertised_subnet,omitempty"`
    // whether to allow ping from vpn into this routed network
    AllowPing                 *bool                                    `json:"allow_ping,omitempty"`
    // Property key may be an IP/Port (i.e. "63.16.0.3:443"), or a port (i.e. ":2222")
    DestinationNat            map[string]NetworkDestinationNatProperty `json:"destination_nat,omitempty"`
    // if `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub, a subnet is required to create and advertise the route to Hub
    NatPool                   *string                                  `json:"nat_pool,omitempty"`
    // toward LAN-side BGP peers
    NoReadvertiseToLanBgp     *bool                                    `json:"no_readvertise_to_lan_bgp,omitempty"`
    // toward LAN-side OSPF peers
    NoReadvertiseToLanOspf    *bool                                    `json:"no_readvertise_to_lan_ospf,omitempty"`
    // toward overlay
    // how HUB should deal with routes it received from Spokes
    NoReadvertiseToOverlay    *bool                                    `json:"no_readvertise_to_overlay,omitempty"`
    // by default, the routes are only readvertised toward the same vrf on spoke
    // to allow it to be leaked to other vrfs
    OtherVrfs                 []string                                 `json:"other_vrfs,omitempty"`
    // whether this network is routable
    Routed                    *bool                                    `json:"routed,omitempty"`
    // if `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub
    SourceNat                 *NetworkSourceNat                        `json:"source_nat,omitempty"`
    // Property key may be an IP Address (i.e. "172.16.0.1"), and IP Address and Port (i.e. "172.16.0.1:8443") or a CIDR (i.e. "172.16.0.12/20")
    StaticNat                 map[string]NetworkStaticNatProperty      `json:"static_nat,omitempty"`
    // toward overlay
    // how HUB should deal with routes it received from Spokes
    SummarizedSubnet          *string                                  `json:"summarized_subnet,omitempty"`
    // toward LAN-side BGP peers
    SummarizedSubnetToLanBgp  *string                                  `json:"summarized_subnet_to_lan_bgp,omitempty"`
    // toward LAN-side OSPF peers
    SummarizedSubnetToLanOspf *string                                  `json:"summarized_subnet_to_lan_ospf,omitempty"`
    AdditionalProperties      map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkVpnAccessConfig.
// It customizes the JSON marshaling process for NetworkVpnAccessConfig objects.
func (n NetworkVpnAccessConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkVpnAccessConfig object to a map representation for JSON marshaling.
func (n NetworkVpnAccessConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.AdvertisedSubnet != nil {
        structMap["advertised_subnet"] = n.AdvertisedSubnet
    }
    if n.AllowPing != nil {
        structMap["allow_ping"] = n.AllowPing
    }
    if n.DestinationNat != nil {
        structMap["destination_nat"] = n.DestinationNat
    }
    if n.NatPool != nil {
        structMap["nat_pool"] = n.NatPool
    }
    if n.NoReadvertiseToLanBgp != nil {
        structMap["no_readvertise_to_lan_bgp"] = n.NoReadvertiseToLanBgp
    }
    if n.NoReadvertiseToLanOspf != nil {
        structMap["no_readvertise_to_lan_ospf"] = n.NoReadvertiseToLanOspf
    }
    if n.NoReadvertiseToOverlay != nil {
        structMap["no_readvertise_to_overlay"] = n.NoReadvertiseToOverlay
    }
    if n.OtherVrfs != nil {
        structMap["other_vrfs"] = n.OtherVrfs
    }
    if n.Routed != nil {
        structMap["routed"] = n.Routed
    }
    if n.SourceNat != nil {
        structMap["source_nat"] = n.SourceNat.toMap()
    }
    if n.StaticNat != nil {
        structMap["static_nat"] = n.StaticNat
    }
    if n.SummarizedSubnet != nil {
        structMap["summarized_subnet"] = n.SummarizedSubnet
    }
    if n.SummarizedSubnetToLanBgp != nil {
        structMap["summarized_subnet_to_lan_bgp"] = n.SummarizedSubnetToLanBgp
    }
    if n.SummarizedSubnetToLanOspf != nil {
        structMap["summarized_subnet_to_lan_ospf"] = n.SummarizedSubnetToLanOspf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkVpnAccessConfig.
// It customizes the JSON unmarshaling process for NetworkVpnAccessConfig objects.
func (n *NetworkVpnAccessConfig) UnmarshalJSON(input []byte) error {
    var temp networkVpnAccessConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "advertised_subnet", "allow_ping", "destination_nat", "nat_pool", "no_readvertise_to_lan_bgp", "no_readvertise_to_lan_ospf", "no_readvertise_to_overlay", "other_vrfs", "routed", "source_nat", "static_nat", "summarized_subnet", "summarized_subnet_to_lan_bgp", "summarized_subnet_to_lan_ospf")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.AdvertisedSubnet = temp.AdvertisedSubnet
    n.AllowPing = temp.AllowPing
    n.DestinationNat = temp.DestinationNat
    n.NatPool = temp.NatPool
    n.NoReadvertiseToLanBgp = temp.NoReadvertiseToLanBgp
    n.NoReadvertiseToLanOspf = temp.NoReadvertiseToLanOspf
    n.NoReadvertiseToOverlay = temp.NoReadvertiseToOverlay
    n.OtherVrfs = temp.OtherVrfs
    n.Routed = temp.Routed
    n.SourceNat = temp.SourceNat
    n.StaticNat = temp.StaticNat
    n.SummarizedSubnet = temp.SummarizedSubnet
    n.SummarizedSubnetToLanBgp = temp.SummarizedSubnetToLanBgp
    n.SummarizedSubnetToLanOspf = temp.SummarizedSubnetToLanOspf
    return nil
}

// networkVpnAccessConfig is a temporary struct used for validating the fields of NetworkVpnAccessConfig.
type networkVpnAccessConfig  struct {
    AdvertisedSubnet          *string                                  `json:"advertised_subnet,omitempty"`
    AllowPing                 *bool                                    `json:"allow_ping,omitempty"`
    DestinationNat            map[string]NetworkDestinationNatProperty `json:"destination_nat,omitempty"`
    NatPool                   *string                                  `json:"nat_pool,omitempty"`
    NoReadvertiseToLanBgp     *bool                                    `json:"no_readvertise_to_lan_bgp,omitempty"`
    NoReadvertiseToLanOspf    *bool                                    `json:"no_readvertise_to_lan_ospf,omitempty"`
    NoReadvertiseToOverlay    *bool                                    `json:"no_readvertise_to_overlay,omitempty"`
    OtherVrfs                 []string                                 `json:"other_vrfs,omitempty"`
    Routed                    *bool                                    `json:"routed,omitempty"`
    SourceNat                 *NetworkSourceNat                        `json:"source_nat,omitempty"`
    StaticNat                 map[string]NetworkStaticNatProperty      `json:"static_nat,omitempty"`
    SummarizedSubnet          *string                                  `json:"summarized_subnet,omitempty"`
    SummarizedSubnetToLanBgp  *string                                  `json:"summarized_subnet_to_lan_bgp,omitempty"`
    SummarizedSubnetToLanOspf *string                                  `json:"summarized_subnet_to_lan_ospf,omitempty"`
}