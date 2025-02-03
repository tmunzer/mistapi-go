package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// Mxedge represents a Mxedge struct.
// MxEdge
type Mxedge struct {
    // When the object has been created, in epoch
    CreatedTime               *float64                              `json:"created_time,omitempty"`
    ForSite                   *bool                                 `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                        *uuid.UUID                            `json:"id,omitempty"`
    Magic                     *string                               `json:"magic,omitempty"`
    Model                     string                                `json:"model"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime              *float64                              `json:"modified_time,omitempty"`
    MxagentRegistered         *bool                                 `json:"mxagent_registered,omitempty"`
    // MxCluster this MxEdge belongs to
    MxclusterId               *uuid.UUID                            `json:"mxcluster_id,omitempty"`
    MxedgeMgmt                *MxedgeMgmt                           `json:"mxedge_mgmt,omitempty"`
    Name                      string                                `json:"name"`
    Note                      *string                               `json:"note,omitempty"`
    NtpServers                []string                              `json:"ntp_servers,omitempty"`
    // IPconfiguration of the Mist Edge out-of_band management interface
    OobIpConfig               *MxedgeOobIpConfig                    `json:"oob_ip_config,omitempty"`
    OrgId                     *uuid.UUID                            `json:"org_id,omitempty"`
    // Proxy Configuration to talk to Mist
    Proxy                     *Proxy                                `json:"proxy,omitempty"`
    // List of services to run, tunterm only for now
    Services                  []string                              `json:"services,omitempty"`
    SiteId                    *uuid.UUID                            `json:"site_id,omitempty"`
    // Global and per-VLAN. Property key is the VLAN ID
    TuntermDhcpdConfig        *MxedgeTuntermDhcpdConfig             `json:"tunterm_dhcpd_config,omitempty"`
    // Property key is a CIDR
    TuntermExtraRoutes        map[string]MxedgeTuntermExtraRoute    `json:"tunterm_extra_routes,omitempty"`
    TuntermIgmpSnoopingConfig *MxedgeTuntermIgmpSnoopingConfig      `json:"tunterm_igmp_snooping_config,omitempty"`
    // IPconfiguration of the Mist Tunnel interface
    TuntermIpConfig           *MxedgeTuntermIpConfig                `json:"tunterm_ip_config,omitempty"`
    TuntermMonitoring         [][]TuntermMonitoringItem             `json:"tunterm_monitoring,omitempty"`
    TuntermMulticastConfig    *MxedgeTuntermMulticastConfig         `json:"tunterm_multicast_config,omitempty"`
    // IPconfigs by VLAN ID. Property key is the VLAN ID
    TuntermOtherIpConfigs     map[string]MxedgeTuntermOtherIpConfig `json:"tunterm_other_ip_configs,omitempty"`
    // Ethernet port configurations
    TuntermPortConfig         *TuntermPortConfig                    `json:"tunterm_port_config,omitempty"`
    TuntermRegistered         *bool                                 `json:"tunterm_registered,omitempty"`
    // If custom vlan settings are desired
    TuntermSwitchConfig       *MxedgeTuntermSwitchConfigs           `json:"tunterm_switch_config,omitempty"`
    Versions                  *MxedgeVersions                       `json:"versions,omitempty"`
    AdditionalProperties      map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for Mxedge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Mxedge) String() string {
    return fmt.Sprintf(
    	"Mxedge[CreatedTime=%v, ForSite=%v, Id=%v, Magic=%v, Model=%v, ModifiedTime=%v, MxagentRegistered=%v, MxclusterId=%v, MxedgeMgmt=%v, Name=%v, Note=%v, NtpServers=%v, OobIpConfig=%v, OrgId=%v, Proxy=%v, Services=%v, SiteId=%v, TuntermDhcpdConfig=%v, TuntermExtraRoutes=%v, TuntermIgmpSnoopingConfig=%v, TuntermIpConfig=%v, TuntermMonitoring=%v, TuntermMulticastConfig=%v, TuntermOtherIpConfigs=%v, TuntermPortConfig=%v, TuntermRegistered=%v, TuntermSwitchConfig=%v, Versions=%v, AdditionalProperties=%v]",
    	m.CreatedTime, m.ForSite, m.Id, m.Magic, m.Model, m.ModifiedTime, m.MxagentRegistered, m.MxclusterId, m.MxedgeMgmt, m.Name, m.Note, m.NtpServers, m.OobIpConfig, m.OrgId, m.Proxy, m.Services, m.SiteId, m.TuntermDhcpdConfig, m.TuntermExtraRoutes, m.TuntermIgmpSnoopingConfig, m.TuntermIpConfig, m.TuntermMonitoring, m.TuntermMulticastConfig, m.TuntermOtherIpConfigs, m.TuntermPortConfig, m.TuntermRegistered, m.TuntermSwitchConfig, m.Versions, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Mxedge.
// It customizes the JSON marshaling process for Mxedge objects.
func (m Mxedge) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "created_time", "for_site", "id", "magic", "model", "modified_time", "mxagent_registered", "mxcluster_id", "mxedge_mgmt", "name", "note", "ntp_servers", "oob_ip_config", "org_id", "proxy", "services", "site_id", "tunterm_dhcpd_config", "tunterm_extra_routes", "tunterm_igmp_snooping_config", "tunterm_ip_config", "tunterm_monitoring", "tunterm_multicast_config", "tunterm_other_ip_configs", "tunterm_port_config", "tunterm_registered", "tunterm_switch_config", "versions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Mxedge object to a map representation for JSON marshaling.
func (m Mxedge) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.CreatedTime != nil {
        structMap["created_time"] = m.CreatedTime
    }
    if m.ForSite != nil {
        structMap["for_site"] = m.ForSite
    }
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.Magic != nil {
        structMap["magic"] = m.Magic
    }
    structMap["model"] = m.Model
    if m.ModifiedTime != nil {
        structMap["modified_time"] = m.ModifiedTime
    }
    if m.MxagentRegistered != nil {
        structMap["mxagent_registered"] = m.MxagentRegistered
    }
    if m.MxclusterId != nil {
        structMap["mxcluster_id"] = m.MxclusterId
    }
    if m.MxedgeMgmt != nil {
        structMap["mxedge_mgmt"] = m.MxedgeMgmt.toMap()
    }
    structMap["name"] = m.Name
    if m.Note != nil {
        structMap["note"] = m.Note
    }
    if m.NtpServers != nil {
        structMap["ntp_servers"] = m.NtpServers
    }
    if m.OobIpConfig != nil {
        structMap["oob_ip_config"] = m.OobIpConfig.toMap()
    }
    if m.OrgId != nil {
        structMap["org_id"] = m.OrgId
    }
    if m.Proxy != nil {
        structMap["proxy"] = m.Proxy.toMap()
    }
    if m.Services != nil {
        structMap["services"] = m.Services
    }
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    if m.TuntermDhcpdConfig != nil {
        structMap["tunterm_dhcpd_config"] = m.TuntermDhcpdConfig.toMap()
    }
    if m.TuntermExtraRoutes != nil {
        structMap["tunterm_extra_routes"] = m.TuntermExtraRoutes
    }
    if m.TuntermIgmpSnoopingConfig != nil {
        structMap["tunterm_igmp_snooping_config"] = m.TuntermIgmpSnoopingConfig.toMap()
    }
    if m.TuntermIpConfig != nil {
        structMap["tunterm_ip_config"] = m.TuntermIpConfig.toMap()
    }
    if m.TuntermMonitoring != nil {
        structMap["tunterm_monitoring"] = m.TuntermMonitoring
    }
    if m.TuntermMulticastConfig != nil {
        structMap["tunterm_multicast_config"] = m.TuntermMulticastConfig.toMap()
    }
    if m.TuntermOtherIpConfigs != nil {
        structMap["tunterm_other_ip_configs"] = m.TuntermOtherIpConfigs
    }
    if m.TuntermPortConfig != nil {
        structMap["tunterm_port_config"] = m.TuntermPortConfig.toMap()
    }
    if m.TuntermRegistered != nil {
        structMap["tunterm_registered"] = m.TuntermRegistered
    }
    if m.TuntermSwitchConfig != nil {
        structMap["tunterm_switch_config"] = m.TuntermSwitchConfig.toMap()
    }
    if m.Versions != nil {
        structMap["versions"] = m.Versions.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Mxedge.
// It customizes the JSON unmarshaling process for Mxedge objects.
func (m *Mxedge) UnmarshalJSON(input []byte) error {
    var temp tempMxedge
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "for_site", "id", "magic", "model", "modified_time", "mxagent_registered", "mxcluster_id", "mxedge_mgmt", "name", "note", "ntp_servers", "oob_ip_config", "org_id", "proxy", "services", "site_id", "tunterm_dhcpd_config", "tunterm_extra_routes", "tunterm_igmp_snooping_config", "tunterm_ip_config", "tunterm_monitoring", "tunterm_multicast_config", "tunterm_other_ip_configs", "tunterm_port_config", "tunterm_registered", "tunterm_switch_config", "versions")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.CreatedTime = temp.CreatedTime
    m.ForSite = temp.ForSite
    m.Id = temp.Id
    m.Magic = temp.Magic
    m.Model = *temp.Model
    m.ModifiedTime = temp.ModifiedTime
    m.MxagentRegistered = temp.MxagentRegistered
    m.MxclusterId = temp.MxclusterId
    m.MxedgeMgmt = temp.MxedgeMgmt
    m.Name = *temp.Name
    m.Note = temp.Note
    m.NtpServers = temp.NtpServers
    m.OobIpConfig = temp.OobIpConfig
    m.OrgId = temp.OrgId
    m.Proxy = temp.Proxy
    m.Services = temp.Services
    m.SiteId = temp.SiteId
    m.TuntermDhcpdConfig = temp.TuntermDhcpdConfig
    m.TuntermExtraRoutes = temp.TuntermExtraRoutes
    m.TuntermIgmpSnoopingConfig = temp.TuntermIgmpSnoopingConfig
    m.TuntermIpConfig = temp.TuntermIpConfig
    m.TuntermMonitoring = temp.TuntermMonitoring
    m.TuntermMulticastConfig = temp.TuntermMulticastConfig
    m.TuntermOtherIpConfigs = temp.TuntermOtherIpConfigs
    m.TuntermPortConfig = temp.TuntermPortConfig
    m.TuntermRegistered = temp.TuntermRegistered
    m.TuntermSwitchConfig = temp.TuntermSwitchConfig
    m.Versions = temp.Versions
    return nil
}

// tempMxedge is a temporary struct used for validating the fields of Mxedge.
type tempMxedge  struct {
    CreatedTime               *float64                              `json:"created_time,omitempty"`
    ForSite                   *bool                                 `json:"for_site,omitempty"`
    Id                        *uuid.UUID                            `json:"id,omitempty"`
    Magic                     *string                               `json:"magic,omitempty"`
    Model                     *string                               `json:"model"`
    ModifiedTime              *float64                              `json:"modified_time,omitempty"`
    MxagentRegistered         *bool                                 `json:"mxagent_registered,omitempty"`
    MxclusterId               *uuid.UUID                            `json:"mxcluster_id,omitempty"`
    MxedgeMgmt                *MxedgeMgmt                           `json:"mxedge_mgmt,omitempty"`
    Name                      *string                               `json:"name"`
    Note                      *string                               `json:"note,omitempty"`
    NtpServers                []string                              `json:"ntp_servers,omitempty"`
    OobIpConfig               *MxedgeOobIpConfig                    `json:"oob_ip_config,omitempty"`
    OrgId                     *uuid.UUID                            `json:"org_id,omitempty"`
    Proxy                     *Proxy                                `json:"proxy,omitempty"`
    Services                  []string                              `json:"services,omitempty"`
    SiteId                    *uuid.UUID                            `json:"site_id,omitempty"`
    TuntermDhcpdConfig        *MxedgeTuntermDhcpdConfig             `json:"tunterm_dhcpd_config,omitempty"`
    TuntermExtraRoutes        map[string]MxedgeTuntermExtraRoute    `json:"tunterm_extra_routes,omitempty"`
    TuntermIgmpSnoopingConfig *MxedgeTuntermIgmpSnoopingConfig      `json:"tunterm_igmp_snooping_config,omitempty"`
    TuntermIpConfig           *MxedgeTuntermIpConfig                `json:"tunterm_ip_config,omitempty"`
    TuntermMonitoring         [][]TuntermMonitoringItem             `json:"tunterm_monitoring,omitempty"`
    TuntermMulticastConfig    *MxedgeTuntermMulticastConfig         `json:"tunterm_multicast_config,omitempty"`
    TuntermOtherIpConfigs     map[string]MxedgeTuntermOtherIpConfig `json:"tunterm_other_ip_configs,omitempty"`
    TuntermPortConfig         *TuntermPortConfig                    `json:"tunterm_port_config,omitempty"`
    TuntermRegistered         *bool                                 `json:"tunterm_registered,omitempty"`
    TuntermSwitchConfig       *MxedgeTuntermSwitchConfigs           `json:"tunterm_switch_config,omitempty"`
    Versions                  *MxedgeVersions                       `json:"versions,omitempty"`
}

func (m *tempMxedge) validate() error {
    var errs []string
    if m.Model == nil {
        errs = append(errs, "required field `model` is missing for type `mxedge`")
    }
    if m.Name == nil {
        errs = append(errs, "required field `name` is missing for type `mxedge`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
