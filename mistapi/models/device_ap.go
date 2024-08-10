package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// DeviceAp represents a DeviceAp struct.
// AP
type DeviceAp struct {
    // Aeroscout AP settings
    Aeroscout            *ApAeroscout            `json:"aeroscout,omitempty"`
    // BLE AP settings
    BleConfig            *BleConfig              `json:"ble_config,omitempty"`
    Centrak              *ApCentrak              `json:"centrak,omitempty"`
    ClientBridge         *ApClientBridge         `json:"client_bridge,omitempty"`
    CreatedTime          *float64                `json:"created_time,omitempty"`
    DeviceprofileId      Optional[uuid.UUID]     `json:"deviceprofile_id"`
    // whether to disable eth1 port
    DisableEth1          *bool                   `json:"disable_eth1,omitempty"`
    // whether to disable eth2 port
    DisableEth2          *bool                   `json:"disable_eth2,omitempty"`
    // whether to disable eth3 port
    DisableEth3          *bool                   `json:"disable_eth3,omitempty"`
    // whether to disable module port
    DisableModule        *bool                   `json:"disable_module,omitempty"`
    EslConfig            *ApEslConfig            `json:"esl_config,omitempty"`
    ForSite              *bool                   `json:"for_site,omitempty"`
    // height, in meters, optional
    Height               *float64                `json:"height,omitempty"`
    Id                   *uuid.UUID              `json:"id,omitempty"`
    Image1Url            Optional[string]        `json:"image1_url"`
    Image2Url            Optional[string]        `json:"image2_url"`
    Image3Url            Optional[string]        `json:"image3_url"`
    // IoT AP settings
    IotConfig            *ApIot                  `json:"iot_config,omitempty"`
    // IP AP settings
    IpConfig             *ApIpConfig             `json:"ip_config,omitempty"`
    // LED AP settings
    Led                  *ApLed                  `json:"led,omitempty"`
    // whether this map is considered locked down
    Locked               *bool                   `json:"locked,omitempty"`
    // device MAC address
    Mac                  *string                 `json:"mac,omitempty"`
    // map where the device belongs to
    MapId                *uuid.UUID              `json:"map_id,omitempty"`
    // Mesh AP settings
    Mesh                 *ApMesh                 `json:"mesh,omitempty"`
    // device Model
    Model                *string                 `json:"model,omitempty"`
    ModifiedTime         *float64                `json:"modified_time,omitempty"`
    Name                 *string                 `json:"name,omitempty"`
    // any notes about this AP
    Notes                *string                 `json:"notes,omitempty"`
    NtpServers           []string                `json:"ntp_servers,omitempty"`
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    // orientation, 0-359, in degrees, up is 0, right is 90.
    Orientation          *int                    `json:"orientation,omitempty"`
    // whether to enable power out through module port (for APH) or eth1 (for APL/BT11)
    PoePassthrough       *bool                   `json:"poe_passthrough,omitempty"`
    // eth0 is not allowed here. Property key is the interface(s) name (e.g. `eth1` or `eth1,eth2`)
    PortConfig           map[string]ApPortConfig `json:"port_config,omitempty"`
    // power related configs
    PwrConfig            *ApPwrConfig            `json:"pwr_config,omitempty"`
    // Radio AP settings
    RadioConfig          *ApRadio                `json:"radio_config,omitempty"`
    // device Serial
    Serial               *string                 `json:"serial,omitempty"`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    // Device Type. enum: `ap`
    Type                 *DeviceTypeApEnum       `json:"type,omitempty"`
    UplinkPortConfig     *ApUplinkPortConfig     `json:"uplink_port_config,omitempty"`
    // USB AP settings
    // Note: if native imagotag is enabled, BLE will be disabled automatically
    // Note: legacy, new config moved to ESL Config.
    UsbConfig            *ApUsb                  `json:"usb_config,omitempty"`
    // a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                 map[string]string       `json:"vars,omitempty"`
    // x in pixel
    X                    *float64                `json:"x,omitempty"`
    // y in pixel
    Y                    *float64                `json:"y,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceAp.
// It customizes the JSON marshaling process for DeviceAp objects.
func (d DeviceAp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceAp object to a map representation for JSON marshaling.
func (d DeviceAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Aeroscout != nil {
        structMap["aeroscout"] = d.Aeroscout.toMap()
    }
    if d.BleConfig != nil {
        structMap["ble_config"] = d.BleConfig.toMap()
    }
    if d.Centrak != nil {
        structMap["centrak"] = d.Centrak.toMap()
    }
    if d.ClientBridge != nil {
        structMap["client_bridge"] = d.ClientBridge.toMap()
    }
    if d.CreatedTime != nil {
        structMap["created_time"] = d.CreatedTime
    }
    if d.DeviceprofileId.IsValueSet() {
        if d.DeviceprofileId.Value() != nil {
            structMap["deviceprofile_id"] = d.DeviceprofileId.Value()
        } else {
            structMap["deviceprofile_id"] = nil
        }
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
    if d.Height != nil {
        structMap["height"] = d.Height
    }
    if d.Id != nil {
        structMap["id"] = d.Id
    }
    if d.Image1Url.IsValueSet() {
        if d.Image1Url.Value() != nil {
            structMap["image1_url"] = d.Image1Url.Value()
        } else {
            structMap["image1_url"] = nil
        }
    }
    if d.Image2Url.IsValueSet() {
        if d.Image2Url.Value() != nil {
            structMap["image2_url"] = d.Image2Url.Value()
        } else {
            structMap["image2_url"] = nil
        }
    }
    if d.Image3Url.IsValueSet() {
        if d.Image3Url.Value() != nil {
            structMap["image3_url"] = d.Image3Url.Value()
        } else {
            structMap["image3_url"] = nil
        }
    }
    if d.IotConfig != nil {
        structMap["iot_config"] = d.IotConfig.toMap()
    }
    if d.IpConfig != nil {
        structMap["ip_config"] = d.IpConfig.toMap()
    }
    if d.Led != nil {
        structMap["led"] = d.Led.toMap()
    }
    if d.Locked != nil {
        structMap["locked"] = d.Locked
    }
    if d.Mac != nil {
        structMap["mac"] = d.Mac
    }
    if d.MapId != nil {
        structMap["map_id"] = d.MapId
    }
    if d.Mesh != nil {
        structMap["mesh"] = d.Mesh.toMap()
    }
    if d.Model != nil {
        structMap["model"] = d.Model
    }
    if d.ModifiedTime != nil {
        structMap["modified_time"] = d.ModifiedTime
    }
    if d.Name != nil {
        structMap["name"] = d.Name
    }
    if d.Notes != nil {
        structMap["notes"] = d.Notes
    }
    if d.NtpServers != nil {
        structMap["ntp_servers"] = d.NtpServers
    }
    if d.OrgId != nil {
        structMap["org_id"] = d.OrgId
    }
    if d.Orientation != nil {
        structMap["orientation"] = d.Orientation
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
    if d.Serial != nil {
        structMap["serial"] = d.Serial
    }
    if d.SiteId != nil {
        structMap["site_id"] = d.SiteId
    }
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.UplinkPortConfig != nil {
        structMap["uplink_port_config"] = d.UplinkPortConfig.toMap()
    }
    if d.UsbConfig != nil {
        structMap["usb_config"] = d.UsbConfig.toMap()
    }
    if d.Vars != nil {
        structMap["vars"] = d.Vars
    }
    if d.X != nil {
        structMap["x"] = d.X
    }
    if d.Y != nil {
        structMap["y"] = d.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceAp.
// It customizes the JSON unmarshaling process for DeviceAp objects.
func (d *DeviceAp) UnmarshalJSON(input []byte) error {
    var temp tempDeviceAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "aeroscout", "ble_config", "centrak", "client_bridge", "created_time", "deviceprofile_id", "disable_eth1", "disable_eth2", "disable_eth3", "disable_module", "esl_config", "for_site", "height", "id", "image1_url", "image2_url", "image3_url", "iot_config", "ip_config", "led", "locked", "mac", "map_id", "mesh", "model", "modified_time", "name", "notes", "ntp_servers", "org_id", "orientation", "poe_passthrough", "port_config", "pwr_config", "radio_config", "serial", "site_id", "type", "uplink_port_config", "usb_config", "vars", "x", "y")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Aeroscout = temp.Aeroscout
    d.BleConfig = temp.BleConfig
    d.Centrak = temp.Centrak
    d.ClientBridge = temp.ClientBridge
    d.CreatedTime = temp.CreatedTime
    d.DeviceprofileId = temp.DeviceprofileId
    d.DisableEth1 = temp.DisableEth1
    d.DisableEth2 = temp.DisableEth2
    d.DisableEth3 = temp.DisableEth3
    d.DisableModule = temp.DisableModule
    d.EslConfig = temp.EslConfig
    d.ForSite = temp.ForSite
    d.Height = temp.Height
    d.Id = temp.Id
    d.Image1Url = temp.Image1Url
    d.Image2Url = temp.Image2Url
    d.Image3Url = temp.Image3Url
    d.IotConfig = temp.IotConfig
    d.IpConfig = temp.IpConfig
    d.Led = temp.Led
    d.Locked = temp.Locked
    d.Mac = temp.Mac
    d.MapId = temp.MapId
    d.Mesh = temp.Mesh
    d.Model = temp.Model
    d.ModifiedTime = temp.ModifiedTime
    d.Name = temp.Name
    d.Notes = temp.Notes
    d.NtpServers = temp.NtpServers
    d.OrgId = temp.OrgId
    d.Orientation = temp.Orientation
    d.PoePassthrough = temp.PoePassthrough
    d.PortConfig = temp.PortConfig
    d.PwrConfig = temp.PwrConfig
    d.RadioConfig = temp.RadioConfig
    d.Serial = temp.Serial
    d.SiteId = temp.SiteId
    d.Type = temp.Type
    d.UplinkPortConfig = temp.UplinkPortConfig
    d.UsbConfig = temp.UsbConfig
    d.Vars = temp.Vars
    d.X = temp.X
    d.Y = temp.Y
    return nil
}

// tempDeviceAp is a temporary struct used for validating the fields of DeviceAp.
type tempDeviceAp  struct {
    Aeroscout        *ApAeroscout            `json:"aeroscout,omitempty"`
    BleConfig        *BleConfig              `json:"ble_config,omitempty"`
    Centrak          *ApCentrak              `json:"centrak,omitempty"`
    ClientBridge     *ApClientBridge         `json:"client_bridge,omitempty"`
    CreatedTime      *float64                `json:"created_time,omitempty"`
    DeviceprofileId  Optional[uuid.UUID]     `json:"deviceprofile_id"`
    DisableEth1      *bool                   `json:"disable_eth1,omitempty"`
    DisableEth2      *bool                   `json:"disable_eth2,omitempty"`
    DisableEth3      *bool                   `json:"disable_eth3,omitempty"`
    DisableModule    *bool                   `json:"disable_module,omitempty"`
    EslConfig        *ApEslConfig            `json:"esl_config,omitempty"`
    ForSite          *bool                   `json:"for_site,omitempty"`
    Height           *float64                `json:"height,omitempty"`
    Id               *uuid.UUID              `json:"id,omitempty"`
    Image1Url        Optional[string]        `json:"image1_url"`
    Image2Url        Optional[string]        `json:"image2_url"`
    Image3Url        Optional[string]        `json:"image3_url"`
    IotConfig        *ApIot                  `json:"iot_config,omitempty"`
    IpConfig         *ApIpConfig             `json:"ip_config,omitempty"`
    Led              *ApLed                  `json:"led,omitempty"`
    Locked           *bool                   `json:"locked,omitempty"`
    Mac              *string                 `json:"mac,omitempty"`
    MapId            *uuid.UUID              `json:"map_id,omitempty"`
    Mesh             *ApMesh                 `json:"mesh,omitempty"`
    Model            *string                 `json:"model,omitempty"`
    ModifiedTime     *float64                `json:"modified_time,omitempty"`
    Name             *string                 `json:"name,omitempty"`
    Notes            *string                 `json:"notes,omitempty"`
    NtpServers       []string                `json:"ntp_servers,omitempty"`
    OrgId            *uuid.UUID              `json:"org_id,omitempty"`
    Orientation      *int                    `json:"orientation,omitempty"`
    PoePassthrough   *bool                   `json:"poe_passthrough,omitempty"`
    PortConfig       map[string]ApPortConfig `json:"port_config,omitempty"`
    PwrConfig        *ApPwrConfig            `json:"pwr_config,omitempty"`
    RadioConfig      *ApRadio                `json:"radio_config,omitempty"`
    Serial           *string                 `json:"serial,omitempty"`
    SiteId           *uuid.UUID              `json:"site_id,omitempty"`
    Type             *DeviceTypeApEnum       `json:"type,omitempty"`
    UplinkPortConfig *ApUplinkPortConfig     `json:"uplink_port_config,omitempty"`
    UsbConfig        *ApUsb                  `json:"usb_config,omitempty"`
    Vars             map[string]string       `json:"vars,omitempty"`
    X                *float64                `json:"x,omitempty"`
    Y                *float64                `json:"y,omitempty"`
}
