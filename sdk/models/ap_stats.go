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
    BleConfig            *ApStatsBleConfig                             `json:"ble_config,omitempty"`
    BleStat              *ApStatsBleStat                               `json:"ble_stat,omitempty"`
    CertExpiry           *float64                                      `json:"cert_expiry,omitempty"`
    // device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage
    EnvStat              *ApStatsEnvStat                               `json:"env_stat,omitempty"`
    EslStat              *ApStatsEslStat                               `json:"esl_stat,omitempty"`
    ExtIp                *string                                       `json:"ext_ip,omitempty"`
    Fwupdate             *ApStatsFwupdate                              `json:"fwupdate,omitempty"`
    IotStat              map[string]ApStatsIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    Ip                   *string                                       `json:"ip,omitempty"`
    // IP AP settings
    IpConfig             *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat               *IpStat                                       `json:"ip_stat,omitempty"`
    // l2tp tunnel status (key is the wxtunnel_id)
    L2tpStat             map[string]ApStatsL2TpStat                    `json:"l2tp_stat,omitempty"`
    // last seen timestamp
    LastSeen             *float64                                      `json:"last_seen,omitempty"`
    // last trouble code of switch
    LastTrouble          *LastTrouble                                  `json:"last_trouble,omitempty"`
    // LED AP settings
    Led                  *ApLed                                        `json:"led,omitempty"`
    // LLDP Stat (neighbor information, power negotiations)
    LldpStat             *ApStatsLldpStat                              `json:"lldp_stat,omitempty"`
    Locating             *bool                                         `json:"locating,omitempty"`
    // whether this AP is considered locked (placement / orientation has been vetted)
    Locked               *bool                                         `json:"locked,omitempty"`
    // device mac
    Mac                  string                                        `json:"mac"`
    MapId                *uuid.UUID                                    `json:"map_id,omitempty"`
    // Property key is the mesh downlink id (e.g `00000000-0000-0000-1000-5c5b35000010`)
    MeshDownlinks        map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink           *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    // device model
    Model                string                                        `json:"model"`
    Mount                *string                                       `json:"mount,omitempty"`
    Name                 *string                                       `json:"name,omitempty"`
    // how many wireless clients are currently connected
    NumClients           *int                                          `json:"num_clients,omitempty"`
    // Property key is the port name (e.g. `eth0`)
    PortStat             map[string]ApStatsPortStat                    `json:"port_stat,omitempty"`
    // in mW, surplus if positive or deficit if negative
    PowerBudget          *float64                                      `json:"power_budget,omitempty"`
    // whether insufficient power
    PowerConstrained     *bool                                         `json:"power_constrained,omitempty"`
    // constrained mode
    PowerOpmode          *string                                       `json:"power_opmode,omitempty"`
    // DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown)
    PowerSrc             *string                                       `json:"power_src,omitempty"`
    RadioConfig          *ApStatsRadioConfig                           `json:"radio_config,omitempty"`
    RadioStat            *ApStatsRadioStat                             `json:"radio_stat,omitempty"`
    RxBps                *float64                                      `json:"rx_bps,omitempty"`
    RxBytes              *int64                                        `json:"rx_bytes,omitempty"`
    RxPkts               *int                                          `json:"rx_pkts,omitempty"`
    // serial
    Serial               *string                                       `json:"serial,omitempty"`
    Status               *ApStatsStatusEnum                            `json:"status,omitempty"`
    SwitchRedundancy     *ApStatsSwitchRedundancy                      `json:"switch_redundancy,omitempty"`
    TxBps                *float64                                      `json:"tx_bps,omitempty"`
    TxBytes              *float64                                      `json:"tx_bytes,omitempty"`
    TxPkts               *float64                                      `json:"tx_pkts,omitempty"`
    // device type, ap / ble
    Type                 *string                                       `json:"type,omitempty"`
    // how long, in seconds, has the device been up (or rebooted)
    Uptime               *float64                                      `json:"uptime,omitempty"`
    UsbStat              *ApStatsUsbStat                               `json:"usb_stat,omitempty"`
    Version              *string                                       `json:"version,omitempty"`
    X                    *float64                                      `json:"x,omitempty"`
    Y                    *float64                                      `json:"y,omitempty"`
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
    if a.BleConfig != nil {
        structMap["ble_config"] = a.BleConfig.toMap()
    }
    if a.BleStat != nil {
        structMap["ble_stat"] = a.BleStat.toMap()
    }
    if a.CertExpiry != nil {
        structMap["cert_expiry"] = a.CertExpiry
    }
    if a.EnvStat != nil {
        structMap["env_stat"] = a.EnvStat.toMap()
    }
    if a.EslStat != nil {
        structMap["esl_stat"] = a.EslStat.toMap()
    }
    if a.ExtIp != nil {
        structMap["ext_ip"] = a.ExtIp
    }
    if a.Fwupdate != nil {
        structMap["fwupdate"] = a.Fwupdate.toMap()
    }
    if a.IotStat != nil {
        structMap["iot_stat"] = a.IotStat
    }
    if a.Ip != nil {
        structMap["ip"] = a.Ip
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
    if a.LastSeen != nil {
        structMap["last_seen"] = a.LastSeen
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
    if a.Locating != nil {
        structMap["locating"] = a.Locating
    }
    if a.Locked != nil {
        structMap["locked"] = a.Locked
    }
    structMap["mac"] = a.Mac
    if a.MapId != nil {
        structMap["map_id"] = a.MapId
    }
    if a.MeshDownlinks != nil {
        structMap["mesh_downlinks"] = a.MeshDownlinks
    }
    if a.MeshUplink != nil {
        structMap["mesh_uplink"] = a.MeshUplink.toMap()
    }
    structMap["model"] = a.Model
    if a.Mount != nil {
        structMap["mount"] = a.Mount
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.NumClients != nil {
        structMap["num_clients"] = a.NumClients
    }
    if a.PortStat != nil {
        structMap["port_stat"] = a.PortStat
    }
    if a.PowerBudget != nil {
        structMap["power_budget"] = a.PowerBudget
    }
    if a.PowerConstrained != nil {
        structMap["power_constrained"] = a.PowerConstrained
    }
    if a.PowerOpmode != nil {
        structMap["power_opmode"] = a.PowerOpmode
    }
    if a.PowerSrc != nil {
        structMap["power_src"] = a.PowerSrc
    }
    if a.RadioConfig != nil {
        structMap["radio_config"] = a.RadioConfig.toMap()
    }
    if a.RadioStat != nil {
        structMap["radio_stat"] = a.RadioStat.toMap()
    }
    if a.RxBps != nil {
        structMap["rx_bps"] = a.RxBps
    }
    if a.RxBytes != nil {
        structMap["rx_bytes"] = a.RxBytes
    }
    if a.RxPkts != nil {
        structMap["rx_pkts"] = a.RxPkts
    }
    if a.Serial != nil {
        structMap["serial"] = a.Serial
    }
    if a.Status != nil {
        structMap["status"] = a.Status
    }
    if a.SwitchRedundancy != nil {
        structMap["switch_redundancy"] = a.SwitchRedundancy.toMap()
    }
    if a.TxBps != nil {
        structMap["tx_bps"] = a.TxBps
    }
    if a.TxBytes != nil {
        structMap["tx_bytes"] = a.TxBytes
    }
    if a.TxPkts != nil {
        structMap["tx_pkts"] = a.TxPkts
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    if a.Uptime != nil {
        structMap["uptime"] = a.Uptime
    }
    if a.UsbStat != nil {
        structMap["usb_stat"] = a.UsbStat.toMap()
    }
    if a.Version != nil {
        structMap["version"] = a.Version
    }
    if a.X != nil {
        structMap["x"] = a.X
    }
    if a.Y != nil {
        structMap["y"] = a.Y
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auto_placement", "ble_config", "ble_stat", "cert_expiry", "env_stat", "esl_stat", "ext_ip", "fwupdate", "iot_stat", "ip", "ip_config", "ip_stat", "l2tp_stat", "last_seen", "last_trouble", "led", "lldp_stat", "locating", "locked", "mac", "map_id", "mesh_downlinks", "mesh_uplink", "model", "mount", "name", "num_clients", "port_stat", "power_budget", "power_constrained", "power_opmode", "power_src", "radio_config", "radio_stat", "rx_bps", "rx_bytes", "rx_pkts", "serial", "status", "switch_redundancy", "tx_bps", "tx_bytes", "tx_pkts", "type", "uptime", "usb_stat", "version", "x", "y")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AutoPlacement = temp.AutoPlacement
    a.BleConfig = temp.BleConfig
    a.BleStat = temp.BleStat
    a.CertExpiry = temp.CertExpiry
    a.EnvStat = temp.EnvStat
    a.EslStat = temp.EslStat
    a.ExtIp = temp.ExtIp
    a.Fwupdate = temp.Fwupdate
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
    a.Mac = *temp.Mac
    a.MapId = temp.MapId
    a.MeshDownlinks = temp.MeshDownlinks
    a.MeshUplink = temp.MeshUplink
    a.Model = *temp.Model
    a.Mount = temp.Mount
    a.Name = temp.Name
    a.NumClients = temp.NumClients
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
    AutoPlacement    *ApStatsAutoPlacement                         `json:"auto_placement,omitempty"`
    BleConfig        *ApStatsBleConfig                             `json:"ble_config,omitempty"`
    BleStat          *ApStatsBleStat                               `json:"ble_stat,omitempty"`
    CertExpiry       *float64                                      `json:"cert_expiry,omitempty"`
    EnvStat          *ApStatsEnvStat                               `json:"env_stat,omitempty"`
    EslStat          *ApStatsEslStat                               `json:"esl_stat,omitempty"`
    ExtIp            *string                                       `json:"ext_ip,omitempty"`
    Fwupdate         *ApStatsFwupdate                              `json:"fwupdate,omitempty"`
    IotStat          map[string]ApStatsIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    Ip               *string                                       `json:"ip,omitempty"`
    IpConfig         *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat           *IpStat                                       `json:"ip_stat,omitempty"`
    L2tpStat         map[string]ApStatsL2TpStat                    `json:"l2tp_stat,omitempty"`
    LastSeen         *float64                                      `json:"last_seen,omitempty"`
    LastTrouble      *LastTrouble                                  `json:"last_trouble,omitempty"`
    Led              *ApLed                                        `json:"led,omitempty"`
    LldpStat         *ApStatsLldpStat                              `json:"lldp_stat,omitempty"`
    Locating         *bool                                         `json:"locating,omitempty"`
    Locked           *bool                                         `json:"locked,omitempty"`
    Mac              *string                                       `json:"mac"`
    MapId            *uuid.UUID                                    `json:"map_id,omitempty"`
    MeshDownlinks    map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink       *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    Model            *string                                       `json:"model"`
    Mount            *string                                       `json:"mount,omitempty"`
    Name             *string                                       `json:"name,omitempty"`
    NumClients       *int                                          `json:"num_clients,omitempty"`
    PortStat         map[string]ApStatsPortStat                    `json:"port_stat,omitempty"`
    PowerBudget      *float64                                      `json:"power_budget,omitempty"`
    PowerConstrained *bool                                         `json:"power_constrained,omitempty"`
    PowerOpmode      *string                                       `json:"power_opmode,omitempty"`
    PowerSrc         *string                                       `json:"power_src,omitempty"`
    RadioConfig      *ApStatsRadioConfig                           `json:"radio_config,omitempty"`
    RadioStat        *ApStatsRadioStat                             `json:"radio_stat,omitempty"`
    RxBps            *float64                                      `json:"rx_bps,omitempty"`
    RxBytes          *int64                                        `json:"rx_bytes,omitempty"`
    RxPkts           *int                                          `json:"rx_pkts,omitempty"`
    Serial           *string                                       `json:"serial,omitempty"`
    Status           *ApStatsStatusEnum                            `json:"status,omitempty"`
    SwitchRedundancy *ApStatsSwitchRedundancy                      `json:"switch_redundancy,omitempty"`
    TxBps            *float64                                      `json:"tx_bps,omitempty"`
    TxBytes          *float64                                      `json:"tx_bytes,omitempty"`
    TxPkts           *float64                                      `json:"tx_pkts,omitempty"`
    Type             *string                                       `json:"type,omitempty"`
    Uptime           *float64                                      `json:"uptime,omitempty"`
    UsbStat          *ApStatsUsbStat                               `json:"usb_stat,omitempty"`
    Version          *string                                       `json:"version,omitempty"`
    X                *float64                                      `json:"x,omitempty"`
    Y                *float64                                      `json:"y,omitempty"`
}

func (a *apStats) validate() error {
    var errs []string
    if a.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Ap_Stats`")
    }
    if a.Model == nil {
        errs = append(errs, "required field `model` is missing for type `Ap_Stats`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
