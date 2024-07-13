package models

import (
    "encoding/json"
)

// ConstDeviceModel represents a ConstDeviceModel struct.
type ConstDeviceModel struct {
    ApType               *string                        `json:"ap_type,omitempty"`
    Band24               *ConstDeviceApBand24           `json:"band24,omitempty"`
    Band5                *ConstDeviceApBand5            `json:"band5,omitempty"`
    Band6                *ConstDeviceApBand5            `json:"band6,omitempty"`
    Band24Usages         *RadioBand24UsageEnum          `json:"band_24_usages,omitempty"`
    CeDfsOk              *bool                          `json:"ce_dfs_ok,omitempty"`
    CiscoPace            *bool                          `json:"cisco_pace,omitempty"`
    Description          *string                        `json:"description,omitempty"`
    // Property key is a list of country codes (e.g. "GB, DE")
    DisallowedChannels   map[string]interface{}         `json:"disallowed_channels,omitempty"`
    Display              *string                        `json:"display,omitempty"`
    // Property key is the GPIO port name (e.g. "D0", "A1")
    Extio                map[string]ConstDeviceApExtios `json:"extio,omitempty"`
    FccDfsOk             *bool                          `json:"fcc_dfs_ok,omitempty"`
    Has11ax              *bool                          `json:"has_11ax,omitempty"`
    HasCompass           *bool                          `json:"has_compass,omitempty"`
    HasExtAnt            *bool                          `json:"has_ext_ant,omitempty"`
    HasExtio             *bool                          `json:"has_extio,omitempty"`
    HasHeight            *bool                          `json:"has_height,omitempty"`
    HasModulePort        *bool                          `json:"has_module_port,omitempty"`
    HasPoeOut            *bool                          `json:"has_poe_out,omitempty"`
    HasScanningRadio     *bool                          `json:"has_scanning_radio,omitempty"`
    HasSelectableRadio   *bool                          `json:"has_selectable_radio,omitempty"`
    HasUsb               *bool                          `json:"has_usb,omitempty"`
    HasVble              *bool                          `json:"has_vble,omitempty"`
    HasWifiBand24        *bool                          `json:"has_wifi_band24,omitempty"`
    HasWifiBand5         *bool                          `json:"has_wifi_band5,omitempty"`
    HasWifiBand6         *bool                          `json:"has_wifi_band6,omitempty"`
    MaxPoeOut            *int                           `json:"max_poe_out,omitempty"`
    MaxWlans             *int                           `json:"max_wlans,omitempty"`
    Model                *string                        `json:"model,omitempty"`
    OtherDfsOk           *bool                          `json:"other_dfs_ok,omitempty"`
    Outdoor              *bool                          `json:"outdoor,omitempty"`
    // Property key is the radio number (e.g. r0, r1, ...). Property value is the RF band (e.g. "24", "5", ...)
    Radios               map[string]string              `json:"radios,omitempty"`
    SharedScanningRadio  *bool                          `json:"shared_scanning_radio,omitempty"`
    Type                 *string                        `json:"type,omitempty"`
    Unmanaged            *bool                          `json:"unmanaged,omitempty"`
    Vble                 *ConstDeviceApVble             `json:"vble,omitempty"`
    Alias                *string                        `json:"alias,omitempty"`
    // Object Key is the interface type name (e.g. "lan_ports", "wan_ports", ...)
    Defaults             *ConstDeviceSwitchDefault1     `json:"defaults,omitempty"`
    EvolvedOs            *bool                          `json:"evolved_os,omitempty"`
    EvpnRiType           *string                        `json:"evpn_ri_type,omitempty"`
    Experimental         *bool                          `json:"experimental,omitempty"`
    FansPluggable        *bool                          `json:"fans_pluggable,omitempty"`
    HasBgp               *bool                          `json:"has_bgp,omitempty"`
    HasEts               *bool                          `json:"has_ets,omitempty"`
    HasEvpn              *bool                          `json:"has_evpn,omitempty"`
    HasIrb               *bool                          `json:"has_irb,omitempty"`
    HasSnapshot          *bool                          `json:"has_snapshot,omitempty"`
    HasVc                *bool                          `json:"has_vc,omitempty"`
    Modular              *bool                          `json:"modular,omitempty"`
    NoShapingRate        *bool                          `json:"no_shaping_rate,omitempty"`
    NumberFans           *int                           `json:"number_fans,omitempty"`
    OcDevice             *bool                          `json:"oc_device,omitempty"`
    OobInterface         *string                        `json:"oob_interface,omitempty"`
    PacketActionDropOnly *bool                          `json:"packet_action_drop_only,omitempty"`
    // Object Key is the PIC number
    Pic                  map[string]string              `json:"pic,omitempty"`
    SubRequired          *string                        `json:"sub_required,omitempty"`
    HaNode0Fpc           *int                           `json:"ha_node0_fpc,omitempty"`
    HaNode1Fpc           *int                           `json:"ha_node1_fpc,omitempty"`
    HasFxp0              *bool                          `json:"has_fxp0,omitempty"`
    HasHaControl         *bool                          `json:"has_ha_control,omitempty"`
    HasHaData            *bool                          `json:"has_ha_data,omitempty"`
    IrbDisabledByDefault *bool                          `json:"irb_disabled_by_default,omitempty"`
    // Object Key is the interface name (e.g. "ge-0/0/1", ...)
    Ports                *ConstDeviceGatewayPorts       `json:"ports,omitempty"`
    T128Device           *bool                          `json:"t128_device,omitempty"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceModel.
// It customizes the JSON marshaling process for ConstDeviceModel objects.
func (c ConstDeviceModel) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceModel object to a map representation for JSON marshaling.
func (c ConstDeviceModel) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ApType != nil {
        structMap["ap_type"] = c.ApType
    }
    if c.Band24 != nil {
        structMap["band24"] = c.Band24.toMap()
    }
    if c.Band5 != nil {
        structMap["band5"] = c.Band5.toMap()
    }
    if c.Band6 != nil {
        structMap["band6"] = c.Band6.toMap()
    }
    if c.Band24Usages != nil {
        structMap["band_24_usages"] = c.Band24Usages
    }
    if c.CeDfsOk != nil {
        structMap["ce_dfs_ok"] = c.CeDfsOk
    }
    if c.CiscoPace != nil {
        structMap["cisco_pace"] = c.CiscoPace
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.DisallowedChannels != nil {
        structMap["disallowed_channels"] = c.DisallowedChannels
    }
    if c.Display != nil {
        structMap["display"] = c.Display
    }
    if c.Extio != nil {
        structMap["extio"] = c.Extio
    }
    if c.FccDfsOk != nil {
        structMap["fcc_dfs_ok"] = c.FccDfsOk
    }
    if c.Has11ax != nil {
        structMap["has_11ax"] = c.Has11ax
    }
    if c.HasCompass != nil {
        structMap["has_compass"] = c.HasCompass
    }
    if c.HasExtAnt != nil {
        structMap["has_ext_ant"] = c.HasExtAnt
    }
    if c.HasExtio != nil {
        structMap["has_extio"] = c.HasExtio
    }
    if c.HasHeight != nil {
        structMap["has_height"] = c.HasHeight
    }
    if c.HasModulePort != nil {
        structMap["has_module_port"] = c.HasModulePort
    }
    if c.HasPoeOut != nil {
        structMap["has_poe_out"] = c.HasPoeOut
    }
    if c.HasScanningRadio != nil {
        structMap["has_scanning_radio"] = c.HasScanningRadio
    }
    if c.HasSelectableRadio != nil {
        structMap["has_selectable_radio"] = c.HasSelectableRadio
    }
    if c.HasUsb != nil {
        structMap["has_usb"] = c.HasUsb
    }
    if c.HasVble != nil {
        structMap["has_vble"] = c.HasVble
    }
    if c.HasWifiBand24 != nil {
        structMap["has_wifi_band24"] = c.HasWifiBand24
    }
    if c.HasWifiBand5 != nil {
        structMap["has_wifi_band5"] = c.HasWifiBand5
    }
    if c.HasWifiBand6 != nil {
        structMap["has_wifi_band6"] = c.HasWifiBand6
    }
    if c.MaxPoeOut != nil {
        structMap["max_poe_out"] = c.MaxPoeOut
    }
    if c.MaxWlans != nil {
        structMap["max_wlans"] = c.MaxWlans
    }
    if c.Model != nil {
        structMap["model"] = c.Model
    }
    if c.OtherDfsOk != nil {
        structMap["other_dfs_ok"] = c.OtherDfsOk
    }
    if c.Outdoor != nil {
        structMap["outdoor"] = c.Outdoor
    }
    if c.Radios != nil {
        structMap["radios"] = c.Radios
    }
    if c.SharedScanningRadio != nil {
        structMap["shared_scanning_radio"] = c.SharedScanningRadio
    }
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    if c.Unmanaged != nil {
        structMap["unmanaged"] = c.Unmanaged
    }
    if c.Vble != nil {
        structMap["vble"] = c.Vble.toMap()
    }
    if c.Alias != nil {
        structMap["alias"] = c.Alias
    }
    if c.Defaults != nil {
        structMap["defaults"] = c.Defaults.toMap()
    }
    if c.EvolvedOs != nil {
        structMap["evolved_os"] = c.EvolvedOs
    }
    if c.EvpnRiType != nil {
        structMap["evpn_ri_type"] = c.EvpnRiType
    }
    if c.Experimental != nil {
        structMap["experimental"] = c.Experimental
    }
    if c.FansPluggable != nil {
        structMap["fans_pluggable"] = c.FansPluggable
    }
    if c.HasBgp != nil {
        structMap["has_bgp"] = c.HasBgp
    }
    if c.HasEts != nil {
        structMap["has_ets"] = c.HasEts
    }
    if c.HasEvpn != nil {
        structMap["has_evpn"] = c.HasEvpn
    }
    if c.HasIrb != nil {
        structMap["has_irb"] = c.HasIrb
    }
    if c.HasSnapshot != nil {
        structMap["has_snapshot"] = c.HasSnapshot
    }
    if c.HasVc != nil {
        structMap["has_vc"] = c.HasVc
    }
    if c.Modular != nil {
        structMap["modular"] = c.Modular
    }
    if c.NoShapingRate != nil {
        structMap["no_shaping_rate"] = c.NoShapingRate
    }
    if c.NumberFans != nil {
        structMap["number_fans"] = c.NumberFans
    }
    if c.OcDevice != nil {
        structMap["oc_device"] = c.OcDevice
    }
    if c.OobInterface != nil {
        structMap["oob_interface"] = c.OobInterface
    }
    if c.PacketActionDropOnly != nil {
        structMap["packet_action_drop_only"] = c.PacketActionDropOnly
    }
    if c.Pic != nil {
        structMap["pic"] = c.Pic
    }
    if c.SubRequired != nil {
        structMap["sub_required"] = c.SubRequired
    }
    if c.HaNode0Fpc != nil {
        structMap["ha_node0_fpc"] = c.HaNode0Fpc
    }
    if c.HaNode1Fpc != nil {
        structMap["ha_node1_fpc"] = c.HaNode1Fpc
    }
    if c.HasFxp0 != nil {
        structMap["has_fxp0"] = c.HasFxp0
    }
    if c.HasHaControl != nil {
        structMap["has_ha_control"] = c.HasHaControl
    }
    if c.HasHaData != nil {
        structMap["has_ha_data"] = c.HasHaData
    }
    if c.IrbDisabledByDefault != nil {
        structMap["irb_disabled_by_default"] = c.IrbDisabledByDefault
    }
    if c.Ports != nil {
        structMap["ports"] = c.Ports.toMap()
    }
    if c.T128Device != nil {
        structMap["t128_device"] = c.T128Device
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceModel.
// It customizes the JSON unmarshaling process for ConstDeviceModel objects.
func (c *ConstDeviceModel) UnmarshalJSON(input []byte) error {
    var temp constDeviceModel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_type", "band24", "band5", "band6", "band_24_usages", "ce_dfs_ok", "cisco_pace", "description", "disallowed_channels", "display", "extio", "fcc_dfs_ok", "has_11ax", "has_compass", "has_ext_ant", "has_extio", "has_height", "has_module_port", "has_poe_out", "has_scanning_radio", "has_selectable_radio", "has_usb", "has_vble", "has_wifi_band24", "has_wifi_band5", "has_wifi_band6", "max_poe_out", "max_wlans", "model", "other_dfs_ok", "outdoor", "radios", "shared_scanning_radio", "type", "unmanaged", "vble", "alias", "defaults", "evolved_os", "evpn_ri_type", "experimental", "fans_pluggable", "has_bgp", "has_ets", "has_evpn", "has_irb", "has_snapshot", "has_vc", "modular", "no_shaping_rate", "number_fans", "oc_device", "oob_interface", "packet_action_drop_only", "pic", "sub_required", "ha_node0_fpc", "ha_node1_fpc", "has_fxp0", "has_ha_control", "has_ha_data", "irb_disabled_by_default", "ports", "t128_device")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ApType = temp.ApType
    c.Band24 = temp.Band24
    c.Band5 = temp.Band5
    c.Band6 = temp.Band6
    c.Band24Usages = temp.Band24Usages
    c.CeDfsOk = temp.CeDfsOk
    c.CiscoPace = temp.CiscoPace
    c.Description = temp.Description
    c.DisallowedChannels = temp.DisallowedChannels
    c.Display = temp.Display
    c.Extio = temp.Extio
    c.FccDfsOk = temp.FccDfsOk
    c.Has11ax = temp.Has11ax
    c.HasCompass = temp.HasCompass
    c.HasExtAnt = temp.HasExtAnt
    c.HasExtio = temp.HasExtio
    c.HasHeight = temp.HasHeight
    c.HasModulePort = temp.HasModulePort
    c.HasPoeOut = temp.HasPoeOut
    c.HasScanningRadio = temp.HasScanningRadio
    c.HasSelectableRadio = temp.HasSelectableRadio
    c.HasUsb = temp.HasUsb
    c.HasVble = temp.HasVble
    c.HasWifiBand24 = temp.HasWifiBand24
    c.HasWifiBand5 = temp.HasWifiBand5
    c.HasWifiBand6 = temp.HasWifiBand6
    c.MaxPoeOut = temp.MaxPoeOut
    c.MaxWlans = temp.MaxWlans
    c.Model = temp.Model
    c.OtherDfsOk = temp.OtherDfsOk
    c.Outdoor = temp.Outdoor
    c.Radios = temp.Radios
    c.SharedScanningRadio = temp.SharedScanningRadio
    c.Type = temp.Type
    c.Unmanaged = temp.Unmanaged
    c.Vble = temp.Vble
    c.Alias = temp.Alias
    c.Defaults = temp.Defaults
    c.EvolvedOs = temp.EvolvedOs
    c.EvpnRiType = temp.EvpnRiType
    c.Experimental = temp.Experimental
    c.FansPluggable = temp.FansPluggable
    c.HasBgp = temp.HasBgp
    c.HasEts = temp.HasEts
    c.HasEvpn = temp.HasEvpn
    c.HasIrb = temp.HasIrb
    c.HasSnapshot = temp.HasSnapshot
    c.HasVc = temp.HasVc
    c.Modular = temp.Modular
    c.NoShapingRate = temp.NoShapingRate
    c.NumberFans = temp.NumberFans
    c.OcDevice = temp.OcDevice
    c.OobInterface = temp.OobInterface
    c.PacketActionDropOnly = temp.PacketActionDropOnly
    c.Pic = temp.Pic
    c.SubRequired = temp.SubRequired
    c.HaNode0Fpc = temp.HaNode0Fpc
    c.HaNode1Fpc = temp.HaNode1Fpc
    c.HasFxp0 = temp.HasFxp0
    c.HasHaControl = temp.HasHaControl
    c.HasHaData = temp.HasHaData
    c.IrbDisabledByDefault = temp.IrbDisabledByDefault
    c.Ports = temp.Ports
    c.T128Device = temp.T128Device
    return nil
}

// constDeviceModel is a temporary struct used for validating the fields of ConstDeviceModel.
type constDeviceModel  struct {
    ApType               *string                        `json:"ap_type,omitempty"`
    Band24               *ConstDeviceApBand24           `json:"band24,omitempty"`
    Band5                *ConstDeviceApBand5            `json:"band5,omitempty"`
    Band6                *ConstDeviceApBand5            `json:"band6,omitempty"`
    Band24Usages         *RadioBand24UsageEnum          `json:"band_24_usages,omitempty"`
    CeDfsOk              *bool                          `json:"ce_dfs_ok,omitempty"`
    CiscoPace            *bool                          `json:"cisco_pace,omitempty"`
    Description          *string                        `json:"description,omitempty"`
    DisallowedChannels   map[string]interface{}         `json:"disallowed_channels,omitempty"`
    Display              *string                        `json:"display,omitempty"`
    Extio                map[string]ConstDeviceApExtios `json:"extio,omitempty"`
    FccDfsOk             *bool                          `json:"fcc_dfs_ok,omitempty"`
    Has11ax              *bool                          `json:"has_11ax,omitempty"`
    HasCompass           *bool                          `json:"has_compass,omitempty"`
    HasExtAnt            *bool                          `json:"has_ext_ant,omitempty"`
    HasExtio             *bool                          `json:"has_extio,omitempty"`
    HasHeight            *bool                          `json:"has_height,omitempty"`
    HasModulePort        *bool                          `json:"has_module_port,omitempty"`
    HasPoeOut            *bool                          `json:"has_poe_out,omitempty"`
    HasScanningRadio     *bool                          `json:"has_scanning_radio,omitempty"`
    HasSelectableRadio   *bool                          `json:"has_selectable_radio,omitempty"`
    HasUsb               *bool                          `json:"has_usb,omitempty"`
    HasVble              *bool                          `json:"has_vble,omitempty"`
    HasWifiBand24        *bool                          `json:"has_wifi_band24,omitempty"`
    HasWifiBand5         *bool                          `json:"has_wifi_band5,omitempty"`
    HasWifiBand6         *bool                          `json:"has_wifi_band6,omitempty"`
    MaxPoeOut            *int                           `json:"max_poe_out,omitempty"`
    MaxWlans             *int                           `json:"max_wlans,omitempty"`
    Model                *string                        `json:"model,omitempty"`
    OtherDfsOk           *bool                          `json:"other_dfs_ok,omitempty"`
    Outdoor              *bool                          `json:"outdoor,omitempty"`
    Radios               map[string]string              `json:"radios,omitempty"`
    SharedScanningRadio  *bool                          `json:"shared_scanning_radio,omitempty"`
    Type                 *string                        `json:"type,omitempty"`
    Unmanaged            *bool                          `json:"unmanaged,omitempty"`
    Vble                 *ConstDeviceApVble             `json:"vble,omitempty"`
    Alias                *string                        `json:"alias,omitempty"`
    Defaults             *ConstDeviceSwitchDefault1     `json:"defaults,omitempty"`
    EvolvedOs            *bool                          `json:"evolved_os,omitempty"`
    EvpnRiType           *string                        `json:"evpn_ri_type,omitempty"`
    Experimental         *bool                          `json:"experimental,omitempty"`
    FansPluggable        *bool                          `json:"fans_pluggable,omitempty"`
    HasBgp               *bool                          `json:"has_bgp,omitempty"`
    HasEts               *bool                          `json:"has_ets,omitempty"`
    HasEvpn              *bool                          `json:"has_evpn,omitempty"`
    HasIrb               *bool                          `json:"has_irb,omitempty"`
    HasSnapshot          *bool                          `json:"has_snapshot,omitempty"`
    HasVc                *bool                          `json:"has_vc,omitempty"`
    Modular              *bool                          `json:"modular,omitempty"`
    NoShapingRate        *bool                          `json:"no_shaping_rate,omitempty"`
    NumberFans           *int                           `json:"number_fans,omitempty"`
    OcDevice             *bool                          `json:"oc_device,omitempty"`
    OobInterface         *string                        `json:"oob_interface,omitempty"`
    PacketActionDropOnly *bool                          `json:"packet_action_drop_only,omitempty"`
    Pic                  map[string]string              `json:"pic,omitempty"`
    SubRequired          *string                        `json:"sub_required,omitempty"`
    HaNode0Fpc           *int                           `json:"ha_node0_fpc,omitempty"`
    HaNode1Fpc           *int                           `json:"ha_node1_fpc,omitempty"`
    HasFxp0              *bool                          `json:"has_fxp0,omitempty"`
    HasHaControl         *bool                          `json:"has_ha_control,omitempty"`
    HasHaData            *bool                          `json:"has_ha_data,omitempty"`
    IrbDisabledByDefault *bool                          `json:"irb_disabled_by_default,omitempty"`
    Ports                *ConstDeviceGatewayPorts       `json:"ports,omitempty"`
    T128Device           *bool                          `json:"t128_device,omitempty"`
}
