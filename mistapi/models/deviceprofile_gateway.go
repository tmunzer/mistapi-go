package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// DeviceprofileGateway represents a DeviceprofileGateway struct.
// Gateway Template is applied to a site for gateway(s) in a site.
type DeviceprofileGateway struct {
	// additional CLI commands to append to the generated Junos config
	// **Note**: no check is done
	AdditionalConfigCmds []string             `json:"additional_config_cmds,omitempty"`
	BgpConfig            map[string]BgpConfig `json:"bgp_config,omitempty"`
	CreatedTime          *float64             `json:"created_time,omitempty"`
	DhcpdConfig          *DhcpdConfig         `json:"dhcpd_config,omitempty"`
	DnsOverride          *bool                `json:"dnsOverride,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// Property key is the destination CIDR (e.g. "10.0.0.0/8")
	ExtraRoutes map[string]GatewayExtraRoute `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
	ExtraRoutes6 map[string]GatewayExtraRoute `json:"extra_routes6,omitempty"`
	// Gateway matching
	GatewayMatching *GatewayMatching `json:"gateway_matching,omitempty"`
	Id              *uuid.UUID       `json:"id,omitempty"`
	// Property key is the profile name
	IdpProfiles map[string]IdpProfile `json:"idp_profiles,omitempty"`
	// Property key is the network name
	IpConfigs    map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
	ModifiedTime *float64                           `json:"modified_time,omitempty"`
	Name         string                             `json:"name"`
	Networks     []Network                          `json:"networks,omitempty"`
	NtpOverride  *bool                              `json:"ntpOverride,omitempty"`
	// list of NTP servers specific to this device. By default, those in Site Settings will be used
	NtpServers []string `json:"ntp_servers,omitempty"`
	// out-of-band (vme/em0/fxp0) IP config
	OobIpConfig *GatewayOobIpConfig `json:"oob_ip_config,omitempty"`
	OrgId       *uuid.UUID          `json:"org_id,omitempty"`
	// Property key is the path name
	PathPreferences map[string]GatewayPathPreferences `json:"path_preferences,omitempty"`
	// Property key is the port(s) name or range (e.g. "ge-0/0/0-10")
	PortConfig map[string]GatewayPortConfig `json:"port_config,omitempty"`
	// auto assigned if not set
	RouterId *string `json:"router_id,omitempty"`
	// Property key is the routing policy name
	RoutingPolicies map[string]RoutingPolicy `json:"routing_policies,omitempty"`
	ServicePolicies []ServicePolicy          `json:"service_policies,omitempty"`
	// Property key is the tunnel name
	TunnelConfigs         map[string]TunnelConfigs `json:"tunnel_configs,omitempty"`
	TunnelProviderOptions *TunnelProviderOptions   `json:"tunnel_provider_options,omitempty"`
	// Device Type. enum: `gateway`
	Type      string     `json:"type"`
	VrfConfig *VrfConfig `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances         map[string]GatewayVrfInstance `json:"vrf_instances,omitempty"`
	AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DeviceprofileGateway.
// It customizes the JSON marshaling process for DeviceprofileGateway objects.
func (d DeviceprofileGateway) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(d.toMap())
}

// toMap converts the DeviceprofileGateway object to a map representation for JSON marshaling.
func (d DeviceprofileGateway) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, d.AdditionalProperties)
	if d.AdditionalConfigCmds != nil {
		structMap["additional_config_cmds"] = d.AdditionalConfigCmds
	}
	if d.BgpConfig != nil {
		structMap["bgp_config"] = d.BgpConfig
	}
	if d.CreatedTime != nil {
		structMap["created_time"] = d.CreatedTime
	}
	if d.DhcpdConfig != nil {
		structMap["dhcpd_config"] = d.DhcpdConfig.toMap()
	}
	if d.DnsOverride != nil {
		structMap["dnsOverride"] = d.DnsOverride
	}
	if d.DnsServers != nil {
		structMap["dns_servers"] = d.DnsServers
	}
	if d.DnsSuffix != nil {
		structMap["dns_suffix"] = d.DnsSuffix
	}
	if d.ExtraRoutes != nil {
		structMap["extra_routes"] = d.ExtraRoutes
	}
	if d.ExtraRoutes6 != nil {
		structMap["extra_routes6"] = d.ExtraRoutes6
	}
	if d.GatewayMatching != nil {
		structMap["gateway_matching"] = d.GatewayMatching.toMap()
	}
	if d.Id != nil {
		structMap["id"] = d.Id
	}
	if d.IdpProfiles != nil {
		structMap["idp_profiles"] = d.IdpProfiles
	}
	if d.IpConfigs != nil {
		structMap["ip_configs"] = d.IpConfigs
	}
	if d.ModifiedTime != nil {
		structMap["modified_time"] = d.ModifiedTime
	}
	structMap["name"] = d.Name
	if d.Networks != nil {
		structMap["networks"] = d.Networks
	}
	if d.NtpOverride != nil {
		structMap["ntpOverride"] = d.NtpOverride
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
	if d.PathPreferences != nil {
		structMap["path_preferences"] = d.PathPreferences
	}
	if d.PortConfig != nil {
		structMap["port_config"] = d.PortConfig
	}
	if d.RouterId != nil {
		structMap["router_id"] = d.RouterId
	}
	if d.RoutingPolicies != nil {
		structMap["routing_policies"] = d.RoutingPolicies
	}
	if d.ServicePolicies != nil {
		structMap["service_policies"] = d.ServicePolicies
	}
	if d.TunnelConfigs != nil {
		structMap["tunnel_configs"] = d.TunnelConfigs
	}
	if d.TunnelProviderOptions != nil {
		structMap["tunnel_provider_options"] = d.TunnelProviderOptions.toMap()
	}
	structMap["type"] = d.Type
	if d.VrfConfig != nil {
		structMap["vrf_config"] = d.VrfConfig.toMap()
	}
	if d.VrfInstances != nil {
		structMap["vrf_instances"] = d.VrfInstances
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeviceprofileGateway.
// It customizes the JSON unmarshaling process for DeviceprofileGateway objects.
func (d *DeviceprofileGateway) UnmarshalJSON(input []byte) error {
	var temp tempDeviceprofileGateway
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "additional_config_cmds", "bgp_config", "created_time", "dhcpd_config", "dnsOverride", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "gateway_matching", "id", "idp_profiles", "ip_configs", "modified_time", "name", "networks", "ntpOverride", "ntp_servers", "oob_ip_config", "org_id", "path_preferences", "port_config", "router_id", "routing_policies", "service_policies", "tunnel_configs", "tunnel_provider_options", "type", "vrf_config", "vrf_instances")
	if err != nil {
		return err
	}

	d.AdditionalProperties = additionalProperties
	d.AdditionalConfigCmds = temp.AdditionalConfigCmds
	d.BgpConfig = temp.BgpConfig
	d.CreatedTime = temp.CreatedTime
	d.DhcpdConfig = temp.DhcpdConfig
	d.DnsOverride = temp.DnsOverride
	d.DnsServers = temp.DnsServers
	d.DnsSuffix = temp.DnsSuffix
	d.ExtraRoutes = temp.ExtraRoutes
	d.ExtraRoutes6 = temp.ExtraRoutes6
	d.GatewayMatching = temp.GatewayMatching
	d.Id = temp.Id
	d.IdpProfiles = temp.IdpProfiles
	d.IpConfigs = temp.IpConfigs
	d.ModifiedTime = temp.ModifiedTime
	d.Name = *temp.Name
	d.Networks = temp.Networks
	d.NtpOverride = temp.NtpOverride
	d.NtpServers = temp.NtpServers
	d.OobIpConfig = temp.OobIpConfig
	d.OrgId = temp.OrgId
	d.PathPreferences = temp.PathPreferences
	d.PortConfig = temp.PortConfig
	d.RouterId = temp.RouterId
	d.RoutingPolicies = temp.RoutingPolicies
	d.ServicePolicies = temp.ServicePolicies
	d.TunnelConfigs = temp.TunnelConfigs
	d.TunnelProviderOptions = temp.TunnelProviderOptions
	d.Type = *temp.Type
	d.VrfConfig = temp.VrfConfig
	d.VrfInstances = temp.VrfInstances
	return nil
}

// tempDeviceprofileGateway is a temporary struct used for validating the fields of DeviceprofileGateway.
type tempDeviceprofileGateway struct {
	AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
	BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
	CreatedTime           *float64                           `json:"created_time,omitempty"`
	DhcpdConfig           *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
	DnsOverride           *bool                              `json:"dnsOverride,omitempty"`
	DnsServers            []string                           `json:"dns_servers,omitempty"`
	DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
	ExtraRoutes           map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
	ExtraRoutes6          map[string]GatewayExtraRoute       `json:"extra_routes6,omitempty"`
	GatewayMatching       *GatewayMatching                   `json:"gateway_matching,omitempty"`
	Id                    *uuid.UUID                         `json:"id,omitempty"`
	IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
	IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
	ModifiedTime          *float64                           `json:"modified_time,omitempty"`
	Name                  *string                            `json:"name"`
	Networks              []Network                          `json:"networks,omitempty"`
	NtpOverride           *bool                              `json:"ntpOverride,omitempty"`
	NtpServers            []string                           `json:"ntp_servers,omitempty"`
	OobIpConfig           *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
	OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
	PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
	PortConfig            map[string]GatewayPortConfig       `json:"port_config,omitempty"`
	RouterId              *string                            `json:"router_id,omitempty"`
	RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
	ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
	TunnelConfigs         map[string]TunnelConfigs           `json:"tunnel_configs,omitempty"`
	TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
	Type                  *string                            `json:"type"`
	VrfConfig             *VrfConfig                         `json:"vrf_config,omitempty"`
	VrfInstances          map[string]GatewayVrfInstance      `json:"vrf_instances,omitempty"`
}

func (d *tempDeviceprofileGateway) validate() error {
	var errs []string
	if d.Name == nil {
		errs = append(errs, "required field `name` is missing for type `deviceprofile_gateway`")
	}
	if d.Type == nil {
		errs = append(errs, "required field `type` is missing for type `deviceprofile_gateway`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
