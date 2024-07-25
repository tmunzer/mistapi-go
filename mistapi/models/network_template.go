package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// NetworkTemplate represents a NetworkTemplate struct.
// Network Template
type NetworkTemplate struct {
    AclPolicies          []AclPolicy                            `json:"acl_policies,omitempty"`
    // ACL Tags to identify traffic source or destination. Key name is the tag name
    AclTags              map[string]AclTag                      `json:"acl_tags,omitempty"`
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    AdditionalConfigCmds []string                               `json:"additional_config_cmds,omitempty"`
    CreatedTime          *float64                               `json:"created_time,omitempty"`
    DhcpSnooping         *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsServers           []string                               `json:"dns_servers,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsSuffix            []string                               `json:"dns_suffix,omitempty"`
    ExtraRoutes          map[string]ExtraRoute                  `json:"extra_routes,omitempty"`
    // Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
    ExtraRoutes6         map[string]ExtraRoute6                 `json:"extra_routes6,omitempty"`
    Id                   *uuid.UUID                             `json:"id,omitempty"`
    // Org Networks that we'd like to import
    ImportOrgNetworks    []string                               `json:"import_org_networks,omitempty"`
    // enable mist_nac to use radsec
    MistNac              *SwitchMistNac                         `json:"mist_nac,omitempty"`
    ModifiedTime         *float64                               `json:"modified_time,omitempty"`
    Name                 *string                                `json:"name,omitempty"`
    // Property key is network name
    Networks             map[string]SwitchNetwork               `json:"networks,omitempty"`
    // list of NTP servers specific to this device. By default, those in Site Settings will be used
    NtpServers           []string                               `json:"ntp_servers,omitempty"`
    OrgId                *uuid.UUID                             `json:"org_id,omitempty"`
    // Property key is the port mirroring instance name
    // port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.
    PortMirroring        map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    PortUsages           map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
    // Junos Radius config
    RadiusConfig         *RadiusConfig                          `json:"radius_config,omitempty"`
    RemoteSyslog         *RemoteSyslog                          `json:"remote_syslog,omitempty"`
    SnmpConfig           *SnmpConfig                            `json:"snmp_config,omitempty"`
    // Switch template
    SwitchMatching       *SwitchMatching                        `json:"switch_matching,omitempty"`
    SwitchMgmt           *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    VrfConfig            *VrfConfig                             `json:"vrf_config,omitempty"`
    // Property key is the network name
    VrfInstances         map[string]VrfInstance                 `json:"vrf_instances,omitempty"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NetworkTemplate.
// It customizes the JSON marshaling process for NetworkTemplate objects.
func (n NetworkTemplate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NetworkTemplate object to a map representation for JSON marshaling.
func (n NetworkTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.AclPolicies != nil {
        structMap["acl_policies"] = n.AclPolicies
    }
    if n.AclTags != nil {
        structMap["acl_tags"] = n.AclTags
    }
    if n.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = n.AdditionalConfigCmds
    }
    if n.CreatedTime != nil {
        structMap["created_time"] = n.CreatedTime
    }
    if n.DhcpSnooping != nil {
        structMap["dhcp_snooping"] = n.DhcpSnooping.toMap()
    }
    if n.DnsServers != nil {
        structMap["dns_servers"] = n.DnsServers
    }
    if n.DnsSuffix != nil {
        structMap["dns_suffix"] = n.DnsSuffix
    }
    if n.ExtraRoutes != nil {
        structMap["extra_routes"] = n.ExtraRoutes
    }
    if n.ExtraRoutes6 != nil {
        structMap["extra_routes6"] = n.ExtraRoutes6
    }
    if n.Id != nil {
        structMap["id"] = n.Id
    }
    if n.ImportOrgNetworks != nil {
        structMap["import_org_networks"] = n.ImportOrgNetworks
    }
    if n.MistNac != nil {
        structMap["mist_nac"] = n.MistNac.toMap()
    }
    if n.ModifiedTime != nil {
        structMap["modified_time"] = n.ModifiedTime
    }
    if n.Name != nil {
        structMap["name"] = n.Name
    }
    if n.Networks != nil {
        structMap["networks"] = n.Networks
    }
    if n.NtpServers != nil {
        structMap["ntp_servers"] = n.NtpServers
    }
    if n.OrgId != nil {
        structMap["org_id"] = n.OrgId
    }
    if n.PortMirroring != nil {
        structMap["port_mirroring"] = n.PortMirroring
    }
    if n.PortUsages != nil {
        structMap["port_usages"] = n.PortUsages
    }
    if n.RadiusConfig != nil {
        structMap["radius_config"] = n.RadiusConfig.toMap()
    }
    if n.RemoteSyslog != nil {
        structMap["remote_syslog"] = n.RemoteSyslog.toMap()
    }
    if n.SnmpConfig != nil {
        structMap["snmp_config"] = n.SnmpConfig.toMap()
    }
    if n.SwitchMatching != nil {
        structMap["switch_matching"] = n.SwitchMatching.toMap()
    }
    if n.SwitchMgmt != nil {
        structMap["switch_mgmt"] = n.SwitchMgmt.toMap()
    }
    if n.VrfConfig != nil {
        structMap["vrf_config"] = n.VrfConfig.toMap()
    }
    if n.VrfInstances != nil {
        structMap["vrf_instances"] = n.VrfInstances
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NetworkTemplate.
// It customizes the JSON unmarshaling process for NetworkTemplate objects.
func (n *NetworkTemplate) UnmarshalJSON(input []byte) error {
    var temp networkTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "acl_policies", "acl_tags", "additional_config_cmds", "created_time", "dhcp_snooping", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "id", "import_org_networks", "mist_nac", "modified_time", "name", "networks", "ntp_servers", "org_id", "port_mirroring", "port_usages", "radius_config", "remote_syslog", "snmp_config", "switch_matching", "switch_mgmt", "vrf_config", "vrf_instances")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.AclPolicies = temp.AclPolicies
    n.AclTags = temp.AclTags
    n.AdditionalConfigCmds = temp.AdditionalConfigCmds
    n.CreatedTime = temp.CreatedTime
    n.DhcpSnooping = temp.DhcpSnooping
    n.DnsServers = temp.DnsServers
    n.DnsSuffix = temp.DnsSuffix
    n.ExtraRoutes = temp.ExtraRoutes
    n.ExtraRoutes6 = temp.ExtraRoutes6
    n.Id = temp.Id
    n.ImportOrgNetworks = temp.ImportOrgNetworks
    n.MistNac = temp.MistNac
    n.ModifiedTime = temp.ModifiedTime
    n.Name = temp.Name
    n.Networks = temp.Networks
    n.NtpServers = temp.NtpServers
    n.OrgId = temp.OrgId
    n.PortMirroring = temp.PortMirroring
    n.PortUsages = temp.PortUsages
    n.RadiusConfig = temp.RadiusConfig
    n.RemoteSyslog = temp.RemoteSyslog
    n.SnmpConfig = temp.SnmpConfig
    n.SwitchMatching = temp.SwitchMatching
    n.SwitchMgmt = temp.SwitchMgmt
    n.VrfConfig = temp.VrfConfig
    n.VrfInstances = temp.VrfInstances
    return nil
}

// networkTemplate is a temporary struct used for validating the fields of NetworkTemplate.
type networkTemplate  struct {
    AclPolicies          []AclPolicy                            `json:"acl_policies,omitempty"`
    AclTags              map[string]AclTag                      `json:"acl_tags,omitempty"`
    AdditionalConfigCmds []string                               `json:"additional_config_cmds,omitempty"`
    CreatedTime          *float64                               `json:"created_time,omitempty"`
    DhcpSnooping         *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
    DnsServers           []string                               `json:"dns_servers,omitempty"`
    DnsSuffix            []string                               `json:"dns_suffix,omitempty"`
    ExtraRoutes          map[string]ExtraRoute                  `json:"extra_routes,omitempty"`
    ExtraRoutes6         map[string]ExtraRoute6                 `json:"extra_routes6,omitempty"`
    Id                   *uuid.UUID                             `json:"id,omitempty"`
    ImportOrgNetworks    []string                               `json:"import_org_networks,omitempty"`
    MistNac              *SwitchMistNac                         `json:"mist_nac,omitempty"`
    ModifiedTime         *float64                               `json:"modified_time,omitempty"`
    Name                 *string                                `json:"name,omitempty"`
    Networks             map[string]SwitchNetwork               `json:"networks,omitempty"`
    NtpServers           []string                               `json:"ntp_servers,omitempty"`
    OrgId                *uuid.UUID                             `json:"org_id,omitempty"`
    PortMirroring        map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    PortUsages           map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
    RadiusConfig         *RadiusConfig                          `json:"radius_config,omitempty"`
    RemoteSyslog         *RemoteSyslog                          `json:"remote_syslog,omitempty"`
    SnmpConfig           *SnmpConfig                            `json:"snmp_config,omitempty"`
    SwitchMatching       *SwitchMatching                        `json:"switch_matching,omitempty"`
    SwitchMgmt           *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    VrfConfig            *VrfConfig                             `json:"vrf_config,omitempty"`
    VrfInstances         map[string]VrfInstance                 `json:"vrf_instances,omitempty"`
}
