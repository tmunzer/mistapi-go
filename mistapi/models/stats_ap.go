package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsAp represents a StatsAp struct.
// AP statistics
type StatsAp struct {
    AutoPlacement        *StatsApAutoPlacement                         `json:"auto_placement,omitempty"`
    AutoUpgradeStat      *StatsApAutoUpgrade                           `json:"auto_upgrade_stat,omitempty"`
    BleStat              *StatsApBle                                   `json:"ble_stat,omitempty"`
    CertExpiry           Optional[float64]                             `json:"cert_expiry"`
    ConfigReverted       Optional[bool]                                `json:"config_reverted"`
    CpuSystem            Optional[int64]                               `json:"cpu_system"`
    CpuUtil              Optional[int]                                 `json:"cpu_util"`
    // when the object has been created, in epoch
    CreatedTime          *float64                                      `json:"created_time,omitempty"`
    DeviceprofileId      Optional[uuid.UUID]                           `json:"deviceprofile_id"`
    // device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage
    EnvStat              *StatsApEnvStat                               `json:"env_stat,omitempty"`
    EslStat              Optional[StatsApEslStat]                      `json:"esl_stat"`
    EvpntopoId           Optional[uuid.UUID]                           `json:"evpntopo_id"`
    ExtIp                Optional[string]                              `json:"ext_ip"`
    Fwupdate             *FwupdateStat                                 `json:"fwupdate,omitempty"`
    Gps                  *StatsApGpsStat                               `json:"gps,omitempty"`
    HwRev                Optional[string]                              `json:"hw_rev"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                                    `json:"id,omitempty"`
    InactiveWiredVlans   []int                                         `json:"inactive_wired_vlans,omitempty"`
    IotStat              map[string]StatsApIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    Ip                   Optional[string]                              `json:"ip"`
    // IP AP settings
    IpConfig             *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat               *IpStat                                       `json:"ip_stat,omitempty"`
    // l2tp tunnel status (key is the wxtunnel_id)
    L2tpStat             map[string]StatsApL2tpStat                    `json:"l2tp_stat,omitempty"`
    // last seen timestamp
    LastSeen             Optional[float64]                             `json:"last_seen"`
    // last trouble code of switch
    LastTrouble          *LastTrouble                                  `json:"last_trouble,omitempty"`
    // LED AP settings
    Led                  *ApLed                                        `json:"led,omitempty"`
    // LLDP Stat (neighbor information, power negotiations)
    LldpStat             *StatsApLldpStat                              `json:"lldp_stat,omitempty"`
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
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64                                      `json:"modified_time,omitempty"`
    Mount                Optional[string]                              `json:"mount"`
    Name                 Optional[string]                              `json:"name"`
    Notes                Optional[string]                              `json:"notes"`
    // how many wireless clients are currently connected
    NumClients           Optional[int]                                 `json:"num_clients"`
    // how many WLANs are applied to the device
    NumWlans             *int                                          `json:"num_wlans,omitempty"`
    OrgId                *uuid.UUID                                    `json:"org_id,omitempty"`
    // Property key is the port name (e.g. `eth0`)
    PortStat             Optional[map[string]StatsApPortStat]          `json:"port_stat"`
    // in mW, surplus if positive or deficit if negative
    PowerBudget          Optional[int]                                 `json:"power_budget"`
    // whether insufficient power
    PowerConstrained     Optional[bool]                                `json:"power_constrained"`
    // constrained mode
    PowerOpmode          Optional[string]                              `json:"power_opmode"`
    // DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown)
    PowerSrc             Optional[string]                              `json:"power_src"`
    RadioConfig          *StatsApRadioConfig                           `json:"radio_config,omitempty"`
    RadioStat            *StatsApRadioStat                             `json:"radio_stat,omitempty"`
    RxBps                Optional[float64]                             `json:"rx_bps"`
    RxBytes              Optional[int64]                               `json:"rx_bytes"`
    RxPkts               Optional[int]                                 `json:"rx_pkts"`
    // serial
    Serial               Optional[string]                              `json:"serial"`
    SiteId               *uuid.UUID                                    `json:"site_id,omitempty"`
    Status               Optional[string]                              `json:"status"`
    SwitchRedundancy     *StatsApSwitchRedundancy                      `json:"switch_redundancy,omitempty"`
    TxBps                Optional[float64]                             `json:"tx_bps"`
    TxBytes              Optional[float64]                             `json:"tx_bytes"`
    TxPkts               Optional[float64]                             `json:"tx_pkts"`
    // Device Type. enum: `ap`
    Type                 string                                        `json:"type"`
    // how long, in seconds, has the device been up (or rebooted)
    Uptime               Optional[float64]                             `json:"uptime"`
    UsbStat              *StatsApUsbStat                               `json:"usb_stat,omitempty"`
    Version              Optional[string]                              `json:"version"`
    X                    Optional[float64]                             `json:"x"`
    Y                    Optional[float64]                             `json:"y"`
    AdditionalProperties map[string]any                                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsAp.
// It customizes the JSON marshaling process for StatsAp objects.
func (s StatsAp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsAp object to a map representation for JSON marshaling.
func (s StatsAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AutoPlacement != nil {
        structMap["auto_placement"] = s.AutoPlacement.toMap()
    }
    if s.AutoUpgradeStat != nil {
        structMap["auto_upgrade_stat"] = s.AutoUpgradeStat.toMap()
    }
    if s.BleStat != nil {
        structMap["ble_stat"] = s.BleStat.toMap()
    }
    if s.CertExpiry.IsValueSet() {
        if s.CertExpiry.Value() != nil {
            structMap["cert_expiry"] = s.CertExpiry.Value()
        } else {
            structMap["cert_expiry"] = nil
        }
    }
    if s.ConfigReverted.IsValueSet() {
        if s.ConfigReverted.Value() != nil {
            structMap["config_reverted"] = s.ConfigReverted.Value()
        } else {
            structMap["config_reverted"] = nil
        }
    }
    if s.CpuSystem.IsValueSet() {
        if s.CpuSystem.Value() != nil {
            structMap["cpu_system"] = s.CpuSystem.Value()
        } else {
            structMap["cpu_system"] = nil
        }
    }
    if s.CpuUtil.IsValueSet() {
        if s.CpuUtil.Value() != nil {
            structMap["cpu_util"] = s.CpuUtil.Value()
        } else {
            structMap["cpu_util"] = nil
        }
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.DeviceprofileId.IsValueSet() {
        if s.DeviceprofileId.Value() != nil {
            structMap["deviceprofile_id"] = s.DeviceprofileId.Value()
        } else {
            structMap["deviceprofile_id"] = nil
        }
    }
    if s.EnvStat != nil {
        structMap["env_stat"] = s.EnvStat.toMap()
    }
    if s.EslStat.IsValueSet() {
        if s.EslStat.Value() != nil {
            structMap["esl_stat"] = s.EslStat.Value().toMap()
        } else {
            structMap["esl_stat"] = nil
        }
    }
    if s.EvpntopoId.IsValueSet() {
        if s.EvpntopoId.Value() != nil {
            structMap["evpntopo_id"] = s.EvpntopoId.Value()
        } else {
            structMap["evpntopo_id"] = nil
        }
    }
    if s.ExtIp.IsValueSet() {
        if s.ExtIp.Value() != nil {
            structMap["ext_ip"] = s.ExtIp.Value()
        } else {
            structMap["ext_ip"] = nil
        }
    }
    if s.Fwupdate != nil {
        structMap["fwupdate"] = s.Fwupdate.toMap()
    }
    if s.Gps != nil {
        structMap["gps"] = s.Gps.toMap()
    }
    if s.HwRev.IsValueSet() {
        if s.HwRev.Value() != nil {
            structMap["hw_rev"] = s.HwRev.Value()
        } else {
            structMap["hw_rev"] = nil
        }
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.InactiveWiredVlans != nil {
        structMap["inactive_wired_vlans"] = s.InactiveWiredVlans
    }
    if s.IotStat != nil {
        structMap["iot_stat"] = s.IotStat
    }
    if s.Ip.IsValueSet() {
        if s.Ip.Value() != nil {
            structMap["ip"] = s.Ip.Value()
        } else {
            structMap["ip"] = nil
        }
    }
    if s.IpConfig != nil {
        structMap["ip_config"] = s.IpConfig.toMap()
    }
    if s.IpStat != nil {
        structMap["ip_stat"] = s.IpStat.toMap()
    }
    if s.L2tpStat != nil {
        structMap["l2tp_stat"] = s.L2tpStat
    }
    if s.LastSeen.IsValueSet() {
        if s.LastSeen.Value() != nil {
            structMap["last_seen"] = s.LastSeen.Value()
        } else {
            structMap["last_seen"] = nil
        }
    }
    if s.LastTrouble != nil {
        structMap["last_trouble"] = s.LastTrouble.toMap()
    }
    if s.Led != nil {
        structMap["led"] = s.Led.toMap()
    }
    if s.LldpStat != nil {
        structMap["lldp_stat"] = s.LldpStat.toMap()
    }
    if s.Locating.IsValueSet() {
        if s.Locating.Value() != nil {
            structMap["locating"] = s.Locating.Value()
        } else {
            structMap["locating"] = nil
        }
    }
    if s.Locked.IsValueSet() {
        if s.Locked.Value() != nil {
            structMap["locked"] = s.Locked.Value()
        } else {
            structMap["locked"] = nil
        }
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    } else {
        structMap["mac"] = nil
    }
    if s.MapId.IsValueSet() {
        if s.MapId.Value() != nil {
            structMap["map_id"] = s.MapId.Value()
        } else {
            structMap["map_id"] = nil
        }
    }
    if s.MemUsedKb.IsValueSet() {
        if s.MemUsedKb.Value() != nil {
            structMap["mem_used_kb"] = s.MemUsedKb.Value()
        } else {
            structMap["mem_used_kb"] = nil
        }
    }
    if s.MeshDownlinks != nil {
        structMap["mesh_downlinks"] = s.MeshDownlinks
    }
    if s.MeshUplink != nil {
        structMap["mesh_uplink"] = s.MeshUplink.toMap()
    }
    if s.Model != nil {
        structMap["model"] = s.Model
    } else {
        structMap["model"] = nil
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.Mount.IsValueSet() {
        if s.Mount.Value() != nil {
            structMap["mount"] = s.Mount.Value()
        } else {
            structMap["mount"] = nil
        }
    }
    if s.Name.IsValueSet() {
        if s.Name.Value() != nil {
            structMap["name"] = s.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if s.Notes.IsValueSet() {
        if s.Notes.Value() != nil {
            structMap["notes"] = s.Notes.Value()
        } else {
            structMap["notes"] = nil
        }
    }
    if s.NumClients.IsValueSet() {
        if s.NumClients.Value() != nil {
            structMap["num_clients"] = s.NumClients.Value()
        } else {
            structMap["num_clients"] = nil
        }
    }
    if s.NumWlans != nil {
        structMap["num_wlans"] = s.NumWlans
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.PortStat.IsValueSet() {
        if s.PortStat.Value() != nil {
            structMap["port_stat"] = s.PortStat.Value()
        } else {
            structMap["port_stat"] = nil
        }
    }
    if s.PowerBudget.IsValueSet() {
        if s.PowerBudget.Value() != nil {
            structMap["power_budget"] = s.PowerBudget.Value()
        } else {
            structMap["power_budget"] = nil
        }
    }
    if s.PowerConstrained.IsValueSet() {
        if s.PowerConstrained.Value() != nil {
            structMap["power_constrained"] = s.PowerConstrained.Value()
        } else {
            structMap["power_constrained"] = nil
        }
    }
    if s.PowerOpmode.IsValueSet() {
        if s.PowerOpmode.Value() != nil {
            structMap["power_opmode"] = s.PowerOpmode.Value()
        } else {
            structMap["power_opmode"] = nil
        }
    }
    if s.PowerSrc.IsValueSet() {
        if s.PowerSrc.Value() != nil {
            structMap["power_src"] = s.PowerSrc.Value()
        } else {
            structMap["power_src"] = nil
        }
    }
    if s.RadioConfig != nil {
        structMap["radio_config"] = s.RadioConfig.toMap()
    }
    if s.RadioStat != nil {
        structMap["radio_stat"] = s.RadioStat.toMap()
    }
    if s.RxBps.IsValueSet() {
        if s.RxBps.Value() != nil {
            structMap["rx_bps"] = s.RxBps.Value()
        } else {
            structMap["rx_bps"] = nil
        }
    }
    if s.RxBytes.IsValueSet() {
        if s.RxBytes.Value() != nil {
            structMap["rx_bytes"] = s.RxBytes.Value()
        } else {
            structMap["rx_bytes"] = nil
        }
    }
    if s.RxPkts.IsValueSet() {
        if s.RxPkts.Value() != nil {
            structMap["rx_pkts"] = s.RxPkts.Value()
        } else {
            structMap["rx_pkts"] = nil
        }
    }
    if s.Serial.IsValueSet() {
        if s.Serial.Value() != nil {
            structMap["serial"] = s.Serial.Value()
        } else {
            structMap["serial"] = nil
        }
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Status.IsValueSet() {
        if s.Status.Value() != nil {
            structMap["status"] = s.Status.Value()
        } else {
            structMap["status"] = nil
        }
    }
    if s.SwitchRedundancy != nil {
        structMap["switch_redundancy"] = s.SwitchRedundancy.toMap()
    }
    if s.TxBps.IsValueSet() {
        if s.TxBps.Value() != nil {
            structMap["tx_bps"] = s.TxBps.Value()
        } else {
            structMap["tx_bps"] = nil
        }
    }
    if s.TxBytes.IsValueSet() {
        if s.TxBytes.Value() != nil {
            structMap["tx_bytes"] = s.TxBytes.Value()
        } else {
            structMap["tx_bytes"] = nil
        }
    }
    if s.TxPkts.IsValueSet() {
        if s.TxPkts.Value() != nil {
            structMap["tx_pkts"] = s.TxPkts.Value()
        } else {
            structMap["tx_pkts"] = nil
        }
    }
    structMap["type"] = s.Type
    if s.Uptime.IsValueSet() {
        if s.Uptime.Value() != nil {
            structMap["uptime"] = s.Uptime.Value()
        } else {
            structMap["uptime"] = nil
        }
    }
    if s.UsbStat != nil {
        structMap["usb_stat"] = s.UsbStat.toMap()
    }
    if s.Version.IsValueSet() {
        if s.Version.Value() != nil {
            structMap["version"] = s.Version.Value()
        } else {
            structMap["version"] = nil
        }
    }
    if s.X.IsValueSet() {
        if s.X.Value() != nil {
            structMap["x"] = s.X.Value()
        } else {
            structMap["x"] = nil
        }
    }
    if s.Y.IsValueSet() {
        if s.Y.Value() != nil {
            structMap["y"] = s.Y.Value()
        } else {
            structMap["y"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsAp.
// It customizes the JSON unmarshaling process for StatsAp objects.
func (s *StatsAp) UnmarshalJSON(input []byte) error {
    var temp tempStatsAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auto_placement", "auto_upgrade_stat", "ble_stat", "cert_expiry", "config_reverted", "cpu_system", "cpu_util", "created_time", "deviceprofile_id", "env_stat", "esl_stat", "evpntopo_id", "ext_ip", "fwupdate", "gps", "hw_rev", "id", "inactive_wired_vlans", "iot_stat", "ip", "ip_config", "ip_stat", "l2tp_stat", "last_seen", "last_trouble", "led", "lldp_stat", "locating", "locked", "mac", "map_id", "mem_used_kb", "mesh_downlinks", "mesh_uplink", "model", "modified_time", "mount", "name", "notes", "num_clients", "num_wlans", "org_id", "port_stat", "power_budget", "power_constrained", "power_opmode", "power_src", "radio_config", "radio_stat", "rx_bps", "rx_bytes", "rx_pkts", "serial", "site_id", "status", "switch_redundancy", "tx_bps", "tx_bytes", "tx_pkts", "type", "uptime", "usb_stat", "version", "x", "y")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AutoPlacement = temp.AutoPlacement
    s.AutoUpgradeStat = temp.AutoUpgradeStat
    s.BleStat = temp.BleStat
    s.CertExpiry = temp.CertExpiry
    s.ConfigReverted = temp.ConfigReverted
    s.CpuSystem = temp.CpuSystem
    s.CpuUtil = temp.CpuUtil
    s.CreatedTime = temp.CreatedTime
    s.DeviceprofileId = temp.DeviceprofileId
    s.EnvStat = temp.EnvStat
    s.EslStat = temp.EslStat
    s.EvpntopoId = temp.EvpntopoId
    s.ExtIp = temp.ExtIp
    s.Fwupdate = temp.Fwupdate
    s.Gps = temp.Gps
    s.HwRev = temp.HwRev
    s.Id = temp.Id
    s.InactiveWiredVlans = temp.InactiveWiredVlans
    s.IotStat = temp.IotStat
    s.Ip = temp.Ip
    s.IpConfig = temp.IpConfig
    s.IpStat = temp.IpStat
    s.L2tpStat = temp.L2tpStat
    s.LastSeen = temp.LastSeen
    s.LastTrouble = temp.LastTrouble
    s.Led = temp.Led
    s.LldpStat = temp.LldpStat
    s.Locating = temp.Locating
    s.Locked = temp.Locked
    s.Mac = temp.Mac
    s.MapId = temp.MapId
    s.MemUsedKb = temp.MemUsedKb
    s.MeshDownlinks = temp.MeshDownlinks
    s.MeshUplink = temp.MeshUplink
    s.Model = temp.Model
    s.ModifiedTime = temp.ModifiedTime
    s.Mount = temp.Mount
    s.Name = temp.Name
    s.Notes = temp.Notes
    s.NumClients = temp.NumClients
    s.NumWlans = temp.NumWlans
    s.OrgId = temp.OrgId
    s.PortStat = temp.PortStat
    s.PowerBudget = temp.PowerBudget
    s.PowerConstrained = temp.PowerConstrained
    s.PowerOpmode = temp.PowerOpmode
    s.PowerSrc = temp.PowerSrc
    s.RadioConfig = temp.RadioConfig
    s.RadioStat = temp.RadioStat
    s.RxBps = temp.RxBps
    s.RxBytes = temp.RxBytes
    s.RxPkts = temp.RxPkts
    s.Serial = temp.Serial
    s.SiteId = temp.SiteId
    s.Status = temp.Status
    s.SwitchRedundancy = temp.SwitchRedundancy
    s.TxBps = temp.TxBps
    s.TxBytes = temp.TxBytes
    s.TxPkts = temp.TxPkts
    s.Type = *temp.Type
    s.Uptime = temp.Uptime
    s.UsbStat = temp.UsbStat
    s.Version = temp.Version
    s.X = temp.X
    s.Y = temp.Y
    return nil
}

// tempStatsAp is a temporary struct used for validating the fields of StatsAp.
type tempStatsAp  struct {
    AutoPlacement      *StatsApAutoPlacement                         `json:"auto_placement,omitempty"`
    AutoUpgradeStat    *StatsApAutoUpgrade                           `json:"auto_upgrade_stat,omitempty"`
    BleStat            *StatsApBle                                   `json:"ble_stat,omitempty"`
    CertExpiry         Optional[float64]                             `json:"cert_expiry"`
    ConfigReverted     Optional[bool]                                `json:"config_reverted"`
    CpuSystem          Optional[int64]                               `json:"cpu_system"`
    CpuUtil            Optional[int]                                 `json:"cpu_util"`
    CreatedTime        *float64                                      `json:"created_time,omitempty"`
    DeviceprofileId    Optional[uuid.UUID]                           `json:"deviceprofile_id"`
    EnvStat            *StatsApEnvStat                               `json:"env_stat,omitempty"`
    EslStat            Optional[StatsApEslStat]                      `json:"esl_stat"`
    EvpntopoId         Optional[uuid.UUID]                           `json:"evpntopo_id"`
    ExtIp              Optional[string]                              `json:"ext_ip"`
    Fwupdate           *FwupdateStat                                 `json:"fwupdate,omitempty"`
    Gps                *StatsApGpsStat                               `json:"gps,omitempty"`
    HwRev              Optional[string]                              `json:"hw_rev"`
    Id                 *uuid.UUID                                    `json:"id,omitempty"`
    InactiveWiredVlans []int                                         `json:"inactive_wired_vlans,omitempty"`
    IotStat            map[string]StatsApIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    Ip                 Optional[string]                              `json:"ip"`
    IpConfig           *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat             *IpStat                                       `json:"ip_stat,omitempty"`
    L2tpStat           map[string]StatsApL2tpStat                    `json:"l2tp_stat,omitempty"`
    LastSeen           Optional[float64]                             `json:"last_seen"`
    LastTrouble        *LastTrouble                                  `json:"last_trouble,omitempty"`
    Led                *ApLed                                        `json:"led,omitempty"`
    LldpStat           *StatsApLldpStat                              `json:"lldp_stat,omitempty"`
    Locating           Optional[bool]                                `json:"locating"`
    Locked             Optional[bool]                                `json:"locked"`
    Mac                *string                                       `json:"mac"`
    MapId              Optional[uuid.UUID]                           `json:"map_id"`
    MemUsedKb          Optional[int64]                               `json:"mem_used_kb"`
    MeshDownlinks      map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink         *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    Model              *string                                       `json:"model"`
    ModifiedTime       *float64                                      `json:"modified_time,omitempty"`
    Mount              Optional[string]                              `json:"mount"`
    Name               Optional[string]                              `json:"name"`
    Notes              Optional[string]                              `json:"notes"`
    NumClients         Optional[int]                                 `json:"num_clients"`
    NumWlans           *int                                          `json:"num_wlans,omitempty"`
    OrgId              *uuid.UUID                                    `json:"org_id,omitempty"`
    PortStat           Optional[map[string]StatsApPortStat]          `json:"port_stat"`
    PowerBudget        Optional[int]                                 `json:"power_budget"`
    PowerConstrained   Optional[bool]                                `json:"power_constrained"`
    PowerOpmode        Optional[string]                              `json:"power_opmode"`
    PowerSrc           Optional[string]                              `json:"power_src"`
    RadioConfig        *StatsApRadioConfig                           `json:"radio_config,omitempty"`
    RadioStat          *StatsApRadioStat                             `json:"radio_stat,omitempty"`
    RxBps              Optional[float64]                             `json:"rx_bps"`
    RxBytes            Optional[int64]                               `json:"rx_bytes"`
    RxPkts             Optional[int]                                 `json:"rx_pkts"`
    Serial             Optional[string]                              `json:"serial"`
    SiteId             *uuid.UUID                                    `json:"site_id,omitempty"`
    Status             Optional[string]                              `json:"status"`
    SwitchRedundancy   *StatsApSwitchRedundancy                      `json:"switch_redundancy,omitempty"`
    TxBps              Optional[float64]                             `json:"tx_bps"`
    TxBytes            Optional[float64]                             `json:"tx_bytes"`
    TxPkts             Optional[float64]                             `json:"tx_pkts"`
    Type               *string                                       `json:"type"`
    Uptime             Optional[float64]                             `json:"uptime"`
    UsbStat            *StatsApUsbStat                               `json:"usb_stat,omitempty"`
    Version            Optional[string]                              `json:"version"`
    X                  Optional[float64]                             `json:"x"`
    Y                  Optional[float64]                             `json:"y"`
}

func (s *tempStatsAp) validate() error {
    var errs []string
    if s.Type == nil {
        errs = append(errs, "required field `type` is missing for type `stats_ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
