package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Mxcluster represents a Mxcluster struct.
// MxCluster
type Mxcluster struct {
    // When the object has been created, in epoch
    CreatedTime               *float64                              `json:"created_time,omitempty"`
    ForSite                   *bool                                 `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                        *uuid.UUID                            `json:"id,omitempty"`
    // Configure cloud-assisted dynamic authorization service on this cluster of mist edges
    MistDas                   *MxedgeDas                            `json:"mist_das,omitempty"`
    MistNac                   *MxclusterNac                         `json:"mist_nac,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime              *float64                              `json:"modified_time,omitempty"`
    MxedgeMgmt                *MxedgeMgmt                           `json:"mxedge_mgmt,omitempty"`
    Name                      *string                               `json:"name,omitempty"`
    OrgId                     *uuid.UUID                            `json:"org_id,omitempty"`
    // Proxy Configuration to talk to Mist
    Proxy                     *Proxy                                `json:"proxy,omitempty"`
    // MxEdge Radsec Configuration
    Radsec                    *MxclusterRadsec                      `json:"radsec,omitempty"`
    RadsecTls                 *MxclusterRadsecTls                   `json:"radsec_tls,omitempty"`
    SiteId                    *uuid.UUID                            `json:"site_id,omitempty"`
    // List of subnets where we allow AP to establish Mist Tunnels from
    TuntermApSubnets          []string                              `json:"tunterm_ap_subnets,omitempty"`
    // DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID
    TuntermDhcpdConfig        *TuntermDhcpdConfig                   `json:"tunterm_dhcpd_config,omitempty"`
    // Extra routes for Mist Tunneled VLANs. Property key is a CIDR
    TuntermExtraRoutes        map[string]MxclusterTuntermExtraRoute `json:"tunterm_extra_routes,omitempty"`
    // Hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP)
    TuntermHosts              []string                              `json:"tunterm_hosts,omitempty"`
    // List of index of tunterm_hosts
    TuntermHostsOrder         []int                                 `json:"tunterm_hosts_order,omitempty"`
    // Ordering of tunterm_hosts for mxedge within the same mxcluster. enum:
    // * `shuffle`: the ordering of tunterm_hosts is randomized by the device''s MAC
    // * `shuffle-by-site`: shuffle by site_id+tunnel_id (so when client connects to a specific Tunnel, it will go to the same (order of) mxedge, and we load-balancing between tunnels)
    // * `ordered`: order decided by tunterm_hosts_order
    TuntermHostsSelection     *MxclusterTuntermHostsSelectionEnum   `json:"tunterm_hosts_selection,omitempty"`
    TuntermMonitoring         [][]TuntermMonitoringItem             `json:"tunterm_monitoring,omitempty"`
    TuntermMonitoringDisabled *bool                                 `json:"tunterm_monitoring_disabled,omitempty"`
    AdditionalProperties      map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for Mxcluster,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Mxcluster) String() string {
    return fmt.Sprintf(
    	"Mxcluster[CreatedTime=%v, ForSite=%v, Id=%v, MistDas=%v, MistNac=%v, ModifiedTime=%v, MxedgeMgmt=%v, Name=%v, OrgId=%v, Proxy=%v, Radsec=%v, RadsecTls=%v, SiteId=%v, TuntermApSubnets=%v, TuntermDhcpdConfig=%v, TuntermExtraRoutes=%v, TuntermHosts=%v, TuntermHostsOrder=%v, TuntermHostsSelection=%v, TuntermMonitoring=%v, TuntermMonitoringDisabled=%v, AdditionalProperties=%v]",
    	m.CreatedTime, m.ForSite, m.Id, m.MistDas, m.MistNac, m.ModifiedTime, m.MxedgeMgmt, m.Name, m.OrgId, m.Proxy, m.Radsec, m.RadsecTls, m.SiteId, m.TuntermApSubnets, m.TuntermDhcpdConfig, m.TuntermExtraRoutes, m.TuntermHosts, m.TuntermHostsOrder, m.TuntermHostsSelection, m.TuntermMonitoring, m.TuntermMonitoringDisabled, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Mxcluster.
// It customizes the JSON marshaling process for Mxcluster objects.
func (m Mxcluster) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "created_time", "for_site", "id", "mist_das", "mist_nac", "modified_time", "mxedge_mgmt", "name", "org_id", "proxy", "radsec", "radsec_tls", "site_id", "tunterm_ap_subnets", "tunterm_dhcpd_config", "tunterm_extra_routes", "tunterm_hosts", "tunterm_hosts_order", "tunterm_hosts_selection", "tunterm_monitoring", "tunterm_monitoring_disabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Mxcluster object to a map representation for JSON marshaling.
func (m Mxcluster) toMap() map[string]any {
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
    if m.MistDas != nil {
        structMap["mist_das"] = m.MistDas.toMap()
    }
    if m.MistNac != nil {
        structMap["mist_nac"] = m.MistNac.toMap()
    }
    if m.ModifiedTime != nil {
        structMap["modified_time"] = m.ModifiedTime
    }
    if m.MxedgeMgmt != nil {
        structMap["mxedge_mgmt"] = m.MxedgeMgmt.toMap()
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.OrgId != nil {
        structMap["org_id"] = m.OrgId
    }
    if m.Proxy != nil {
        structMap["proxy"] = m.Proxy.toMap()
    }
    if m.Radsec != nil {
        structMap["radsec"] = m.Radsec.toMap()
    }
    if m.RadsecTls != nil {
        structMap["radsec_tls"] = m.RadsecTls.toMap()
    }
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    if m.TuntermApSubnets != nil {
        structMap["tunterm_ap_subnets"] = m.TuntermApSubnets
    }
    if m.TuntermDhcpdConfig != nil {
        structMap["tunterm_dhcpd_config"] = m.TuntermDhcpdConfig.toMap()
    }
    if m.TuntermExtraRoutes != nil {
        structMap["tunterm_extra_routes"] = m.TuntermExtraRoutes
    }
    if m.TuntermHosts != nil {
        structMap["tunterm_hosts"] = m.TuntermHosts
    }
    if m.TuntermHostsOrder != nil {
        structMap["tunterm_hosts_order"] = m.TuntermHostsOrder
    }
    if m.TuntermHostsSelection != nil {
        structMap["tunterm_hosts_selection"] = m.TuntermHostsSelection
    }
    if m.TuntermMonitoring != nil {
        structMap["tunterm_monitoring"] = m.TuntermMonitoring
    }
    if m.TuntermMonitoringDisabled != nil {
        structMap["tunterm_monitoring_disabled"] = m.TuntermMonitoringDisabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Mxcluster.
// It customizes the JSON unmarshaling process for Mxcluster objects.
func (m *Mxcluster) UnmarshalJSON(input []byte) error {
    var temp tempMxcluster
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "for_site", "id", "mist_das", "mist_nac", "modified_time", "mxedge_mgmt", "name", "org_id", "proxy", "radsec", "radsec_tls", "site_id", "tunterm_ap_subnets", "tunterm_dhcpd_config", "tunterm_extra_routes", "tunterm_hosts", "tunterm_hosts_order", "tunterm_hosts_selection", "tunterm_monitoring", "tunterm_monitoring_disabled")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.CreatedTime = temp.CreatedTime
    m.ForSite = temp.ForSite
    m.Id = temp.Id
    m.MistDas = temp.MistDas
    m.MistNac = temp.MistNac
    m.ModifiedTime = temp.ModifiedTime
    m.MxedgeMgmt = temp.MxedgeMgmt
    m.Name = temp.Name
    m.OrgId = temp.OrgId
    m.Proxy = temp.Proxy
    m.Radsec = temp.Radsec
    m.RadsecTls = temp.RadsecTls
    m.SiteId = temp.SiteId
    m.TuntermApSubnets = temp.TuntermApSubnets
    m.TuntermDhcpdConfig = temp.TuntermDhcpdConfig
    m.TuntermExtraRoutes = temp.TuntermExtraRoutes
    m.TuntermHosts = temp.TuntermHosts
    m.TuntermHostsOrder = temp.TuntermHostsOrder
    m.TuntermHostsSelection = temp.TuntermHostsSelection
    m.TuntermMonitoring = temp.TuntermMonitoring
    m.TuntermMonitoringDisabled = temp.TuntermMonitoringDisabled
    return nil
}

// tempMxcluster is a temporary struct used for validating the fields of Mxcluster.
type tempMxcluster  struct {
    CreatedTime               *float64                              `json:"created_time,omitempty"`
    ForSite                   *bool                                 `json:"for_site,omitempty"`
    Id                        *uuid.UUID                            `json:"id,omitempty"`
    MistDas                   *MxedgeDas                            `json:"mist_das,omitempty"`
    MistNac                   *MxclusterNac                         `json:"mist_nac,omitempty"`
    ModifiedTime              *float64                              `json:"modified_time,omitempty"`
    MxedgeMgmt                *MxedgeMgmt                           `json:"mxedge_mgmt,omitempty"`
    Name                      *string                               `json:"name,omitempty"`
    OrgId                     *uuid.UUID                            `json:"org_id,omitempty"`
    Proxy                     *Proxy                                `json:"proxy,omitempty"`
    Radsec                    *MxclusterRadsec                      `json:"radsec,omitempty"`
    RadsecTls                 *MxclusterRadsecTls                   `json:"radsec_tls,omitempty"`
    SiteId                    *uuid.UUID                            `json:"site_id,omitempty"`
    TuntermApSubnets          []string                              `json:"tunterm_ap_subnets,omitempty"`
    TuntermDhcpdConfig        *TuntermDhcpdConfig                   `json:"tunterm_dhcpd_config,omitempty"`
    TuntermExtraRoutes        map[string]MxclusterTuntermExtraRoute `json:"tunterm_extra_routes,omitempty"`
    TuntermHosts              []string                              `json:"tunterm_hosts,omitempty"`
    TuntermHostsOrder         []int                                 `json:"tunterm_hosts_order,omitempty"`
    TuntermHostsSelection     *MxclusterTuntermHostsSelectionEnum   `json:"tunterm_hosts_selection,omitempty"`
    TuntermMonitoring         [][]TuntermMonitoringItem             `json:"tunterm_monitoring,omitempty"`
    TuntermMonitoringDisabled *bool                                 `json:"tunterm_monitoring_disabled,omitempty"`
}
