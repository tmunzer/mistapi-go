package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ConstDeviceAp represents a ConstDeviceAp struct.
type ConstDeviceAp struct {
    ApType               string                         `json:"ap_type"`
    Band24               *ConstDeviceApBand24           `json:"band24,omitempty"`
    Band5                *ConstDeviceApBand5            `json:"band5,omitempty"`
    Band6                *ConstDeviceApBand5            `json:"band6,omitempty"`
    Band24Usages         []ConstDeviceApBand24UsageEnum `json:"band_24_usages,omitempty"`
    CeDfsOk              *bool                          `json:"ce_dfs_ok,omitempty"`
    CiscoPace            *bool                          `json:"cisco_pace,omitempty"`
    Description          *string                        `json:"description,omitempty"`
    // Property key is a list of country codes (e.g. "GB, DE")
    DisallowedChannels   map[string][]int               `json:"disallowed_channels,omitempty"`
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
    // Device Type. enum: `ap`
    Type                 string                         `json:"type"`
    Unmanaged            *bool                          `json:"unmanaged,omitempty"`
    Vble                 *ConstDeviceApVble             `json:"vble,omitempty"`
    AdditionalProperties map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for ConstDeviceAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstDeviceAp) String() string {
    return fmt.Sprintf(
    	"ConstDeviceAp[ApType=%v, Band24=%v, Band5=%v, Band6=%v, Band24Usages=%v, CeDfsOk=%v, CiscoPace=%v, Description=%v, DisallowedChannels=%v, Display=%v, Extio=%v, FccDfsOk=%v, Has11ax=%v, HasCompass=%v, HasExtAnt=%v, HasExtio=%v, HasHeight=%v, HasModulePort=%v, HasPoeOut=%v, HasScanningRadio=%v, HasSelectableRadio=%v, HasUsb=%v, HasVble=%v, HasWifiBand24=%v, HasWifiBand5=%v, HasWifiBand6=%v, MaxPoeOut=%v, MaxWlans=%v, Model=%v, OtherDfsOk=%v, Outdoor=%v, Radios=%v, SharedScanningRadio=%v, Type=%v, Unmanaged=%v, Vble=%v, AdditionalProperties=%v]",
    	c.ApType, c.Band24, c.Band5, c.Band6, c.Band24Usages, c.CeDfsOk, c.CiscoPace, c.Description, c.DisallowedChannels, c.Display, c.Extio, c.FccDfsOk, c.Has11ax, c.HasCompass, c.HasExtAnt, c.HasExtio, c.HasHeight, c.HasModulePort, c.HasPoeOut, c.HasScanningRadio, c.HasSelectableRadio, c.HasUsb, c.HasVble, c.HasWifiBand24, c.HasWifiBand5, c.HasWifiBand6, c.MaxPoeOut, c.MaxWlans, c.Model, c.OtherDfsOk, c.Outdoor, c.Radios, c.SharedScanningRadio, c.Type, c.Unmanaged, c.Vble, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceAp.
// It customizes the JSON marshaling process for ConstDeviceAp objects.
func (c ConstDeviceAp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap_type", "band24", "band5", "band6", "band_24_usages", "ce_dfs_ok", "cisco_pace", "description", "disallowed_channels", "display", "extio", "fcc_dfs_ok", "has_11ax", "has_compass", "has_ext_ant", "has_extio", "has_height", "has_module_port", "has_poe_out", "has_scanning_radio", "has_selectable_radio", "has_usb", "has_vble", "has_wifi_band24", "has_wifi_band5", "has_wifi_band6", "max_poe_out", "max_wlans", "model", "other_dfs_ok", "outdoor", "radios", "shared_scanning_radio", "type", "unmanaged", "vble"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceAp object to a map representation for JSON marshaling.
func (c ConstDeviceAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["ap_type"] = c.ApType
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
    structMap["type"] = c.Type
    if c.Unmanaged != nil {
        structMap["unmanaged"] = c.Unmanaged
    }
    if c.Vble != nil {
        structMap["vble"] = c.Vble.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceAp.
// It customizes the JSON unmarshaling process for ConstDeviceAp objects.
func (c *ConstDeviceAp) UnmarshalJSON(input []byte) error {
    var temp tempConstDeviceAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_type", "band24", "band5", "band6", "band_24_usages", "ce_dfs_ok", "cisco_pace", "description", "disallowed_channels", "display", "extio", "fcc_dfs_ok", "has_11ax", "has_compass", "has_ext_ant", "has_extio", "has_height", "has_module_port", "has_poe_out", "has_scanning_radio", "has_selectable_radio", "has_usb", "has_vble", "has_wifi_band24", "has_wifi_band5", "has_wifi_band6", "max_poe_out", "max_wlans", "model", "other_dfs_ok", "outdoor", "radios", "shared_scanning_radio", "type", "unmanaged", "vble")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ApType = *temp.ApType
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
    c.Type = *temp.Type
    c.Unmanaged = temp.Unmanaged
    c.Vble = temp.Vble
    return nil
}

// tempConstDeviceAp is a temporary struct used for validating the fields of ConstDeviceAp.
type tempConstDeviceAp  struct {
    ApType              *string                        `json:"ap_type"`
    Band24              *ConstDeviceApBand24           `json:"band24,omitempty"`
    Band5               *ConstDeviceApBand5            `json:"band5,omitempty"`
    Band6               *ConstDeviceApBand5            `json:"band6,omitempty"`
    Band24Usages        []ConstDeviceApBand24UsageEnum `json:"band_24_usages,omitempty"`
    CeDfsOk             *bool                          `json:"ce_dfs_ok,omitempty"`
    CiscoPace           *bool                          `json:"cisco_pace,omitempty"`
    Description         *string                        `json:"description,omitempty"`
    DisallowedChannels  map[string][]int               `json:"disallowed_channels,omitempty"`
    Display             *string                        `json:"display,omitempty"`
    Extio               map[string]ConstDeviceApExtios `json:"extio,omitempty"`
    FccDfsOk            *bool                          `json:"fcc_dfs_ok,omitempty"`
    Has11ax             *bool                          `json:"has_11ax,omitempty"`
    HasCompass          *bool                          `json:"has_compass,omitempty"`
    HasExtAnt           *bool                          `json:"has_ext_ant,omitempty"`
    HasExtio            *bool                          `json:"has_extio,omitempty"`
    HasHeight           *bool                          `json:"has_height,omitempty"`
    HasModulePort       *bool                          `json:"has_module_port,omitempty"`
    HasPoeOut           *bool                          `json:"has_poe_out,omitempty"`
    HasScanningRadio    *bool                          `json:"has_scanning_radio,omitempty"`
    HasSelectableRadio  *bool                          `json:"has_selectable_radio,omitempty"`
    HasUsb              *bool                          `json:"has_usb,omitempty"`
    HasVble             *bool                          `json:"has_vble,omitempty"`
    HasWifiBand24       *bool                          `json:"has_wifi_band24,omitempty"`
    HasWifiBand5        *bool                          `json:"has_wifi_band5,omitempty"`
    HasWifiBand6        *bool                          `json:"has_wifi_band6,omitempty"`
    MaxPoeOut           *int                           `json:"max_poe_out,omitempty"`
    MaxWlans            *int                           `json:"max_wlans,omitempty"`
    Model               *string                        `json:"model,omitempty"`
    OtherDfsOk          *bool                          `json:"other_dfs_ok,omitempty"`
    Outdoor             *bool                          `json:"outdoor,omitempty"`
    Radios              map[string]string              `json:"radios,omitempty"`
    SharedScanningRadio *bool                          `json:"shared_scanning_radio,omitempty"`
    Type                *string                        `json:"type"`
    Unmanaged           *bool                          `json:"unmanaged,omitempty"`
    Vble                *ConstDeviceApVble             `json:"vble,omitempty"`
}

func (c *tempConstDeviceAp) validate() error {
    var errs []string
    if c.ApType == nil {
        errs = append(errs, "required field `ap_type` is missing for type `const_device_ap`")
    }
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `const_device_ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
