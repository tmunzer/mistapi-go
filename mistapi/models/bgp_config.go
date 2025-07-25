// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// BgpConfig represents a BgpConfig struct.
// BFD is enabled when either bfd_minimum_interval or bfd_multiplier is configured
type BgpConfig struct {
    AuthKey                *string                       `json:"auth_key,omitempty"`
    // When bfd_multiplier is configured alone. Default:
    // * 1000 if `type`==`external`
    // * 350 `type`==`internal`
    BfdMinimumInterval     Optional[int]                 `json:"bfd_minimum_interval"`
    // When bfd_minimum_interval_is_configured alone
    BfdMultiplier          Optional[int]                 `json:"bfd_multiplier"`
    // BFD provides faster path failure detection and is enabled by default
    DisableBfd             *bool                         `json:"disable_bfd,omitempty"`
    Export                 *string                       `json:"export,omitempty"`
    // Default export policies if no per-neighbor policies defined
    ExportPolicy           *string                       `json:"export_policy,omitempty"`
    // By default, either inet/net6 unicast depending on neighbor IP family (v4 or v6). For v6 neighbors, to exchange v4 nexthop, which allows dual-stack support, enable this
    ExtendedV4Nexthop      *bool                         `json:"extended_v4_nexthop,omitempty"`
    // `0` means disable
    GracefulRestartTime    *int                          `json:"graceful_restart_time,omitempty"`
    HoldTime               *int                          `json:"hold_time,omitempty"`
    Import                 *string                       `json:"import,omitempty"`
    // Default import policies if no per-neighbor policies defined
    ImportPolicy           *string                       `json:"import_policy,omitempty"`
    // BGP AS, value in range 1-4294967295
    LocalAs                *BgpAs                        `json:"local_as,omitempty"`
    // BGP AS, value in range 1-4294967295
    NeighborAs             *BgpAs                        `json:"neighbor_as,omitempty"`
    // If per-neighbor as is desired. Property key is the neighbor address
    Neighbors              map[string]BgpConfigNeighbors `json:"neighbors,omitempty"`
    // If `type`!=`external`or `via`==`wan`networks where we expect BGP neighbor to connect to/from
    Networks               []string                      `json:"networks,omitempty"`
    NoPrivateAs            *bool                         `json:"no_private_as,omitempty"`
    // By default, we'll re-advertise all learned BGP routers toward overlay
    NoReadvertiseToOverlay *bool                         `json:"no_readvertise_to_overlay,omitempty"`
    // If `type`==`tunnel`
    TunnelName             *string                       `json:"tunnel_name,omitempty"`
    // enum: `external`, `internal`
    Type                   *BgpConfigTypeEnum            `json:"type,omitempty"`
    // network name. enum: `lan`, `tunnel`, `vpn`, `wan`
    Via                    *BgpConfigViaEnum             `json:"via,omitempty"`
    VpnName                *string                       `json:"vpn_name,omitempty"`
    // If `via`==`wan`
    WanName                *string                       `json:"wan_name,omitempty"`
    AdditionalProperties   map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for BgpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BgpConfig) String() string {
    return fmt.Sprintf(
    	"BgpConfig[AuthKey=%v, BfdMinimumInterval=%v, BfdMultiplier=%v, DisableBfd=%v, Export=%v, ExportPolicy=%v, ExtendedV4Nexthop=%v, GracefulRestartTime=%v, HoldTime=%v, Import=%v, ImportPolicy=%v, LocalAs=%v, NeighborAs=%v, Neighbors=%v, Networks=%v, NoPrivateAs=%v, NoReadvertiseToOverlay=%v, TunnelName=%v, Type=%v, Via=%v, VpnName=%v, WanName=%v, AdditionalProperties=%v]",
    	b.AuthKey, b.BfdMinimumInterval, b.BfdMultiplier, b.DisableBfd, b.Export, b.ExportPolicy, b.ExtendedV4Nexthop, b.GracefulRestartTime, b.HoldTime, b.Import, b.ImportPolicy, b.LocalAs, b.NeighborAs, b.Neighbors, b.Networks, b.NoPrivateAs, b.NoReadvertiseToOverlay, b.TunnelName, b.Type, b.Via, b.VpnName, b.WanName, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BgpConfig.
// It customizes the JSON marshaling process for BgpConfig objects.
func (b BgpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "auth_key", "bfd_minimum_interval", "bfd_multiplier", "disable_bfd", "export", "export_policy", "extended_v4_nexthop", "graceful_restart_time", "hold_time", "import", "import_policy", "local_as", "neighbor_as", "neighbors", "networks", "no_private_as", "no_readvertise_to_overlay", "tunnel_name", "type", "via", "vpn_name", "wan_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BgpConfig object to a map representation for JSON marshaling.
func (b BgpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
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
        structMap["local_as"] = b.LocalAs.toMap()
    }
    if b.NeighborAs != nil {
        structMap["neighbor_as"] = b.NeighborAs.toMap()
    }
    if b.Neighbors != nil {
        structMap["neighbors"] = b.Neighbors
    }
    if b.Networks != nil {
        structMap["networks"] = b.Networks
    }
    if b.NoPrivateAs != nil {
        structMap["no_private_as"] = b.NoPrivateAs
    }
    if b.NoReadvertiseToOverlay != nil {
        structMap["no_readvertise_to_overlay"] = b.NoReadvertiseToOverlay
    }
    if b.TunnelName != nil {
        structMap["tunnel_name"] = b.TunnelName
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
    var temp tempBgpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_key", "bfd_minimum_interval", "bfd_multiplier", "disable_bfd", "export", "export_policy", "extended_v4_nexthop", "graceful_restart_time", "hold_time", "import", "import_policy", "local_as", "neighbor_as", "neighbors", "networks", "no_private_as", "no_readvertise_to_overlay", "tunnel_name", "type", "via", "vpn_name", "wan_name")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.AuthKey = temp.AuthKey
    b.BfdMinimumInterval = temp.BfdMinimumInterval
    b.BfdMultiplier = temp.BfdMultiplier
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
    b.NoPrivateAs = temp.NoPrivateAs
    b.NoReadvertiseToOverlay = temp.NoReadvertiseToOverlay
    b.TunnelName = temp.TunnelName
    b.Type = temp.Type
    b.Via = temp.Via
    b.VpnName = temp.VpnName
    b.WanName = temp.WanName
    return nil
}

// tempBgpConfig is a temporary struct used for validating the fields of BgpConfig.
type tempBgpConfig  struct {
    AuthKey                *string                       `json:"auth_key,omitempty"`
    BfdMinimumInterval     Optional[int]                 `json:"bfd_minimum_interval"`
    BfdMultiplier          Optional[int]                 `json:"bfd_multiplier"`
    DisableBfd             *bool                         `json:"disable_bfd,omitempty"`
    Export                 *string                       `json:"export,omitempty"`
    ExportPolicy           *string                       `json:"export_policy,omitempty"`
    ExtendedV4Nexthop      *bool                         `json:"extended_v4_nexthop,omitempty"`
    GracefulRestartTime    *int                          `json:"graceful_restart_time,omitempty"`
    HoldTime               *int                          `json:"hold_time,omitempty"`
    Import                 *string                       `json:"import,omitempty"`
    ImportPolicy           *string                       `json:"import_policy,omitempty"`
    LocalAs                *BgpAs                        `json:"local_as,omitempty"`
    NeighborAs             *BgpAs                        `json:"neighbor_as,omitempty"`
    Neighbors              map[string]BgpConfigNeighbors `json:"neighbors,omitempty"`
    Networks               []string                      `json:"networks,omitempty"`
    NoPrivateAs            *bool                         `json:"no_private_as,omitempty"`
    NoReadvertiseToOverlay *bool                         `json:"no_readvertise_to_overlay,omitempty"`
    TunnelName             *string                       `json:"tunnel_name,omitempty"`
    Type                   *BgpConfigTypeEnum            `json:"type,omitempty"`
    Via                    *BgpConfigViaEnum             `json:"via,omitempty"`
    VpnName                *string                       `json:"vpn_name,omitempty"`
    WanName                *string                       `json:"wan_name,omitempty"`
}
