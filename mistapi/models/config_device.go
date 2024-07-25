package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ConfigDevice represents a ConfigDevice struct.
type ConfigDevice struct {
    // Aeroscout AP settings
    Aeroscout             *ApAeroscout                       `json:"aeroscout,omitempty"`
    // BLE AP settings
    BleConfig             *BleConfig                         `json:"ble_config,omitempty"`
    Centrak               *ApCentrak                         `json:"centrak,omitempty"`
    ClientBridge          *ApClientBridge                    `json:"client_bridge,omitempty"`
    CreatedTime           *float64                           `json:"created_time,omitempty"`
    DeviceprofileId       Optional[uuid.UUID]                `json:"deviceprofile_id"`
    // whether to disable eth1 port
    DisableEth1           *bool                              `json:"disable_eth1,omitempty"`
    // whether to disable eth2 port
    DisableEth2           *bool                              `json:"disable_eth2,omitempty"`
    // whether to disable eth3 port
    DisableEth3           *bool                              `json:"disable_eth3,omitempty"`
    // whether to disable module port
    DisableModule         *bool                              `json:"disable_module,omitempty"`
    EslConfig             *ApEslConfig                       `json:"esl_config,omitempty"`
    ForSite               *bool                              `json:"for_site,omitempty"`
    // height, in meters, optional
    Height                *float64                           `json:"height,omitempty"`
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    Image1Url             Optional[string]                   `json:"image1_url"`
    Image2Url             Optional[string]                   `json:"image2_url"`
    Image3Url             Optional[string]                   `json:"image3_url"`
    // IoT AP settings
    IotConfig             *ApIot                             `json:"iot_config,omitempty"`
    // IP AP settings
    IpConfig              *ApIpConfig1                       `json:"ip_config,omitempty"`
    // LED AP settings
    Led                   *ApLed                             `json:"led,omitempty"`
    // whether this map is considered locked down
    Locked                *bool                              `json:"locked,omitempty"`
    // device MAC address
    Mac                   *string                            `json:"mac,omitempty"`
    // map where the device belongs to
    MapId                 *uuid.UUID                         `json:"map_id,omitempty"`
    // Mesh AP settings
    Mesh                  *ApMesh                            `json:"mesh,omitempty"`
    // device Model
    Model                 *string                            `json:"model,omitempty"`
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    Name                  *string                            `json:"name,omitempty"`
    // any notes about this AP
    Notes                 *string                            `json:"notes,omitempty"`
    // list of NTP servers specific to this device. By default, those in Site Settings will be used
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    // orientation, 0-359, in degrees, up is 0, right is 90.
    Orientation           *int                               `json:"orientation,omitempty"`
    // whether to enable power out through module port (for APH) or eth1 (for APL/BT11)
    PoePassthrough        *bool                              `json:"poe_passthrough,omitempty"`
    // eth0 is not allowed here. 
    // Property key is the interface(s) name (e.g. "eth1" or"eth1,eth2")
    PortConfig            *PortConfig1                       `json:"port_config,omitempty"`
    // power related configs
    PwrConfig             *ApPwrConfig                       `json:"pwr_config,omitempty"`
    // Radio AP settings
    RadioConfig           *ApRadio                           `json:"radio_config,omitempty"`
    // device Serial
    Serial                *string                            `json:"serial,omitempty"`
    SiteId                *uuid.UUID                         `json:"site_id,omitempty"`
    // Device Type
    Type                  *DeviceTypeApEnum                  `json:"type,omitempty"`
    UplinkPortConfig      *ApUplinkPortConfig                `json:"uplink_port_config,omitempty"`
    // USB AP settings
    // Note: if native imagotag is enabled, BLE will be disabled automatically
    // Note: legacy, new config moved to ESL Config.
    UsbConfig             *ApUsb                             `json:"usb_config,omitempty"`
    // a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                  map[string]string                  `json:"vars,omitempty"`
    // x in pixel
    X                     *float64                           `json:"x,omitempty"`
    // y in pixel
    Y                     *float64                           `json:"y,omitempty"`
    AclPolicies           []AclPolicy                        `json:"acl_policies,omitempty"`
    // ACL Tags to identify traffic source or destination. Key name is the tag name
    AclTags               map[string]AclTag                  `json:"acl_tags,omitempty"`
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    DhcpSnooping          *DhcpSnooping                      `json:"dhcp_snooping,omitempty"`
    DhcpdConfig           *DhcpdConfig1                      `json:"dhcpd_config,omitempty"`
    // for a claimed switch, we control the configs by default. This option (disables the behavior)
    DisableAutoConfig     *bool                              `json:"disable_auto_config,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    // EVPN Junos settings
    EvpnConfig            *EvpnConfig                        `json:"evpn_config,omitempty"`
    // Property key is the destination CIDR (e.g. "10.0.0.0/8")
    ExtraRoutes           *ExtraRoutes                       `json:"extra_routes,omitempty"`
    // Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
    ExtraRoutes6          *ExtraRoutes6                      `json:"extra_routes6,omitempty"`
    // for an adopted switch, we don’t overwrite their existing configs automatically
    Managed               *bool                              `json:"managed,omitempty"`
    // enable mist_nac to use radsec
    MistNac               *SwitchMistNac                     `json:"mist_nac,omitempty"`
    // Property key is network name
    Networks              map[string]SwitchNetwork           `json:"networks,omitempty"`
    // - If HA configuration: key parameter will be nodeX (eg: node1)
    // - If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1`
    OobIpConfig           *SwitchOobIpConfig1                `json:"oob_ip_config,omitempty"`
    // Junos OSPF config
    OspfConfig            *OspfConfig                        `json:"ospf_config,omitempty"`
    // Property key is the network name
    OtherIpConfigs        map[string]JunosOtherIpConfig      `json:"other_ip_configs,omitempty"`
    // Property key is the port mirroring instance name
    // port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.
    PortMirroring         *SwitchPortMirroring               `json:"port_mirroring,omitempty"`
    PortUsages            map[string]SwitchPortUsage         `json:"port_usages,omitempty"`
    // Junos Radius config
    RadiusConfig          *RadiusConfig                      `json:"radius_config,omitempty"`
    RemoteSyslog          *RemoteSyslog                      `json:"remote_syslog,omitempty"`
    Role                  *string                            `json:"role,omitempty"`
    // used for OSPF / BGP / EVPN
    RouterId              *string                            `json:"router_id,omitempty"`
    SnmpConfig            *SnmpConfig                        `json:"snmp_config,omitempty"`
    StpConfig             *SwitchStpConfig                   `json:"stp_config,omitempty"`
    SwitchMgmt            *SwitchMgmt                        `json:"switch_mgmt,omitempty"`
    // whether to use it for snmp / syslog / tacplus / radius
    UseRouterIdAsSourceIp *bool                              `json:"use_router_id_as_source_ip,omitempty"`
    // required for preprovisioned Virtual Chassis
    VirtualChassis        *SwitchVirtualChassis              `json:"virtual_chassis,omitempty"`
    VrfConfig             *VrfConfig                         `json:"vrf_config,omitempty"`
    // Property key is the network name
    VrfInstances          map[string]VrfInstance             `json:"vrf_instances,omitempty"`
    // Junos VRRP config
    VrrpConfig            *VrrpConfig                        `json:"vrrp_config,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    // Property key is the profile name
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    // Property key is the network name
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    MspId                 *uuid.UUID                         `json:"msp_id,omitempty"`
    // Property key is the path name
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    // Property key is the routing policy name
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    // Property key is the tunnel name
    TunnelConfigs         map[string]TunnelConfigs           `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
    AdditionalProperties  map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConfigDevice.
// It customizes the JSON marshaling process for ConfigDevice objects.
func (c ConfigDevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConfigDevice object to a map representation for JSON marshaling.
func (c ConfigDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Aeroscout != nil {
        structMap["aeroscout"] = c.Aeroscout.toMap()
    }
    if c.BleConfig != nil {
        structMap["ble_config"] = c.BleConfig.toMap()
    }
    if c.Centrak != nil {
        structMap["centrak"] = c.Centrak.toMap()
    }
    if c.ClientBridge != nil {
        structMap["client_bridge"] = c.ClientBridge.toMap()
    }
    if c.CreatedTime != nil {
        structMap["created_time"] = c.CreatedTime
    }
    if c.DeviceprofileId.IsValueSet() {
        if c.DeviceprofileId.Value() != nil {
            structMap["deviceprofile_id"] = c.DeviceprofileId.Value()
        } else {
            structMap["deviceprofile_id"] = nil
        }
    }
    if c.DisableEth1 != nil {
        structMap["disable_eth1"] = c.DisableEth1
    }
    if c.DisableEth2 != nil {
        structMap["disable_eth2"] = c.DisableEth2
    }
    if c.DisableEth3 != nil {
        structMap["disable_eth3"] = c.DisableEth3
    }
    if c.DisableModule != nil {
        structMap["disable_module"] = c.DisableModule
    }
    if c.EslConfig != nil {
        structMap["esl_config"] = c.EslConfig.toMap()
    }
    if c.ForSite != nil {
        structMap["for_site"] = c.ForSite
    }
    if c.Height != nil {
        structMap["height"] = c.Height
    }
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Image1Url.IsValueSet() {
        if c.Image1Url.Value() != nil {
            structMap["image1_url"] = c.Image1Url.Value()
        } else {
            structMap["image1_url"] = nil
        }
    }
    if c.Image2Url.IsValueSet() {
        if c.Image2Url.Value() != nil {
            structMap["image2_url"] = c.Image2Url.Value()
        } else {
            structMap["image2_url"] = nil
        }
    }
    if c.Image3Url.IsValueSet() {
        if c.Image3Url.Value() != nil {
            structMap["image3_url"] = c.Image3Url.Value()
        } else {
            structMap["image3_url"] = nil
        }
    }
    if c.IotConfig != nil {
        structMap["iot_config"] = c.IotConfig.toMap()
    }
    if c.IpConfig != nil {
        structMap["ip_config"] = c.IpConfig.toMap()
    }
    if c.Led != nil {
        structMap["led"] = c.Led.toMap()
    }
    if c.Locked != nil {
        structMap["locked"] = c.Locked
    }
    if c.Mac != nil {
        structMap["mac"] = c.Mac
    }
    if c.MapId != nil {
        structMap["map_id"] = c.MapId
    }
    if c.Mesh != nil {
        structMap["mesh"] = c.Mesh.toMap()
    }
    if c.Model != nil {
        structMap["model"] = c.Model
    }
    if c.ModifiedTime != nil {
        structMap["modified_time"] = c.ModifiedTime
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Notes != nil {
        structMap["notes"] = c.Notes
    }
    if c.NtpServers != nil {
        structMap["ntp_servers"] = c.NtpServers
    }
    if c.OrgId != nil {
        structMap["org_id"] = c.OrgId
    }
    if c.Orientation != nil {
        structMap["orientation"] = c.Orientation
    }
    if c.PoePassthrough != nil {
        structMap["poe_passthrough"] = c.PoePassthrough
    }
    if c.PortConfig != nil {
        structMap["port_config"] = c.PortConfig.toMap()
    }
    if c.PwrConfig != nil {
        structMap["pwr_config"] = c.PwrConfig.toMap()
    }
    if c.RadioConfig != nil {
        structMap["radio_config"] = c.RadioConfig.toMap()
    }
    if c.Serial != nil {
        structMap["serial"] = c.Serial
    }
    if c.SiteId != nil {
        structMap["site_id"] = c.SiteId
    }
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    if c.UplinkPortConfig != nil {
        structMap["uplink_port_config"] = c.UplinkPortConfig.toMap()
    }
    if c.UsbConfig != nil {
        structMap["usb_config"] = c.UsbConfig.toMap()
    }
    if c.Vars != nil {
        structMap["vars"] = c.Vars
    }
    if c.X != nil {
        structMap["x"] = c.X
    }
    if c.Y != nil {
        structMap["y"] = c.Y
    }
    if c.AclPolicies != nil {
        structMap["acl_policies"] = c.AclPolicies
    }
    if c.AclTags != nil {
        structMap["acl_tags"] = c.AclTags
    }
    if c.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = c.AdditionalConfigCmds
    }
    if c.DhcpSnooping != nil {
        structMap["dhcp_snooping"] = c.DhcpSnooping.toMap()
    }
    if c.DhcpdConfig != nil {
        structMap["dhcpd_config"] = c.DhcpdConfig.toMap()
    }
    if c.DisableAutoConfig != nil {
        structMap["disable_auto_config"] = c.DisableAutoConfig
    }
    if c.DnsServers != nil {
        structMap["dns_servers"] = c.DnsServers
    }
    if c.DnsSuffix != nil {
        structMap["dns_suffix"] = c.DnsSuffix
    }
    if c.EvpnConfig != nil {
        structMap["evpn_config"] = c.EvpnConfig.toMap()
    }
    if c.ExtraRoutes != nil {
        structMap["extra_routes"] = c.ExtraRoutes.toMap()
    }
    if c.ExtraRoutes6 != nil {
        structMap["extra_routes6"] = c.ExtraRoutes6.toMap()
    }
    if c.Managed != nil {
        structMap["managed"] = c.Managed
    }
    if c.MistNac != nil {
        structMap["mist_nac"] = c.MistNac.toMap()
    }
    if c.Networks != nil {
        structMap["networks"] = c.Networks
    }
    if c.OobIpConfig != nil {
        structMap["oob_ip_config"] = c.OobIpConfig.toMap()
    }
    if c.OspfConfig != nil {
        structMap["ospf_config"] = c.OspfConfig.toMap()
    }
    if c.OtherIpConfigs != nil {
        structMap["other_ip_configs"] = c.OtherIpConfigs
    }
    if c.PortMirroring != nil {
        structMap["port_mirroring"] = c.PortMirroring.toMap()
    }
    if c.PortUsages != nil {
        structMap["port_usages"] = c.PortUsages
    }
    if c.RadiusConfig != nil {
        structMap["radius_config"] = c.RadiusConfig.toMap()
    }
    if c.RemoteSyslog != nil {
        structMap["remote_syslog"] = c.RemoteSyslog.toMap()
    }
    if c.Role != nil {
        structMap["role"] = c.Role
    }
    if c.RouterId != nil {
        structMap["router_id"] = c.RouterId
    }
    if c.SnmpConfig != nil {
        structMap["snmp_config"] = c.SnmpConfig.toMap()
    }
    if c.StpConfig != nil {
        structMap["stp_config"] = c.StpConfig.toMap()
    }
    if c.SwitchMgmt != nil {
        structMap["switch_mgmt"] = c.SwitchMgmt.toMap()
    }
    if c.UseRouterIdAsSourceIp != nil {
        structMap["use_router_id_as_source_ip"] = c.UseRouterIdAsSourceIp
    }
    if c.VirtualChassis != nil {
        structMap["virtual_chassis"] = c.VirtualChassis.toMap()
    }
    if c.VrfConfig != nil {
        structMap["vrf_config"] = c.VrfConfig.toMap()
    }
    if c.VrfInstances != nil {
        structMap["vrf_instances"] = c.VrfInstances
    }
    if c.VrrpConfig != nil {
        structMap["vrrp_config"] = c.VrrpConfig.toMap()
    }
    if c.BgpConfig != nil {
        structMap["bgp_config"] = c.BgpConfig
    }
    if c.IdpProfiles != nil {
        structMap["idp_profiles"] = c.IdpProfiles
    }
    if c.IpConfigs != nil {
        structMap["ip_configs"] = c.IpConfigs
    }
    if c.MspId != nil {
        structMap["msp_id"] = c.MspId
    }
    if c.PathPreferences != nil {
        structMap["path_preferences"] = c.PathPreferences
    }
    if c.RoutingPolicies != nil {
        structMap["routing_policies"] = c.RoutingPolicies
    }
    if c.ServicePolicies != nil {
        structMap["service_policies"] = c.ServicePolicies
    }
    if c.TunnelConfigs != nil {
        structMap["tunnel_configs"] = c.TunnelConfigs
    }
    if c.TunnelProviderOptions != nil {
        structMap["tunnel_provider_options"] = c.TunnelProviderOptions.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConfigDevice.
// It customizes the JSON unmarshaling process for ConfigDevice objects.
func (c *ConfigDevice) UnmarshalJSON(input []byte) error {
    var temp configDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "aeroscout", "ble_config", "centrak", "client_bridge", "created_time", "deviceprofile_id", "disable_eth1", "disable_eth2", "disable_eth3", "disable_module", "esl_config", "for_site", "height", "id", "image1_url", "image2_url", "image3_url", "iot_config", "ip_config", "led", "locked", "mac", "map_id", "mesh", "model", "modified_time", "name", "notes", "ntp_servers", "org_id", "orientation", "poe_passthrough", "port_config", "pwr_config", "radio_config", "serial", "site_id", "type", "uplink_port_config", "usb_config", "vars", "x", "y", "acl_policies", "acl_tags", "additional_config_cmds", "dhcp_snooping", "dhcpd_config", "disable_auto_config", "dns_servers", "dns_suffix", "evpn_config", "extra_routes", "extra_routes6", "managed", "mist_nac", "networks", "oob_ip_config", "ospf_config", "other_ip_configs", "port_mirroring", "port_usages", "radius_config", "remote_syslog", "role", "router_id", "snmp_config", "stp_config", "switch_mgmt", "use_router_id_as_source_ip", "virtual_chassis", "vrf_config", "vrf_instances", "vrrp_config", "bgp_config", "idp_profiles", "ip_configs", "msp_id", "path_preferences", "routing_policies", "service_policies", "tunnel_configs", "tunnel_provider_options")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Aeroscout = temp.Aeroscout
    c.BleConfig = temp.BleConfig
    c.Centrak = temp.Centrak
    c.ClientBridge = temp.ClientBridge
    c.CreatedTime = temp.CreatedTime
    c.DeviceprofileId = temp.DeviceprofileId
    c.DisableEth1 = temp.DisableEth1
    c.DisableEth2 = temp.DisableEth2
    c.DisableEth3 = temp.DisableEth3
    c.DisableModule = temp.DisableModule
    c.EslConfig = temp.EslConfig
    c.ForSite = temp.ForSite
    c.Height = temp.Height
    c.Id = temp.Id
    c.Image1Url = temp.Image1Url
    c.Image2Url = temp.Image2Url
    c.Image3Url = temp.Image3Url
    c.IotConfig = temp.IotConfig
    c.IpConfig = temp.IpConfig
    c.Led = temp.Led
    c.Locked = temp.Locked
    c.Mac = temp.Mac
    c.MapId = temp.MapId
    c.Mesh = temp.Mesh
    c.Model = temp.Model
    c.ModifiedTime = temp.ModifiedTime
    c.Name = temp.Name
    c.Notes = temp.Notes
    c.NtpServers = temp.NtpServers
    c.OrgId = temp.OrgId
    c.Orientation = temp.Orientation
    c.PoePassthrough = temp.PoePassthrough
    c.PortConfig = temp.PortConfig
    c.PwrConfig = temp.PwrConfig
    c.RadioConfig = temp.RadioConfig
    c.Serial = temp.Serial
    c.SiteId = temp.SiteId
    c.Type = temp.Type
    c.UplinkPortConfig = temp.UplinkPortConfig
    c.UsbConfig = temp.UsbConfig
    c.Vars = temp.Vars
    c.X = temp.X
    c.Y = temp.Y
    c.AclPolicies = temp.AclPolicies
    c.AclTags = temp.AclTags
    c.AdditionalConfigCmds = temp.AdditionalConfigCmds
    c.DhcpSnooping = temp.DhcpSnooping
    c.DhcpdConfig = temp.DhcpdConfig
    c.DisableAutoConfig = temp.DisableAutoConfig
    c.DnsServers = temp.DnsServers
    c.DnsSuffix = temp.DnsSuffix
    c.EvpnConfig = temp.EvpnConfig
    c.ExtraRoutes = temp.ExtraRoutes
    c.ExtraRoutes6 = temp.ExtraRoutes6
    c.Managed = temp.Managed
    c.MistNac = temp.MistNac
    c.Networks = temp.Networks
    c.OobIpConfig = temp.OobIpConfig
    c.OspfConfig = temp.OspfConfig
    c.OtherIpConfigs = temp.OtherIpConfigs
    c.PortMirroring = temp.PortMirroring
    c.PortUsages = temp.PortUsages
    c.RadiusConfig = temp.RadiusConfig
    c.RemoteSyslog = temp.RemoteSyslog
    c.Role = temp.Role
    c.RouterId = temp.RouterId
    c.SnmpConfig = temp.SnmpConfig
    c.StpConfig = temp.StpConfig
    c.SwitchMgmt = temp.SwitchMgmt
    c.UseRouterIdAsSourceIp = temp.UseRouterIdAsSourceIp
    c.VirtualChassis = temp.VirtualChassis
    c.VrfConfig = temp.VrfConfig
    c.VrfInstances = temp.VrfInstances
    c.VrrpConfig = temp.VrrpConfig
    c.BgpConfig = temp.BgpConfig
    c.IdpProfiles = temp.IdpProfiles
    c.IpConfigs = temp.IpConfigs
    c.MspId = temp.MspId
    c.PathPreferences = temp.PathPreferences
    c.RoutingPolicies = temp.RoutingPolicies
    c.ServicePolicies = temp.ServicePolicies
    c.TunnelConfigs = temp.TunnelConfigs
    c.TunnelProviderOptions = temp.TunnelProviderOptions
    return nil
}

// configDevice is a temporary struct used for validating the fields of ConfigDevice.
type configDevice  struct {
    Aeroscout             *ApAeroscout                       `json:"aeroscout,omitempty"`
    BleConfig             *BleConfig                         `json:"ble_config,omitempty"`
    Centrak               *ApCentrak                         `json:"centrak,omitempty"`
    ClientBridge          *ApClientBridge                    `json:"client_bridge,omitempty"`
    CreatedTime           *float64                           `json:"created_time,omitempty"`
    DeviceprofileId       Optional[uuid.UUID]                `json:"deviceprofile_id"`
    DisableEth1           *bool                              `json:"disable_eth1,omitempty"`
    DisableEth2           *bool                              `json:"disable_eth2,omitempty"`
    DisableEth3           *bool                              `json:"disable_eth3,omitempty"`
    DisableModule         *bool                              `json:"disable_module,omitempty"`
    EslConfig             *ApEslConfig                       `json:"esl_config,omitempty"`
    ForSite               *bool                              `json:"for_site,omitempty"`
    Height                *float64                           `json:"height,omitempty"`
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    Image1Url             Optional[string]                   `json:"image1_url"`
    Image2Url             Optional[string]                   `json:"image2_url"`
    Image3Url             Optional[string]                   `json:"image3_url"`
    IotConfig             *ApIot                             `json:"iot_config,omitempty"`
    IpConfig              *ApIpConfig1                       `json:"ip_config,omitempty"`
    Led                   *ApLed                             `json:"led,omitempty"`
    Locked                *bool                              `json:"locked,omitempty"`
    Mac                   *string                            `json:"mac,omitempty"`
    MapId                 *uuid.UUID                         `json:"map_id,omitempty"`
    Mesh                  *ApMesh                            `json:"mesh,omitempty"`
    Model                 *string                            `json:"model,omitempty"`
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    Name                  *string                            `json:"name,omitempty"`
    Notes                 *string                            `json:"notes,omitempty"`
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    Orientation           *int                               `json:"orientation,omitempty"`
    PoePassthrough        *bool                              `json:"poe_passthrough,omitempty"`
    PortConfig            *PortConfig1                       `json:"port_config,omitempty"`
    PwrConfig             *ApPwrConfig                       `json:"pwr_config,omitempty"`
    RadioConfig           *ApRadio                           `json:"radio_config,omitempty"`
    Serial                *string                            `json:"serial,omitempty"`
    SiteId                *uuid.UUID                         `json:"site_id,omitempty"`
    Type                  *DeviceTypeApEnum                  `json:"type,omitempty"`
    UplinkPortConfig      *ApUplinkPortConfig                `json:"uplink_port_config,omitempty"`
    UsbConfig             *ApUsb                             `json:"usb_config,omitempty"`
    Vars                  map[string]string                  `json:"vars,omitempty"`
    X                     *float64                           `json:"x,omitempty"`
    Y                     *float64                           `json:"y,omitempty"`
    AclPolicies           []AclPolicy                        `json:"acl_policies,omitempty"`
    AclTags               map[string]AclTag                  `json:"acl_tags,omitempty"`
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    DhcpSnooping          *DhcpSnooping                      `json:"dhcp_snooping,omitempty"`
    DhcpdConfig           *DhcpdConfig1                      `json:"dhcpd_config,omitempty"`
    DisableAutoConfig     *bool                              `json:"disable_auto_config,omitempty"`
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    EvpnConfig            *EvpnConfig                        `json:"evpn_config,omitempty"`
    ExtraRoutes           *ExtraRoutes                       `json:"extra_routes,omitempty"`
    ExtraRoutes6          *ExtraRoutes6                      `json:"extra_routes6,omitempty"`
    Managed               *bool                              `json:"managed,omitempty"`
    MistNac               *SwitchMistNac                     `json:"mist_nac,omitempty"`
    Networks              map[string]SwitchNetwork           `json:"networks,omitempty"`
    OobIpConfig           *SwitchOobIpConfig1                `json:"oob_ip_config,omitempty"`
    OspfConfig            *OspfConfig                        `json:"ospf_config,omitempty"`
    OtherIpConfigs        map[string]JunosOtherIpConfig      `json:"other_ip_configs,omitempty"`
    PortMirroring         *SwitchPortMirroring               `json:"port_mirroring,omitempty"`
    PortUsages            map[string]SwitchPortUsage         `json:"port_usages,omitempty"`
    RadiusConfig          *RadiusConfig                      `json:"radius_config,omitempty"`
    RemoteSyslog          *RemoteSyslog                      `json:"remote_syslog,omitempty"`
    Role                  *string                            `json:"role,omitempty"`
    RouterId              *string                            `json:"router_id,omitempty"`
    SnmpConfig            *SnmpConfig                        `json:"snmp_config,omitempty"`
    StpConfig             *SwitchStpConfig                   `json:"stp_config,omitempty"`
    SwitchMgmt            *SwitchMgmt                        `json:"switch_mgmt,omitempty"`
    UseRouterIdAsSourceIp *bool                              `json:"use_router_id_as_source_ip,omitempty"`
    VirtualChassis        *SwitchVirtualChassis              `json:"virtual_chassis,omitempty"`
    VrfConfig             *VrfConfig                         `json:"vrf_config,omitempty"`
    VrfInstances          map[string]VrfInstance             `json:"vrf_instances,omitempty"`
    VrrpConfig            *VrrpConfig                        `json:"vrrp_config,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    MspId                 *uuid.UUID                         `json:"msp_id,omitempty"`
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    TunnelConfigs         map[string]TunnelConfigs           `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
}
