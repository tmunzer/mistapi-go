package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// StatsDevice represents a StatsDevice struct.
type StatsDevice struct {
    AutoPlacement        *ApStatsAutoPlacement                         `json:"auto_placement,omitempty"`
    AutoUpgradeStat      *ApStatsAutoUpgrade                           `json:"auto_upgrade_stat,omitempty"`
    BleStat              *ApStatsBle                                   `json:"ble_stat,omitempty"`
    CertExpiry           Optional[int64]                               `json:"cert_expiry"`
    ConfigReverted       Optional[bool]                                `json:"config_reverted"`
    CpuSystem            Optional[int64]                               `json:"cpu_system"`
    CpuUtil              Optional[int]                                 `json:"cpu_util"`
    CreatedTime          *float64                                      `json:"created_time,omitempty"`
    DeviceprofileId      Optional[uuid.UUID]                           `json:"deviceprofile_id"`
    // device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage
    EnvStat              *ApStatsEnvStat                               `json:"env_stat,omitempty"`
    EslStat              Optional[ApStatsEslStat]                      `json:"esl_stat"`
    EvpntopoId           Optional[uuid.UUID]                           `json:"evpntopo_id"`
    // IP address
    ExtIp                Optional[string]                              `json:"ext_ip"`
    Fwupdate             *FwupdateStat                                 `json:"fwupdate,omitempty"`
    // device hardware revision number
    HwRev                Optional[string]                              `json:"hw_rev"`
    // serial
    Id                   *uuid.UUID                                    `json:"id,omitempty"`
    InactiveWiredVlans   []int                                         `json:"inactive_wired_vlans,omitempty"`
    IotStat              map[string]ApStatsIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    // IP address
    Ip                   Optional[string]                              `json:"ip"`
    // IP AP settings
    IpConfig             *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat               *IpStat                                       `json:"ip_stat,omitempty"`
    // l2tp tunnel status (key is the wxtunnel_id)
    L2tpStat             map[string]ApStatsL2tpStat                    `json:"l2tp_stat,omitempty"`
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
    Mac                  Optional[string]                              `json:"mac"`
    // serial
    MapId                Optional[uuid.UUID]                           `json:"map_id"`
    MemUsedKb            Optional[int64]                               `json:"mem_used_kb"`
    // Property key is the mesh downlink id (e.g `00000000-0000-0000-1000-5c5b35000010`)
    MeshDownlinks        map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink           *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    // device model
    Model                Optional[string]                              `json:"model"`
    ModifiedTime         *float64                                      `json:"modified_time,omitempty"`
    Mount                Optional[string]                              `json:"mount"`
    // device name if configured
    Name                 Optional[string]                              `json:"name"`
    Notes                Optional[string]                              `json:"notes"`
    // how many wireless clients are currently connected
    NumClients           Optional[int]                                 `json:"num_clients"`
    // serial
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
    // serial
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
    ApRedundancy         *SwitchStatsApRedundancy1                     `json:"ap_redundancy,omitempty"`
    ArpTableStats        *ArpTableStats                                `json:"arp_table_stats,omitempty"`
    Clients              []SwitchStatsClientItem                       `json:"clients,omitempty"`
    ClientsStats         *SwitchStatsClientsStats                      `json:"clients_stats,omitempty"`
    ConfigStatus         *string                                       `json:"config_status,omitempty"`
    CpuStat              *CpuStat                                      `json:"cpu_stat,omitempty"`
    // Property key is the network name
    DhcpdStat            *DhcpdStat                                    `json:"dhcpd_stat,omitempty"`
    FwVersionsOutofsync  *bool                                         `json:"fw_versions_outofsync,omitempty"`
    // whether the switch supports packet capture
    HasPcap              *bool                                         `json:"has_pcap,omitempty"`
    // hostname reported by the device
    Hostname             *string                                       `json:"hostname,omitempty"`
    // Property key is the interface name
    IfStat               *IfStat                                       `json:"if_stat,omitempty"`
    MacTableStats        *MacTableStats                                `json:"mac_table_stats,omitempty"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    MemoryStat           *MemoryStat                                   `json:"memory_stat,omitempty"`
    ModuleStat           []ModuleStatItem1                             `json:"module_stat,omitempty"`
    Ports                []DeviceStatsPort                             `json:"ports,omitempty"`
    RouteSummaryStats    *RouteSummaryStats                            `json:"route_summary_stats,omitempty"`
    ServiceStat          *ServiceStatProperty                          `json:"service_stat,omitempty"`
    VcMac                Optional[string]                              `json:"vc_mac"`
    VcSetupInfo          *SwitchStatsVcSetupInfo                       `json:"vc_setup_info,omitempty"`
    ClusterConfig        *ClusterConfigStats                           `json:"cluster_config,omitempty"`
    ClusterStat          *GatewayStatsCluster                          `json:"cluster_stat,omitempty"`
    ConductorName        *string                                       `json:"conductor_name,omitempty"`
    Cpu2Stat             *CpuStat                                      `json:"cpu2_stat,omitempty"`
    // Property key is the network name
    Dhcpd2Stat           map[string]DhcpdStatLan                       `json:"dhcpd2_stat,omitempty"`
    // Property key is the interface name
    If2Stat              map[string]IfStatProperty                     `json:"if2_stat,omitempty"`
    Ip2Stat              *IpStat                                       `json:"ip2_stat,omitempty"`
    IsHa                 Optional[bool]                                `json:"is_ha"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    Memory2Stat          *MemoryStat                                   `json:"memory2_stat,omitempty"`
    Module2Stat          []ModuleStatItem                              `json:"module2_stat,omitempty"`
    NodeName             *string                                       `json:"node_name,omitempty"`
    // device name if configured
    RouterName           *string                                       `json:"router_name,omitempty"`
    Service2Stat         map[string]ServiceStatProperty                `json:"service2_stat,omitempty"`
    ServiceStatus        *GatewayStatsServiceStatus                    `json:"service_status,omitempty"`
    Spu2Stat             []GatewayStatsSpuItem                         `json:"spu2_stat,omitempty"`
    SpuStat              []GatewayStatsSpuItem                         `json:"spu_stat,omitempty"`
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
    if s.Mac.IsValueSet() {
        if s.Mac.Value() != nil {
            structMap["mac"] = s.Mac.Value()
        } else {
            structMap["mac"] = nil
        }
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
    if s.Model.IsValueSet() {
        if s.Model.Value() != nil {
            structMap["model"] = s.Model.Value()
        } else {
            structMap["model"] = nil
        }
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
    if s.OrgId.IsValueSet() {
        if s.OrgId.Value() != nil {
            structMap["org_id"] = s.OrgId.Value()
        } else {
            structMap["org_id"] = nil
        }
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
    if s.Type.IsValueSet() {
        if s.Type.Value() != nil {
            structMap["type"] = s.Type.Value()
        } else {
            structMap["type"] = nil
        }
    }
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
    if s.ApRedundancy != nil {
        structMap["ap_redundancy"] = s.ApRedundancy.toMap()
    }
    if s.ArpTableStats != nil {
        structMap["arp_table_stats"] = s.ArpTableStats.toMap()
    }
    if s.Clients != nil {
        structMap["clients"] = s.Clients
    }
    if s.ClientsStats != nil {
        structMap["clients_stats"] = s.ClientsStats.toMap()
    }
    if s.ConfigStatus != nil {
        structMap["config_status"] = s.ConfigStatus
    }
    if s.CpuStat != nil {
        structMap["cpu_stat"] = s.CpuStat.toMap()
    }
    if s.DhcpdStat != nil {
        structMap["dhcpd_stat"] = s.DhcpdStat.toMap()
    }
    if s.FwVersionsOutofsync != nil {
        structMap["fw_versions_outofsync"] = s.FwVersionsOutofsync
    }
    if s.HasPcap != nil {
        structMap["has_pcap"] = s.HasPcap
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.IfStat != nil {
        structMap["if_stat"] = s.IfStat.toMap()
    }
    if s.MacTableStats != nil {
        structMap["mac_table_stats"] = s.MacTableStats.toMap()
    }
    if s.MemoryStat != nil {
        structMap["memory_stat"] = s.MemoryStat.toMap()
    }
    if s.ModuleStat != nil {
        structMap["module_stat"] = s.ModuleStat
    }
    if s.Ports != nil {
        structMap["ports"] = s.Ports
    }
    if s.RouteSummaryStats != nil {
        structMap["route_summary_stats"] = s.RouteSummaryStats.toMap()
    }
    if s.ServiceStat != nil {
        structMap["service_stat"] = s.ServiceStat.toMap()
    }
    if s.VcMac.IsValueSet() {
        if s.VcMac.Value() != nil {
            structMap["vc_mac"] = s.VcMac.Value()
        } else {
            structMap["vc_mac"] = nil
        }
    }
    if s.VcSetupInfo != nil {
        structMap["vc_setup_info"] = s.VcSetupInfo.toMap()
    }
    if s.ClusterConfig != nil {
        structMap["cluster_config"] = s.ClusterConfig.toMap()
    }
    if s.ClusterStat != nil {
        structMap["cluster_stat"] = s.ClusterStat.toMap()
    }
    if s.ConductorName != nil {
        structMap["conductor_name"] = s.ConductorName
    }
    if s.Cpu2Stat != nil {
        structMap["cpu2_stat"] = s.Cpu2Stat.toMap()
    }
    if s.Dhcpd2Stat != nil {
        structMap["dhcpd2_stat"] = s.Dhcpd2Stat
    }
    if s.If2Stat != nil {
        structMap["if2_stat"] = s.If2Stat
    }
    if s.Ip2Stat != nil {
        structMap["ip2_stat"] = s.Ip2Stat.toMap()
    }
    if s.IsHa.IsValueSet() {
        if s.IsHa.Value() != nil {
            structMap["is_ha"] = s.IsHa.Value()
        } else {
            structMap["is_ha"] = nil
        }
    }
    if s.Memory2Stat != nil {
        structMap["memory2_stat"] = s.Memory2Stat.toMap()
    }
    if s.Module2Stat != nil {
        structMap["module2_stat"] = s.Module2Stat
    }
    if s.NodeName != nil {
        structMap["node_name"] = s.NodeName
    }
    if s.RouterName != nil {
        structMap["router_name"] = s.RouterName
    }
    if s.Service2Stat != nil {
        structMap["service2_stat"] = s.Service2Stat
    }
    if s.ServiceStatus != nil {
        structMap["service_status"] = s.ServiceStatus.toMap()
    }
    if s.Spu2Stat != nil {
        structMap["spu2_stat"] = s.Spu2Stat
    }
    if s.SpuStat != nil {
        structMap["spu_stat"] = s.SpuStat
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsDevice.
// It customizes the JSON unmarshaling process for StatsDevice objects.
func (s *StatsDevice) UnmarshalJSON(input []byte) error {
    var temp tempStatsDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auto_placement", "auto_upgrade_stat", "ble_stat", "cert_expiry", "config_reverted", "cpu_system", "cpu_util", "created_time", "deviceprofile_id", "env_stat", "esl_stat", "evpntopo_id", "ext_ip", "fwupdate", "hw_rev", "id", "inactive_wired_vlans", "iot_stat", "ip", "ip_config", "ip_stat", "l2tp_stat", "last_seen", "last_trouble", "led", "lldp_stat", "locating", "locked", "mac", "map_id", "mem_used_kb", "mesh_downlinks", "mesh_uplink", "model", "modified_time", "mount", "name", "notes", "num_clients", "org_id", "port_stat", "power_budget", "power_constrained", "power_opmode", "power_src", "radio_config", "radio_stat", "rx_bps", "rx_bytes", "rx_pkts", "serial", "site_id", "status", "switch_redundancy", "tx_bps", "tx_bytes", "tx_pkts", "type", "uptime", "usb_stat", "version", "x", "y", "ap_redundancy", "arp_table_stats", "clients", "clients_stats", "config_status", "cpu_stat", "dhcpd_stat", "fw_versions_outofsync", "has_pcap", "hostname", "if_stat", "mac_table_stats", "memory_stat", "module_stat", "ports", "route_summary_stats", "service_stat", "vc_mac", "vc_setup_info", "cluster_config", "cluster_stat", "conductor_name", "cpu2_stat", "dhcpd2_stat", "if2_stat", "ip2_stat", "is_ha", "memory2_stat", "module2_stat", "node_name", "router_name", "service2_stat", "service_status", "spu2_stat", "spu_stat")
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
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    s.UsbStat = temp.UsbStat
    s.Version = temp.Version
    s.X = temp.X
    s.Y = temp.Y
    s.ApRedundancy = temp.ApRedundancy
    s.ArpTableStats = temp.ArpTableStats
    s.Clients = temp.Clients
    s.ClientsStats = temp.ClientsStats
    s.ConfigStatus = temp.ConfigStatus
    s.CpuStat = temp.CpuStat
    s.DhcpdStat = temp.DhcpdStat
    s.FwVersionsOutofsync = temp.FwVersionsOutofsync
    s.HasPcap = temp.HasPcap
    s.Hostname = temp.Hostname
    s.IfStat = temp.IfStat
    s.MacTableStats = temp.MacTableStats
    s.MemoryStat = temp.MemoryStat
    s.ModuleStat = temp.ModuleStat
    s.Ports = temp.Ports
    s.RouteSummaryStats = temp.RouteSummaryStats
    s.ServiceStat = temp.ServiceStat
    s.VcMac = temp.VcMac
    s.VcSetupInfo = temp.VcSetupInfo
    s.ClusterConfig = temp.ClusterConfig
    s.ClusterStat = temp.ClusterStat
    s.ConductorName = temp.ConductorName
    s.Cpu2Stat = temp.Cpu2Stat
    s.Dhcpd2Stat = temp.Dhcpd2Stat
    s.If2Stat = temp.If2Stat
    s.Ip2Stat = temp.Ip2Stat
    s.IsHa = temp.IsHa
    s.Memory2Stat = temp.Memory2Stat
    s.Module2Stat = temp.Module2Stat
    s.NodeName = temp.NodeName
    s.RouterName = temp.RouterName
    s.Service2Stat = temp.Service2Stat
    s.ServiceStatus = temp.ServiceStatus
    s.Spu2Stat = temp.Spu2Stat
    s.SpuStat = temp.SpuStat
    return nil
}

// tempStatsDevice is a temporary struct used for validating the fields of StatsDevice.
type tempStatsDevice  struct {
    AutoPlacement       *ApStatsAutoPlacement                         `json:"auto_placement,omitempty"`
    AutoUpgradeStat     *ApStatsAutoUpgrade                           `json:"auto_upgrade_stat,omitempty"`
    BleStat             *ApStatsBle                                   `json:"ble_stat,omitempty"`
    CertExpiry          Optional[int64]                               `json:"cert_expiry"`
    ConfigReverted      Optional[bool]                                `json:"config_reverted"`
    CpuSystem           Optional[int64]                               `json:"cpu_system"`
    CpuUtil             Optional[int]                                 `json:"cpu_util"`
    CreatedTime         *float64                                      `json:"created_time,omitempty"`
    DeviceprofileId     Optional[uuid.UUID]                           `json:"deviceprofile_id"`
    EnvStat             *ApStatsEnvStat                               `json:"env_stat,omitempty"`
    EslStat             Optional[ApStatsEslStat]                      `json:"esl_stat"`
    EvpntopoId          Optional[uuid.UUID]                           `json:"evpntopo_id"`
    ExtIp               Optional[string]                              `json:"ext_ip"`
    Fwupdate            *FwupdateStat                                 `json:"fwupdate,omitempty"`
    HwRev               Optional[string]                              `json:"hw_rev"`
    Id                  *uuid.UUID                                    `json:"id,omitempty"`
    InactiveWiredVlans  []int                                         `json:"inactive_wired_vlans,omitempty"`
    IotStat             map[string]ApStatsIotStatAdditionalProperties `json:"iot_stat,omitempty"`
    Ip                  Optional[string]                              `json:"ip"`
    IpConfig            *ApIpConfig                                   `json:"ip_config,omitempty"`
    IpStat              *IpStat                                       `json:"ip_stat,omitempty"`
    L2tpStat            map[string]ApStatsL2tpStat                    `json:"l2tp_stat,omitempty"`
    LastSeen            Optional[float64]                             `json:"last_seen"`
    LastTrouble         *LastTrouble                                  `json:"last_trouble,omitempty"`
    Led                 *ApLed                                        `json:"led,omitempty"`
    LldpStat            *ApStatsLldpStat                              `json:"lldp_stat,omitempty"`
    Locating            Optional[bool]                                `json:"locating"`
    Locked              Optional[bool]                                `json:"locked"`
    Mac                 Optional[string]                              `json:"mac"`
    MapId               Optional[uuid.UUID]                           `json:"map_id"`
    MemUsedKb           Optional[int64]                               `json:"mem_used_kb"`
    MeshDownlinks       map[string]ApStatMeshDownlink                 `json:"mesh_downlinks,omitempty"`
    MeshUplink          *ApStatMeshUplink                             `json:"mesh_uplink,omitempty"`
    Model               Optional[string]                              `json:"model"`
    ModifiedTime        *float64                                      `json:"modified_time,omitempty"`
    Mount               Optional[string]                              `json:"mount"`
    Name                Optional[string]                              `json:"name"`
    Notes               Optional[string]                              `json:"notes"`
    NumClients          Optional[int]                                 `json:"num_clients"`
    OrgId               Optional[uuid.UUID]                           `json:"org_id"`
    PortStat            Optional[map[string]ApStatsPortStat]          `json:"port_stat"`
    PowerBudget         Optional[int]                                 `json:"power_budget"`
    PowerConstrained    Optional[bool]                                `json:"power_constrained"`
    PowerOpmode         Optional[string]                              `json:"power_opmode"`
    PowerSrc            Optional[string]                              `json:"power_src"`
    RadioConfig         *ApStatsRadioConfig                           `json:"radio_config,omitempty"`
    RadioStat           *ApStatsRadioStat                             `json:"radio_stat,omitempty"`
    RxBps               Optional[float64]                             `json:"rx_bps"`
    RxBytes             Optional[int64]                               `json:"rx_bytes"`
    RxPkts              Optional[int]                                 `json:"rx_pkts"`
    Serial              Optional[string]                              `json:"serial"`
    SiteId              *uuid.UUID                                    `json:"site_id,omitempty"`
    Status              Optional[string]                              `json:"status"`
    SwitchRedundancy    *ApStatsSwitchRedundancy                      `json:"switch_redundancy,omitempty"`
    TxBps               Optional[float64]                             `json:"tx_bps"`
    TxBytes             Optional[float64]                             `json:"tx_bytes"`
    TxPkts              Optional[float64]                             `json:"tx_pkts"`
    Type                Optional[string]                              `json:"type"`
    Uptime              Optional[float64]                             `json:"uptime"`
    UsbStat             *ApStatsUsbStat                               `json:"usb_stat,omitempty"`
    Version             Optional[string]                              `json:"version"`
    X                   Optional[float64]                             `json:"x"`
    Y                   Optional[float64]                             `json:"y"`
    ApRedundancy        *SwitchStatsApRedundancy1                     `json:"ap_redundancy,omitempty"`
    ArpTableStats       *ArpTableStats                                `json:"arp_table_stats,omitempty"`
    Clients             []SwitchStatsClientItem                       `json:"clients,omitempty"`
    ClientsStats        *SwitchStatsClientsStats                      `json:"clients_stats,omitempty"`
    ConfigStatus        *string                                       `json:"config_status,omitempty"`
    CpuStat             *CpuStat                                      `json:"cpu_stat,omitempty"`
    DhcpdStat           *DhcpdStat                                    `json:"dhcpd_stat,omitempty"`
    FwVersionsOutofsync *bool                                         `json:"fw_versions_outofsync,omitempty"`
    HasPcap             *bool                                         `json:"has_pcap,omitempty"`
    Hostname            *string                                       `json:"hostname,omitempty"`
    IfStat              *IfStat                                       `json:"if_stat,omitempty"`
    MacTableStats       *MacTableStats                                `json:"mac_table_stats,omitempty"`
    MemoryStat          *MemoryStat                                   `json:"memory_stat,omitempty"`
    ModuleStat          []ModuleStatItem1                             `json:"module_stat,omitempty"`
    Ports               []DeviceStatsPort                             `json:"ports,omitempty"`
    RouteSummaryStats   *RouteSummaryStats                            `json:"route_summary_stats,omitempty"`
    ServiceStat         *ServiceStatProperty                          `json:"service_stat,omitempty"`
    VcMac               Optional[string]                              `json:"vc_mac"`
    VcSetupInfo         *SwitchStatsVcSetupInfo                       `json:"vc_setup_info,omitempty"`
    ClusterConfig       *ClusterConfigStats                           `json:"cluster_config,omitempty"`
    ClusterStat         *GatewayStatsCluster                          `json:"cluster_stat,omitempty"`
    ConductorName       *string                                       `json:"conductor_name,omitempty"`
    Cpu2Stat            *CpuStat                                      `json:"cpu2_stat,omitempty"`
    Dhcpd2Stat          map[string]DhcpdStatLan                       `json:"dhcpd2_stat,omitempty"`
    If2Stat             map[string]IfStatProperty                     `json:"if2_stat,omitempty"`
    Ip2Stat             *IpStat                                       `json:"ip2_stat,omitempty"`
    IsHa                Optional[bool]                                `json:"is_ha"`
    Memory2Stat         *MemoryStat                                   `json:"memory2_stat,omitempty"`
    Module2Stat         []ModuleStatItem                              `json:"module2_stat,omitempty"`
    NodeName            *string                                       `json:"node_name,omitempty"`
    RouterName          *string                                       `json:"router_name,omitempty"`
    Service2Stat        map[string]ServiceStatProperty                `json:"service2_stat,omitempty"`
    ServiceStatus       *GatewayStatsServiceStatus                    `json:"service_status,omitempty"`
    Spu2Stat            []GatewayStatsSpuItem                         `json:"spu2_stat,omitempty"`
    SpuStat             []GatewayStatsSpuItem                         `json:"spu_stat,omitempty"`
}
