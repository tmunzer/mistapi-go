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

// DeviceprofileSwitch represents a DeviceprofileSwitch struct.
// Switch Device Profiles can be applied to one or multiple switches. The settings from the Device Profile will override the settings from the Switch Template and the Site Settings.
type DeviceprofileSwitch struct {
	AclPolicies []AclPolicy `json:"acl_policies,omitempty"`
	// ACL Tags to identify traffic source or destination. Key name is the tag name
	AclTags map[string]AclTag `json:"acl_tags,omitempty"`
	// additional CLI commands to append to the generated Junos config. **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	// Property key is the destination subnet (e.g. "172.16.3.0/24")
	AggregateRoutes map[string]AggregateRoute `json:"aggregate_routes,omitempty"`
	// Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64")
	AggregateRoutes6 map[string]AggregateRoute `json:"aggregate_routes6,omitempty"`
	// When the object has been created, in epoch
	CreatedTime  *float64           `json:"created_time,omitempty"`
	DhcpSnooping *DhcpSnooping      `json:"dhcp_snooping,omitempty"`
	DhcpdConfig  *SwitchDhcpdConfig `json:"dhcpd_config,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// EVPN Junos settings
	EvpnConfig *EvpnConfig `json:"evpn_config,omitempty"`
	// Property key is the destination CIDR (e.g. "10.0.0.0/8")
	ExtraRoutes map[string]ExtraRoute `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
	ExtraRoutes6 map[string]ExtraRoute6 `json:"extra_routes6,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// Property Key is the IOT port name, e.g.:
	// * `IN0` or `IN1` for the FPC0 input port with 5V triggered inputs
	// * `OUT1` for the FPC0 output port (can only be triggered by either IN0 or IN1)
	// * "X/IN0`, `X/IN1` and `X/OUT` are used to define IOT ports on VC members
	IotConfig map[string]SwitchIotPort `json:"iot_config,omitempty"`
	// Junos IP Config
	IpConfig *JunosIpConfig `json:"ip_config,omitempty"`
	// Enable mist_nac to use RadSec
	MistNac *SwitchMistNac `json:"mist_nac,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	Name         string   `json:"name"`
	// Property key is network name
	Networks map[string]SwitchNetwork `json:"networks,omitempty"`
	// List of NTP servers specific to this device. By default, those in Site Settings will be used
	NtpServers []string `json:"ntp_servers,omitempty"`
	// Switch OOB IP Config:
	// - If HA configuration: key parameter will be nodeX (eg: node1)
	// - If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1`
	OobIpConfig *SwitchOobIpConfig `json:"oob_ip_config,omitempty"`
	OrgId       *uuid.UUID         `json:"org_id,omitempty"`
	// Junos OSPF areas. Property key is the OSPF Area (Area should be a number (0-255) / IP address)
	OspfAreas map[string]OspfArea `json:"ospf_areas,omitempty"`
	// Property key is the network name. Defines the additional IP Addresses configured on the device.
	OtherIpConfigs map[string]JunosOtherIpConfig `json:"other_ip_configs,omitempty"`
	// Property key is the port name or range (e.g. "ge-0/0/0-10")
	PortConfig map[string]JunosPortConfig `json:"port_config,omitempty"`
	// Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed
	PortMirroring map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
	// Property key is the port usage name. Defines the profiles of port configuration configured on the switch
	PortUsages map[string]SwitchPortUsage `json:"port_usages,omitempty"`
	// Junos Radius config
	RadiusConfig *SwitchRadiusConfig `json:"radius_config,omitempty"`
	RemoteSyslog *RemoteSyslog       `json:"remote_syslog,omitempty"`
	// Property key is the routing policy name
	RoutingPolicies map[string]SwRoutingPolicy `json:"routing_policies,omitempty"`
	SiteId          *uuid.UUID                 `json:"site_id,omitempty"`
	SnmpConfig      *SnmpConfig                `json:"snmp_config,omitempty"`
	StpConfig       *SwitchStpConfig           `json:"stp_config,omitempty"`
	// Switch settings
	SwitchMgmt *SwitchMgmt `json:"switch_mgmt,omitempty"`
	// Device Type. enum: `switch`
	Type string `json:"type"`
	// Whether to use it for snmp / syslog / tacplus / radius
	UseRouterIdAsSourceIp *bool      `json:"use_router_id_as_source_ip,omitempty"`
	VrfConfig             *VrfConfig `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances map[string]SwitchVrfInstance `json:"vrf_instances,omitempty"`
	// Junos VRRP config
	VrrpConfig           *VrrpConfig            `json:"vrrp_config,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeviceprofileSwitch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeviceprofileSwitch) String() string {
	return fmt.Sprintf(
		"DeviceprofileSwitch[AclPolicies=%v, AclTags=%v, AdditionalConfigCmds=%v, AggregateRoutes=%v, AggregateRoutes6=%v, CreatedTime=%v, DhcpSnooping=%v, DhcpdConfig=%v, DnsServers=%v, DnsSuffix=%v, EvpnConfig=%v, ExtraRoutes=%v, ExtraRoutes6=%v, Id=%v, IotConfig=%v, IpConfig=%v, MistNac=%v, ModifiedTime=%v, Name=%v, Networks=%v, NtpServers=%v, OobIpConfig=%v, OrgId=%v, OspfAreas=%v, OtherIpConfigs=%v, PortConfig=%v, PortMirroring=%v, PortUsages=%v, RadiusConfig=%v, RemoteSyslog=%v, RoutingPolicies=%v, SiteId=%v, SnmpConfig=%v, StpConfig=%v, SwitchMgmt=%v, Type=%v, UseRouterIdAsSourceIp=%v, VrfConfig=%v, VrfInstances=%v, VrrpConfig=%v, AdditionalProperties=%v]",
		d.AclPolicies, d.AclTags, d.AdditionalConfigCmds, d.AggregateRoutes, d.AggregateRoutes6, d.CreatedTime, d.DhcpSnooping, d.DhcpdConfig, d.DnsServers, d.DnsSuffix, d.EvpnConfig, d.ExtraRoutes, d.ExtraRoutes6, d.Id, d.IotConfig, d.IpConfig, d.MistNac, d.ModifiedTime, d.Name, d.Networks, d.NtpServers, d.OobIpConfig, d.OrgId, d.OspfAreas, d.OtherIpConfigs, d.PortConfig, d.PortMirroring, d.PortUsages, d.RadiusConfig, d.RemoteSyslog, d.RoutingPolicies, d.SiteId, d.SnmpConfig, d.StpConfig, d.SwitchMgmt, d.Type, d.UseRouterIdAsSourceIp, d.VrfConfig, d.VrfInstances, d.VrrpConfig, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeviceprofileSwitch.
// It customizes the JSON marshaling process for DeviceprofileSwitch objects.
func (d DeviceprofileSwitch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"acl_policies", "acl_tags", "additional_config_cmds", "aggregate_routes", "aggregate_routes6", "created_time", "dhcp_snooping", "dhcpd_config", "dns_servers", "dns_suffix", "evpn_config", "extra_routes", "extra_routes6", "id", "iot_config", "ip_config", "mist_nac", "modified_time", "name", "networks", "ntp_servers", "oob_ip_config", "org_id", "ospf_areas", "other_ip_configs", "port_config", "port_mirroring", "port_usages", "radius_config", "remote_syslog", "routing_policies", "site_id", "snmp_config", "stp_config", "switch_mgmt", "type", "use_router_id_as_source_ip", "vrf_config", "vrf_instances", "vrrp_config"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DeviceprofileSwitch object to a map representation for JSON marshaling.
func (d DeviceprofileSwitch) toMap() map[string]any {
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
	if d.DhcpSnooping != nil {
		structMap["dhcp_snooping"] = d.DhcpSnooping.toMap()
	}
	if d.DhcpdConfig != nil {
		structMap["dhcpd_config"] = d.DhcpdConfig.toMap()
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
	if d.IotConfig != nil {
		structMap["iot_config"] = d.IotConfig
	}
	if d.IpConfig != nil {
		structMap["ip_config"] = d.IpConfig.toMap()
	}
	if d.MistNac != nil {
		structMap["mist_nac"] = d.MistNac.toMap()
	}
	if d.ModifiedTime != nil {
		structMap["modified_time"] = d.ModifiedTime
	}
	structMap["name"] = d.Name
	if d.Networks != nil {
		structMap["networks"] = d.Networks
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
	if d.RoutingPolicies != nil {
		structMap["routing_policies"] = d.RoutingPolicies
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
	if d.VrfConfig != nil {
		structMap["vrf_config"] = d.VrfConfig.toMap()
	}
	if d.VrfInstances != nil {
		structMap["vrf_instances"] = d.VrfInstances
	}
	if d.VrrpConfig != nil {
		structMap["vrrp_config"] = d.VrrpConfig.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceprofileSwitch.
// It customizes the JSON unmarshaling process for DeviceprofileSwitch objects.
func (d *DeviceprofileSwitch) UnmarshalJSON(input []byte) error {
	var temp tempDeviceprofileSwitch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acl_policies", "acl_tags", "additional_config_cmds", "aggregate_routes", "aggregate_routes6", "created_time", "dhcp_snooping", "dhcpd_config", "dns_servers", "dns_suffix", "evpn_config", "extra_routes", "extra_routes6", "id", "iot_config", "ip_config", "mist_nac", "modified_time", "name", "networks", "ntp_servers", "oob_ip_config", "org_id", "ospf_areas", "other_ip_configs", "port_config", "port_mirroring", "port_usages", "radius_config", "remote_syslog", "routing_policies", "site_id", "snmp_config", "stp_config", "switch_mgmt", "type", "use_router_id_as_source_ip", "vrf_config", "vrf_instances", "vrrp_config")
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
	d.DhcpSnooping = temp.DhcpSnooping
	d.DhcpdConfig = temp.DhcpdConfig
	d.DnsServers = temp.DnsServers
	d.DnsSuffix = temp.DnsSuffix
	d.EvpnConfig = temp.EvpnConfig
	d.ExtraRoutes = temp.ExtraRoutes
	d.ExtraRoutes6 = temp.ExtraRoutes6
	d.Id = temp.Id
	d.IotConfig = temp.IotConfig
	d.IpConfig = temp.IpConfig
	d.MistNac = temp.MistNac
	d.ModifiedTime = temp.ModifiedTime
	d.Name = *temp.Name
	d.Networks = temp.Networks
	d.NtpServers = temp.NtpServers
	d.OobIpConfig = temp.OobIpConfig
	d.OrgId = temp.OrgId
	d.OspfAreas = temp.OspfAreas
	d.OtherIpConfigs = temp.OtherIpConfigs
	d.PortConfig = temp.PortConfig
	d.PortMirroring = temp.PortMirroring
	d.PortUsages = temp.PortUsages
	d.RadiusConfig = temp.RadiusConfig
	d.RemoteSyslog = temp.RemoteSyslog
	d.RoutingPolicies = temp.RoutingPolicies
	d.SiteId = temp.SiteId
	d.SnmpConfig = temp.SnmpConfig
	d.StpConfig = temp.StpConfig
	d.SwitchMgmt = temp.SwitchMgmt
	d.Type = *temp.Type
	d.UseRouterIdAsSourceIp = temp.UseRouterIdAsSourceIp
	d.VrfConfig = temp.VrfConfig
	d.VrfInstances = temp.VrfInstances
	d.VrrpConfig = temp.VrrpConfig
	return nil
}

// tempDeviceprofileSwitch is a temporary struct used for validating the fields of DeviceprofileSwitch.
type tempDeviceprofileSwitch struct {
	AclPolicies           []AclPolicy                            `json:"acl_policies,omitempty"`
	AclTags               map[string]AclTag                      `json:"acl_tags,omitempty"`
	AdditionalConfigCmds  []string                               `json:"additional_config_cmds,omitempty"`
	AggregateRoutes       map[string]AggregateRoute              `json:"aggregate_routes,omitempty"`
	AggregateRoutes6      map[string]AggregateRoute              `json:"aggregate_routes6,omitempty"`
	CreatedTime           *float64                               `json:"created_time,omitempty"`
	DhcpSnooping          *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
	DhcpdConfig           *SwitchDhcpdConfig                     `json:"dhcpd_config,omitempty"`
	DnsServers            []string                               `json:"dns_servers,omitempty"`
	DnsSuffix             []string                               `json:"dns_suffix,omitempty"`
	EvpnConfig            *EvpnConfig                            `json:"evpn_config,omitempty"`
	ExtraRoutes           map[string]ExtraRoute                  `json:"extra_routes,omitempty"`
	ExtraRoutes6          map[string]ExtraRoute6                 `json:"extra_routes6,omitempty"`
	Id                    *uuid.UUID                             `json:"id,omitempty"`
	IotConfig             map[string]SwitchIotPort               `json:"iot_config,omitempty"`
	IpConfig              *JunosIpConfig                         `json:"ip_config,omitempty"`
	MistNac               *SwitchMistNac                         `json:"mist_nac,omitempty"`
	ModifiedTime          *float64                               `json:"modified_time,omitempty"`
	Name                  *string                                `json:"name"`
	Networks              map[string]SwitchNetwork               `json:"networks,omitempty"`
	NtpServers            []string                               `json:"ntp_servers,omitempty"`
	OobIpConfig           *SwitchOobIpConfig                     `json:"oob_ip_config,omitempty"`
	OrgId                 *uuid.UUID                             `json:"org_id,omitempty"`
	OspfAreas             map[string]OspfArea                    `json:"ospf_areas,omitempty"`
	OtherIpConfigs        map[string]JunosOtherIpConfig          `json:"other_ip_configs,omitempty"`
	PortConfig            map[string]JunosPortConfig             `json:"port_config,omitempty"`
	PortMirroring         map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
	PortUsages            map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
	RadiusConfig          *SwitchRadiusConfig                    `json:"radius_config,omitempty"`
	RemoteSyslog          *RemoteSyslog                          `json:"remote_syslog,omitempty"`
	RoutingPolicies       map[string]SwRoutingPolicy             `json:"routing_policies,omitempty"`
	SiteId                *uuid.UUID                             `json:"site_id,omitempty"`
	SnmpConfig            *SnmpConfig                            `json:"snmp_config,omitempty"`
	StpConfig             *SwitchStpConfig                       `json:"stp_config,omitempty"`
	SwitchMgmt            *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
	Type                  *string                                `json:"type"`
	UseRouterIdAsSourceIp *bool                                  `json:"use_router_id_as_source_ip,omitempty"`
	VrfConfig             *VrfConfig                             `json:"vrf_config,omitempty"`
	VrfInstances          map[string]SwitchVrfInstance           `json:"vrf_instances,omitempty"`
	VrrpConfig            *VrrpConfig                            `json:"vrrp_config,omitempty"`
}

func (d *tempDeviceprofileSwitch) validate() error {
	var errs []string
	if d.Name == nil {
		errs = append(errs, "required field `name` is missing for type `deviceprofile_switch`")
	}
	if d.Type == nil {
		errs = append(errs, "required field `type` is missing for type `deviceprofile_switch`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
