package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// StatsDevice represents a StatsDevice struct.
type StatsDevice struct {
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
    // IP address
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
    Mac                  *MemoryStat1                                  `json:"mac,omitempty"`
    MapId                *uuid.UUID                                    `json:"map_id,omitempty"`
    // Property key is the mesh downlink id (e.g `00000000-0000-0000-1000-5c5b35000010`)
    MeshDownlinks        map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink           *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    // device model
    Model                *string                                       `json:"model,omitempty"`
    Mount                *string                                       `json:"mount,omitempty"`
    // device name if configured
    Name                 *string                                       `json:"name,omitempty"`
    // how many wireless clients are currently connected
    NumClients           *SwitchStatsNumClients1                       `json:"num_clients,omitempty"`
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
    ApRedundancy         *SwitchStatsApRedundancy1                     `json:"ap_redundancy,omitempty"`
    Clients              []SwitchStatsClient                           `json:"clients,omitempty"`
    CpuStat              *CpuStat                                      `json:"cpu_stat,omitempty"`
    // whether the switch supports packet capture
    HasPcap              *bool                                         `json:"has_pcap,omitempty"`
    // hostname reported by the device
    Hostname             *string                                       `json:"hostname,omitempty"`
    // Property key is the interface name
    IfStat               map[string]SwitchStatsIfStat                  `json:"if_stat,omitempty"`
    ModuleStat           []ModuleStat                                  `json:"module_stat,omitempty"`
    ClusterStat          map[string]GatewayStatsClusterStat            `json:"cluster_stat,omitempty"`
    Cpu2Stat             *CpuStat                                      `json:"cpu2_stat,omitempty"`
    // Property key is the network name
    Dhcpd2Stat           map[string]GatewayStatsDhcpdStatLan           `json:"dhcpd2_stat,omitempty"`
    // Property key is the network name
    DhcpdStat            map[string]GatewayStatsDhcpdStatLan           `json:"dhcpd_stat,omitempty"`
    Ip2Stat              *IpStat                                       `json:"ip2_stat,omitempty"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    Memory2Stat          *MemoryStat                                   `json:"memory2_stat,omitempty"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    MemoryStat           *MemoryStat                                   `json:"memory_stat,omitempty"`
    Module2Stat          []ModuleStat                                  `json:"module2_stat,omitempty"`
    Spu2Stat             *GatewayStatsSpuStat                          `json:"spu2_stat,omitempty"`
    SpuStat              *GatewayStatsSpuStat                          `json:"spu_stat,omitempty"`
    AdditionalProperties map[string]any                                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsDevice.
// It customizes the JSON marshaling process for StatsDevice objects.
func (s StatsDevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsDevice object to a map representation for JSON marshaling.
func (s StatsDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AutoPlacement != nil {
        structMap["auto_placement"] = s.AutoPlacement.toMap()
    }
    if s.BleConfig != nil {
        structMap["ble_config"] = s.BleConfig.toMap()
    }
    if s.BleStat != nil {
        structMap["ble_stat"] = s.BleStat.toMap()
    }
    if s.CertExpiry != nil {
        structMap["cert_expiry"] = s.CertExpiry
    }
    if s.EnvStat != nil {
        structMap["env_stat"] = s.EnvStat.toMap()
    }
    if s.EslStat != nil {
        structMap["esl_stat"] = s.EslStat.toMap()
    }
    if s.ExtIp != nil {
        structMap["ext_ip"] = s.ExtIp
    }
    if s.Fwupdate != nil {
        structMap["fwupdate"] = s.Fwupdate.toMap()
    }
    if s.IotStat != nil {
        structMap["iot_stat"] = s.IotStat
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
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
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
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
    if s.Locating != nil {
        structMap["locating"] = s.Locating
    }
    if s.Locked != nil {
        structMap["locked"] = s.Locked
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac.toMap()
    }
    if s.MapId != nil {
        structMap["map_id"] = s.MapId
    }
    if s.MeshDownlinks != nil {
        structMap["mesh_downlinks"] = s.MeshDownlinks
    }
    if s.MeshUplink != nil {
        structMap["mesh_uplink"] = s.MeshUplink.toMap()
    }
    if s.Model != nil {
        structMap["model"] = s.Model
    }
    if s.Mount != nil {
        structMap["mount"] = s.Mount
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.NumClients != nil {
        structMap["num_clients"] = s.NumClients.toMap()
    }
    if s.PortStat != nil {
        structMap["port_stat"] = s.PortStat
    }
    if s.PowerBudget != nil {
        structMap["power_budget"] = s.PowerBudget
    }
    if s.PowerConstrained != nil {
        structMap["power_constrained"] = s.PowerConstrained
    }
    if s.PowerOpmode != nil {
        structMap["power_opmode"] = s.PowerOpmode
    }
    if s.PowerSrc != nil {
        structMap["power_src"] = s.PowerSrc
    }
    if s.RadioConfig != nil {
        structMap["radio_config"] = s.RadioConfig.toMap()
    }
    if s.RadioStat != nil {
        structMap["radio_stat"] = s.RadioStat.toMap()
    }
    if s.RxBps != nil {
        structMap["rx_bps"] = s.RxBps
    }
    if s.RxBytes != nil {
        structMap["rx_bytes"] = s.RxBytes
    }
    if s.RxPkts != nil {
        structMap["rx_pkts"] = s.RxPkts
    }
    if s.Serial != nil {
        structMap["serial"] = s.Serial
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.SwitchRedundancy != nil {
        structMap["switch_redundancy"] = s.SwitchRedundancy.toMap()
    }
    if s.TxBps != nil {
        structMap["tx_bps"] = s.TxBps
    }
    if s.TxBytes != nil {
        structMap["tx_bytes"] = s.TxBytes
    }
    if s.TxPkts != nil {
        structMap["tx_pkts"] = s.TxPkts
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.UsbStat != nil {
        structMap["usb_stat"] = s.UsbStat.toMap()
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    if s.X != nil {
        structMap["x"] = s.X
    }
    if s.Y != nil {
        structMap["y"] = s.Y
    }
    if s.ApRedundancy != nil {
        structMap["ap_redundancy"] = s.ApRedundancy.toMap()
    }
    if s.Clients != nil {
        structMap["clients"] = s.Clients
    }
    if s.CpuStat != nil {
        structMap["cpu_stat"] = s.CpuStat.toMap()
    }
    if s.HasPcap != nil {
        structMap["has_pcap"] = s.HasPcap
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.IfStat != nil {
        structMap["if_stat"] = s.IfStat
    }
    if s.ModuleStat != nil {
        structMap["module_stat"] = s.ModuleStat
    }
    if s.ClusterStat != nil {
        structMap["cluster_stat"] = s.ClusterStat
    }
    if s.Cpu2Stat != nil {
        structMap["cpu2_stat"] = s.Cpu2Stat.toMap()
    }
    if s.Dhcpd2Stat != nil {
        structMap["dhcpd2_stat"] = s.Dhcpd2Stat
    }
    if s.DhcpdStat != nil {
        structMap["dhcpd_stat"] = s.DhcpdStat
    }
    if s.Ip2Stat != nil {
        structMap["ip2_stat"] = s.Ip2Stat.toMap()
    }
    if s.Memory2Stat != nil {
        structMap["memory2_stat"] = s.Memory2Stat.toMap()
    }
    if s.MemoryStat != nil {
        structMap["memory_stat"] = s.MemoryStat.toMap()
    }
    if s.Module2Stat != nil {
        structMap["module2_stat"] = s.Module2Stat
    }
    if s.Spu2Stat != nil {
        structMap["spu2_stat"] = s.Spu2Stat.toMap()
    }
    if s.SpuStat != nil {
        structMap["spu_stat"] = s.SpuStat.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDevice.
// It customizes the JSON unmarshaling process for StatsDevice objects.
func (s *StatsDevice) UnmarshalJSON(input []byte) error {
    var temp statsDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auto_placement", "ble_config", "ble_stat", "cert_expiry", "env_stat", "esl_stat", "ext_ip", "fwupdate", "iot_stat", "ip", "ip_config", "ip_stat", "l2tp_stat", "last_seen", "last_trouble", "led", "lldp_stat", "locating", "locked", "mac", "map_id", "mesh_downlinks", "mesh_uplink", "model", "mount", "name", "num_clients", "port_stat", "power_budget", "power_constrained", "power_opmode", "power_src", "radio_config", "radio_stat", "rx_bps", "rx_bytes", "rx_pkts", "serial", "status", "switch_redundancy", "tx_bps", "tx_bytes", "tx_pkts", "type", "uptime", "usb_stat", "version", "x", "y", "ap_redundancy", "clients", "cpu_stat", "has_pcap", "hostname", "if_stat", "module_stat", "cluster_stat", "cpu2_stat", "dhcpd2_stat", "dhcpd_stat", "ip2_stat", "memory2_stat", "memory_stat", "module2_stat", "spu2_stat", "spu_stat")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AutoPlacement = temp.AutoPlacement
    s.BleConfig = temp.BleConfig
    s.BleStat = temp.BleStat
    s.CertExpiry = temp.CertExpiry
    s.EnvStat = temp.EnvStat
    s.EslStat = temp.EslStat
    s.ExtIp = temp.ExtIp
    s.Fwupdate = temp.Fwupdate
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
    s.MeshDownlinks = temp.MeshDownlinks
    s.MeshUplink = temp.MeshUplink
    s.Model = temp.Model
    s.Mount = temp.Mount
    s.Name = temp.Name
    s.NumClients = temp.NumClients
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
    s.Status = temp.Status
    s.SwitchRedundancy = temp.SwitchRedundancy
    s.TxBps = temp.TxBps
    s.TxBytes = temp.TxBytes
    s.TxPkts = temp.TxPkts
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    s.UsbStat = temp.UsbStat
    s.Version = temp.Version
    s.X = temp.X
    s.Y = temp.Y
    s.ApRedundancy = temp.ApRedundancy
    s.Clients = temp.Clients
    s.CpuStat = temp.CpuStat
    s.HasPcap = temp.HasPcap
    s.Hostname = temp.Hostname
    s.IfStat = temp.IfStat
    s.ModuleStat = temp.ModuleStat
    s.ClusterStat = temp.ClusterStat
    s.Cpu2Stat = temp.Cpu2Stat
    s.Dhcpd2Stat = temp.Dhcpd2Stat
    s.DhcpdStat = temp.DhcpdStat
    s.Ip2Stat = temp.Ip2Stat
    s.Memory2Stat = temp.Memory2Stat
    s.MemoryStat = temp.MemoryStat
    s.Module2Stat = temp.Module2Stat
    s.Spu2Stat = temp.Spu2Stat
    s.SpuStat = temp.SpuStat
    return nil
}

// statsDevice is a temporary struct used for validating the fields of StatsDevice.
type statsDevice  struct {
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
    Mac              *MemoryStat1                                  `json:"mac,omitempty"`
    MapId            *uuid.UUID                                    `json:"map_id,omitempty"`
    MeshDownlinks    map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink       *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    Model            *string                                       `json:"model,omitempty"`
    Mount            *string                                       `json:"mount,omitempty"`
    Name             *string                                       `json:"name,omitempty"`
    NumClients       *SwitchStatsNumClients1                       `json:"num_clients,omitempty"`
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
    ApRedundancy     *SwitchStatsApRedundancy1                     `json:"ap_redundancy,omitempty"`
    Clients          []SwitchStatsClient                           `json:"clients,omitempty"`
    CpuStat          *CpuStat                                      `json:"cpu_stat,omitempty"`
    HasPcap          *bool                                         `json:"has_pcap,omitempty"`
    Hostname         *string                                       `json:"hostname,omitempty"`
    IfStat           map[string]SwitchStatsIfStat                  `json:"if_stat,omitempty"`
    ModuleStat       []ModuleStat                                  `json:"module_stat,omitempty"`
    ClusterStat      map[string]GatewayStatsClusterStat            `json:"cluster_stat,omitempty"`
    Cpu2Stat         *CpuStat                                      `json:"cpu2_stat,omitempty"`
    Dhcpd2Stat       map[string]GatewayStatsDhcpdStatLan           `json:"dhcpd2_stat,omitempty"`
    DhcpdStat        map[string]GatewayStatsDhcpdStatLan           `json:"dhcpd_stat,omitempty"`
    Ip2Stat          *IpStat                                       `json:"ip2_stat,omitempty"`
    Memory2Stat      *MemoryStat                                   `json:"memory2_stat,omitempty"`
    MemoryStat       *MemoryStat                                   `json:"memory_stat,omitempty"`
    Module2Stat      []ModuleStat                                  `json:"module2_stat,omitempty"`
    Spu2Stat         *GatewayStatsSpuStat                          `json:"spu2_stat,omitempty"`
    SpuStat          *GatewayStatsSpuStat                          `json:"spu_stat,omitempty"`
}
