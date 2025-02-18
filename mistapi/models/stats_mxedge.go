package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// StatsMxedge represents a StatsMxedge struct.
type StatsMxedge struct {
    // CPU/core stats list
    CpuStat              *StatsMxedgeCpuStat               `json:"cpu_stat,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64                          `json:"created_time,omitempty"`
    // Indicate fips configuration on the device
    FipsEnabled          *bool                             `json:"fips_enabled,omitempty"`
    ForSite              *bool                             `json:"for_site,omitempty"`
    Fwupdate             *FwupdateStat                     `json:"fwupdate,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID                        `json:"id,omitempty"`
    // IDRAC version of the mist edge device
    IdracVersion         *string                           `json:"idrac_version,omitempty"`
    // IP stats
    IpStat               *StatsMxedgeIpStat                `json:"ip_stat,omitempty"`
    // Stat for LAG (Link Aggregation Group). Property key is the LAG name
    LagStat              map[string]StatsMxedgeLagStat     `json:"lag_stat,omitempty"`
    LastSeen             *float64                          `json:"last_seen,omitempty"`
    Mac                  *string                           `json:"mac,omitempty"`
    // Memory usage
    MemoryStat           *StatsMxedgeMemoryStat            `json:"memory_stat,omitempty"`
    Model                *string                           `json:"model,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                          `json:"modified_time,omitempty"`
    MxagentRegistered    *bool                             `json:"mxagent_registered,omitempty"`
    MxclusterId          *uuid.UUID                        `json:"mxcluster_id,omitempty"`
    // The name of the tunnel
    Name                 *string                           `json:"name,omitempty"`
    NumTunnels           *int                              `json:"num_tunnels,omitempty"`
    // IPconfiguration of the Mist Edge out-of_band management interface
    OobIpConfig          *MxedgeOobIpConfig                `json:"oob_ip_config,omitempty"`
    OobIpStat            *StatsMxedgeOobIpStat             `json:"oob_ip_stat,omitempty"`
    OrgId                *uuid.UUID                        `json:"org_id,omitempty"`
    PortStat             map[string]StatsMxedgePortStat    `json:"port_stat,omitempty"`
    Serial               Optional[string]                  `json:"serial"`
    // Stat for each services
    ServiceStat          map[string]StatsMxedgeServiceStat `json:"service_stat,omitempty"`
    Services             []string                          `json:"services,omitempty"`
    SiteId               *uuid.UUID                        `json:"site_id,omitempty"`
    Status               *string                           `json:"status,omitempty"`
    TuntermIpConfig      *StatsMxedgeTuntermIpConfig       `json:"tunterm_ip_config,omitempty"`
    TuntermPortConfig    *StatsMxedgeTuntermPortConfig     `json:"tunterm_port_config,omitempty"`
    TuntermRegistered    *bool                             `json:"tunterm_registered,omitempty"`
    TuntermStat          *StatsMxedgeTuntermStat           `json:"tunterm_stat,omitempty"`
    Uptime               *int                              `json:"uptime,omitempty"`
    // Virtualization environment
    VirtualizationType   *string                           `json:"virtualization_type,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for StatsMxedge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsMxedge) String() string {
    return fmt.Sprintf(
    	"StatsMxedge[CpuStat=%v, CreatedTime=%v, FipsEnabled=%v, ForSite=%v, Fwupdate=%v, Id=%v, IdracVersion=%v, IpStat=%v, LagStat=%v, LastSeen=%v, Mac=%v, MemoryStat=%v, Model=%v, ModifiedTime=%v, MxagentRegistered=%v, MxclusterId=%v, Name=%v, NumTunnels=%v, OobIpConfig=%v, OobIpStat=%v, OrgId=%v, PortStat=%v, Serial=%v, ServiceStat=%v, Services=%v, SiteId=%v, Status=%v, TuntermIpConfig=%v, TuntermPortConfig=%v, TuntermRegistered=%v, TuntermStat=%v, Uptime=%v, VirtualizationType=%v, AdditionalProperties=%v]",
    	s.CpuStat, s.CreatedTime, s.FipsEnabled, s.ForSite, s.Fwupdate, s.Id, s.IdracVersion, s.IpStat, s.LagStat, s.LastSeen, s.Mac, s.MemoryStat, s.Model, s.ModifiedTime, s.MxagentRegistered, s.MxclusterId, s.Name, s.NumTunnels, s.OobIpConfig, s.OobIpStat, s.OrgId, s.PortStat, s.Serial, s.ServiceStat, s.Services, s.SiteId, s.Status, s.TuntermIpConfig, s.TuntermPortConfig, s.TuntermRegistered, s.TuntermStat, s.Uptime, s.VirtualizationType, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsMxedge.
// It customizes the JSON marshaling process for StatsMxedge objects.
func (s StatsMxedge) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "cpu_stat", "created_time", "fips_enabled", "for_site", "fwupdate", "id", "idrac_version", "ip_stat", "lag_stat", "last_seen", "mac", "memory_stat", "model", "modified_time", "mxagent_registered", "mxcluster_id", "name", "num_tunnels", "oob_ip_config", "oob_ip_stat", "org_id", "port_stat", "serial", "service_stat", "services", "site_id", "status", "tunterm_ip_config", "tunterm_port_config", "tunterm_registered", "tunterm_stat", "uptime", "virtualization_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxedge object to a map representation for JSON marshaling.
func (s StatsMxedge) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CpuStat != nil {
        structMap["cpu_stat"] = s.CpuStat.toMap()
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.FipsEnabled != nil {
        structMap["fips_enabled"] = s.FipsEnabled
    }
    if s.ForSite != nil {
        structMap["for_site"] = s.ForSite
    }
    if s.Fwupdate != nil {
        structMap["fwupdate"] = s.Fwupdate.toMap()
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.IdracVersion != nil {
        structMap["idrac_version"] = s.IdracVersion
    }
    if s.IpStat != nil {
        structMap["ip_stat"] = s.IpStat.toMap()
    }
    if s.LagStat != nil {
        structMap["lag_stat"] = s.LagStat
    }
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    }
    if s.Mac != nil {
        structMap["mac"] = s.Mac
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
    if s.MxagentRegistered != nil {
        structMap["mxagent_registered"] = s.MxagentRegistered
    }
    if s.MxclusterId != nil {
        structMap["mxcluster_id"] = s.MxclusterId
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.NumTunnels != nil {
        structMap["num_tunnels"] = s.NumTunnels
    }
    if s.OobIpConfig != nil {
        structMap["oob_ip_config"] = s.OobIpConfig.toMap()
    }
    if s.OobIpStat != nil {
        structMap["oob_ip_stat"] = s.OobIpStat.toMap()
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.PortStat != nil {
        structMap["port_stat"] = s.PortStat
    }
    if s.Serial.IsValueSet() {
        if s.Serial.Value() != nil {
            structMap["serial"] = s.Serial.Value()
        } else {
            structMap["serial"] = nil
        }
    }
    if s.ServiceStat != nil {
        structMap["service_stat"] = s.ServiceStat
    }
    if s.Services != nil {
        structMap["services"] = s.Services
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.TuntermIpConfig != nil {
        structMap["tunterm_ip_config"] = s.TuntermIpConfig.toMap()
    }
    if s.TuntermPortConfig != nil {
        structMap["tunterm_port_config"] = s.TuntermPortConfig.toMap()
    }
    if s.TuntermRegistered != nil {
        structMap["tunterm_registered"] = s.TuntermRegistered
    }
    if s.TuntermStat != nil {
        structMap["tunterm_stat"] = s.TuntermStat.toMap()
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.VirtualizationType != nil {
        structMap["virtualization_type"] = s.VirtualizationType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxedge.
// It customizes the JSON unmarshaling process for StatsMxedge objects.
func (s *StatsMxedge) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxedge
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cpu_stat", "created_time", "fips_enabled", "for_site", "fwupdate", "id", "idrac_version", "ip_stat", "lag_stat", "last_seen", "mac", "memory_stat", "model", "modified_time", "mxagent_registered", "mxcluster_id", "name", "num_tunnels", "oob_ip_config", "oob_ip_stat", "org_id", "port_stat", "serial", "service_stat", "services", "site_id", "status", "tunterm_ip_config", "tunterm_port_config", "tunterm_registered", "tunterm_stat", "uptime", "virtualization_type")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CpuStat = temp.CpuStat
    s.CreatedTime = temp.CreatedTime
    s.FipsEnabled = temp.FipsEnabled
    s.ForSite = temp.ForSite
    s.Fwupdate = temp.Fwupdate
    s.Id = temp.Id
    s.IdracVersion = temp.IdracVersion
    s.IpStat = temp.IpStat
    s.LagStat = temp.LagStat
    s.LastSeen = temp.LastSeen
    s.Mac = temp.Mac
    s.MemoryStat = temp.MemoryStat
    s.Model = temp.Model
    s.ModifiedTime = temp.ModifiedTime
    s.MxagentRegistered = temp.MxagentRegistered
    s.MxclusterId = temp.MxclusterId
    s.Name = temp.Name
    s.NumTunnels = temp.NumTunnels
    s.OobIpConfig = temp.OobIpConfig
    s.OobIpStat = temp.OobIpStat
    s.OrgId = temp.OrgId
    s.PortStat = temp.PortStat
    s.Serial = temp.Serial
    s.ServiceStat = temp.ServiceStat
    s.Services = temp.Services
    s.SiteId = temp.SiteId
    s.Status = temp.Status
    s.TuntermIpConfig = temp.TuntermIpConfig
    s.TuntermPortConfig = temp.TuntermPortConfig
    s.TuntermRegistered = temp.TuntermRegistered
    s.TuntermStat = temp.TuntermStat
    s.Uptime = temp.Uptime
    s.VirtualizationType = temp.VirtualizationType
    return nil
}

// tempStatsMxedge is a temporary struct used for validating the fields of StatsMxedge.
type tempStatsMxedge  struct {
    CpuStat            *StatsMxedgeCpuStat               `json:"cpu_stat,omitempty"`
    CreatedTime        *float64                          `json:"created_time,omitempty"`
    FipsEnabled        *bool                             `json:"fips_enabled,omitempty"`
    ForSite            *bool                             `json:"for_site,omitempty"`
    Fwupdate           *FwupdateStat                     `json:"fwupdate,omitempty"`
    Id                 *uuid.UUID                        `json:"id,omitempty"`
    IdracVersion       *string                           `json:"idrac_version,omitempty"`
    IpStat             *StatsMxedgeIpStat                `json:"ip_stat,omitempty"`
    LagStat            map[string]StatsMxedgeLagStat     `json:"lag_stat,omitempty"`
    LastSeen           *float64                          `json:"last_seen,omitempty"`
    Mac                *string                           `json:"mac,omitempty"`
    MemoryStat         *StatsMxedgeMemoryStat            `json:"memory_stat,omitempty"`
    Model              *string                           `json:"model,omitempty"`
    ModifiedTime       *float64                          `json:"modified_time,omitempty"`
    MxagentRegistered  *bool                             `json:"mxagent_registered,omitempty"`
    MxclusterId        *uuid.UUID                        `json:"mxcluster_id,omitempty"`
    Name               *string                           `json:"name,omitempty"`
    NumTunnels         *int                              `json:"num_tunnels,omitempty"`
    OobIpConfig        *MxedgeOobIpConfig                `json:"oob_ip_config,omitempty"`
    OobIpStat          *StatsMxedgeOobIpStat             `json:"oob_ip_stat,omitempty"`
    OrgId              *uuid.UUID                        `json:"org_id,omitempty"`
    PortStat           map[string]StatsMxedgePortStat    `json:"port_stat,omitempty"`
    Serial             Optional[string]                  `json:"serial"`
    ServiceStat        map[string]StatsMxedgeServiceStat `json:"service_stat,omitempty"`
    Services           []string                          `json:"services,omitempty"`
    SiteId             *uuid.UUID                        `json:"site_id,omitempty"`
    Status             *string                           `json:"status,omitempty"`
    TuntermIpConfig    *StatsMxedgeTuntermIpConfig       `json:"tunterm_ip_config,omitempty"`
    TuntermPortConfig  *StatsMxedgeTuntermPortConfig     `json:"tunterm_port_config,omitempty"`
    TuntermRegistered  *bool                             `json:"tunterm_registered,omitempty"`
    TuntermStat        *StatsMxedgeTuntermStat           `json:"tunterm_stat,omitempty"`
    Uptime             *int                              `json:"uptime,omitempty"`
    VirtualizationType *string                           `json:"virtualization_type,omitempty"`
}
