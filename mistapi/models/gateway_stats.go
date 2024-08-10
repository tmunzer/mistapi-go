package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// GatewayStats represents a GatewayStats struct.
// Gateway statistics
type GatewayStats struct {
    ApRedundancy         *ApRedundancy                  `json:"ap_redundancy,omitempty"`
    ArpTableStats        *ArpTableStats                 `json:"arp_table_stats,omitempty"`
    CertExpiry           *int64                         `json:"cert_expiry,omitempty"`
    ClusterConfig        *ClusterConfigStats            `json:"cluster_config,omitempty"`
    ClusterStat          *GatewayStatsCluster           `json:"cluster_stat,omitempty"`
    ConductorName        *string                        `json:"conductor_name,omitempty"`
    ConfigStatus         *string                        `json:"config_status,omitempty"`
    Cpu2Stat             *CpuStat                       `json:"cpu2_stat,omitempty"`
    CpuStat              *CpuStat                       `json:"cpu_stat,omitempty"`
    CreatedTime          *float64                       `json:"created_time,omitempty"`
    DeviceprofileId      Optional[uuid.UUID]            `json:"deviceprofile_id"`
    // Property key is the network name
    Dhcpd2Stat           map[string]DhcpdStatLan        `json:"dhcpd2_stat,omitempty"`
    // Property key is the network name
    DhcpdStat            map[string]DhcpdStatLan        `json:"dhcpd_stat,omitempty"`
    EvpntopoId           Optional[uuid.UUID]            `json:"evpntopo_id"`
    // IP address
    ExtIp                Optional[string]               `json:"ext_ip"`
    Fwupdate             *FwupdateStat                  `json:"fwupdate,omitempty"`
    HasPcap              Optional[bool]                 `json:"has_pcap"`
    // hostname reported by the device
    Hostname             *string                        `json:"hostname,omitempty"`
    // serial
    Id                   *uuid.UUID                     `json:"id,omitempty"`
    // Property key is the interface name
    If2Stat              map[string]IfStatProperty      `json:"if2_stat,omitempty"`
    // Property key is the interface name
    IfStat               map[string]IfStatProperty      `json:"if_stat,omitempty"`
    // IP address
    Ip                   Optional[string]               `json:"ip"`
    Ip2Stat              *IpStat                        `json:"ip2_stat,omitempty"`
    IpStat               *IpStat                        `json:"ip_stat,omitempty"`
    IsHa                 Optional[bool]                 `json:"is_ha"`
    // last seen timestamp
    LastSeen             *float64                       `json:"last_seen,omitempty"`
    // device mac
    Mac                  string                         `json:"mac"`
    // serial
    MapId                Optional[uuid.UUID]            `json:"map_id"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    Memory2Stat          *MemoryStat                    `json:"memory2_stat,omitempty"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    MemoryStat           *MemoryStat                    `json:"memory_stat,omitempty"`
    // device model
    Model                *string                        `json:"model,omitempty"`
    ModifiedTime         *float64                       `json:"modified_time,omitempty"`
    Module2Stat          []ModuleStatItem               `json:"module2_stat,omitempty"`
    ModuleStat           []ModuleStatItem               `json:"module_stat,omitempty"`
    // device name if configured
    Name                 *string                        `json:"name,omitempty"`
    NodeName             *string                        `json:"node_name,omitempty"`
    // serial
    OrgId                *uuid.UUID                     `json:"org_id,omitempty"`
    Ports                []DeviceStatsPort              `json:"ports,omitempty"`
    RouteSummaryStats    *RouteSummaryStats             `json:"route_summary_stats,omitempty"`
    // device name if configured
    RouterName           *string                        `json:"router_name,omitempty"`
    // serial
    Serial               *string                        `json:"serial,omitempty"`
    Service2Stat         map[string]ServiceStatProperty `json:"service2_stat,omitempty"`
    ServiceStat          map[string]ServiceStatProperty `json:"service_stat,omitempty"`
    ServiceStatus        *GatewayStatsServiceStatus     `json:"service_status,omitempty"`
    // serial
    SiteId               *uuid.UUID                     `json:"site_id,omitempty"`
    Spu2Stat             []GatewayStatsSpuItem          `json:"spu2_stat,omitempty"`
    SpuStat              []GatewayStatsSpuItem          `json:"spu_stat,omitempty"`
    Status               *string                        `json:"status,omitempty"`
    Type                 *string                        `json:"type,omitempty"`
    Uptime               *float64                       `json:"uptime,omitempty"`
    Version              *string                        `json:"version,omitempty"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GatewayStats.
// It customizes the JSON marshaling process for GatewayStats objects.
func (g GatewayStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GatewayStats object to a map representation for JSON marshaling.
func (g GatewayStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.ApRedundancy != nil {
        structMap["ap_redundancy"] = g.ApRedundancy.toMap()
    }
    if g.ArpTableStats != nil {
        structMap["arp_table_stats"] = g.ArpTableStats.toMap()
    }
    if g.CertExpiry != nil {
        structMap["cert_expiry"] = g.CertExpiry
    }
    if g.ClusterConfig != nil {
        structMap["cluster_config"] = g.ClusterConfig.toMap()
    }
    if g.ClusterStat != nil {
        structMap["cluster_stat"] = g.ClusterStat.toMap()
    }
    if g.ConductorName != nil {
        structMap["conductor_name"] = g.ConductorName
    }
    if g.ConfigStatus != nil {
        structMap["config_status"] = g.ConfigStatus
    }
    if g.Cpu2Stat != nil {
        structMap["cpu2_stat"] = g.Cpu2Stat.toMap()
    }
    if g.CpuStat != nil {
        structMap["cpu_stat"] = g.CpuStat.toMap()
    }
    if g.CreatedTime != nil {
        structMap["created_time"] = g.CreatedTime
    }
    if g.DeviceprofileId.IsValueSet() {
        if g.DeviceprofileId.Value() != nil {
            structMap["deviceprofile_id"] = g.DeviceprofileId.Value()
        } else {
            structMap["deviceprofile_id"] = nil
        }
    }
    if g.Dhcpd2Stat != nil {
        structMap["dhcpd2_stat"] = g.Dhcpd2Stat
    }
    if g.DhcpdStat != nil {
        structMap["dhcpd_stat"] = g.DhcpdStat
    }
    if g.EvpntopoId.IsValueSet() {
        if g.EvpntopoId.Value() != nil {
            structMap["evpntopo_id"] = g.EvpntopoId.Value()
        } else {
            structMap["evpntopo_id"] = nil
        }
    }
    if g.ExtIp.IsValueSet() {
        if g.ExtIp.Value() != nil {
            structMap["ext_ip"] = g.ExtIp.Value()
        } else {
            structMap["ext_ip"] = nil
        }
    }
    if g.Fwupdate != nil {
        structMap["fwupdate"] = g.Fwupdate.toMap()
    }
    if g.HasPcap.IsValueSet() {
        if g.HasPcap.Value() != nil {
            structMap["has_pcap"] = g.HasPcap.Value()
        } else {
            structMap["has_pcap"] = nil
        }
    }
    if g.Hostname != nil {
        structMap["hostname"] = g.Hostname
    }
    if g.Id != nil {
        structMap["id"] = g.Id
    }
    if g.If2Stat != nil {
        structMap["if2_stat"] = g.If2Stat
    }
    if g.IfStat != nil {
        structMap["if_stat"] = g.IfStat
    }
    if g.Ip.IsValueSet() {
        if g.Ip.Value() != nil {
            structMap["ip"] = g.Ip.Value()
        } else {
            structMap["ip"] = nil
        }
    }
    if g.Ip2Stat != nil {
        structMap["ip2_stat"] = g.Ip2Stat.toMap()
    }
    if g.IpStat != nil {
        structMap["ip_stat"] = g.IpStat.toMap()
    }
    if g.IsHa.IsValueSet() {
        if g.IsHa.Value() != nil {
            structMap["is_ha"] = g.IsHa.Value()
        } else {
            structMap["is_ha"] = nil
        }
    }
    if g.LastSeen != nil {
        structMap["last_seen"] = g.LastSeen
    }
    structMap["mac"] = g.Mac
    if g.MapId.IsValueSet() {
        if g.MapId.Value() != nil {
            structMap["map_id"] = g.MapId.Value()
        } else {
            structMap["map_id"] = nil
        }
    }
    if g.Memory2Stat != nil {
        structMap["memory2_stat"] = g.Memory2Stat.toMap()
    }
    if g.MemoryStat != nil {
        structMap["memory_stat"] = g.MemoryStat.toMap()
    }
    if g.Model != nil {
        structMap["model"] = g.Model
    }
    if g.ModifiedTime != nil {
        structMap["modified_time"] = g.ModifiedTime
    }
    if g.Module2Stat != nil {
        structMap["module2_stat"] = g.Module2Stat
    }
    if g.ModuleStat != nil {
        structMap["module_stat"] = g.ModuleStat
    }
    if g.Name != nil {
        structMap["name"] = g.Name
    }
    if g.NodeName != nil {
        structMap["node_name"] = g.NodeName
    }
    if g.OrgId != nil {
        structMap["org_id"] = g.OrgId
    }
    if g.Ports != nil {
        structMap["ports"] = g.Ports
    }
    if g.RouteSummaryStats != nil {
        structMap["route_summary_stats"] = g.RouteSummaryStats.toMap()
    }
    if g.RouterName != nil {
        structMap["router_name"] = g.RouterName
    }
    if g.Serial != nil {
        structMap["serial"] = g.Serial
    }
    if g.Service2Stat != nil {
        structMap["service2_stat"] = g.Service2Stat
    }
    if g.ServiceStat != nil {
        structMap["service_stat"] = g.ServiceStat
    }
    if g.ServiceStatus != nil {
        structMap["service_status"] = g.ServiceStatus.toMap()
    }
    if g.SiteId != nil {
        structMap["site_id"] = g.SiteId
    }
    if g.Spu2Stat != nil {
        structMap["spu2_stat"] = g.Spu2Stat
    }
    if g.SpuStat != nil {
        structMap["spu_stat"] = g.SpuStat
    }
    if g.Status != nil {
        structMap["status"] = g.Status
    }
    if g.Type != nil {
        structMap["type"] = g.Type
    }
    if g.Uptime != nil {
        structMap["uptime"] = g.Uptime
    }
    if g.Version != nil {
        structMap["version"] = g.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GatewayStats.
// It customizes the JSON unmarshaling process for GatewayStats objects.
func (g *GatewayStats) UnmarshalJSON(input []byte) error {
    var temp tempGatewayStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_redundancy", "arp_table_stats", "cert_expiry", "cluster_config", "cluster_stat", "conductor_name", "config_status", "cpu2_stat", "cpu_stat", "created_time", "deviceprofile_id", "dhcpd2_stat", "dhcpd_stat", "evpntopo_id", "ext_ip", "fwupdate", "has_pcap", "hostname", "id", "if2_stat", "if_stat", "ip", "ip2_stat", "ip_stat", "is_ha", "last_seen", "mac", "map_id", "memory2_stat", "memory_stat", "model", "modified_time", "module2_stat", "module_stat", "name", "node_name", "org_id", "ports", "route_summary_stats", "router_name", "serial", "service2_stat", "service_stat", "service_status", "site_id", "spu2_stat", "spu_stat", "status", "type", "uptime", "version")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.ApRedundancy = temp.ApRedundancy
    g.ArpTableStats = temp.ArpTableStats
    g.CertExpiry = temp.CertExpiry
    g.ClusterConfig = temp.ClusterConfig
    g.ClusterStat = temp.ClusterStat
    g.ConductorName = temp.ConductorName
    g.ConfigStatus = temp.ConfigStatus
    g.Cpu2Stat = temp.Cpu2Stat
    g.CpuStat = temp.CpuStat
    g.CreatedTime = temp.CreatedTime
    g.DeviceprofileId = temp.DeviceprofileId
    g.Dhcpd2Stat = temp.Dhcpd2Stat
    g.DhcpdStat = temp.DhcpdStat
    g.EvpntopoId = temp.EvpntopoId
    g.ExtIp = temp.ExtIp
    g.Fwupdate = temp.Fwupdate
    g.HasPcap = temp.HasPcap
    g.Hostname = temp.Hostname
    g.Id = temp.Id
    g.If2Stat = temp.If2Stat
    g.IfStat = temp.IfStat
    g.Ip = temp.Ip
    g.Ip2Stat = temp.Ip2Stat
    g.IpStat = temp.IpStat
    g.IsHa = temp.IsHa
    g.LastSeen = temp.LastSeen
    g.Mac = *temp.Mac
    g.MapId = temp.MapId
    g.Memory2Stat = temp.Memory2Stat
    g.MemoryStat = temp.MemoryStat
    g.Model = temp.Model
    g.ModifiedTime = temp.ModifiedTime
    g.Module2Stat = temp.Module2Stat
    g.ModuleStat = temp.ModuleStat
    g.Name = temp.Name
    g.NodeName = temp.NodeName
    g.OrgId = temp.OrgId
    g.Ports = temp.Ports
    g.RouteSummaryStats = temp.RouteSummaryStats
    g.RouterName = temp.RouterName
    g.Serial = temp.Serial
    g.Service2Stat = temp.Service2Stat
    g.ServiceStat = temp.ServiceStat
    g.ServiceStatus = temp.ServiceStatus
    g.SiteId = temp.SiteId
    g.Spu2Stat = temp.Spu2Stat
    g.SpuStat = temp.SpuStat
    g.Status = temp.Status
    g.Type = temp.Type
    g.Uptime = temp.Uptime
    g.Version = temp.Version
    return nil
}

// tempGatewayStats is a temporary struct used for validating the fields of GatewayStats.
type tempGatewayStats  struct {
    ApRedundancy      *ApRedundancy                  `json:"ap_redundancy,omitempty"`
    ArpTableStats     *ArpTableStats                 `json:"arp_table_stats,omitempty"`
    CertExpiry        *int64                         `json:"cert_expiry,omitempty"`
    ClusterConfig     *ClusterConfigStats            `json:"cluster_config,omitempty"`
    ClusterStat       *GatewayStatsCluster           `json:"cluster_stat,omitempty"`
    ConductorName     *string                        `json:"conductor_name,omitempty"`
    ConfigStatus      *string                        `json:"config_status,omitempty"`
    Cpu2Stat          *CpuStat                       `json:"cpu2_stat,omitempty"`
    CpuStat           *CpuStat                       `json:"cpu_stat,omitempty"`
    CreatedTime       *float64                       `json:"created_time,omitempty"`
    DeviceprofileId   Optional[uuid.UUID]            `json:"deviceprofile_id"`
    Dhcpd2Stat        map[string]DhcpdStatLan        `json:"dhcpd2_stat,omitempty"`
    DhcpdStat         map[string]DhcpdStatLan        `json:"dhcpd_stat,omitempty"`
    EvpntopoId        Optional[uuid.UUID]            `json:"evpntopo_id"`
    ExtIp             Optional[string]               `json:"ext_ip"`
    Fwupdate          *FwupdateStat                  `json:"fwupdate,omitempty"`
    HasPcap           Optional[bool]                 `json:"has_pcap"`
    Hostname          *string                        `json:"hostname,omitempty"`
    Id                *uuid.UUID                     `json:"id,omitempty"`
    If2Stat           map[string]IfStatProperty      `json:"if2_stat,omitempty"`
    IfStat            map[string]IfStatProperty      `json:"if_stat,omitempty"`
    Ip                Optional[string]               `json:"ip"`
    Ip2Stat           *IpStat                        `json:"ip2_stat,omitempty"`
    IpStat            *IpStat                        `json:"ip_stat,omitempty"`
    IsHa              Optional[bool]                 `json:"is_ha"`
    LastSeen          *float64                       `json:"last_seen,omitempty"`
    Mac               *string                        `json:"mac"`
    MapId             Optional[uuid.UUID]            `json:"map_id"`
    Memory2Stat       *MemoryStat                    `json:"memory2_stat,omitempty"`
    MemoryStat        *MemoryStat                    `json:"memory_stat,omitempty"`
    Model             *string                        `json:"model,omitempty"`
    ModifiedTime      *float64                       `json:"modified_time,omitempty"`
    Module2Stat       []ModuleStatItem               `json:"module2_stat,omitempty"`
    ModuleStat        []ModuleStatItem               `json:"module_stat,omitempty"`
    Name              *string                        `json:"name,omitempty"`
    NodeName          *string                        `json:"node_name,omitempty"`
    OrgId             *uuid.UUID                     `json:"org_id,omitempty"`
    Ports             []DeviceStatsPort              `json:"ports,omitempty"`
    RouteSummaryStats *RouteSummaryStats             `json:"route_summary_stats,omitempty"`
    RouterName        *string                        `json:"router_name,omitempty"`
    Serial            *string                        `json:"serial,omitempty"`
    Service2Stat      map[string]ServiceStatProperty `json:"service2_stat,omitempty"`
    ServiceStat       map[string]ServiceStatProperty `json:"service_stat,omitempty"`
    ServiceStatus     *GatewayStatsServiceStatus     `json:"service_status,omitempty"`
    SiteId            *uuid.UUID                     `json:"site_id,omitempty"`
    Spu2Stat          []GatewayStatsSpuItem          `json:"spu2_stat,omitempty"`
    SpuStat           []GatewayStatsSpuItem          `json:"spu_stat,omitempty"`
    Status            *string                        `json:"status,omitempty"`
    Type              *string                        `json:"type,omitempty"`
    Uptime            *float64                       `json:"uptime,omitempty"`
    Version           *string                        `json:"version,omitempty"`
}

func (g *tempGatewayStats) validate() error {
    var errs []string
    if g.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `gateway_stats`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
