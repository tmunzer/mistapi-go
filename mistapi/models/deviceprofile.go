package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// Deviceprofile represents a Deviceprofile struct.
type Deviceprofile struct {
    // Aeroscout AP settings
    Aeroscout             *ApAeroscout                       `json:"aeroscout,omitempty"`
    // BLE AP settings
    BleConfig             *BleConfig                         `json:"ble_config,omitempty"`
    CreatedTime           *float64                           `json:"created_time,omitempty"`
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
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    // IoT AP settings
    IotConfig             *ApIot                             `json:"iot_config,omitempty"`
    // IP AP settings
    IpConfig              *ApIpConfig                        `json:"ip_config,omitempty"`
    // LED AP settings
    Led                   *ApLed                             `json:"led,omitempty"`
    // Mesh AP settings
    Mesh                  *ApMesh                            `json:"mesh,omitempty"`
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    Name                  Optional[string]                   `json:"name"`
    // list of NTP servers specific to this device. By default, those in Site Settings will be used
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    // whether to enable power out through module port (for APH) or eth1 (for APL/BT11)
    PoePassthrough        *bool                              `json:"poe_passthrough,omitempty"`
    // Property key is the interface(s) name (e.g. "eth1,eth2")
    PortConfig            *PortConfig1                       `json:"port_config,omitempty"`
    // power related configs
    PwrConfig             *ApPwrConfig                       `json:"pwr_config,omitempty"`
    // Radio AP settings
    RadioConfig           *ApRadio                           `json:"radio_config,omitempty"`
    SiteId                *uuid.UUID                         `json:"site_id,omitempty"`
    // for people who want to fully control the vlans (advanced)
    SwitchConfig          *ApSwitch                          `json:"switch_config,omitempty"`           // Deprecated
    // Device Type. enum: `ap`
    Type                  *DeviceTypeApEnum                  `json:"type,omitempty"`
    UplinkPortConfig      *ApUplinkPortConfig                `json:"uplink_port_config,omitempty"`
    // USB AP settings
    // Note: if native imagotag is enabled, BLE will be disabled automatically
    // Note: legacy, new config moved to ESL Config.
    UsbConfig             *ApUsb                             `json:"usb_config,omitempty"`
    // a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                  map[string]string                  `json:"vars,omitempty"`
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    DhcpdConfig           *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
    DnsOverride           *bool                              `json:"dnsOverride,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    // Property key is the destination CIDR (e.g. "10.0.0.0/8")
    ExtraRoutes           map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
    // Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
    ExtraRoutes6          map[string]GatewayExtraRoute       `json:"extra_routes6,omitempty"`
    // Gateway matching
    GatewayMatching       *GatewayMatching                   `json:"gateway_matching,omitempty"`
    // Property key is the profile name
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    // Property key is the network name
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    Networks              []Network                          `json:"networks,omitempty"`
    NtpOverride           *bool                              `json:"ntpOverride,omitempty"`
    // out-of-band (vme/em0/fxp0) IP config
    OobIpConfig           *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
    // Property key is the path name
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    // auto assigned if not set
    RouterId              *string                            `json:"router_id,omitempty"`
    // Property key is the routing policy name
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    // Property key is the tunnel name
    TunnelConfigs         map[string]TunnelConfigs           `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
    VrfConfig             *VrfConfig                         `json:"vrf_config,omitempty"`
    // Property key is the network name
    VrfInstances          map[string]GatewayVrfInstance      `json:"vrf_instances,omitempty"`
    AdditionalProperties  map[string]any                     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Deviceprofile.
// It customizes the JSON marshaling process for Deviceprofile objects.
func (d Deviceprofile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the Deviceprofile object to a map representation for JSON marshaling.
func (d Deviceprofile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Aeroscout != nil {
        structMap["aeroscout"] = d.Aeroscout.toMap()
    }
    if d.BleConfig != nil {
        structMap["ble_config"] = d.BleConfig.toMap()
    }
    if d.CreatedTime != nil {
        structMap["created_time"] = d.CreatedTime
    }
    if d.DisableEth1 != nil {
        structMap["disable_eth1"] = d.DisableEth1
    }
    if d.DisableEth2 != nil {
        structMap["disable_eth2"] = d.DisableEth2
    }
    if d.DisableEth3 != nil {
        structMap["disable_eth3"] = d.DisableEth3
    }
    if d.DisableModule != nil {
        structMap["disable_module"] = d.DisableModule
    }
    if d.EslConfig != nil {
        structMap["esl_config"] = d.EslConfig.toMap()
    }
    if d.ForSite != nil {
        structMap["for_site"] = d.ForSite
    }
    if d.Id != nil {
        structMap["id"] = d.Id
    }
    if d.IotConfig != nil {
        structMap["iot_config"] = d.IotConfig.toMap()
    }
    if d.IpConfig != nil {
        structMap["ip_config"] = d.IpConfig.toMap()
    }
    if d.Led != nil {
        structMap["led"] = d.Led.toMap()
    }
    if d.Mesh != nil {
        structMap["mesh"] = d.Mesh.toMap()
    }
    if d.ModifiedTime != nil {
        structMap["modified_time"] = d.ModifiedTime
    }
    if d.Name.IsValueSet() {
        if d.Name.Value() != nil {
            structMap["name"] = d.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if d.NtpServers != nil {
        structMap["ntp_servers"] = d.NtpServers
    }
    if d.OrgId != nil {
        structMap["org_id"] = d.OrgId
    }
    if d.PoePassthrough != nil {
        structMap["poe_passthrough"] = d.PoePassthrough
    }
    if d.PortConfig != nil {
        structMap["port_config"] = d.PortConfig.toMap()
    }
    if d.PwrConfig != nil {
        structMap["pwr_config"] = d.PwrConfig.toMap()
    }
    if d.RadioConfig != nil {
        structMap["radio_config"] = d.RadioConfig.toMap()
    }
    if d.SiteId != nil {
        structMap["site_id"] = d.SiteId
    }
    if d.SwitchConfig != nil {
        structMap["switch_config"] = d.SwitchConfig.toMap()
    }
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.UplinkPortConfig != nil {
        structMap["uplink_port_config"] = d.UplinkPortConfig.toMap()
    }
    if d.UsbConfig != nil {
        structMap["usb_config"] = d.UsbConfig.toMap()
    }
    if d.Vars != nil {
        structMap["vars"] = d.Vars
    }
    if d.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = d.AdditionalConfigCmds
    }
    if d.BgpConfig != nil {
        structMap["bgp_config"] = d.BgpConfig
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
    if d.IdpProfiles != nil {
        structMap["idp_profiles"] = d.IdpProfiles
    }
    if d.IpConfigs != nil {
        structMap["ip_configs"] = d.IpConfigs
    }
    if d.Networks != nil {
        structMap["networks"] = d.Networks
    }
    if d.NtpOverride != nil {
        structMap["ntpOverride"] = d.NtpOverride
    }
    if d.OobIpConfig != nil {
        structMap["oob_ip_config"] = d.OobIpConfig.toMap()
    }
    if d.PathPreferences != nil {
        structMap["path_preferences"] = d.PathPreferences
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
    if d.VrfConfig != nil {
        structMap["vrf_config"] = d.VrfConfig.toMap()
    }
    if d.VrfInstances != nil {
        structMap["vrf_instances"] = d.VrfInstances
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Deviceprofile.
// It customizes the JSON unmarshaling process for Deviceprofile objects.
func (d *Deviceprofile) UnmarshalJSON(input []byte) error {
    var temp tempDeviceprofile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "aeroscout", "ble_config", "created_time", "disable_eth1", "disable_eth2", "disable_eth3", "disable_module", "esl_config", "for_site", "id", "iot_config", "ip_config", "led", "mesh", "modified_time", "name", "ntp_servers", "org_id", "poe_passthrough", "port_config", "pwr_config", "radio_config", "site_id", "switch_config", "type", "uplink_port_config", "usb_config", "vars", "additional_config_cmds", "bgp_config", "dhcpd_config", "dnsOverride", "dns_servers", "dns_suffix", "extra_routes", "extra_routes6", "gateway_matching", "idp_profiles", "ip_configs", "networks", "ntpOverride", "oob_ip_config", "path_preferences", "router_id", "routing_policies", "service_policies", "tunnel_configs", "tunnel_provider_options", "vrf_config", "vrf_instances")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Aeroscout = temp.Aeroscout
    d.BleConfig = temp.BleConfig
    d.CreatedTime = temp.CreatedTime
    d.DisableEth1 = temp.DisableEth1
    d.DisableEth2 = temp.DisableEth2
    d.DisableEth3 = temp.DisableEth3
    d.DisableModule = temp.DisableModule
    d.EslConfig = temp.EslConfig
    d.ForSite = temp.ForSite
    d.Id = temp.Id
    d.IotConfig = temp.IotConfig
    d.IpConfig = temp.IpConfig
    d.Led = temp.Led
    d.Mesh = temp.Mesh
    d.ModifiedTime = temp.ModifiedTime
    d.Name = temp.Name
    d.NtpServers = temp.NtpServers
    d.OrgId = temp.OrgId
    d.PoePassthrough = temp.PoePassthrough
    d.PortConfig = temp.PortConfig
    d.PwrConfig = temp.PwrConfig
    d.RadioConfig = temp.RadioConfig
    d.SiteId = temp.SiteId
    d.SwitchConfig = temp.SwitchConfig
    d.Type = temp.Type
    d.UplinkPortConfig = temp.UplinkPortConfig
    d.UsbConfig = temp.UsbConfig
    d.Vars = temp.Vars
    d.AdditionalConfigCmds = temp.AdditionalConfigCmds
    d.BgpConfig = temp.BgpConfig
    d.DhcpdConfig = temp.DhcpdConfig
    d.DnsOverride = temp.DnsOverride
    d.DnsServers = temp.DnsServers
    d.DnsSuffix = temp.DnsSuffix
    d.ExtraRoutes = temp.ExtraRoutes
    d.ExtraRoutes6 = temp.ExtraRoutes6
    d.GatewayMatching = temp.GatewayMatching
    d.IdpProfiles = temp.IdpProfiles
    d.IpConfigs = temp.IpConfigs
    d.Networks = temp.Networks
    d.NtpOverride = temp.NtpOverride
    d.OobIpConfig = temp.OobIpConfig
    d.PathPreferences = temp.PathPreferences
    d.RouterId = temp.RouterId
    d.RoutingPolicies = temp.RoutingPolicies
    d.ServicePolicies = temp.ServicePolicies
    d.TunnelConfigs = temp.TunnelConfigs
    d.TunnelProviderOptions = temp.TunnelProviderOptions
    d.VrfConfig = temp.VrfConfig
    d.VrfInstances = temp.VrfInstances
    return nil
}

// tempDeviceprofile is a temporary struct used for validating the fields of Deviceprofile.
type tempDeviceprofile  struct {
    Aeroscout             *ApAeroscout                       `json:"aeroscout,omitempty"`
    BleConfig             *BleConfig                         `json:"ble_config,omitempty"`
    CreatedTime           *float64                           `json:"created_time,omitempty"`
    DisableEth1           *bool                              `json:"disable_eth1,omitempty"`
    DisableEth2           *bool                              `json:"disable_eth2,omitempty"`
    DisableEth3           *bool                              `json:"disable_eth3,omitempty"`
    DisableModule         *bool                              `json:"disable_module,omitempty"`
    EslConfig             *ApEslConfig                       `json:"esl_config,omitempty"`
    ForSite               *bool                              `json:"for_site,omitempty"`
    Id                    *uuid.UUID                         `json:"id,omitempty"`
    IotConfig             *ApIot                             `json:"iot_config,omitempty"`
    IpConfig              *ApIpConfig                        `json:"ip_config,omitempty"`
    Led                   *ApLed                             `json:"led,omitempty"`
    Mesh                  *ApMesh                            `json:"mesh,omitempty"`
    ModifiedTime          *float64                           `json:"modified_time,omitempty"`
    Name                  Optional[string]                   `json:"name"`
    NtpServers            []string                           `json:"ntp_servers,omitempty"`
    OrgId                 *uuid.UUID                         `json:"org_id,omitempty"`
    PoePassthrough        *bool                              `json:"poe_passthrough,omitempty"`
    PortConfig            *PortConfig1                       `json:"port_config,omitempty"`
    PwrConfig             *ApPwrConfig                       `json:"pwr_config,omitempty"`
    RadioConfig           *ApRadio                           `json:"radio_config,omitempty"`
    SiteId                *uuid.UUID                         `json:"site_id,omitempty"`
    SwitchConfig          *ApSwitch                          `json:"switch_config,omitempty"`
    Type                  *DeviceTypeApEnum                  `json:"type,omitempty"`
    UplinkPortConfig      *ApUplinkPortConfig                `json:"uplink_port_config,omitempty"`
    UsbConfig             *ApUsb                             `json:"usb_config,omitempty"`
    Vars                  map[string]string                  `json:"vars,omitempty"`
    AdditionalConfigCmds  []string                           `json:"additional_config_cmds,omitempty"`
    BgpConfig             map[string]BgpConfig               `json:"bgp_config,omitempty"`
    DhcpdConfig           *DhcpdConfig                       `json:"dhcpd_config,omitempty"`
    DnsOverride           *bool                              `json:"dnsOverride,omitempty"`
    DnsServers            []string                           `json:"dns_servers,omitempty"`
    DnsSuffix             []string                           `json:"dns_suffix,omitempty"`
    ExtraRoutes           map[string]GatewayExtraRoute       `json:"extra_routes,omitempty"`
    ExtraRoutes6          map[string]GatewayExtraRoute       `json:"extra_routes6,omitempty"`
    GatewayMatching       *GatewayMatching                   `json:"gateway_matching,omitempty"`
    IdpProfiles           map[string]IdpProfile              `json:"idp_profiles,omitempty"`
    IpConfigs             map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
    Networks              []Network                          `json:"networks,omitempty"`
    NtpOverride           *bool                              `json:"ntpOverride,omitempty"`
    OobIpConfig           *GatewayOobIpConfig                `json:"oob_ip_config,omitempty"`
    PathPreferences       map[string]GatewayPathPreferences  `json:"path_preferences,omitempty"`
    RouterId              *string                            `json:"router_id,omitempty"`
    RoutingPolicies       map[string]RoutingPolicy           `json:"routing_policies,omitempty"`
    ServicePolicies       []ServicePolicy                    `json:"service_policies,omitempty"`
    TunnelConfigs         map[string]TunnelConfigs           `json:"tunnel_configs,omitempty"`
    TunnelProviderOptions *TunnelProviderOptions             `json:"tunnel_provider_options,omitempty"`
    VrfConfig             *VrfConfig                         `json:"vrf_config,omitempty"`
    VrfInstances          map[string]GatewayVrfInstance      `json:"vrf_instances,omitempty"`
}
