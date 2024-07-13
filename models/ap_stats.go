package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ApStats represents a ApStats struct.
// AP statistics
type ApStats struct {
    AutoPlacement        *ApStatsAutoPlacement                         `json:"auto_placement,omitempty"`
    AutoUpgradeStat      *ApStatsAutoUpgrade                           `json:"auto_upgrade_stat,omitempty"`
    BleStat              *ApStatsBle                                   `json:"ble_stat,omitempty"`
    CertExpiry           Optional[float64]                             `json:"cert_expiry"`
    ConfigReverted       Optional[bool]                                `json:"config_reverted"`
    CpuSystem            Optional[int64]                               `json:"cpu_system"`
    CpuUtil              Optional[int]                                 `json:"cpu_util"`
    CreatedTime          Optional[int64]                               `json:"created_time"`
    DeviceprofileId      Optional[uuid.UUID]                           `json:"deviceprofile_id"`
    // device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage
    EnvStat              *ApStatsEnvStat                               `json:"env_stat,omitempty"`
    EslStat              Optional[ApStatsEslStat]                      `json:"esl_stat"`
    EvpntopoId           Optional[uuid.UUID]                           `json:"evpntopo_id"`
    ExtIp                Optional[string]                              `json:"ext_ip"`
    Fwupdate             *FwupdateStat                                 `json:"fwupdate,omitempty"`
    HwRev                Optional[string]                              `json:"hw_rev"`
    Id                   *uuid.UUID                                    `json:"id,omitempty"`
    InactiveWiredVlans   []int                                         `json:"inactive_wired_vlans,omitempty"`
    IotStat              map[string]ApStatsIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    Ip                   Optional[string]                              `json:"ip"`
    // IP AP settings
    IpConfig             *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat               *IpStat                                       `json:"ip_stat,omitempty"`
    // l2tp tunnel status (key is the wxtunnel_id)
    L2tpStat             map[string]ApStatsL2TpStat                    `json:"l2tp_stat,omitempty"`
    // last seen timestamp
    LastSeen             Optional[float64]                             `json:"last_seen"`
    // last trouble code of switch
    LastTrouble          *LastTrouble                                  `json:"last_trouble,omitempty"`
    // LED AP settings
    Led                  *ApLed                                        `json:"led,omitempty"`
    // LLDP Stat (neighbor information, power negotiations)
    LldpStat             *ApStatsLldpStat                              `json:"lldp_stat,omitempty"`
    Locating             Optional[bool]                                `json:"locating"`
    // whether this AP is considered locked (placement / orientation has been vetted)
    Locked               Optional[bool]                                `json:"locked"`
    // device mac
    Mac                  *string                                       `json:"mac"`
    MapId                Optional[uuid.UUID]                           `json:"map_id"`
    MemUsedKb            Optional[int64]                               `json:"mem_used_kb"`
    // Property key is the mesh downlink id (e.g `00000000-0000-0000-1000-5c5b35000010`)
    MeshDownlinks        map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink           *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    // device model
    Model                *string                                       `json:"model"`
    ModifiedTime         Optional[int64]                               `json:"modified_time"`
    Mount                Optional[string]                              `json:"mount"`
    Name                 Optional[string]                              `json:"name"`
    Notes                Optional[string]                              `json:"notes"`
    // how many wireless clients are currently connected
    NumClients           Optional[int]                                 `json:"num_clients"`
    OrgId                Optional[uuid.UUID]                           `json:"org_id"`
    // Property key is the port name (e.g. `eth0`)
    PortStat             Optional[map[string]ApStatsPortStat]          `json:"port_stat"`
    // in mW, surplus if positive or deficit if negative
    PowerBudget          Optional[int]                                 `json:"power_budget"`
    // whether insufficient power
    PowerConstrained     Optional[bool]                                `json:"power_constrained"`
    // constrained mode
    PowerOpmode          Optional[string]                              `json:"power_opmode"`
    // DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown)
    PowerSrc             Optional[string]                              `json:"power_src"`
    RadioConfig          *ApStatsRadioConfig                           `json:"radio_config,omitempty"`
    RadioStat            *ApStatsRadioStat                             `json:"radio_stat,omitempty"`
    RxBps                Optional[float64]                             `json:"rx_bps"`
    RxBytes              Optional[int64]                               `json:"rx_bytes"`
    RxPkts               Optional[int]                                 `json:"rx_pkts"`
    // serial
    Serial               Optional[string]                              `json:"serial"`
    SiteId               *uuid.UUID                                    `json:"site_id,omitempty"`
    Status               Optional[string]                              `json:"status"`
    SwitchRedundancy     *ApStatsSwitchRedundancy                      `json:"switch_redundancy,omitempty"`
    TxBps                Optional[float64]                             `json:"tx_bps"`
    TxBytes              Optional[float64]                             `json:"tx_bytes"`
    TxPkts               Optional[float64]                             `json:"tx_pkts"`
    // device type, ap / ble /switch / gateway
    Type                 Optional[string]                              `json:"type"`
    // how long, in seconds, has the device been up (or rebooted)
    Uptime               Optional[float64]                             `json:"uptime"`
    UsbStat              *ApStatsUsbStat                               `json:"usb_stat,omitempty"`
    Version              Optional[string]                              `json:"version"`
    X                    Optional[float64]                             `json:"x"`
    Y                    Optional[float64]                             `json:"y"`
    AdditionalProperties map[string]any                                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStats.
// It customizes the JSON marshaling process for ApStats objects.
func (a ApStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStats object to a map representation for JSON marshaling.
func (a ApStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AutoPlacement != nil {
        structMap["auto_placement"] = a.AutoPlacement.toMap()
    }
    if a.AutoUpgradeStat != nil {
        structMap["auto_upgrade_stat"] = a.AutoUpgradeStat.toMap()
    }
    if a.BleStat != nil {
        structMap["ble_stat"] = a.BleStat.toMap()
    }
    if a.CertExpiry.IsValueSet() {
        if a.CertExpiry.Value() != nil {
            structMap["cert_expiry"] = a.CertExpiry.Value()
        } else {
            structMap["cert_expiry"] = nil
        }
    }
    if a.ConfigReverted.IsValueSet() {
        if a.ConfigReverted.Value() != nil {
            structMap["config_reverted"] = a.ConfigReverted.Value()
        } else {
            structMap["config_reverted"] = nil
        }
    }
    if a.CpuSystem.IsValueSet() {
        if a.CpuSystem.Value() != nil {
            structMap["cpu_system"] = a.CpuSystem.Value()
        } else {
            structMap["cpu_system"] = nil
        }
    }
    if a.CpuUtil.IsValueSet() {
        if a.CpuUtil.Value() != nil {
            structMap["cpu_util"] = a.CpuUtil.Value()
        } else {
            structMap["cpu_util"] = nil
        }
    }
    if a.CreatedTime.IsValueSet() {
        if a.CreatedTime.Value() != nil {
            structMap["created_time"] = a.CreatedTime.Value()
        } else {
            structMap["created_time"] = nil
        }
    }
    if a.DeviceprofileId.IsValueSet() {
        if a.DeviceprofileId.Value() != nil {
            structMap["deviceprofile_id"] = a.DeviceprofileId.Value()
        } else {
            structMap["deviceprofile_id"] = nil
        }
    }
    if a.EnvStat != nil {
        structMap["env_stat"] = a.EnvStat.toMap()
    }
    if a.EslStat.IsValueSet() {
        if a.EslStat.Value() != nil {
            structMap["esl_stat"] = a.EslStat.Value().toMap()
        } else {
            structMap["esl_stat"] = nil
        }
    }
    if a.EvpntopoId.IsValueSet() {
        if a.EvpntopoId.Value() != nil {
            structMap["evpntopo_id"] = a.EvpntopoId.Value()
        } else {
            structMap["evpntopo_id"] = nil
        }
    }
    if a.ExtIp.IsValueSet() {
        if a.ExtIp.Value() != nil {
            structMap["ext_ip"] = a.ExtIp.Value()
        } else {
            structMap["ext_ip"] = nil
        }
    }
    if a.Fwupdate != nil {
        structMap["fwupdate"] = a.Fwupdate.toMap()
    }
    if a.HwRev.IsValueSet() {
        if a.HwRev.Value() != nil {
            structMap["hw_rev"] = a.HwRev.Value()
        } else {
            structMap["hw_rev"] = nil
        }
    }
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.InactiveWiredVlans != nil {
        structMap["inactive_wired_vlans"] = a.InactiveWiredVlans
    }
    if a.IotStat != nil {
        structMap["iot_stat"] = a.IotStat
    }
    if a.Ip.IsValueSet() {
        if a.Ip.Value() != nil {
            structMap["ip"] = a.Ip.Value()
        } else {
            structMap["ip"] = nil
        }
    }
    if a.IpConfig != nil {
        structMap["ip_config"] = a.IpConfig.toMap()
    }
    if a.IpStat != nil {
        structMap["ip_stat"] = a.IpStat.toMap()
    }
    if a.L2tpStat != nil {
        structMap["l2tp_stat"] = a.L2tpStat
    }
    if a.LastSeen.IsValueSet() {
        if a.LastSeen.Value() != nil {
            structMap["last_seen"] = a.LastSeen.Value()
        } else {
            structMap["last_seen"] = nil
        }
    }
    if a.LastTrouble != nil {
        structMap["last_trouble"] = a.LastTrouble.toMap()
    }
    if a.Led != nil {
        structMap["led"] = a.Led.toMap()
    }
    if a.LldpStat != nil {
        structMap["lldp_stat"] = a.LldpStat.toMap()
    }
    if a.Locating.IsValueSet() {
        if a.Locating.Value() != nil {
            structMap["locating"] = a.Locating.Value()
        } else {
            structMap["locating"] = nil
        }
    }
    if a.Locked.IsValueSet() {
        if a.Locked.Value() != nil {
            structMap["locked"] = a.Locked.Value()
        } else {
            structMap["locked"] = nil
        }
    }
    if a.Mac != nil {
        structMap["mac"] = a.Mac
    } else {
        structMap["mac"] = nil
    }
    if a.MapId.IsValueSet() {
        if a.MapId.Value() != nil {
            structMap["map_id"] = a.MapId.Value()
        } else {
            structMap["map_id"] = nil
        }
    }
    if a.MemUsedKb.IsValueSet() {
        if a.MemUsedKb.Value() != nil {
            structMap["mem_used_kb"] = a.MemUsedKb.Value()
        } else {
            structMap["mem_used_kb"] = nil
        }
    }
    if a.MeshDownlinks != nil {
        structMap["mesh_downlinks"] = a.MeshDownlinks
    }
    if a.MeshUplink != nil {
        structMap["mesh_uplink"] = a.MeshUplink.toMap()
    }
    if a.Model != nil {
        structMap["model"] = a.Model
    } else {
        structMap["model"] = nil
    }
    if a.ModifiedTime.IsValueSet() {
        if a.ModifiedTime.Value() != nil {
            structMap["modified_time"] = a.ModifiedTime.Value()
        } else {
            structMap["modified_time"] = nil
        }
    }
    if a.Mount.IsValueSet() {
        if a.Mount.Value() != nil {
            structMap["mount"] = a.Mount.Value()
        } else {
            structMap["mount"] = nil
        }
    }
    if a.Name.IsValueSet() {
        if a.Name.Value() != nil {
            structMap["name"] = a.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if a.Notes.IsValueSet() {
        if a.Notes.Value() != nil {
            structMap["notes"] = a.Notes.Value()
        } else {
            structMap["notes"] = nil
        }
    }
    if a.NumClients.IsValueSet() {
        if a.NumClients.Value() != nil {
            structMap["num_clients"] = a.NumClients.Value()
        } else {
            structMap["num_clients"] = nil
        }
    }
    if a.OrgId.IsValueSet() {
        if a.OrgId.Value() != nil {
            structMap["org_id"] = a.OrgId.Value()
        } else {
            structMap["org_id"] = nil
        }
    }
    if a.PortStat.IsValueSet() {
        if a.PortStat.Value() != nil {
            structMap["port_stat"] = a.PortStat.Value()
        } else {
            structMap["port_stat"] = nil
        }
    }
    if a.PowerBudget.IsValueSet() {
        if a.PowerBudget.Value() != nil {
            structMap["power_budget"] = a.PowerBudget.Value()
        } else {
            structMap["power_budget"] = nil
        }
    }
    if a.PowerConstrained.IsValueSet() {
        if a.PowerConstrained.Value() != nil {
            structMap["power_constrained"] = a.PowerConstrained.Value()
        } else {
            structMap["power_constrained"] = nil
        }
    }
    if a.PowerOpmode.IsValueSet() {
        if a.PowerOpmode.Value() != nil {
            structMap["power_opmode"] = a.PowerOpmode.Value()
        } else {
            structMap["power_opmode"] = nil
        }
    }
    if a.PowerSrc.IsValueSet() {
        if a.PowerSrc.Value() != nil {
            structMap["power_src"] = a.PowerSrc.Value()
        } else {
            structMap["power_src"] = nil
        }
    }
    if a.RadioConfig != nil {
        structMap["radio_config"] = a.RadioConfig.toMap()
    }
    if a.RadioStat != nil {
        structMap["radio_stat"] = a.RadioStat.toMap()
    }
    if a.RxBps.IsValueSet() {
        if a.RxBps.Value() != nil {
            structMap["rx_bps"] = a.RxBps.Value()
        } else {
            structMap["rx_bps"] = nil
        }
    }
    if a.RxBytes.IsValueSet() {
        if a.RxBytes.Value() != nil {
            structMap["rx_bytes"] = a.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if a.RxPkts.IsValueSet() {
        if a.RxPkts.Value() != nil {
            structMap["rx_pkts"] = a.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
        }
    }
    if a.Serial.IsValueSet() {
        if a.Serial.Value() != nil {
            structMap["serial"] = a.Serial.Value()
        } else {
            structMap["serial"] = nil
        }
    }
    if a.SiteId != nil {
        structMap["site_id"] = a.SiteId
    }
    if a.Status.IsValueSet() {
        if a.Status.Value() != nil {
            structMap["status"] = a.Status.Value()
        } else {
            structMap["status"] = nil
        }
    }
    if a.SwitchRedundancy != nil {
        structMap["switch_redundancy"] = a.SwitchRedundancy.toMap()
    }
    if a.TxBps.IsValueSet() {
        if a.TxBps.Value() != nil {
            structMap["tx_bps"] = a.TxBps.Value()
        } else {
            structMap["tx_bps"] = nil
        }
    }
    if a.TxBytes.IsValueSet() {
        if a.TxBytes.Value() != nil {
            structMap["tx_bytes"] = a.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if a.TxPkts.IsValueSet() {
        if a.TxPkts.Value() != nil {
            structMap["tx_pkts"] = a.TxPkts.Value()
        } else {
            structMap["tx_pkts"] = nil
        }
    }
    if a.Type.IsValueSet() {
        if a.Type.Value() != nil {
            structMap["type"] = a.Type.Value()
        } else {
            structMap["type"] = nil
        }
    }
    if a.Uptime.IsValueSet() {
        if a.Uptime.Value() != nil {
            structMap["uptime"] = a.Uptime.Value()
        } else {
            structMap["uptime"] = nil
        }
    }
    if a.UsbStat != nil {
        structMap["usb_stat"] = a.UsbStat.toMap()
    }
    if a.Version.IsValueSet() {
        if a.Version.Value() != nil {
            structMap["version"] = a.Version.Value()
        } else {
            structMap["version"] = nil
        }
    }
    if a.X.IsValueSet() {
        if a.X.Value() != nil {
            structMap["x"] = a.X.Value()
        } else {
            structMap["x"] = nil
        }
    }
    if a.Y.IsValueSet() {
        if a.Y.Value() != nil {
            structMap["y"] = a.Y.Value()
        } else {
            structMap["y"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStats.
// It customizes the JSON unmarshaling process for ApStats objects.
func (a *ApStats) UnmarshalJSON(input []byte) error {
    var temp apStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auto_placement", "auto_upgrade_stat", "ble_stat", "cert_expiry", "config_reverted", "cpu_system", "cpu_util", "created_time", "deviceprofile_id", "env_stat", "esl_stat", "evpntopo_id", "ext_ip", "fwupdate", "hw_rev", "id", "inactive_wired_vlans", "iot_stat", "ip", "ip_config", "ip_stat", "l2tp_stat", "last_seen", "last_trouble", "led", "lldp_stat", "locating", "locked", "mac", "map_id", "mem_used_kb", "mesh_downlinks", "mesh_uplink", "model", "modified_time", "mount", "name", "notes", "num_clients", "org_id", "port_stat", "power_budget", "power_constrained", "power_opmode", "power_src", "radio_config", "radio_stat", "rx_bps", "rx_bytes", "rx_pkts", "serial", "site_id", "status", "switch_redundancy", "tx_bps", "tx_bytes", "tx_pkts", "type", "uptime", "usb_stat", "version", "x", "y")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AutoPlacement = temp.AutoPlacement
    a.AutoUpgradeStat = temp.AutoUpgradeStat
    a.BleStat = temp.BleStat
    a.CertExpiry = temp.CertExpiry
    a.ConfigReverted = temp.ConfigReverted
    a.CpuSystem = temp.CpuSystem
    a.CpuUtil = temp.CpuUtil
    a.CreatedTime = temp.CreatedTime
    a.DeviceprofileId = temp.DeviceprofileId
    a.EnvStat = temp.EnvStat
    a.EslStat = temp.EslStat
    a.EvpntopoId = temp.EvpntopoId
    a.ExtIp = temp.ExtIp
    a.Fwupdate = temp.Fwupdate
    a.HwRev = temp.HwRev
    a.Id = temp.Id
    a.InactiveWiredVlans = temp.InactiveWiredVlans
    a.IotStat = temp.IotStat
    a.Ip = temp.Ip
    a.IpConfig = temp.IpConfig
    a.IpStat = temp.IpStat
    a.L2tpStat = temp.L2tpStat
    a.LastSeen = temp.LastSeen
    a.LastTrouble = temp.LastTrouble
    a.Led = temp.Led
    a.LldpStat = temp.LldpStat
    a.Locating = temp.Locating
    a.Locked = temp.Locked
    a.Mac = temp.Mac
    a.MapId = temp.MapId
    a.MemUsedKb = temp.MemUsedKb
    a.MeshDownlinks = temp.MeshDownlinks
    a.MeshUplink = temp.MeshUplink
    a.Model = temp.Model
    a.ModifiedTime = temp.ModifiedTime
    a.Mount = temp.Mount
    a.Name = temp.Name
    a.Notes = temp.Notes
    a.NumClients = temp.NumClients
    a.OrgId = temp.OrgId
    a.PortStat = temp.PortStat
    a.PowerBudget = temp.PowerBudget
    a.PowerConstrained = temp.PowerConstrained
    a.PowerOpmode = temp.PowerOpmode
    a.PowerSrc = temp.PowerSrc
    a.RadioConfig = temp.RadioConfig
    a.RadioStat = temp.RadioStat
    a.RxBps = temp.RxBps
    a.RxBytes = temp.RxBytes
    a.RxPkts = temp.RxPkts
    a.Serial = temp.Serial
    a.SiteId = temp.SiteId
    a.Status = temp.Status
    a.SwitchRedundancy = temp.SwitchRedundancy
    a.TxBps = temp.TxBps
    a.TxBytes = temp.TxBytes
    a.TxPkts = temp.TxPkts
    a.Type = temp.Type
    a.Uptime = temp.Uptime
    a.UsbStat = temp.UsbStat
    a.Version = temp.Version
    a.X = temp.X
    a.Y = temp.Y
    return nil
}

// apStats is a temporary struct used for validating the fields of ApStats.
type apStats  struct {
    AutoPlacement      *ApStatsAutoPlacement                         `json:"auto_placement,omitempty"`
    AutoUpgradeStat    *ApStatsAutoUpgrade                           `json:"auto_upgrade_stat,omitempty"`
    BleStat            *ApStatsBle                                   `json:"ble_stat,omitempty"`
    CertExpiry         Optional[float64]                             `json:"cert_expiry"`
    ConfigReverted     Optional[bool]                                `json:"config_reverted"`
    CpuSystem          Optional[int64]                               `json:"cpu_system"`
    CpuUtil            Optional[int]                                 `json:"cpu_util"`
    CreatedTime        Optional[int64]                               `json:"created_time"`
    DeviceprofileId    Optional[uuid.UUID]                           `json:"deviceprofile_id"`
    EnvStat            *ApStatsEnvStat                               `json:"env_stat,omitempty"`
    EslStat            Optional[ApStatsEslStat]                      `json:"esl_stat"`
    EvpntopoId         Optional[uuid.UUID]                           `json:"evpntopo_id"`
    ExtIp              Optional[string]                              `json:"ext_ip"`
    Fwupdate           *FwupdateStat                                 `json:"fwupdate,omitempty"`
    HwRev              Optional[string]                              `json:"hw_rev"`
    Id                 *uuid.UUID                                    `json:"id,omitempty"`
    InactiveWiredVlans []int                                         `json:"inactive_wired_vlans,omitempty"`
    IotStat            map[string]ApStatsIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    Ip                 Optional[string]                              `json:"ip"`
    IpConfig           *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat             *IpStat                                       `json:"ip_stat,omitempty"`
    L2tpStat           map[string]ApStatsL2TpStat                    `json:"l2tp_stat,omitempty"`
    LastSeen           Optional[float64]                             `json:"last_seen"`
    LastTrouble        *LastTrouble                                  `json:"last_trouble,omitempty"`
    Led                *ApLed                                        `json:"led,omitempty"`
    LldpStat           *ApStatsLldpStat                              `json:"lldp_stat,omitempty"`
    Locating           Optional[bool]                                `json:"locating"`
    Locked             Optional[bool]                                `json:"locked"`
    Mac                *string                                       `json:"mac"`
    MapId              Optional[uuid.UUID]                           `json:"map_id"`
    MemUsedKb          Optional[int64]                               `json:"mem_used_kb"`
    MeshDownlinks      map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink         *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    Model              *string                                       `json:"model"`
    ModifiedTime       Optional[int64]                               `json:"modified_time"`
    Mount              Optional[string]                              `json:"mount"`
    Name               Optional[string]                              `json:"name"`
    Notes              Optional[string]                              `json:"notes"`
    NumClients         Optional[int]                                 `json:"num_clients"`
    OrgId              Optional[uuid.UUID]                           `json:"org_id"`
    PortStat           Optional[map[string]ApStatsPortStat]          `json:"port_stat"`
    PowerBudget        Optional[int]                                 `json:"power_budget"`
    PowerConstrained   Optional[bool]                                `json:"power_constrained"`
    PowerOpmode        Optional[string]                              `json:"power_opmode"`
    PowerSrc           Optional[string]                              `json:"power_src"`
    RadioConfig        *ApStatsRadioConfig                           `json:"radio_config,omitempty"`
    RadioStat          *ApStatsRadioStat                             `json:"radio_stat,omitempty"`
    RxBps              Optional[float64]                             `json:"rx_bps"`
    RxBytes            Optional[int64]                               `json:"rx_bytes"`
    RxPkts             Optional[int]                                 `json:"rx_pkts"`
    Serial             Optional[string]                              `json:"serial"`
    SiteId             *uuid.UUID                                    `json:"site_id,omitempty"`
    Status             Optional[string]                              `json:"status"`
    SwitchRedundancy   *ApStatsSwitchRedundancy                      `json:"switch_redundancy,omitempty"`
    TxBps              Optional[float64]                             `json:"tx_bps"`
    TxBytes            Optional[float64]                             `json:"tx_bytes"`
    TxPkts             Optional[float64]                             `json:"tx_pkts"`
    Type               Optional[string]                              `json:"type"`
    Uptime             Optional[float64]                             `json:"uptime"`
    UsbStat            *ApStatsUsbStat                               `json:"usb_stat,omitempty"`
    Version            Optional[string]                              `json:"version"`
    X                  Optional[float64]                             `json:"x"`
    Y                  Optional[float64]                             `json:"y"`
}

func (a *apStats) validate() error {
    var errs []string
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
