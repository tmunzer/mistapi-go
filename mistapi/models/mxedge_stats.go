package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// MxedgeStats represents a MxedgeStats struct.
type MxedgeStats struct {
    // CPU/core stats list
    CpuStat              *MxedgeStatsCpuStat               `json:"cpu_stat,omitempty"`
    CreatedTime          *float64                          `json:"created_time,omitempty"`
    // alue indicating fips configuration on the device
    FipsEnabled          *bool                             `json:"fips_enabled,omitempty"`
    ForSite              *bool                             `json:"for_site,omitempty"`
    Id                   *uuid.UUID                        `json:"id,omitempty"`
    // OOBM IP stats
    IpStat               *MxedgeStatsIpStat                `json:"ip_stat,omitempty"`
    // Stat for LAG (Link Aggregation Group). Property key is the LAG name
    LagStat              map[string]MxedgeStatsLagStat     `json:"lag_stat,omitempty"`
    LastSeen             *float64                          `json:"last_seen,omitempty"`
    Mac                  *string                           `json:"mac,omitempty"`
    // Memory usage
    MemoryStat           *MxedgeStatsMemoryStat            `json:"memory_stat,omitempty"`
    Model                *string                           `json:"model,omitempty"`
    ModifiedTime         *float64                          `json:"modified_time,omitempty"`
    MxagentRegistered    *bool                             `json:"mxagent_registered,omitempty"`
    MxclusterId          *uuid.UUID                        `json:"mxcluster_id,omitempty"`
    // The name of the tunnel
    Name                 *string                           `json:"name,omitempty"`
    NumTunnels           *int                              `json:"num_tunnels,omitempty"`
    // ip configuration of the Mist Edge out-of_band management interface
    OobIpConfig          *MxedgeOobIpConfig                `json:"oob_ip_config,omitempty"`
    OrgId                *uuid.UUID                        `json:"org_id,omitempty"`
    PortStat             map[string]MxedgeStatsPortStat    `json:"port_stat,omitempty"`
    SensorStat           *interface{}                      `json:"sensor_stat,omitempty"`
    Serial               Optional[string]                  `json:"serial"`
    // stat for each services
    ServiceStat          map[string]MxedgeStatsServiceStat `json:"service_stat,omitempty"`
    Services             []string                          `json:"services,omitempty"`
    SiteId               *uuid.UUID                        `json:"site_id,omitempty"`
    Status               *string                           `json:"status,omitempty"`
    TuntermIpConfig      *MxedgeStatsTuntermIpConfig       `json:"tunterm_ip_config,omitempty"`
    TuntermPortConfig    *MxedgeStatsTuntermPortConfig     `json:"tunterm_port_config,omitempty"`
    TuntermRegistered    *bool                             `json:"tunterm_registered,omitempty"`
    TuntermStat          *MxedgeStatsTuntermStat           `json:"tunterm_stat,omitempty"`
    Uptime               *int                              `json:"uptime,omitempty"`
    // Virtualization environment
    VirtualizationType   *string                           `json:"virtualization_type,omitempty"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStats.
// It customizes the JSON marshaling process for MxedgeStats objects.
func (m MxedgeStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStats object to a map representation for JSON marshaling.
func (m MxedgeStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.CpuStat != nil {
        structMap["cpu_stat"] = m.CpuStat.toMap()
    }
    if m.CreatedTime != nil {
        structMap["created_time"] = m.CreatedTime
    }
    if m.FipsEnabled != nil {
        structMap["fips_enabled"] = m.FipsEnabled
    }
    if m.ForSite != nil {
        structMap["for_site"] = m.ForSite
    }
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.IpStat != nil {
        structMap["ip_stat"] = m.IpStat.toMap()
    }
    if m.LagStat != nil {
        structMap["lag_stat"] = m.LagStat
    }
    if m.LastSeen != nil {
        structMap["last_seen"] = m.LastSeen
    }
    if m.Mac != nil {
        structMap["mac"] = m.Mac
    }
    if m.MemoryStat != nil {
        structMap["memory_stat"] = m.MemoryStat.toMap()
    }
    if m.Model != nil {
        structMap["model"] = m.Model
    }
    if m.ModifiedTime != nil {
        structMap["modified_time"] = m.ModifiedTime
    }
    if m.MxagentRegistered != nil {
        structMap["mxagent_registered"] = m.MxagentRegistered
    }
    if m.MxclusterId != nil {
        structMap["mxcluster_id"] = m.MxclusterId
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.NumTunnels != nil {
        structMap["num_tunnels"] = m.NumTunnels
    }
    if m.OobIpConfig != nil {
        structMap["oob_ip_config"] = m.OobIpConfig.toMap()
    }
    if m.OrgId != nil {
        structMap["org_id"] = m.OrgId
    }
    if m.PortStat != nil {
        structMap["port_stat"] = m.PortStat
    }
    if m.SensorStat != nil {
        structMap["sensor_stat"] = m.SensorStat
    }
    if m.Serial.IsValueSet() {
        if m.Serial.Value() != nil {
            structMap["serial"] = m.Serial.Value()
        } else {
            structMap["serial"] = nil
        }
    }
    if m.ServiceStat != nil {
        structMap["service_stat"] = m.ServiceStat
    }
    if m.Services != nil {
        structMap["services"] = m.Services
    }
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    if m.Status != nil {
        structMap["status"] = m.Status
    }
    if m.TuntermIpConfig != nil {
        structMap["tunterm_ip_config"] = m.TuntermIpConfig.toMap()
    }
    if m.TuntermPortConfig != nil {
        structMap["tunterm_port_config"] = m.TuntermPortConfig.toMap()
    }
    if m.TuntermRegistered != nil {
        structMap["tunterm_registered"] = m.TuntermRegistered
    }
    if m.TuntermStat != nil {
        structMap["tunterm_stat"] = m.TuntermStat.toMap()
    }
    if m.Uptime != nil {
        structMap["uptime"] = m.Uptime
    }
    if m.VirtualizationType != nil {
        structMap["virtualization_type"] = m.VirtualizationType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStats.
// It customizes the JSON unmarshaling process for MxedgeStats objects.
func (m *MxedgeStats) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cpu_stat", "created_time", "fips_enabled", "for_site", "id", "ip_stat", "lag_stat", "last_seen", "mac", "memory_stat", "model", "modified_time", "mxagent_registered", "mxcluster_id", "name", "num_tunnels", "oob_ip_config", "org_id", "port_stat", "sensor_stat", "serial", "service_stat", "services", "site_id", "status", "tunterm_ip_config", "tunterm_port_config", "tunterm_registered", "tunterm_stat", "uptime", "virtualization_type")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.CpuStat = temp.CpuStat
    m.CreatedTime = temp.CreatedTime
    m.FipsEnabled = temp.FipsEnabled
    m.ForSite = temp.ForSite
    m.Id = temp.Id
    m.IpStat = temp.IpStat
    m.LagStat = temp.LagStat
    m.LastSeen = temp.LastSeen
    m.Mac = temp.Mac
    m.MemoryStat = temp.MemoryStat
    m.Model = temp.Model
    m.ModifiedTime = temp.ModifiedTime
    m.MxagentRegistered = temp.MxagentRegistered
    m.MxclusterId = temp.MxclusterId
    m.Name = temp.Name
    m.NumTunnels = temp.NumTunnels
    m.OobIpConfig = temp.OobIpConfig
    m.OrgId = temp.OrgId
    m.PortStat = temp.PortStat
    m.SensorStat = temp.SensorStat
    m.Serial = temp.Serial
    m.ServiceStat = temp.ServiceStat
    m.Services = temp.Services
    m.SiteId = temp.SiteId
    m.Status = temp.Status
    m.TuntermIpConfig = temp.TuntermIpConfig
    m.TuntermPortConfig = temp.TuntermPortConfig
    m.TuntermRegistered = temp.TuntermRegistered
    m.TuntermStat = temp.TuntermStat
    m.Uptime = temp.Uptime
    m.VirtualizationType = temp.VirtualizationType
    return nil
}

// tempMxedgeStats is a temporary struct used for validating the fields of MxedgeStats.
type tempMxedgeStats  struct {
    CpuStat            *MxedgeStatsCpuStat               `json:"cpu_stat,omitempty"`
    CreatedTime        *float64                          `json:"created_time,omitempty"`
    FipsEnabled        *bool                             `json:"fips_enabled,omitempty"`
    ForSite            *bool                             `json:"for_site,omitempty"`
    Id                 *uuid.UUID                        `json:"id,omitempty"`
    IpStat             *MxedgeStatsIpStat                `json:"ip_stat,omitempty"`
    LagStat            map[string]MxedgeStatsLagStat     `json:"lag_stat,omitempty"`
    LastSeen           *float64                          `json:"last_seen,omitempty"`
    Mac                *string                           `json:"mac,omitempty"`
    MemoryStat         *MxedgeStatsMemoryStat            `json:"memory_stat,omitempty"`
    Model              *string                           `json:"model,omitempty"`
    ModifiedTime       *float64                          `json:"modified_time,omitempty"`
    MxagentRegistered  *bool                             `json:"mxagent_registered,omitempty"`
    MxclusterId        *uuid.UUID                        `json:"mxcluster_id,omitempty"`
    Name               *string                           `json:"name,omitempty"`
    NumTunnels         *int                              `json:"num_tunnels,omitempty"`
    OobIpConfig        *MxedgeOobIpConfig                `json:"oob_ip_config,omitempty"`
    OrgId              *uuid.UUID                        `json:"org_id,omitempty"`
    PortStat           map[string]MxedgeStatsPortStat    `json:"port_stat,omitempty"`
    SensorStat         *interface{}                      `json:"sensor_stat,omitempty"`
    Serial             Optional[string]                  `json:"serial"`
    ServiceStat        map[string]MxedgeStatsServiceStat `json:"service_stat,omitempty"`
    Services           []string                          `json:"services,omitempty"`
    SiteId             *uuid.UUID                        `json:"site_id,omitempty"`
    Status             *string                           `json:"status,omitempty"`
    TuntermIpConfig    *MxedgeStatsTuntermIpConfig       `json:"tunterm_ip_config,omitempty"`
    TuntermPortConfig  *MxedgeStatsTuntermPortConfig     `json:"tunterm_port_config,omitempty"`
    TuntermRegistered  *bool                             `json:"tunterm_registered,omitempty"`
    TuntermStat        *MxedgeStatsTuntermStat           `json:"tunterm_stat,omitempty"`
    Uptime             *int                              `json:"uptime,omitempty"`
    VirtualizationType *string                           `json:"virtualization_type,omitempty"`
}
