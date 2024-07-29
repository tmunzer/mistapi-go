package models

import (
    "encoding/json"
)

// BgpConfig represents a BgpConfig struct.
// BFD is enabled when either bfd_minimum_interval or bfd_multiplier is configured
type BgpConfig struct {
    AuthKey                *string                       `json:"auth_key,omitempty"`
    // when bfd_multiplier is configured alone. Default:
    //   * 1000 if `type`==`external`
    //   * 350 `type`==`internal`
    BfdMinimumInterval     Optional[int]                 `json:"bfd_minimum_interval"`
    // when bfd_minimum_interval_is_configured alone
    BfdMultiplier          Optional[int]                 `json:"bfd_multiplier"`
    Communities            []BgpConfigCommunity          `json:"communities,omitempty"`
    // BFD provides faster path failure detection and is enabled by default
    DisableBfd             *bool                         `json:"disable_bfd,omitempty"`
    Export                 *string                       `json:"export,omitempty"`
    // default export policies if no per-neighbor policies defined
    ExportPolicy           *string                       `json:"export_policy,omitempty"`
    // by default, either inet/net6 unicast depending on neighbor IP family (v4 or v6)
    // for v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this
    ExtendedV4Nexthop      *bool                         `json:"extended_v4_nexthop,omitempty"`
    // `0` means disable
    GracefulRestartTime    *int                          `json:"graceful_restart_time,omitempty"`
    HoldTime               *int                          `json:"hold_time,omitempty"`
    Import                 *string                       `json:"import,omitempty"`
    // default import policies if no per-neighbor policies defined
    ImportPolicy           *string                       `json:"import_policy,omitempty"`
    LocalAs                *int                          `json:"local_as,omitempty"`
    NeighborAs             *int                          `json:"neighbor_as,omitempty"`
    // if per-neighbor as is desired. Property key is the neighbor address
    Neighbors              map[string]BgpConfigNeighbors `json:"neighbors,omitempty"`
    // if `type`!=`external`or `via`==`wan`networks where we expect BGP neighbor to connect to/from
    Networks               []string                      `json:"networks,omitempty"`
    // by default, we'll re-advertise all learned BGP routers toward overlay
    NoReadvertiseToOverlay *bool                         `json:"no_readvertise_to_overlay,omitempty"`
    // enum: `external`, `internal`
    Type                   *BgpConfigTypeEnum            `json:"type,omitempty"`
    // network name. enum: `lan`, `vpn`, `wan`
    Via                    *BgpConfigViaEnum             `json:"via,omitempty"`
    VpnName                *string                       `json:"vpn_name,omitempty"`
    // if `via`==`wan`
    WanName                *string                       `json:"wan_name,omitempty"`
    AdditionalProperties   map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BgpConfig.
// It customizes the JSON marshaling process for BgpConfig objects.
func (b BgpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BgpConfig object to a map representation for JSON marshaling.
func (b BgpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.AuthKey != nil {
        structMap["auth_key"] = b.AuthKey
    }
    if b.BfdMinimumInterval.IsValueSet() {
        if b.BfdMinimumInterval.Value() != nil {
            structMap["bfd_minimum_interval"] = b.BfdMinimumInterval.Value()
        } else {
            structMap["bfd_minimum_interval"] = nil
        }
    }
    if b.BfdMultiplier.IsValueSet() {
        if b.BfdMultiplier.Value() != nil {
            structMap["bfd_multiplier"] = b.BfdMultiplier.Value()
        } else {
            structMap["bfd_multiplier"] = nil
        }
    }
    if b.Communities != nil {
        structMap["communities"] = b.Communities
    }
    if b.DisableBfd != nil {
        structMap["disable_bfd"] = b.DisableBfd
    }
    if b.Export != nil {
        structMap["export"] = b.Export
    }
    if b.ExportPolicy != nil {
        structMap["export_policy"] = b.ExportPolicy
    }
    if b.ExtendedV4Nexthop != nil {
        structMap["extended_v4_nexthop"] = b.ExtendedV4Nexthop
    }
    if b.GracefulRestartTime != nil {
        structMap["graceful_restart_time"] = b.GracefulRestartTime
    }
    if b.HoldTime != nil {
        structMap["hold_time"] = b.HoldTime
    }
    if b.Import != nil {
        structMap["import"] = b.Import
    }
    if b.ImportPolicy != nil {
        structMap["import_policy"] = b.ImportPolicy
    }
    if b.LocalAs != nil {
        structMap["local_as"] = b.LocalAs
    }
    if b.NeighborAs != nil {
        structMap["neighbor_as"] = b.NeighborAs
    }
    if b.Neighbors != nil {
        structMap["neighbors"] = b.Neighbors
    }
    if b.Networks != nil {
        structMap["networks"] = b.Networks
    }
    if b.NoReadvertiseToOverlay != nil {
        structMap["no_readvertise_to_overlay"] = b.NoReadvertiseToOverlay
    }
    if b.Type != nil {
        structMap["type"] = b.Type
    }
    if b.Via != nil {
        structMap["via"] = b.Via
    }
    if b.VpnName != nil {
        structMap["vpn_name"] = b.VpnName
    }
    if b.WanName != nil {
        structMap["wan_name"] = b.WanName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BgpConfig.
// It customizes the JSON unmarshaling process for BgpConfig objects.
func (b *BgpConfig) UnmarshalJSON(input []byte) error {
    var temp bgpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auth_key", "bfd_minimum_interval", "bfd_multiplier", "communities", "disable_bfd", "export", "export_policy", "extended_v4_nexthop", "graceful_restart_time", "hold_time", "import", "import_policy", "local_as", "neighbor_as", "neighbors", "networks", "no_readvertise_to_overlay", "type", "via", "vpn_name", "wan_name")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.AuthKey = temp.AuthKey
    b.BfdMinimumInterval = temp.BfdMinimumInterval
    b.BfdMultiplier = temp.BfdMultiplier
    b.Communities = temp.Communities
    b.DisableBfd = temp.DisableBfd
    b.Export = temp.Export
    b.ExportPolicy = temp.ExportPolicy
    b.ExtendedV4Nexthop = temp.ExtendedV4Nexthop
    b.GracefulRestartTime = temp.GracefulRestartTime
    b.HoldTime = temp.HoldTime
    b.Import = temp.Import
    b.ImportPolicy = temp.ImportPolicy
    b.LocalAs = temp.LocalAs
    b.NeighborAs = temp.NeighborAs
    b.Neighbors = temp.Neighbors
    b.Networks = temp.Networks
    b.NoReadvertiseToOverlay = temp.NoReadvertiseToOverlay
    b.Type = temp.Type
    b.Via = temp.Via
    b.VpnName = temp.VpnName
    b.WanName = temp.WanName
    return nil
}

// bgpConfig is a temporary struct used for validating the fields of BgpConfig.
type bgpConfig  struct {
    AuthKey                *string                       `json:"auth_key,omitempty"`
    BfdMinimumInterval     Optional[int]                 `json:"bfd_minimum_interval"`
    BfdMultiplier          Optional[int]                 `json:"bfd_multiplier"`
    Communities            []BgpConfigCommunity          `json:"communities,omitempty"`
    DisableBfd             *bool                         `json:"disable_bfd,omitempty"`
    Export                 *string                       `json:"export,omitempty"`
    ExportPolicy           *string                       `json:"export_policy,omitempty"`
    ExtendedV4Nexthop      *bool                         `json:"extended_v4_nexthop,omitempty"`
    GracefulRestartTime    *int                          `json:"graceful_restart_time,omitempty"`
    HoldTime               *int                          `json:"hold_time,omitempty"`
    Import                 *string                       `json:"import,omitempty"`
    ImportPolicy           *string                       `json:"import_policy,omitempty"`
    LocalAs                *int                          `json:"local_as,omitempty"`
    NeighborAs             *int                          `json:"neighbor_as,omitempty"`
    Neighbors              map[string]BgpConfigNeighbors `json:"neighbors,omitempty"`
    Networks               []string                      `json:"networks,omitempty"`
    NoReadvertiseToOverlay *bool                         `json:"no_readvertise_to_overlay,omitempty"`
    Type                   *BgpConfigTypeEnum            `json:"type,omitempty"`
    Via                    *BgpConfigViaEnum             `json:"via,omitempty"`
    VpnName                *string                       `json:"vpn_name,omitempty"`
    WanName                *string                       `json:"wan_name,omitempty"`
}
