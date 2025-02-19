package models

import (
    "encoding/json"
    "fmt"
)

// EvpnOptions represents a EvpnOptions struct.
// EVPN Options
type EvpnOptions struct {
    // Optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides
    AutoLoopbackSubnet   *string                          `json:"auto_loopback_subnet,omitempty"`
    // Optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides
    AutoLoopbackSubnet6  *string                          `json:"auto_loopback_subnet6,omitempty"`
    // Optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored
    AutoRouterIdSubnet   *string                          `json:"auto_router_id_subnet,omitempty"`
    // Optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored
    AutoRouterIdSubnet6  *string                          `json:"auto_router_id_subnet6,omitempty"`
    // Optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway. When `routed_at` != `core`, whether to do virtual-gateway at core as well
    CoreAsBorder         *bool                            `json:"core_as_border,omitempty"`
    Overlay              *EvpnOptionsOverlay              `json:"overlay,omitempty"`
    // Only for by Core-Distribution architecture when `evpn_options.routed_at`==`core`. By default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address's v4_mac. If enabled, 00-00-5e-00-XX-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)
    PerVlanVgaV4Mac      *bool                            `json:"per_vlan_vga_v4_mac,omitempty"`
    // optional, where virtual-gateway should reside. enum: `core`, `distribution`, `edge`
    RoutedAt             *EvpnOptionsRoutedAtEnum         `json:"routed_at,omitempty"`
    Underlay             *EvpnOptionsUnderlay             `json:"underlay,omitempty"`
    // Optional, for EX9200 only to seggregate virtual-switches
    VsInstances          map[string]EvpnOptionsVsInstance `json:"vs_instances,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnOptions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnOptions) String() string {
    return fmt.Sprintf(
    	"EvpnOptions[AutoLoopbackSubnet=%v, AutoLoopbackSubnet6=%v, AutoRouterIdSubnet=%v, AutoRouterIdSubnet6=%v, CoreAsBorder=%v, Overlay=%v, PerVlanVgaV4Mac=%v, RoutedAt=%v, Underlay=%v, VsInstances=%v, AdditionalProperties=%v]",
    	e.AutoLoopbackSubnet, e.AutoLoopbackSubnet6, e.AutoRouterIdSubnet, e.AutoRouterIdSubnet6, e.CoreAsBorder, e.Overlay, e.PerVlanVgaV4Mac, e.RoutedAt, e.Underlay, e.VsInstances, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnOptions.
// It customizes the JSON marshaling process for EvpnOptions objects.
func (e EvpnOptions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "auto_loopback_subnet", "auto_loopback_subnet6", "auto_router_id_subnet", "auto_router_id_subnet6", "core_as_border", "overlay", "per_vlan_vga_v4_mac", "routed_at", "underlay", "vs_instances"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnOptions object to a map representation for JSON marshaling.
func (e EvpnOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.AutoLoopbackSubnet != nil {
        structMap["auto_loopback_subnet"] = e.AutoLoopbackSubnet
    }
    if e.AutoLoopbackSubnet6 != nil {
        structMap["auto_loopback_subnet6"] = e.AutoLoopbackSubnet6
    }
    if e.AutoRouterIdSubnet != nil {
        structMap["auto_router_id_subnet"] = e.AutoRouterIdSubnet
    }
    if e.AutoRouterIdSubnet6 != nil {
        structMap["auto_router_id_subnet6"] = e.AutoRouterIdSubnet6
    }
    if e.CoreAsBorder != nil {
        structMap["core_as_border"] = e.CoreAsBorder
    }
    if e.Overlay != nil {
        structMap["overlay"] = e.Overlay.toMap()
    }
    if e.PerVlanVgaV4Mac != nil {
        structMap["per_vlan_vga_v4_mac"] = e.PerVlanVgaV4Mac
    }
    if e.RoutedAt != nil {
        structMap["routed_at"] = e.RoutedAt
    }
    if e.Underlay != nil {
        structMap["underlay"] = e.Underlay.toMap()
    }
    if e.VsInstances != nil {
        structMap["vs_instances"] = e.VsInstances
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnOptions.
// It customizes the JSON unmarshaling process for EvpnOptions objects.
func (e *EvpnOptions) UnmarshalJSON(input []byte) error {
    var temp tempEvpnOptions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_loopback_subnet", "auto_loopback_subnet6", "auto_router_id_subnet", "auto_router_id_subnet6", "core_as_border", "overlay", "per_vlan_vga_v4_mac", "routed_at", "underlay", "vs_instances")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.AutoLoopbackSubnet = temp.AutoLoopbackSubnet
    e.AutoLoopbackSubnet6 = temp.AutoLoopbackSubnet6
    e.AutoRouterIdSubnet = temp.AutoRouterIdSubnet
    e.AutoRouterIdSubnet6 = temp.AutoRouterIdSubnet6
    e.CoreAsBorder = temp.CoreAsBorder
    e.Overlay = temp.Overlay
    e.PerVlanVgaV4Mac = temp.PerVlanVgaV4Mac
    e.RoutedAt = temp.RoutedAt
    e.Underlay = temp.Underlay
    e.VsInstances = temp.VsInstances
    return nil
}

// tempEvpnOptions is a temporary struct used for validating the fields of EvpnOptions.
type tempEvpnOptions  struct {
    AutoLoopbackSubnet  *string                          `json:"auto_loopback_subnet,omitempty"`
    AutoLoopbackSubnet6 *string                          `json:"auto_loopback_subnet6,omitempty"`
    AutoRouterIdSubnet  *string                          `json:"auto_router_id_subnet,omitempty"`
    AutoRouterIdSubnet6 *string                          `json:"auto_router_id_subnet6,omitempty"`
    CoreAsBorder        *bool                            `json:"core_as_border,omitempty"`
    Overlay             *EvpnOptionsOverlay              `json:"overlay,omitempty"`
    PerVlanVgaV4Mac     *bool                            `json:"per_vlan_vga_v4_mac,omitempty"`
    RoutedAt            *EvpnOptionsRoutedAtEnum         `json:"routed_at,omitempty"`
    Underlay            *EvpnOptionsUnderlay             `json:"underlay,omitempty"`
    VsInstances         map[string]EvpnOptionsVsInstance `json:"vs_instances,omitempty"`
}
