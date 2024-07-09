package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GatewayStats represents a GatewayStats struct.
// Gateway statistics
type GatewayStats struct {
    ApRedundancy         *ApRedundancy                       `json:"ap_redundancy,omitempty"`
    ClusterStat          map[string]GatewayStatsClusterStat  `json:"cluster_stat,omitempty"`
    Cpu2Stat             *CpuStat                            `json:"cpu2_stat,omitempty"`
    CpuStat              *CpuStat                            `json:"cpu_stat,omitempty"`
    // Property key is the network name
    Dhcpd2Stat           map[string]GatewayStatsDhcpdStatLan `json:"dhcpd2_stat,omitempty"`
    // Property key is the network name
    DhcpdStat            map[string]GatewayStatsDhcpdStatLan `json:"dhcpd_stat,omitempty"`
    // hostname reported by the device
    Hostname             *string                             `json:"hostname,omitempty"`
    // IP address
    Ip                   *string                             `json:"ip,omitempty"`
    Ip2Stat              *IpStat                             `json:"ip2_stat,omitempty"`
    IpStat               *IpStat                             `json:"ip_stat,omitempty"`
    // last seen timestamp
    LastSeen             *float64                            `json:"last_seen,omitempty"`
    // device mac
    Mac                  string                              `json:"mac"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    Memory2Stat          *MemoryStat                         `json:"memory2_stat,omitempty"`
    // memory usage stat (for virtual chassis, memory usage of master RE)
    MemoryStat           *MemoryStat                         `json:"memory_stat,omitempty"`
    // device model
    Model                *string                             `json:"model,omitempty"`
    Module2Stat          []ModuleStat                        `json:"module2_stat,omitempty"`
    ModuleStat           []ModuleStat                        `json:"module_stat,omitempty"`
    // device name if configured
    Name                 *string                             `json:"name,omitempty"`
    // serial
    Serial               *string                             `json:"serial,omitempty"`
    Spu2Stat             *GatewayStatsSpuStat                `json:"spu2_stat,omitempty"`
    SpuStat              *GatewayStatsSpuStat                `json:"spu_stat,omitempty"`
    Status               *string                             `json:"status,omitempty"`
    Type                 *string                             `json:"type,omitempty"`
    Uptime               *float64                            `json:"uptime,omitempty"`
    Version              *string                             `json:"version,omitempty"`
    AdditionalProperties map[string]any                      `json:"_"`
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
    if g.ClusterStat != nil {
        structMap["cluster_stat"] = g.ClusterStat
    }
    if g.Cpu2Stat != nil {
        structMap["cpu2_stat"] = g.Cpu2Stat.toMap()
    }
    if g.CpuStat != nil {
        structMap["cpu_stat"] = g.CpuStat.toMap()
    }
    if g.Dhcpd2Stat != nil {
        structMap["dhcpd2_stat"] = g.Dhcpd2Stat
    }
    if g.DhcpdStat != nil {
        structMap["dhcpd_stat"] = g.DhcpdStat
    }
    if g.Hostname != nil {
        structMap["hostname"] = g.Hostname
    }
    if g.Ip != nil {
        structMap["ip"] = g.Ip
    }
    if g.Ip2Stat != nil {
        structMap["ip2_stat"] = g.Ip2Stat.toMap()
    }
    if g.IpStat != nil {
        structMap["ip_stat"] = g.IpStat.toMap()
    }
    if g.LastSeen != nil {
        structMap["last_seen"] = g.LastSeen
    }
    structMap["mac"] = g.Mac
    if g.Memory2Stat != nil {
        structMap["memory2_stat"] = g.Memory2Stat.toMap()
    }
    if g.MemoryStat != nil {
        structMap["memory_stat"] = g.MemoryStat.toMap()
    }
    if g.Model != nil {
        structMap["model"] = g.Model
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
    if g.Serial != nil {
        structMap["serial"] = g.Serial
    }
    if g.Spu2Stat != nil {
        structMap["spu2_stat"] = g.Spu2Stat.toMap()
    }
    if g.SpuStat != nil {
        structMap["spu_stat"] = g.SpuStat.toMap()
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
    var temp gatewayStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_redundancy", "cluster_stat", "cpu2_stat", "cpu_stat", "dhcpd2_stat", "dhcpd_stat", "hostname", "ip", "ip2_stat", "ip_stat", "last_seen", "mac", "memory2_stat", "memory_stat", "model", "module2_stat", "module_stat", "name", "serial", "spu2_stat", "spu_stat", "status", "type", "uptime", "version")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.ApRedundancy = temp.ApRedundancy
    g.ClusterStat = temp.ClusterStat
    g.Cpu2Stat = temp.Cpu2Stat
    g.CpuStat = temp.CpuStat
    g.Dhcpd2Stat = temp.Dhcpd2Stat
    g.DhcpdStat = temp.DhcpdStat
    g.Hostname = temp.Hostname
    g.Ip = temp.Ip
    g.Ip2Stat = temp.Ip2Stat
    g.IpStat = temp.IpStat
    g.LastSeen = temp.LastSeen
    g.Mac = *temp.Mac
    g.Memory2Stat = temp.Memory2Stat
    g.MemoryStat = temp.MemoryStat
    g.Model = temp.Model
    g.Module2Stat = temp.Module2Stat
    g.ModuleStat = temp.ModuleStat
    g.Name = temp.Name
    g.Serial = temp.Serial
    g.Spu2Stat = temp.Spu2Stat
    g.SpuStat = temp.SpuStat
    g.Status = temp.Status
    g.Type = temp.Type
    g.Uptime = temp.Uptime
    g.Version = temp.Version
    return nil
}

// gatewayStats is a temporary struct used for validating the fields of GatewayStats.
type gatewayStats  struct {
    ApRedundancy *ApRedundancy                       `json:"ap_redundancy,omitempty"`
    ClusterStat  map[string]GatewayStatsClusterStat  `json:"cluster_stat,omitempty"`
    Cpu2Stat     *CpuStat                            `json:"cpu2_stat,omitempty"`
    CpuStat      *CpuStat                            `json:"cpu_stat,omitempty"`
    Dhcpd2Stat   map[string]GatewayStatsDhcpdStatLan `json:"dhcpd2_stat,omitempty"`
    DhcpdStat    map[string]GatewayStatsDhcpdStatLan `json:"dhcpd_stat,omitempty"`
    Hostname     *string                             `json:"hostname,omitempty"`
    Ip           *string                             `json:"ip,omitempty"`
    Ip2Stat      *IpStat                             `json:"ip2_stat,omitempty"`
    IpStat       *IpStat                             `json:"ip_stat,omitempty"`
    LastSeen     *float64                            `json:"last_seen,omitempty"`
    Mac          *string                             `json:"mac"`
    Memory2Stat  *MemoryStat                         `json:"memory2_stat,omitempty"`
    MemoryStat   *MemoryStat                         `json:"memory_stat,omitempty"`
    Model        *string                             `json:"model,omitempty"`
    Module2Stat  []ModuleStat                        `json:"module2_stat,omitempty"`
    ModuleStat   []ModuleStat                        `json:"module_stat,omitempty"`
    Name         *string                             `json:"name,omitempty"`
    Serial       *string                             `json:"serial,omitempty"`
    Spu2Stat     *GatewayStatsSpuStat                `json:"spu2_stat,omitempty"`
    SpuStat      *GatewayStatsSpuStat                `json:"spu_stat,omitempty"`
    Status       *string                             `json:"status,omitempty"`
    Type         *string                             `json:"type,omitempty"`
    Uptime       *float64                            `json:"uptime,omitempty"`
    Version      *string                             `json:"version,omitempty"`
}

func (g *gatewayStats) validate() error {
    var errs []string
    if g.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Gateway_Stats`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
