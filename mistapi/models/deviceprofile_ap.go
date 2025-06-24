package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// DeviceprofileAp represents a DeviceprofileAp struct.
// Device Profile
type DeviceprofileAp struct {
    // Aeroscout AP settings
    Aeroscout            *ApAeroscout            `json:"aeroscout,omitempty"`
    // BLE AP settings
    BleConfig            *BleConfig              `json:"ble_config,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64                `json:"created_time,omitempty"`
    // Whether to disable eth1 port
    DisableEth1          *bool                   `json:"disable_eth1,omitempty"`
    // Whether to disable eth2 port
    DisableEth2          *bool                   `json:"disable_eth2,omitempty"`
    // Whether to disable eth3 port
    DisableEth3          *bool                   `json:"disable_eth3,omitempty"`
    // Whether to disable module port
    DisableModule        *bool                   `json:"disable_module,omitempty"`
    EslConfig            *ApEslConfig            `json:"esl_config,omitempty"`
    ForSite              *bool                   `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID              `json:"id,omitempty"`
    // IoT AP settings
    IotConfig            *ApIot                  `json:"iot_config,omitempty"`
    // IP AP settings
    IpConfig             *ApIpConfig             `json:"ip_config,omitempty"`
    LacpConfig           *DeviceApLacpConfig     `json:"lacp_config,omitempty"`
    // LED AP settings
    Led                  *ApLed                  `json:"led,omitempty"`
    // Mesh AP settings
    Mesh                 *ApMesh                 `json:"mesh,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                `json:"modified_time,omitempty"`
    Name                 *string                 `json:"name"`
    NtpServers           []string                `json:"ntp_servers,omitempty"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // Whether to enable power out through module port (for APH) or eth1 (for APL/BT11)
    PoePassthrough       *bool                   `json:"poe_passthrough,omitempty"`
    // eth0 is not allowed here. Property key is the interface(s) name (e.g. `eth1` or `eth1,eth2`). If spcified, this takes predecence over switch_config (deprecated)
    PortConfig           map[string]ApPortConfig `json:"port_config,omitempty"`
    // Power related configs
    PwrConfig            *ApPwrConfig            `json:"pwr_config,omitempty"`
    // Radio AP settings
    RadioConfig          *ApRadio                `json:"radio_config,omitempty"`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    // For people who want to fully control the vlans (advanced)
    SwitchConfig         *ApSwitch               `json:"switch_config,omitempty"`      // Deprecated
    // Device Type. enum: `ap`
    Type                 string                  `json:"type"`
    // AP Uplink port configuration
    UplinkPortConfig     *ApUplinkPortConfig     `json:"uplink_port_config,omitempty"`
    // USB AP settings
    // - Note: if native imagotag is enabled, BLE will be disabled automatically
    // - Note: legacy, new config moved to ESL Config.
    UsbConfig            *ApUsb                  `json:"usb_config,omitempty"`
    // Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                 map[string]string       `json:"vars,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceprofileAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceprofileAp) String() string {
    return fmt.Sprintf(
    	"DeviceprofileAp[Aeroscout=%v, BleConfig=%v, CreatedTime=%v, DisableEth1=%v, DisableEth2=%v, DisableEth3=%v, DisableModule=%v, EslConfig=%v, ForSite=%v, Id=%v, IotConfig=%v, IpConfig=%v, LacpConfig=%v, Led=%v, Mesh=%v, ModifiedTime=%v, Name=%v, NtpServers=%v, OrgId=%v, PoePassthrough=%v, PortConfig=%v, PwrConfig=%v, RadioConfig=%v, SiteId=%v, SwitchConfig=%v, Type=%v, UplinkPortConfig=%v, UsbConfig=%v, Vars=%v, AdditionalProperties=%v]",
    	d.Aeroscout, d.BleConfig, d.CreatedTime, d.DisableEth1, d.DisableEth2, d.DisableEth3, d.DisableModule, d.EslConfig, d.ForSite, d.Id, d.IotConfig, d.IpConfig, d.LacpConfig, d.Led, d.Mesh, d.ModifiedTime, d.Name, d.NtpServers, d.OrgId, d.PoePassthrough, d.PortConfig, d.PwrConfig, d.RadioConfig, d.SiteId, d.SwitchConfig, d.Type, d.UplinkPortConfig, d.UsbConfig, d.Vars, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceprofileAp.
// It customizes the JSON marshaling process for DeviceprofileAp objects.
func (d DeviceprofileAp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "aeroscout", "ble_config", "created_time", "disable_eth1", "disable_eth2", "disable_eth3", "disable_module", "esl_config", "for_site", "id", "iot_config", "ip_config", "lacp_config", "led", "mesh", "modified_time", "name", "ntp_servers", "org_id", "poe_passthrough", "port_config", "pwr_config", "radio_config", "site_id", "switch_config", "type", "uplink_port_config", "usb_config", "vars"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceprofileAp object to a map representation for JSON marshaling.
func (d DeviceprofileAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Aeroscout != nil {
        structMap["aeroscout"] = d.Aeroscout.toMap()
    }
    if d.BleConfig != nil {
        structMap["ble_config"] = d.BleConfig.toMap()
    }
    if d.CreatedTime != nil {
        structMap["created_time"] = d.CreatedTime
    }
    if d.DisableEth1 != nil {
        structMap["disable_eth1"] = d.DisableEth1
    }
    if d.DisableEth2 != nil {
        structMap["disable_eth2"] = d.DisableEth2
    }
    if d.DisableEth3 != nil {
        structMap["disable_eth3"] = d.DisableEth3
    }
    if d.DisableModule != nil {
        structMap["disable_module"] = d.DisableModule
    }
    if d.EslConfig != nil {
        structMap["esl_config"] = d.EslConfig.toMap()
    }
    if d.ForSite != nil {
        structMap["for_site"] = d.ForSite
    }
    if d.Id != nil {
        structMap["id"] = d.Id
    }
    if d.IotConfig != nil {
        structMap["iot_config"] = d.IotConfig.toMap()
    }
    if d.IpConfig != nil {
        structMap["ip_config"] = d.IpConfig.toMap()
    }
    if d.LacpConfig != nil {
        structMap["lacp_config"] = d.LacpConfig.toMap()
    }
    if d.Led != nil {
        structMap["led"] = d.Led.toMap()
    }
    if d.Mesh != nil {
        structMap["mesh"] = d.Mesh.toMap()
    }
    if d.ModifiedTime != nil {
        structMap["modified_time"] = d.ModifiedTime
    }
    if d.Name != nil {
        structMap["name"] = d.Name
    } else {
        structMap["name"] = nil
    }
    if d.NtpServers != nil {
        structMap["ntp_servers"] = d.NtpServers
    }
    if d.OrgId != nil {
        structMap["org_id"] = d.OrgId
    }
    if d.PoePassthrough != nil {
        structMap["poe_passthrough"] = d.PoePassthrough
    }
    if d.PortConfig != nil {
        structMap["port_config"] = d.PortConfig
    }
    if d.PwrConfig != nil {
        structMap["pwr_config"] = d.PwrConfig.toMap()
    }
    if d.RadioConfig != nil {
        structMap["radio_config"] = d.RadioConfig.toMap()
    }
    if d.SiteId != nil {
        structMap["site_id"] = d.SiteId
    }
    if d.SwitchConfig != nil {
        structMap["switch_config"] = d.SwitchConfig.toMap()
    }
    structMap["type"] = d.Type
    if d.UplinkPortConfig != nil {
        structMap["uplink_port_config"] = d.UplinkPortConfig.toMap()
    }
    if d.UsbConfig != nil {
        structMap["usb_config"] = d.UsbConfig.toMap()
    }
    if d.Vars != nil {
        structMap["vars"] = d.Vars
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceprofileAp.
// It customizes the JSON unmarshaling process for DeviceprofileAp objects.
func (d *DeviceprofileAp) UnmarshalJSON(input []byte) error {
    var temp tempDeviceprofileAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aeroscout", "ble_config", "created_time", "disable_eth1", "disable_eth2", "disable_eth3", "disable_module", "esl_config", "for_site", "id", "iot_config", "ip_config", "lacp_config", "led", "mesh", "modified_time", "name", "ntp_servers", "org_id", "poe_passthrough", "port_config", "pwr_config", "radio_config", "site_id", "switch_config", "type", "uplink_port_config", "usb_config", "vars")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Aeroscout = temp.Aeroscout
    d.BleConfig = temp.BleConfig
    d.CreatedTime = temp.CreatedTime
    d.DisableEth1 = temp.DisableEth1
    d.DisableEth2 = temp.DisableEth2
    d.DisableEth3 = temp.DisableEth3
    d.DisableModule = temp.DisableModule
    d.EslConfig = temp.EslConfig
    d.ForSite = temp.ForSite
    d.Id = temp.Id
    d.IotConfig = temp.IotConfig
    d.IpConfig = temp.IpConfig
    d.LacpConfig = temp.LacpConfig
    d.Led = temp.Led
    d.Mesh = temp.Mesh
    d.ModifiedTime = temp.ModifiedTime
    d.Name = temp.Name
    d.NtpServers = temp.NtpServers
    d.OrgId = temp.OrgId
    d.PoePassthrough = temp.PoePassthrough
    d.PortConfig = temp.PortConfig
    d.PwrConfig = temp.PwrConfig
    d.RadioConfig = temp.RadioConfig
    d.SiteId = temp.SiteId
    d.SwitchConfig = temp.SwitchConfig
    d.Type = *temp.Type
    d.UplinkPortConfig = temp.UplinkPortConfig
    d.UsbConfig = temp.UsbConfig
    d.Vars = temp.Vars
    return nil
}

// tempDeviceprofileAp is a temporary struct used for validating the fields of DeviceprofileAp.
type tempDeviceprofileAp  struct {
    Aeroscout        *ApAeroscout            `json:"aeroscout,omitempty"`
    BleConfig        *BleConfig              `json:"ble_config,omitempty"`
    CreatedTime      *float64                `json:"created_time,omitempty"`
    DisableEth1      *bool                   `json:"disable_eth1,omitempty"`
    DisableEth2      *bool                   `json:"disable_eth2,omitempty"`
    DisableEth3      *bool                   `json:"disable_eth3,omitempty"`
    DisableModule    *bool                   `json:"disable_module,omitempty"`
    EslConfig        *ApEslConfig            `json:"esl_config,omitempty"`
    ForSite          *bool                   `json:"for_site,omitempty"`
    Id               *uuid.UUID              `json:"id,omitempty"`
    IotConfig        *ApIot                  `json:"iot_config,omitempty"`
    IpConfig         *ApIpConfig             `json:"ip_config,omitempty"`
    LacpConfig       *DeviceApLacpConfig     `json:"lacp_config,omitempty"`
    Led              *ApLed                  `json:"led,omitempty"`
    Mesh             *ApMesh                 `json:"mesh,omitempty"`
    ModifiedTime     *float64                `json:"modified_time,omitempty"`
    Name             *string                 `json:"name"`
    NtpServers       []string                `json:"ntp_servers,omitempty"`
    OrgId            *uuid.UUID              `json:"org_id,omitempty"`
    PoePassthrough   *bool                   `json:"poe_passthrough,omitempty"`
    PortConfig       map[string]ApPortConfig `json:"port_config,omitempty"`
    PwrConfig        *ApPwrConfig            `json:"pwr_config,omitempty"`
    RadioConfig      *ApRadio                `json:"radio_config,omitempty"`
    SiteId           *uuid.UUID              `json:"site_id,omitempty"`
    SwitchConfig     *ApSwitch               `json:"switch_config,omitempty"`
    Type             *string                 `json:"type"`
    UplinkPortConfig *ApUplinkPortConfig     `json:"uplink_port_config,omitempty"`
    UsbConfig        *ApUsb                  `json:"usb_config,omitempty"`
    Vars             map[string]string       `json:"vars,omitempty"`
}

func (d *tempDeviceprofileAp) validate() error {
    var errs []string
    if d.Type == nil {
        errs = append(errs, "required field `type` is missing for type `deviceprofile_ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
