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

// DeviceSwitch represents a DeviceSwitch struct.
// You can configure `port_usages` and `networks` settings at the device level, but most of the time it's better use the Site Setting to achieve better consistency and be able to re-use the same settings across switches entries defined here will "replace" those defined in Site Setting/Network Template
type DeviceSwitch struct {
    AclPolicies           []AclPolicy                            `json:"acl_policies,omitempty"`
    // ACL Tags to identify traffic source or destination. Key name is the tag name
    AclTags               map[string]AclTag                      `json:"acl_tags,omitempty"`
    // additional CLI commands to append to the generated Junos config. **Note**: no check is done
    AdditionalConfigCmds  []string                               `json:"additional_config_cmds,omitempty"`
    // Property key is the destination subnet (e.g. "172.16.3.0/24")
    AggregateRoutes       map[string]AggregateRoute              `json:"aggregate_routes,omitempty"`
    // Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64")
    AggregateRoutes6      map[string]AggregateRoute              `json:"aggregate_routes6,omitempty"`
    // When the object has been created, in epoch
    CreatedTime           *float64                               `json:"created_time,omitempty"`
    DeviceprofileId       *uuid.UUID                             `json:"deviceprofile_id,omitempty"`
    DhcpSnooping          *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
    DhcpdConfig           *SwitchDhcpdConfig                     `json:"dhcpd_config,omitempty"`
    // This disables the default behavior of a cloud-ready switch/gateway being managed/configured by Mist. Setting this to `true` means you want to disable the default behavior and do not want the device to be Mist-managed.
    DisableAutoConfig     *bool                                  `json:"disable_auto_config,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsServers            []string                               `json:"dns_servers,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsSuffix             []string                               `json:"dns_suffix,omitempty"`
    // EVPN Junos settings
    EvpnConfig            *EvpnConfig                            `json:"evpn_config,omitempty"`
    // Property key is the destination CIDR (e.g. "10.0.0.0/8")
    ExtraRoutes           map[string]ExtraRoute                  `json:"extra_routes,omitempty"`
    // Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
    ExtraRoutes6          map[string]ExtraRoute6                 `json:"extra_routes6,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                    *uuid.UUID                             `json:"id,omitempty"`
    Image1Url             Optional[string]                       `json:"image1_url"`
    Image2Url             Optional[string]                       `json:"image2_url"`
    Image3Url             Optional[string]                       `json:"image3_url"`
    // Property Key is the IOT port name, e.g.:
    // * `IN0` or `IN1` for the FPC0 input port with 5V triggered inputs
    // * `OUT1` for the FPC0 output port (can only be triggered by either IN0 or IN1)
    // * "X/IN0`, `X/IN1` and `X/OUT` are used to define IOT ports on VC members
    IotConfig             map[string]SwitchIotPort               `json:"iot_config,omitempty"`
    // Junos IP Config
    IpConfig              *JunosIpConfig                         `json:"ip_config,omitempty"`
    // Local port override, overriding the port configuration from `port_config`. Property key is the port name or range (e.g. "ge-0/0/0-10")
    LocalPortConfig       map[string]JunosLocalPortConfig        `json:"local_port_config,omitempty"`
    // Device MAC address
    Mac                   *string                                `json:"mac,omitempty"`
    // An adopted switch/gateway will not be managed/configured by Mist by default. Setting this parameter to `true` enables the adopted switch/gateway to be managed/configured by Mist.
    Managed               *bool                                  `json:"managed,omitempty"`
    // Map where the device belongs to
    MapId                 *uuid.UUID                             `json:"map_id,omitempty"`
    // Enable mist_nac to use RadSec
    MistNac               *SwitchMistNac                         `json:"mist_nac,omitempty"`
    // Device Model
    Model                 *string                                `json:"model,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime          *float64                               `json:"modified_time,omitempty"`
    Name                  *string                                `json:"name,omitempty"`
    // Property key is network name
    Networks              map[string]SwitchNetwork               `json:"networks,omitempty"`
    Notes                 *string                                `json:"notes,omitempty"`
    // List of NTP servers specific to this device. By default, those in Site Settings will be used
    NtpServers            []string                               `json:"ntp_servers,omitempty"`
    // Switch OOB IP Config:
    // - If HA configuration: key parameter will be nodeX (eg: node1)
    // - If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1`
    OobIpConfig           *SwitchOobIpConfig                     `json:"oob_ip_config,omitempty"`
    OrgId                 *uuid.UUID                             `json:"org_id,omitempty"`
    // Junos OSPF areas. Property key is the OSPF Area (Area should be a number (0-255) / IP address)
    OspfAreas             map[string]OspfArea                    `json:"ospf_areas,omitempty"`
    OspfConfig            *SwitchOspfConfig                      `json:"ospf_config,omitempty"`
    // Property key is the network name. Defines the additional IP Addresses configured on the device.
    OtherIpConfigs        map[string]JunosOtherIpConfig          `json:"other_ip_configs,omitempty"`
    // Property key is the port name or range (e.g. "ge-0/0/0-10")
    PortConfig            map[string]JunosPortConfig             `json:"port_config,omitempty"`
    // Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed
    PortMirroring         map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    // Property key is the port usage name. Defines the profiles of port configuration configured on the switch
    PortUsages            map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
    // Junos Radius config
    RadiusConfig          *SwitchRadiusConfig                    `json:"radius_config,omitempty"`
    RemoteSyslog          *RemoteSyslog                          `json:"remote_syslog,omitempty"`
    Role                  *string                                `json:"role,omitempty"`
    // Used for OSPF / BGP / EVPN
    RouterId              *string                                `json:"router_id,omitempty"`
    // Device Serial
    Serial                *string                                `json:"serial,omitempty"`
    SiteId                *uuid.UUID                             `json:"site_id,omitempty"`
    SnmpConfig            *SnmpConfig                            `json:"snmp_config,omitempty"`
    StpConfig             *SwitchStpConfig                       `json:"stp_config,omitempty"`
    // Switch settings
    SwitchMgmt            *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    // Device Type. enum: `switch`
    Type                  string                                 `json:"type"`
    // Whether to use it for snmp / syslog / tacplus / radius
    UseRouterIdAsSourceIp *bool                                  `json:"use_router_id_as_source_ip,omitempty"`
    // Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                  map[string]string                      `json:"vars,omitempty"`
    // Required for preprovisioned Virtual Chassis
    VirtualChassis        *SwitchVirtualChassis                  `json:"virtual_chassis,omitempty"`
    VrfConfig             *VrfConfig                             `json:"vrf_config,omitempty"`
    // Property key is the network name
    VrfInstances          map[string]SwitchVrfInstance           `json:"vrf_instances,omitempty"`
    // Junos VRRP config
    VrrpConfig            *VrrpConfig                            `json:"vrrp_config,omitempty"`
    // X in pixel
    X                     *float64                               `json:"x,omitempty"`
    // Y in pixel
    Y                     *float64                               `json:"y,omitempty"`
    AdditionalProperties  map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceSwitch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceSwitch) String() string {
    return fmt.Sprintf(
    	"DeviceSwitch[AclPolicies=%v, AclTags=%v, AdditionalConfigCmds=%v, AggregateRoutes=%v, AggregateRoutes6=%v, CreatedTime=%v, DeviceprofileId=%v, DhcpSnooping=%v, DhcpdConfig=%v, DisableAutoConfig=%v, DnsServers=%v, DnsSuffix=%v, EvpnConfig=%v, ExtraRoutes=%v, ExtraRoutes6=%v, Id=%v, Image1Url=%v, Image2Url=%v, Image3Url=%v, IotConfig=%v, IpConfig=%v, LocalPortConfig=%v, Mac=%v, Managed=%v, MapId=%v, MistNac=%v, Model=%v, ModifiedTime=%v, Name=%v, Networks=%v, Notes=%v, NtpServers=%v, OobIpConfig=%v, OrgId=%v, OspfAreas=%v, OspfConfig=%v, OtherIpConfigs=%v, PortConfig=%v, PortMirroring=%v, PortUsages=%v, RadiusConfig=%v, RemoteSyslog=%v, Role=%v, RouterId=%v, Serial=%v, SiteId=%v, SnmpConfig=%v, StpConfig=%v, SwitchMgmt=%v, Type=%v, UseRouterIdAsSourceIp=%v, Vars=%v, VirtualChassis=%v, VrfConfig=%v, VrfInstances=%v, VrrpConfig=%v, X=%v, Y=%v, AdditionalProperties=%v]",
    	d.AclPolicies, d.AclTags, d.AdditionalConfigCmds, d.AggregateRoutes, d.AggregateRoutes6, d.CreatedTime, d.DeviceprofileId, d.DhcpSnooping, d.DhcpdConfig, d.DisableAutoConfig, d.DnsServers, d.DnsSuffix, d.EvpnConfig, d.ExtraRoutes, d.ExtraRoutes6, d.Id, d.Image1Url, d.Image2Url, d.Image3Url, d.IotConfig, d.IpConfig, d.LocalPortConfig, d.Mac, d.Managed, d.MapId, d.MistNac, d.Model, d.ModifiedTime, d.Name, d.Networks, d.Notes, d.NtpServers, d.OobIpConfig, d.OrgId, d.OspfAreas, d.OspfConfig, d.OtherIpConfigs, d.PortConfig, d.PortMirroring, d.PortUsages, d.RadiusConfig, d.RemoteSyslog, d.Role, d.RouterId, d.Serial, d.SiteId, d.SnmpConfig, d.StpConfig, d.SwitchMgmt, d.Type, d.UseRouterIdAsSourceIp, d.Vars, d.VirtualChassis, d.VrfConfig, d.VrfInstances, d.VrrpConfig, d.X, d.Y, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceSwitch.
// It customizes the JSON marshaling process for DeviceSwitch objects.
func (d DeviceSwitch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "acl_policies", "acl_tags", "additional_config_cmds", "aggregate_routes", "aggregate_routes6", "created_time", "deviceprofile_id", "dhcp_snooping", "dhcpd_config", "disable_auto_config", "dns_servers", "dns_suffix", "evpn_config", "extra_routes", "extra_routes6", "id", "image1_url", "image2_url", "image3_url", "iot_config", "ip_config", "local_port_config", "mac", "managed", "map_id", "mist_nac", "model", "modified_time", "name", "networks", "notes", "ntp_servers", "oob_ip_config", "org_id", "ospf_areas", "ospf_config", "other_ip_configs", "port_config", "port_mirroring", "port_usages", "radius_config", "remote_syslog", "role", "router_id", "serial", "site_id", "snmp_config", "stp_config", "switch_mgmt", "type", "use_router_id_as_source_ip", "vars", "virtual_chassis", "vrf_config", "vrf_instances", "vrrp_config", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeviceSwitch object to a map representation for JSON marshaling.
func (d DeviceSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.AclPolicies != nil {
        structMap["acl_policies"] = d.AclPolicies
    }
    if d.AclTags != nil {
        structMap["acl_tags"] = d.AclTags
    }
    if d.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = d.AdditionalConfigCmds
    }
    if d.AggregateRoutes != nil {
        structMap["aggregate_routes"] = d.AggregateRoutes
    }
    if d.AggregateRoutes6 != nil {
        structMap["aggregate_routes6"] = d.AggregateRoutes6
    }
    if d.CreatedTime != nil {
        structMap["created_time"] = d.CreatedTime
    }
    if d.DeviceprofileId != nil {
        structMap["deviceprofile_id"] = d.DeviceprofileId
    }
    if d.DhcpSnooping != nil {
        structMap["dhcp_snooping"] = d.DhcpSnooping.toMap()
    }
    if d.DhcpdConfig != nil {
        structMap["dhcpd_config"] = d.DhcpdConfig.toMap()
    }
    if d.DisableAutoConfig != nil {
        structMap["disable_auto_config"] = d.DisableAutoConfig
    }
    if d.DnsServers != nil {
        structMap["dns_servers"] = d.DnsServers
    }
    if d.DnsSuffix != nil {
        structMap["dns_suffix"] = d.DnsSuffix
    }
    if d.EvpnConfig != nil {
        structMap["evpn_config"] = d.EvpnConfig.toMap()
    }
    if d.ExtraRoutes != nil {
        structMap["extra_routes"] = d.ExtraRoutes
    }
    if d.ExtraRoutes6 != nil {
        structMap["extra_routes6"] = d.ExtraRoutes6
    }
    if d.Id != nil {
        structMap["id"] = d.Id
    }
    if d.Image1Url.IsValueSet() {
        if d.Image1Url.Value() != nil {
            structMap["image1_url"] = d.Image1Url.Value()
        } else {
            structMap["image1_url"] = nil
        }
    }
    if d.Image2Url.IsValueSet() {
        if d.Image2Url.Value() != nil {
            structMap["image2_url"] = d.Image2Url.Value()
        } else {
            structMap["image2_url"] = nil
        }
    }
    if d.Image3Url.IsValueSet() {
        if d.Image3Url.Value() != nil {
            structMap["image3_url"] = d.Image3Url.Value()
        } else {
            structMap["image3_url"] = nil
        }
    }
    if d.IotConfig != nil {
        structMap["iot_config"] = d.IotConfig
    }
    if d.IpConfig != nil {
        structMap["ip_config"] = d.IpConfig.toMap()
    }
    if d.LocalPortConfig != nil {
        structMap["local_port_config"] = d.LocalPortConfig
    }
    if d.Mac != nil {
        structMap["mac"] = d.Mac
    }
    if d.Managed != nil {
        structMap["managed"] = d.Managed
    }
    if d.MapId != nil {
        structMap["map_id"] = d.MapId
    }
    if d.MistNac != nil {
        structMap["mist_nac"] = d.MistNac.toMap()
    }
    if d.Model != nil {
        structMap["model"] = d.Model
    }
    if d.ModifiedTime != nil {
        structMap["modified_time"] = d.ModifiedTime
    }
    if d.Name != nil {
        structMap["name"] = d.Name
    }
    if d.Networks != nil {
        structMap["networks"] = d.Networks
    }
    if d.Notes != nil {
        structMap["notes"] = d.Notes
    }
    if d.NtpServers != nil {
        structMap["ntp_servers"] = d.NtpServers
    }
    if d.OobIpConfig != nil {
        structMap["oob_ip_config"] = d.OobIpConfig.toMap()
    }
    if d.OrgId != nil {
        structMap["org_id"] = d.OrgId
    }
    if d.OspfAreas != nil {
        structMap["ospf_areas"] = d.OspfAreas
    }
    if d.OspfConfig != nil {
        structMap["ospf_config"] = d.OspfConfig.toMap()
    }
    if d.OtherIpConfigs != nil {
        structMap["other_ip_configs"] = d.OtherIpConfigs
    }
    if d.PortConfig != nil {
        structMap["port_config"] = d.PortConfig
    }
    if d.PortMirroring != nil {
        structMap["port_mirroring"] = d.PortMirroring
    }
    if d.PortUsages != nil {
        structMap["port_usages"] = d.PortUsages
    }
    if d.RadiusConfig != nil {
        structMap["radius_config"] = d.RadiusConfig.toMap()
    }
    if d.RemoteSyslog != nil {
        structMap["remote_syslog"] = d.RemoteSyslog.toMap()
    }
    if d.Role != nil {
        structMap["role"] = d.Role
    }
    if d.RouterId != nil {
        structMap["router_id"] = d.RouterId
    }
    if d.Serial != nil {
        structMap["serial"] = d.Serial
    }
    if d.SiteId != nil {
        structMap["site_id"] = d.SiteId
    }
    if d.SnmpConfig != nil {
        structMap["snmp_config"] = d.SnmpConfig.toMap()
    }
    if d.StpConfig != nil {
        structMap["stp_config"] = d.StpConfig.toMap()
    }
    if d.SwitchMgmt != nil {
        structMap["switch_mgmt"] = d.SwitchMgmt.toMap()
    }
    structMap["type"] = d.Type
    if d.UseRouterIdAsSourceIp != nil {
        structMap["use_router_id_as_source_ip"] = d.UseRouterIdAsSourceIp
    }
    if d.Vars != nil {
        structMap["vars"] = d.Vars
    }
    if d.VirtualChassis != nil {
        structMap["virtual_chassis"] = d.VirtualChassis.toMap()
    }
    if d.VrfConfig != nil {
        structMap["vrf_config"] = d.VrfConfig.toMap()
    }
    if d.VrfInstances != nil {
        structMap["vrf_instances"] = d.VrfInstances
    }
    if d.VrrpConfig != nil {
        structMap["vrrp_config"] = d.VrrpConfig.toMap()
    }
    if d.X != nil {
        structMap["x"] = d.X
    }
    if d.Y != nil {
        structMap["y"] = d.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceSwitch.
// It customizes the JSON unmarshaling process for DeviceSwitch objects.
func (d *DeviceSwitch) UnmarshalJSON(input []byte) error {
    var temp tempDeviceSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acl_policies", "acl_tags", "additional_config_cmds", "aggregate_routes", "aggregate_routes6", "created_time", "deviceprofile_id", "dhcp_snooping", "dhcpd_config", "disable_auto_config", "dns_servers", "dns_suffix", "evpn_config", "extra_routes", "extra_routes6", "id", "image1_url", "image2_url", "image3_url", "iot_config", "ip_config", "local_port_config", "mac", "managed", "map_id", "mist_nac", "model", "modified_time", "name", "networks", "notes", "ntp_servers", "oob_ip_config", "org_id", "ospf_areas", "ospf_config", "other_ip_configs", "port_config", "port_mirroring", "port_usages", "radius_config", "remote_syslog", "role", "router_id", "serial", "site_id", "snmp_config", "stp_config", "switch_mgmt", "type", "use_router_id_as_source_ip", "vars", "virtual_chassis", "vrf_config", "vrf_instances", "vrrp_config", "x", "y")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.AclPolicies = temp.AclPolicies
    d.AclTags = temp.AclTags
    d.AdditionalConfigCmds = temp.AdditionalConfigCmds
    d.AggregateRoutes = temp.AggregateRoutes
    d.AggregateRoutes6 = temp.AggregateRoutes6
    d.CreatedTime = temp.CreatedTime
    d.DeviceprofileId = temp.DeviceprofileId
    d.DhcpSnooping = temp.DhcpSnooping
    d.DhcpdConfig = temp.DhcpdConfig
    d.DisableAutoConfig = temp.DisableAutoConfig
    d.DnsServers = temp.DnsServers
    d.DnsSuffix = temp.DnsSuffix
    d.EvpnConfig = temp.EvpnConfig
    d.ExtraRoutes = temp.ExtraRoutes
    d.ExtraRoutes6 = temp.ExtraRoutes6
    d.Id = temp.Id
    d.Image1Url = temp.Image1Url
    d.Image2Url = temp.Image2Url
    d.Image3Url = temp.Image3Url
    d.IotConfig = temp.IotConfig
    d.IpConfig = temp.IpConfig
    d.LocalPortConfig = temp.LocalPortConfig
    d.Mac = temp.Mac
    d.Managed = temp.Managed
    d.MapId = temp.MapId
    d.MistNac = temp.MistNac
    d.Model = temp.Model
    d.ModifiedTime = temp.ModifiedTime
    d.Name = temp.Name
    d.Networks = temp.Networks
    d.Notes = temp.Notes
    d.NtpServers = temp.NtpServers
    d.OobIpConfig = temp.OobIpConfig
    d.OrgId = temp.OrgId
    d.OspfAreas = temp.OspfAreas
    d.OspfConfig = temp.OspfConfig
    d.OtherIpConfigs = temp.OtherIpConfigs
    d.PortConfig = temp.PortConfig
    d.PortMirroring = temp.PortMirroring
    d.PortUsages = temp.PortUsages
    d.RadiusConfig = temp.RadiusConfig
    d.RemoteSyslog = temp.RemoteSyslog
    d.Role = temp.Role
    d.RouterId = temp.RouterId
    d.Serial = temp.Serial
    d.SiteId = temp.SiteId
    d.SnmpConfig = temp.SnmpConfig
    d.StpConfig = temp.StpConfig
    d.SwitchMgmt = temp.SwitchMgmt
    d.Type = *temp.Type
    d.UseRouterIdAsSourceIp = temp.UseRouterIdAsSourceIp
    d.Vars = temp.Vars
    d.VirtualChassis = temp.VirtualChassis
    d.VrfConfig = temp.VrfConfig
    d.VrfInstances = temp.VrfInstances
    d.VrrpConfig = temp.VrrpConfig
    d.X = temp.X
    d.Y = temp.Y
    return nil
}

// tempDeviceSwitch is a temporary struct used for validating the fields of DeviceSwitch.
type tempDeviceSwitch  struct {
    AclPolicies           []AclPolicy                            `json:"acl_policies,omitempty"`
    AclTags               map[string]AclTag                      `json:"acl_tags,omitempty"`
    AdditionalConfigCmds  []string                               `json:"additional_config_cmds,omitempty"`
    AggregateRoutes       map[string]AggregateRoute              `json:"aggregate_routes,omitempty"`
    AggregateRoutes6      map[string]AggregateRoute              `json:"aggregate_routes6,omitempty"`
    CreatedTime           *float64                               `json:"created_time,omitempty"`
    DeviceprofileId       *uuid.UUID                             `json:"deviceprofile_id,omitempty"`
    DhcpSnooping          *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
    DhcpdConfig           *SwitchDhcpdConfig                     `json:"dhcpd_config,omitempty"`
    DisableAutoConfig     *bool                                  `json:"disable_auto_config,omitempty"`
    DnsServers            []string                               `json:"dns_servers,omitempty"`
    DnsSuffix             []string                               `json:"dns_suffix,omitempty"`
    EvpnConfig            *EvpnConfig                            `json:"evpn_config,omitempty"`
    ExtraRoutes           map[string]ExtraRoute                  `json:"extra_routes,omitempty"`
    ExtraRoutes6          map[string]ExtraRoute6                 `json:"extra_routes6,omitempty"`
    Id                    *uuid.UUID                             `json:"id,omitempty"`
    Image1Url             Optional[string]                       `json:"image1_url"`
    Image2Url             Optional[string]                       `json:"image2_url"`
    Image3Url             Optional[string]                       `json:"image3_url"`
    IotConfig             map[string]SwitchIotPort               `json:"iot_config,omitempty"`
    IpConfig              *JunosIpConfig                         `json:"ip_config,omitempty"`
    LocalPortConfig       map[string]JunosLocalPortConfig        `json:"local_port_config,omitempty"`
    Mac                   *string                                `json:"mac,omitempty"`
    Managed               *bool                                  `json:"managed,omitempty"`
    MapId                 *uuid.UUID                             `json:"map_id,omitempty"`
    MistNac               *SwitchMistNac                         `json:"mist_nac,omitempty"`
    Model                 *string                                `json:"model,omitempty"`
    ModifiedTime          *float64                               `json:"modified_time,omitempty"`
    Name                  *string                                `json:"name,omitempty"`
    Networks              map[string]SwitchNetwork               `json:"networks,omitempty"`
    Notes                 *string                                `json:"notes,omitempty"`
    NtpServers            []string                               `json:"ntp_servers,omitempty"`
    OobIpConfig           *SwitchOobIpConfig                     `json:"oob_ip_config,omitempty"`
    OrgId                 *uuid.UUID                             `json:"org_id,omitempty"`
    OspfAreas             map[string]OspfArea                    `json:"ospf_areas,omitempty"`
    OspfConfig            *SwitchOspfConfig                      `json:"ospf_config,omitempty"`
    OtherIpConfigs        map[string]JunosOtherIpConfig          `json:"other_ip_configs,omitempty"`
    PortConfig            map[string]JunosPortConfig             `json:"port_config,omitempty"`
    PortMirroring         map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    PortUsages            map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
    RadiusConfig          *SwitchRadiusConfig                    `json:"radius_config,omitempty"`
    RemoteSyslog          *RemoteSyslog                          `json:"remote_syslog,omitempty"`
    Role                  *string                                `json:"role,omitempty"`
    RouterId              *string                                `json:"router_id,omitempty"`
    Serial                *string                                `json:"serial,omitempty"`
    SiteId                *uuid.UUID                             `json:"site_id,omitempty"`
    SnmpConfig            *SnmpConfig                            `json:"snmp_config,omitempty"`
    StpConfig             *SwitchStpConfig                       `json:"stp_config,omitempty"`
    SwitchMgmt            *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    Type                  *string                                `json:"type"`
    UseRouterIdAsSourceIp *bool                                  `json:"use_router_id_as_source_ip,omitempty"`
    Vars                  map[string]string                      `json:"vars,omitempty"`
    VirtualChassis        *SwitchVirtualChassis                  `json:"virtual_chassis,omitempty"`
    VrfConfig             *VrfConfig                             `json:"vrf_config,omitempty"`
    VrfInstances          map[string]SwitchVrfInstance           `json:"vrf_instances,omitempty"`
    VrrpConfig            *VrrpConfig                            `json:"vrrp_config,omitempty"`
    X                     *float64                               `json:"x,omitempty"`
    Y                     *float64                               `json:"y,omitempty"`
}

func (d *tempDeviceSwitch) validate() error {
    var errs []string
    if d.Type == nil {
        errs = append(errs, "required field `type` is missing for type `device_switch`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
