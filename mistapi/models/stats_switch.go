// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsSwitch represents a StatsSwitch struct.
// Switch statistics
type StatsSwitch struct {
    ApRedundancy         *StatsSwitchApRedundancy       `json:"ap_redundancy,omitempty"`
    ArpTableStats        *ArpTableStats                 `json:"arp_table_stats,omitempty"`
    CertExpiry           *int64                         `json:"cert_expiry,omitempty"`
    Clients              []StatsSwitchClientItem        `json:"clients,omitempty"`
    ClientsStats         *StatsSwitchClientsStats       `json:"clients_stats,omitempty"`
    ConfigStatus         *string                        `json:"config_status,omitempty"`
    CpuStat              *CpuStat                       `json:"cpu_stat,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64                       `json:"created_time,omitempty"`
    DeviceprofileId      Optional[uuid.UUID]            `json:"deviceprofile_id"`
    // Property key is the network name
    DhcpdStat            map[string]DhcpdStatLan        `json:"dhcpd_stat,omitempty"`
    EvpntopoId           Optional[uuid.UUID]            `json:"evpntopo_id"`
    FwVersionsOutofsync  *bool                          `json:"fw_versions_outofsync,omitempty"`
    Fwupdate             *FwupdateStat                  `json:"fwupdate,omitempty"`
    // Whether the switch supports packet capture
    HasPcap              *bool                          `json:"has_pcap,omitempty"`
    // Hostname reported by the device
    Hostname             *string                        `json:"hostname,omitempty"`
    // Device hardware revision number
    HwRev                *string                        `json:"hw_rev,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID                     `json:"id,omitempty"`
    // Property key is the interface name
    IfStat               map[string]IfStatProperty      `json:"if_stat,omitempty"`
    Ip                   *string                        `json:"ip,omitempty"`
    IpStat               *IpStat                        `json:"ip_stat,omitempty"`
    // Last seen timestamp
    LastSeen             Optional[float64]              `json:"last_seen"`
    // Last trouble code of switch
    LastTrouble          *LastTrouble                   `json:"last_trouble,omitempty"`
    Mac                  *string                        `json:"mac,omitempty"`
    MacTableStats        *MacTableStats                 `json:"mac_table_stats,omitempty"`
    MapId                Optional[uuid.UUID]            `json:"map_id"`
    // Memory usage stat (for virtual chassis, memory usage of master RE)
    MemoryStat           *MemoryStat                    `json:"memory_stat,omitempty"`
    Model                *string                        `json:"model,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                       `json:"modified_time,omitempty"`
    ModuleStat           []StatsSwitchModuleStatItem    `json:"module_stat,omitempty"`
    // Device name if configured
    Name                 *string                        `json:"name,omitempty"`
    OrgId                *uuid.UUID                     `json:"org_id,omitempty"`
    Ports                []StatsSwitchPort              `json:"ports,omitempty"`
    RouteSummaryStats    *RouteSummaryStats             `json:"route_summary_stats,omitempty"`
    Serial               *string                        `json:"serial,omitempty"`
    ServiceStat          map[string]ServiceStatProperty `json:"service_stat,omitempty"`
    SiteId               *uuid.UUID                     `json:"site_id,omitempty"`
    Status               *string                        `json:"status,omitempty"`
    // Device Type. enum: `switch`
    Type                 string                         `json:"type"`
    Uptime               Optional[float64]              `json:"uptime"`
    VcMac                Optional[string]               `json:"vc_mac"`
    VcSetupInfo          *StatsSwitchVcSetupInfo        `json:"vc_setup_info,omitempty"`
    Version              Optional[string]               `json:"version"`
    AdditionalProperties map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSwitch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSwitch) String() string {
    return fmt.Sprintf(
    	"StatsSwitch[ApRedundancy=%v, ArpTableStats=%v, CertExpiry=%v, Clients=%v, ClientsStats=%v, ConfigStatus=%v, CpuStat=%v, CreatedTime=%v, DeviceprofileId=%v, DhcpdStat=%v, EvpntopoId=%v, FwVersionsOutofsync=%v, Fwupdate=%v, HasPcap=%v, Hostname=%v, HwRev=%v, Id=%v, IfStat=%v, Ip=%v, IpStat=%v, LastSeen=%v, LastTrouble=%v, Mac=%v, MacTableStats=%v, MapId=%v, MemoryStat=%v, Model=%v, ModifiedTime=%v, ModuleStat=%v, Name=%v, OrgId=%v, Ports=%v, RouteSummaryStats=%v, Serial=%v, ServiceStat=%v, SiteId=%v, Status=%v, Type=%v, Uptime=%v, VcMac=%v, VcSetupInfo=%v, Version=%v, AdditionalProperties=%v]",
    	s.ApRedundancy, s.ArpTableStats, s.CertExpiry, s.Clients, s.ClientsStats, s.ConfigStatus, s.CpuStat, s.CreatedTime, s.DeviceprofileId, s.DhcpdStat, s.EvpntopoId, s.FwVersionsOutofsync, s.Fwupdate, s.HasPcap, s.Hostname, s.HwRev, s.Id, s.IfStat, s.Ip, s.IpStat, s.LastSeen, s.LastTrouble, s.Mac, s.MacTableStats, s.MapId, s.MemoryStat, s.Model, s.ModifiedTime, s.ModuleStat, s.Name, s.OrgId, s.Ports, s.RouteSummaryStats, s.Serial, s.ServiceStat, s.SiteId, s.Status, s.Type, s.Uptime, s.VcMac, s.VcSetupInfo, s.Version, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSwitch.
// It customizes the JSON marshaling process for StatsSwitch objects.
func (s StatsSwitch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ap_redundancy", "arp_table_stats", "cert_expiry", "clients", "clients_stats", "config_status", "cpu_stat", "created_time", "deviceprofile_id", "dhcpd_stat", "evpntopo_id", "fw_versions_outofsync", "fwupdate", "has_pcap", "hostname", "hw_rev", "id", "if_stat", "ip", "ip_stat", "last_seen", "last_trouble", "mac", "mac_table_stats", "map_id", "memory_stat", "model", "modified_time", "module_stat", "name", "org_id", "ports", "route_summary_stats", "serial", "service_stat", "site_id", "status", "type", "uptime", "vc_mac", "vc_setup_info", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSwitch object to a map representation for JSON marshaling.
func (s StatsSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ApRedundancy != nil {
        structMap["ap_redundancy"] = s.ApRedundancy.toMap()
    }
    if s.ArpTableStats != nil {
        structMap["arp_table_stats"] = s.ArpTableStats.toMap()
    }
    if s.CertExpiry != nil {
        structMap["cert_expiry"] = s.CertExpiry
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
    if s.FwVersionsOutofsync != nil {
        structMap["fw_versions_outofsync"] = s.FwVersionsOutofsync
    }
    if s.Fwupdate != nil {
        structMap["fwupdate"] = s.Fwupdate.toMap()
    }
    if s.HasPcap != nil {
        structMap["has_pcap"] = s.HasPcap
    }
    if s.Hostname != nil {
        structMap["hostname"] = s.Hostname
    }
    if s.HwRev != nil {
        structMap["hw_rev"] = s.HwRev
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.IfStat != nil {
        structMap["if_stat"] = s.IfStat
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    if s.IpStat != nil {
        structMap["ip_stat"] = s.IpStat.toMap()
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
    if s.Mac != nil {
        structMap["mac"] = s.Mac
    }
    if s.MacTableStats != nil {
        structMap["mac_table_stats"] = s.MacTableStats.toMap()
    }
    if s.MapId.IsValueSet() {
        if s.MapId.Value() != nil {
            structMap["map_id"] = s.MapId.Value()
        } else {
            structMap["map_id"] = nil
        }
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
    if s.ModuleStat != nil {
        structMap["module_stat"] = s.ModuleStat
    }
    if s.Name != nil {
        structMap["name"] = s.Name
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
    if s.Serial != nil {
        structMap["serial"] = s.Serial
    }
    if s.ServiceStat != nil {
        structMap["service_stat"] = s.ServiceStat
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    structMap["type"] = s.Type
    if s.Uptime.IsValueSet() {
        if s.Uptime.Value() != nil {
            structMap["uptime"] = s.Uptime.Value()
        } else {
            structMap["uptime"] = nil
        }
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
    if s.Version.IsValueSet() {
        if s.Version.Value() != nil {
            structMap["version"] = s.Version.Value()
        } else {
            structMap["version"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSwitch.
// It customizes the JSON unmarshaling process for StatsSwitch objects.
func (s *StatsSwitch) UnmarshalJSON(input []byte) error {
    var temp tempStatsSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_redundancy", "arp_table_stats", "cert_expiry", "clients", "clients_stats", "config_status", "cpu_stat", "created_time", "deviceprofile_id", "dhcpd_stat", "evpntopo_id", "fw_versions_outofsync", "fwupdate", "has_pcap", "hostname", "hw_rev", "id", "if_stat", "ip", "ip_stat", "last_seen", "last_trouble", "mac", "mac_table_stats", "map_id", "memory_stat", "model", "modified_time", "module_stat", "name", "org_id", "ports", "route_summary_stats", "serial", "service_stat", "site_id", "status", "type", "uptime", "vc_mac", "vc_setup_info", "version")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ApRedundancy = temp.ApRedundancy
    s.ArpTableStats = temp.ArpTableStats
    s.CertExpiry = temp.CertExpiry
    s.Clients = temp.Clients
    s.ClientsStats = temp.ClientsStats
    s.ConfigStatus = temp.ConfigStatus
    s.CpuStat = temp.CpuStat
    s.CreatedTime = temp.CreatedTime
    s.DeviceprofileId = temp.DeviceprofileId
    s.DhcpdStat = temp.DhcpdStat
    s.EvpntopoId = temp.EvpntopoId
    s.FwVersionsOutofsync = temp.FwVersionsOutofsync
    s.Fwupdate = temp.Fwupdate
    s.HasPcap = temp.HasPcap
    s.Hostname = temp.Hostname
    s.HwRev = temp.HwRev
    s.Id = temp.Id
    s.IfStat = temp.IfStat
    s.Ip = temp.Ip
    s.IpStat = temp.IpStat
    s.LastSeen = temp.LastSeen
    s.LastTrouble = temp.LastTrouble
    s.Mac = temp.Mac
    s.MacTableStats = temp.MacTableStats
    s.MapId = temp.MapId
    s.MemoryStat = temp.MemoryStat
    s.Model = temp.Model
    s.ModifiedTime = temp.ModifiedTime
    s.ModuleStat = temp.ModuleStat
    s.Name = temp.Name
    s.OrgId = temp.OrgId
    s.Ports = temp.Ports
    s.RouteSummaryStats = temp.RouteSummaryStats
    s.Serial = temp.Serial
    s.ServiceStat = temp.ServiceStat
    s.SiteId = temp.SiteId
    s.Status = temp.Status
    s.Type = *temp.Type
    s.Uptime = temp.Uptime
    s.VcMac = temp.VcMac
    s.VcSetupInfo = temp.VcSetupInfo
    s.Version = temp.Version
    return nil
}

// tempStatsSwitch is a temporary struct used for validating the fields of StatsSwitch.
type tempStatsSwitch  struct {
    ApRedundancy        *StatsSwitchApRedundancy       `json:"ap_redundancy,omitempty"`
    ArpTableStats       *ArpTableStats                 `json:"arp_table_stats,omitempty"`
    CertExpiry          *int64                         `json:"cert_expiry,omitempty"`
    Clients             []StatsSwitchClientItem        `json:"clients,omitempty"`
    ClientsStats        *StatsSwitchClientsStats       `json:"clients_stats,omitempty"`
    ConfigStatus        *string                        `json:"config_status,omitempty"`
    CpuStat             *CpuStat                       `json:"cpu_stat,omitempty"`
    CreatedTime         *float64                       `json:"created_time,omitempty"`
    DeviceprofileId     Optional[uuid.UUID]            `json:"deviceprofile_id"`
    DhcpdStat           map[string]DhcpdStatLan        `json:"dhcpd_stat,omitempty"`
    EvpntopoId          Optional[uuid.UUID]            `json:"evpntopo_id"`
    FwVersionsOutofsync *bool                          `json:"fw_versions_outofsync,omitempty"`
    Fwupdate            *FwupdateStat                  `json:"fwupdate,omitempty"`
    HasPcap             *bool                          `json:"has_pcap,omitempty"`
    Hostname            *string                        `json:"hostname,omitempty"`
    HwRev               *string                        `json:"hw_rev,omitempty"`
    Id                  *uuid.UUID                     `json:"id,omitempty"`
    IfStat              map[string]IfStatProperty      `json:"if_stat,omitempty"`
    Ip                  *string                        `json:"ip,omitempty"`
    IpStat              *IpStat                        `json:"ip_stat,omitempty"`
    LastSeen            Optional[float64]              `json:"last_seen"`
    LastTrouble         *LastTrouble                   `json:"last_trouble,omitempty"`
    Mac                 *string                        `json:"mac,omitempty"`
    MacTableStats       *MacTableStats                 `json:"mac_table_stats,omitempty"`
    MapId               Optional[uuid.UUID]            `json:"map_id"`
    MemoryStat          *MemoryStat                    `json:"memory_stat,omitempty"`
    Model               *string                        `json:"model,omitempty"`
    ModifiedTime        *float64                       `json:"modified_time,omitempty"`
    ModuleStat          []StatsSwitchModuleStatItem    `json:"module_stat,omitempty"`
    Name                *string                        `json:"name,omitempty"`
    OrgId               *uuid.UUID                     `json:"org_id,omitempty"`
    Ports               []StatsSwitchPort              `json:"ports,omitempty"`
    RouteSummaryStats   *RouteSummaryStats             `json:"route_summary_stats,omitempty"`
    Serial              *string                        `json:"serial,omitempty"`
    ServiceStat         map[string]ServiceStatProperty `json:"service_stat,omitempty"`
    SiteId              *uuid.UUID                     `json:"site_id,omitempty"`
    Status              *string                        `json:"status,omitempty"`
    Type                *string                        `json:"type"`
    Uptime              Optional[float64]              `json:"uptime"`
    VcMac               Optional[string]               `json:"vc_mac"`
    VcSetupInfo         *StatsSwitchVcSetupInfo        `json:"vc_setup_info,omitempty"`
    Version             Optional[string]               `json:"version"`
}

func (s *tempStatsSwitch) validate() error {
    var errs []string
    if s.Type == nil {
        errs = append(errs, "required field `type` is missing for type `stats_switch`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
