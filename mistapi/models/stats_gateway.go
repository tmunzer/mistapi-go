package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsGateway represents a StatsGateway struct.
// Gateway statistics
type StatsGateway struct {
    ApRedundancy         *ApRedundancy                  `json:"ap_redundancy,omitempty"`
    ArpTableStats        *ArpTableStats                 `json:"arp_table_stats,omitempty"`
    CertExpiry           *int64                         `json:"cert_expiry,omitempty"`
    ClusterConfig        *StatsClusterConfig            `json:"cluster_config,omitempty"`
    ClusterStat          *StatsGatewayCluster           `json:"cluster_stat,omitempty"`
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
    Ports                []StatsDevicePort              `json:"ports,omitempty"`
    RouteSummaryStats    *RouteSummaryStats             `json:"route_summary_stats,omitempty"`
    // device name if configured
    RouterName           *string                        `json:"router_name,omitempty"`
    // serial
    Serial               *string                        `json:"serial,omitempty"`
    Service2Stat         map[string]ServiceStatProperty `json:"service2_stat,omitempty"`
    ServiceStat          map[string]ServiceStatProperty `json:"service_stat,omitempty"`
    ServiceStatus        *StatsGatewayServiceStatus     `json:"service_status,omitempty"`
    // serial
    SiteId               *uuid.UUID                     `json:"site_id,omitempty"`
    Spu2Stat             []StatsGatewaySpuItem          `json:"spu2_stat,omitempty"`
    SpuStat              []StatsGatewaySpuItem          `json:"spu_stat,omitempty"`
    Status               *string                        `json:"status,omitempty"`
    Type                 *string                        `json:"type,omitempty"`
    Uptime               *float64                       `json:"uptime,omitempty"`
    Version              *string                        `json:"version,omitempty"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsGateway.
// It customizes the JSON marshaling process for StatsGateway objects.
func (s StatsGateway) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsGateway object to a map representation for JSON marshaling.
func (s StatsGateway) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ApRedundancy != nil {
        structMap["ap_redundancy"] = s.ApRedundancy.toMap()
    }
    if s.ArpTableStats != nil {
        structMap["arp_table_stats"] = s.ArpTableStats.toMap()
    }
    if s.CertExpiry != nil {
        structMap["cert_expiry"] = s.CertExpiry
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
    if s.ConfigStatus != nil {
        structMap["config_status"] = s.ConfigStatus
    }
    if s.Cpu2Stat != nil {
        structMap["cpu2_stat"] = s.Cpu2Stat.toMap()
    }
    if s.CpuStat != nil {
        structMap["cpu_stat"] = s.CpuStat.toMap()
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
    if s.Dhcpd2Stat != nil {
        structMap["dhcpd2_stat"] = s.Dhcpd2Stat
    }
    if s.DhcpdStat != nil {
        structMap["dhcpd_stat"] = s.DhcpdStat
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
    if s.HasPcap.IsValueSet() {
        if s.HasPcap.Value() != nil {
            structMap["has_pcap"] = s.HasPcap.Value()
        } else {
            structMap["has_pcap"] = nil
        }
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.If2Stat != nil {
        structMap["if2_stat"] = s.If2Stat
    }
    if s.IfStat != nil {
        structMap["if_stat"] = s.IfStat
    }
    if s.Ip.IsValueSet() {
        if s.Ip.Value() != nil {
            structMap["ip"] = s.Ip.Value()
        } else {
            structMap["ip"] = nil
        }
    }
    if s.Ip2Stat != nil {
        structMap["ip2_stat"] = s.Ip2Stat.toMap()
    }
    if s.IpStat != nil {
        structMap["ip_stat"] = s.IpStat.toMap()
    }
    if s.IsHa.IsValueSet() {
        if s.IsHa.Value() != nil {
            structMap["is_ha"] = s.IsHa.Value()
        } else {
            structMap["is_ha"] = nil
        }
    }
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    }
    structMap["mac"] = s.Mac
    if s.MapId.IsValueSet() {
        if s.MapId.Value() != nil {
            structMap["map_id"] = s.MapId.Value()
        } else {
            structMap["map_id"] = nil
        }
    }
    if s.Memory2Stat != nil {
        structMap["memory2_stat"] = s.Memory2Stat.toMap()
    }
    if s.MemoryStat != nil {
        structMap["memory_stat"] = s.MemoryStat.toMap()
    }
    if s.Model != nil {
        structMap["model"] = s.Model
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.Module2Stat != nil {
        structMap["module2_stat"] = s.Module2Stat
    }
    if s.ModuleStat != nil {
        structMap["module_stat"] = s.ModuleStat
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.NodeName != nil {
        structMap["node_name"] = s.NodeName
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.Ports != nil {
        structMap["ports"] = s.Ports
    }
    if s.RouteSummaryStats != nil {
        structMap["route_summary_stats"] = s.RouteSummaryStats.toMap()
    }
    if s.RouterName != nil {
        structMap["router_name"] = s.RouterName
    }
    if s.Serial != nil {
        structMap["serial"] = s.Serial
    }
    if s.Service2Stat != nil {
        structMap["service2_stat"] = s.Service2Stat
    }
    if s.ServiceStat != nil {
        structMap["service_stat"] = s.ServiceStat
    }
    if s.ServiceStatus != nil {
        structMap["service_status"] = s.ServiceStatus.toMap()
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Spu2Stat != nil {
        structMap["spu2_stat"] = s.Spu2Stat
    }
    if s.SpuStat != nil {
        structMap["spu_stat"] = s.SpuStat
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.Version != nil {
        structMap["version"] = s.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsGateway.
// It customizes the JSON unmarshaling process for StatsGateway objects.
func (s *StatsGateway) UnmarshalJSON(input []byte) error {
    var temp tempStatsGateway
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
    
    s.AdditionalProperties = additionalProperties
    s.ApRedundancy = temp.ApRedundancy
    s.ArpTableStats = temp.ArpTableStats
    s.CertExpiry = temp.CertExpiry
    s.ClusterConfig = temp.ClusterConfig
    s.ClusterStat = temp.ClusterStat
    s.ConductorName = temp.ConductorName
    s.ConfigStatus = temp.ConfigStatus
    s.Cpu2Stat = temp.Cpu2Stat
    s.CpuStat = temp.CpuStat
    s.CreatedTime = temp.CreatedTime
    s.DeviceprofileId = temp.DeviceprofileId
    s.Dhcpd2Stat = temp.Dhcpd2Stat
    s.DhcpdStat = temp.DhcpdStat
    s.EvpntopoId = temp.EvpntopoId
    s.ExtIp = temp.ExtIp
    s.Fwupdate = temp.Fwupdate
    s.HasPcap = temp.HasPcap
    s.Hostname = temp.Hostname
    s.Id = temp.Id
    s.If2Stat = temp.If2Stat
    s.IfStat = temp.IfStat
    s.Ip = temp.Ip
    s.Ip2Stat = temp.Ip2Stat
    s.IpStat = temp.IpStat
    s.IsHa = temp.IsHa
    s.LastSeen = temp.LastSeen
    s.Mac = *temp.Mac
    s.MapId = temp.MapId
    s.Memory2Stat = temp.Memory2Stat
    s.MemoryStat = temp.MemoryStat
    s.Model = temp.Model
    s.ModifiedTime = temp.ModifiedTime
    s.Module2Stat = temp.Module2Stat
    s.ModuleStat = temp.ModuleStat
    s.Name = temp.Name
    s.NodeName = temp.NodeName
    s.OrgId = temp.OrgId
    s.Ports = temp.Ports
    s.RouteSummaryStats = temp.RouteSummaryStats
    s.RouterName = temp.RouterName
    s.Serial = temp.Serial
    s.Service2Stat = temp.Service2Stat
    s.ServiceStat = temp.ServiceStat
    s.ServiceStatus = temp.ServiceStatus
    s.SiteId = temp.SiteId
    s.Spu2Stat = temp.Spu2Stat
    s.SpuStat = temp.SpuStat
    s.Status = temp.Status
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    s.Version = temp.Version
    return nil
}

// tempStatsGateway is a temporary struct used for validating the fields of StatsGateway.
type tempStatsGateway  struct {
    ApRedundancy      *ApRedundancy                  `json:"ap_redundancy,omitempty"`
    ArpTableStats     *ArpTableStats                 `json:"arp_table_stats,omitempty"`
    CertExpiry        *int64                         `json:"cert_expiry,omitempty"`
    ClusterConfig     *StatsClusterConfig            `json:"cluster_config,omitempty"`
    ClusterStat       *StatsGatewayCluster           `json:"cluster_stat,omitempty"`
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
    Ports             []StatsDevicePort              `json:"ports,omitempty"`
    RouteSummaryStats *RouteSummaryStats             `json:"route_summary_stats,omitempty"`
    RouterName        *string                        `json:"router_name,omitempty"`
    Serial            *string                        `json:"serial,omitempty"`
    Service2Stat      map[string]ServiceStatProperty `json:"service2_stat,omitempty"`
    ServiceStat       map[string]ServiceStatProperty `json:"service_stat,omitempty"`
    ServiceStatus     *StatsGatewayServiceStatus     `json:"service_status,omitempty"`
    SiteId            *uuid.UUID                     `json:"site_id,omitempty"`
    Spu2Stat          []StatsGatewaySpuItem          `json:"spu2_stat,omitempty"`
    SpuStat           []StatsGatewaySpuItem          `json:"spu_stat,omitempty"`
    Status            *string                        `json:"status,omitempty"`
    Type              *string                        `json:"type,omitempty"`
    Uptime            *float64                       `json:"uptime,omitempty"`
    Version           *string                        `json:"version,omitempty"`
}

func (s *tempStatsGateway) validate() error {
    var errs []string
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `stats_gateway`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
