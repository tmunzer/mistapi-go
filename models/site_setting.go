package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// SiteSetting represents a SiteSetting struct.
// Site Settings
type SiteSetting struct {
    AclPolicies                     []AclPolicy                            `json:"acl_policies,omitempty"`
    // ACL Tags to identify traffic source or destination. Key name is the tag name
    AclTags                         map[string]AclTag                      `json:"acl_tags,omitempty"`
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    AdditionalConfigCmds            []string                               `json:"additional_config_cmds,omitempty"`
    Analytic                        *SiteSettingAnalytic                   `json:"analytic,omitempty"`
    ApMatching                      *SiteSettingApMatching                 `json:"ap_matching,omitempty"`
    ApPortConfig                    *SiteSettingApPortConfig               `json:"ap_port_config,omitempty"`
    // enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.
    ApUpdownThreshold               Optional[int]                          `json:"ap_updown_threshold"`
    // if we're able to determine its x/y/orientation, this will be populated
    AutoPlacement                   *SiteSettingAutoPlacement              `json:"auto_placement,omitempty"`
    // Auto Upgrade Settings
    AutoUpgrade                     *SiteSettingAutoUpgrade                `json:"auto_upgrade,omitempty"`
    BlacklistUrl                    *string                                `json:"blacklist_url,omitempty"`
    // BLE AP settings
    BleConfig                       *BleConfig                             `json:"ble_config,omitempty"`
    // whether to enable ap auto config revert
    ConfigAutoRevert                *bool                                  `json:"config_auto_revert,omitempty"`
    // mist also uses some heuristic rules to prevent destructive configs from being pushed
    ConfigPushPolicy                *SiteSettingConfigPushPolicy           `json:"config_push_policy,omitempty"`
    CreatedTime                     *float64                               `json:"created_time,omitempty"`
    // you can define some URLs that's critical to site operaitons the latency will be captured and considered for site health
    CriticalUrlMonitoring           *SiteSettingCriticalUrlMonitoring      `json:"critical_url_monitoring,omitempty"`
    // sending AP_DISCONNECTED event in device-updowns only if AP_CONNECTED is not seen within the threshold, in minutes
    DeviceUpdownThreshold           *int                                   `json:"device_updown_threshold,omitempty"`
    DhcpSnooping                    *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
    // if some system-default port usages are not desired - namely, ap / iot / uplink
    DisabledSystemDefinedPortUsages []string                               `json:"disabled_system_defined_port_usages,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsServers                      []string                               `json:"dns_servers,omitempty"`
    // Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
    DnsSuffix                       []string                               `json:"dns_suffix,omitempty"`
    // **Note**: if hours does not exist, it’s treated as everyday of the week, 00:00-23:59. Currently we don’t allow multiple ranges for the same day
    // **Note**: default values for `dwell_tags`: passerby (1,300) bounce (301, 14400) engaged (14401, 28800) stationed (28801, 42000)
    // **Note**: default values for `dwell_tag_names`: passerby = “Passerby”, bounce = “Visitor”, engaged = “Associates”, stationed = “Assets”
    Engagement                      *SiteEngagement                        `json:"engagement,omitempty"`
    // EVPN Options
    EvpnOptions                     *EvpnOptions                           `json:"evpn_options,omitempty"`
    ExtraRoutes                     map[string]ExtraRouteProperties        `json:"extra_routes,omitempty"`
    // Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
    ExtraRoutes6                    map[string]ExtraRoute6Properties       `json:"extra_routes6,omitempty"`
    // name/val pair objects for location engine to use
    Flags                           map[string]string                      `json:"flags,omitempty"`
    ForSite                         *bool                                  `json:"for_site,omitempty"`
    // Gateway Template is applied to a site for gateway(s) in a site.
    Gateway                         *GatewayTemplate                       `json:"gateway,omitempty"`
    // additional CLI commands to append to the generated Junos config
    // **Note**: no check is done
    GatewayAdditionalConfigCmds     []string                               `json:"gateway_additional_config_cmds,omitempty"`
    // Gateway Site settings
    GatewayMgmt                     *SiteSettingGatewayMgmt                `json:"gateway_mgmt,omitempty"`
    // enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.
    GatewayUpdownThreshold          Optional[int]                          `json:"gateway_updown_threshold"`
    Id                              *uuid.UUID                             `json:"id,omitempty"`
    // LED AP settings
    Led                             *ApLed                                 `json:"led,omitempty"`
    // enable mist_nac to use radsec
    MistNac                         *SwitchMistNac                         `json:"mist_nac,omitempty"`
    ModifiedTime                    *float64                               `json:"modified_time,omitempty"`
    // site mist edges form a cluster of radsecproxy servers
    Mxedge                          *SiteSettingMxedge                     `json:"mxedge,omitempty"`
    MxedgeMgmt                      *MxedgeMgmt                            `json:"mxedge_mgmt,omitempty"`
    // Site MxTunnel
    Mxtunnels                       *SiteMxtunnel                          `json:"mxtunnels,omitempty"`
    // Property key is network name
    Networks                        map[string]SwitchNetwork               `json:"networks,omitempty"`
    // list of NTP servers
    NtpServers                      []string                               `json:"ntp_servers,omitempty"`
    // Occupancy Analytics settings
    Occupancy                       *SiteOccupancyAnalytics                `json:"occupancy,omitempty"`
    OrgId                           *uuid.UUID                             `json:"org_id,omitempty"`
    // OSPF Areas can be configured both in Site Level (Switch/Gateway Template) or Device (Switch/Gateway) Level
    OspfAreas                       map[string]OspfAreas                   `json:"ospf_areas,omitempty"`
    PaloaltoNetworks                *SiteSettingPaloaltoNetworks           `json:"paloalto_networks,omitempty"`
    // whether to store the config on AP
    PersistConfigOnDevice           *bool                                  `json:"persist_config_on_device,omitempty"`
    // Property key is the port mirroring instance name
    // port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output.
    PortMirroring                   map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    PortUsages                      map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
    // Proxy Configuration to talk to Mist
    Proxy                           *Proxy                                 `json:"proxy,omitempty"`
    // Radio AP settings
    RadioConfig                     *ApRadio                               `json:"radio_config,omitempty"`
    // Junos Radius config
    RadiusConfig                    *RadiusConfig                          `json:"radius_config,omitempty"`
    RemoteSyslog                    *RemoteSyslog                          `json:"remote_syslog,omitempty"`
    // whether AP should periodically connect to BLE devices and report GATT device info (device name, manufacturer name, serial number, battery %, temperature, humidity)
    ReportGatt                      *bool                                  `json:"report_gatt,omitempty"`
    // Rogue site settings
    Rogue                           *SiteRogue                             `json:"rogue,omitempty"`
    // managed mobility
    Rtsa                            *SiteSettingRtsa                       `json:"rtsa,omitempty"`
    // Set of heuristic rules will be enabled when marvis subscription is not available.
    // It triggers when, in a Z minute window, there are more than Y distinct client encountring over X failures
    SimpleAlert                     *SimpleAlert                           `json:"simple_alert,omitempty"`
    SiteId                          *uuid.UUID                             `json:"site_id,omitempty"`
    Skyatp                          *SiteSettingSkyatp                     `json:"skyatp,omitempty"`
    SnmpConfig                      *SnmpConfig                            `json:"snmp_config,omitempty"`
    SrxApp                          *SiteSettingSrxApp                     `json:"srx_app,omitempty"`
    // when limit_ssh_access = true in Org Setting, list of SSH public keys provided by Mist Support to install onto APs (see Org:Setting)
    SshKeys                         []string                               `json:"ssh_keys,omitempty"`
    Ssr                             *SiteSettingSsr                        `json:"ssr,omitempty"`
    StatusPortal                    *SiteSettingStatusPortal               `json:"status_portal,omitempty"`
    // Network Template
    Switch                          *NetworkTemplate                       `json:"switch,omitempty"`
    // Switch template
    SwitchMatching                  *SwitchMatching                        `json:"switch_matching,omitempty"`
    SwitchMgmt                      *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    // enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.
    SwitchUpdownThreshold           Optional[int]                          `json:"switch_updown_threshold"`
    SyntheticTest                   *SynthetictestConfig                   `json:"synthetic_test,omitempty"`
    // whether to track anonymous BLE assets (requires ‘track_asset’ enabled)
    TrackAnonymousDevices           *bool                                  `json:"track_anonymous_devices,omitempty"`
    TuntermMonitoring               []TuntermMonitoringItem                `json:"tunterm_monitoring,omitempty"`
    TuntermMonitoringDisabled       *bool                                  `json:"tunterm_monitoring_disabled,omitempty"`
    TuntermMulticastConfig          *SiteSettingTuntermMulticastConfig     `json:"tunterm_multicast_config,omitempty"`
    // a dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
    Vars                            map[string]string                      `json:"vars,omitempty"`
    Vna                             *SiteSettingVna                        `json:"vna,omitempty"`
    VrfConfig                       *VrfConfig                             `json:"vrf_config,omitempty"`
    // Property key is the network name
    VrfInstances                    map[string]VrfInstance                 `json:"vrf_instances,omitempty"`
    // Property key is the vrrp group
    VrrpGroups                      map[string]VrrpGroup                   `json:"vrrp_groups,omitempty"`
    VsInstance                      *VsInstance                            `json:"vs_instance,omitempty"`
    WanVna                          *SiteSettingWanVna                     `json:"wan_vna,omitempty"`
    WatchedStationUrl               *string                                `json:"watched_station_url,omitempty"`
    WhitelistUrl                    *string                                `json:"whitelist_url,omitempty"`
    // WIDS site settings
    Wids                            *SiteWids                              `json:"wids,omitempty"`
    // Wi-Fi site settings
    Wifi                            *SiteWifi                              `json:"wifi,omitempty"`
    WiredVna                        *SiteSettingWiredVna                   `json:"wired_vna,omitempty"`
    // Zone Occupancy alert site settings
    ZoneOccupancyAlert              *SiteZoneOccupancyAlert                `json:"zone_occupancy_alert,omitempty"`
    AdditionalProperties            map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSetting.
// It customizes the JSON marshaling process for SiteSetting objects.
func (s SiteSetting) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSetting object to a map representation for JSON marshaling.
func (s SiteSetting) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AclPolicies != nil {
        structMap["acl_policies"] = s.AclPolicies
    }
    if s.AclTags != nil {
        structMap["acl_tags"] = s.AclTags
    }
    if s.AdditionalConfigCmds != nil {
        structMap["additional_config_cmds"] = s.AdditionalConfigCmds
    }
    if s.Analytic != nil {
        structMap["analytic"] = s.Analytic.toMap()
    }
    if s.ApMatching != nil {
        structMap["ap_matching"] = s.ApMatching.toMap()
    }
    if s.ApPortConfig != nil {
        structMap["ap_port_config"] = s.ApPortConfig.toMap()
    }
    if s.ApUpdownThreshold.IsValueSet() {
        if s.ApUpdownThreshold.Value() != nil {
            structMap["ap_updown_threshold"] = s.ApUpdownThreshold.Value()
        } else {
            structMap["ap_updown_threshold"] = nil
        }
    }
    if s.AutoPlacement != nil {
        structMap["auto_placement"] = s.AutoPlacement.toMap()
    }
    if s.AutoUpgrade != nil {
        structMap["auto_upgrade"] = s.AutoUpgrade.toMap()
    }
    if s.BlacklistUrl != nil {
        structMap["blacklist_url"] = s.BlacklistUrl
    }
    if s.BleConfig != nil {
        structMap["ble_config"] = s.BleConfig.toMap()
    }
    if s.ConfigAutoRevert != nil {
        structMap["config_auto_revert"] = s.ConfigAutoRevert
    }
    if s.ConfigPushPolicy != nil {
        structMap["config_push_policy"] = s.ConfigPushPolicy.toMap()
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.CriticalUrlMonitoring != nil {
        structMap["critical_url_monitoring"] = s.CriticalUrlMonitoring.toMap()
    }
    if s.DeviceUpdownThreshold != nil {
        structMap["device_updown_threshold"] = s.DeviceUpdownThreshold
    }
    if s.DhcpSnooping != nil {
        structMap["dhcp_snooping"] = s.DhcpSnooping.toMap()
    }
    if s.DisabledSystemDefinedPortUsages != nil {
        structMap["disabled_system_defined_port_usages"] = s.DisabledSystemDefinedPortUsages
    }
    if s.DnsServers != nil {
        structMap["dns_servers"] = s.DnsServers
    }
    if s.DnsSuffix != nil {
        structMap["dns_suffix"] = s.DnsSuffix
    }
    if s.Engagement != nil {
        structMap["engagement"] = s.Engagement.toMap()
    }
    if s.EvpnOptions != nil {
        structMap["evpn_options"] = s.EvpnOptions.toMap()
    }
    if s.ExtraRoutes != nil {
        structMap["extra_routes"] = s.ExtraRoutes
    }
    if s.ExtraRoutes6 != nil {
        structMap["extra_routes6"] = s.ExtraRoutes6
    }
    if s.Flags != nil {
        structMap["flags"] = s.Flags
    }
    if s.ForSite != nil {
        structMap["for_site"] = s.ForSite
    }
    if s.Gateway != nil {
        structMap["gateway"] = s.Gateway.toMap()
    }
    if s.GatewayAdditionalConfigCmds != nil {
        structMap["gateway_additional_config_cmds"] = s.GatewayAdditionalConfigCmds
    }
    if s.GatewayMgmt != nil {
        structMap["gateway_mgmt"] = s.GatewayMgmt.toMap()
    }
    if s.GatewayUpdownThreshold.IsValueSet() {
        if s.GatewayUpdownThreshold.Value() != nil {
            structMap["gateway_updown_threshold"] = s.GatewayUpdownThreshold.Value()
        } else {
            structMap["gateway_updown_threshold"] = nil
        }
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Led != nil {
        structMap["led"] = s.Led.toMap()
    }
    if s.MistNac != nil {
        structMap["mist_nac"] = s.MistNac.toMap()
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.Mxedge != nil {
        structMap["mxedge"] = s.Mxedge.toMap()
    }
    if s.MxedgeMgmt != nil {
        structMap["mxedge_mgmt"] = s.MxedgeMgmt.toMap()
    }
    if s.Mxtunnels != nil {
        structMap["mxtunnels"] = s.Mxtunnels.toMap()
    }
    if s.Networks != nil {
        structMap["networks"] = s.Networks
    }
    if s.NtpServers != nil {
        structMap["ntp_servers"] = s.NtpServers
    }
    if s.Occupancy != nil {
        structMap["occupancy"] = s.Occupancy.toMap()
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.OspfAreas != nil {
        structMap["ospf_areas"] = s.OspfAreas
    }
    if s.PaloaltoNetworks != nil {
        structMap["paloalto_networks"] = s.PaloaltoNetworks.toMap()
    }
    if s.PersistConfigOnDevice != nil {
        structMap["persist_config_on_device"] = s.PersistConfigOnDevice
    }
    if s.PortMirroring != nil {
        structMap["port_mirroring"] = s.PortMirroring
    }
    if s.PortUsages != nil {
        structMap["port_usages"] = s.PortUsages
    }
    if s.Proxy != nil {
        structMap["proxy"] = s.Proxy.toMap()
    }
    if s.RadioConfig != nil {
        structMap["radio_config"] = s.RadioConfig.toMap()
    }
    if s.RadiusConfig != nil {
        structMap["radius_config"] = s.RadiusConfig.toMap()
    }
    if s.RemoteSyslog != nil {
        structMap["remote_syslog"] = s.RemoteSyslog.toMap()
    }
    if s.ReportGatt != nil {
        structMap["report_gatt"] = s.ReportGatt
    }
    if s.Rogue != nil {
        structMap["rogue"] = s.Rogue.toMap()
    }
    if s.Rtsa != nil {
        structMap["rtsa"] = s.Rtsa.toMap()
    }
    if s.SimpleAlert != nil {
        structMap["simple_alert"] = s.SimpleAlert.toMap()
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Skyatp != nil {
        structMap["skyatp"] = s.Skyatp.toMap()
    }
    if s.SnmpConfig != nil {
        structMap["snmp_config"] = s.SnmpConfig.toMap()
    }
    if s.SrxApp != nil {
        structMap["srx_app"] = s.SrxApp.toMap()
    }
    if s.SshKeys != nil {
        structMap["ssh_keys"] = s.SshKeys
    }
    if s.Ssr != nil {
        structMap["ssr"] = s.Ssr.toMap()
    }
    if s.StatusPortal != nil {
        structMap["status_portal"] = s.StatusPortal.toMap()
    }
    if s.Switch != nil {
        structMap["switch"] = s.Switch.toMap()
    }
    if s.SwitchMatching != nil {
        structMap["switch_matching"] = s.SwitchMatching.toMap()
    }
    if s.SwitchMgmt != nil {
        structMap["switch_mgmt"] = s.SwitchMgmt.toMap()
    }
    if s.SwitchUpdownThreshold.IsValueSet() {
        if s.SwitchUpdownThreshold.Value() != nil {
            structMap["switch_updown_threshold"] = s.SwitchUpdownThreshold.Value()
        } else {
            structMap["switch_updown_threshold"] = nil
        }
    }
    if s.SyntheticTest != nil {
        structMap["synthetic_test"] = s.SyntheticTest.toMap()
    }
    if s.TrackAnonymousDevices != nil {
        structMap["track_anonymous_devices"] = s.TrackAnonymousDevices
    }
    if s.TuntermMonitoring != nil {
        structMap["tunterm_monitoring"] = s.TuntermMonitoring
    }
    if s.TuntermMonitoringDisabled != nil {
        structMap["tunterm_monitoring_disabled"] = s.TuntermMonitoringDisabled
    }
    if s.TuntermMulticastConfig != nil {
        structMap["tunterm_multicast_config"] = s.TuntermMulticastConfig.toMap()
    }
    if s.Vars != nil {
        structMap["vars"] = s.Vars
    }
    if s.Vna != nil {
        structMap["vna"] = s.Vna.toMap()
    }
    if s.VrfConfig != nil {
        structMap["vrf_config"] = s.VrfConfig.toMap()
    }
    if s.VrfInstances != nil {
        structMap["vrf_instances"] = s.VrfInstances
    }
    if s.VrrpGroups != nil {
        structMap["vrrp_groups"] = s.VrrpGroups
    }
    if s.VsInstance != nil {
        structMap["vs_instance"] = s.VsInstance.toMap()
    }
    if s.WanVna != nil {
        structMap["wan_vna"] = s.WanVna.toMap()
    }
    if s.WatchedStationUrl != nil {
        structMap["watched_station_url"] = s.WatchedStationUrl
    }
    if s.WhitelistUrl != nil {
        structMap["whitelist_url"] = s.WhitelistUrl
    }
    if s.Wids != nil {
        structMap["wids"] = s.Wids.toMap()
    }
    if s.Wifi != nil {
        structMap["wifi"] = s.Wifi.toMap()
    }
    if s.WiredVna != nil {
        structMap["wired_vna"] = s.WiredVna.toMap()
    }
    if s.ZoneOccupancyAlert != nil {
        structMap["zone_occupancy_alert"] = s.ZoneOccupancyAlert.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSetting.
// It customizes the JSON unmarshaling process for SiteSetting objects.
func (s *SiteSetting) UnmarshalJSON(input []byte) error {
    var temp siteSetting
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "acl_policies", "acl_tags", "additional_config_cmds", "analytic", "ap_matching", "ap_port_config", "ap_updown_threshold", "auto_placement", "auto_upgrade", "blacklist_url", "ble_config", "config_auto_revert", "config_push_policy", "created_time", "critical_url_monitoring", "device_updown_threshold", "dhcp_snooping", "disabled_system_defined_port_usages", "dns_servers", "dns_suffix", "engagement", "evpn_options", "extra_routes", "extra_routes6", "flags", "for_site", "gateway", "gateway_additional_config_cmds", "gateway_mgmt", "gateway_updown_threshold", "id", "led", "mist_nac", "modified_time", "mxedge", "mxedge_mgmt", "mxtunnels", "networks", "ntp_servers", "occupancy", "org_id", "ospf_areas", "paloalto_networks", "persist_config_on_device", "port_mirroring", "port_usages", "proxy", "radio_config", "radius_config", "remote_syslog", "report_gatt", "rogue", "rtsa", "simple_alert", "site_id", "skyatp", "snmp_config", "srx_app", "ssh_keys", "ssr", "status_portal", "switch", "switch_matching", "switch_mgmt", "switch_updown_threshold", "synthetic_test", "track_anonymous_devices", "tunterm_monitoring", "tunterm_monitoring_disabled", "tunterm_multicast_config", "vars", "vna", "vrf_config", "vrf_instances", "vrrp_groups", "vs_instance", "wan_vna", "watched_station_url", "whitelist_url", "wids", "wifi", "wired_vna", "zone_occupancy_alert")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AclPolicies = temp.AclPolicies
    s.AclTags = temp.AclTags
    s.AdditionalConfigCmds = temp.AdditionalConfigCmds
    s.Analytic = temp.Analytic
    s.ApMatching = temp.ApMatching
    s.ApPortConfig = temp.ApPortConfig
    s.ApUpdownThreshold = temp.ApUpdownThreshold
    s.AutoPlacement = temp.AutoPlacement
    s.AutoUpgrade = temp.AutoUpgrade
    s.BlacklistUrl = temp.BlacklistUrl
    s.BleConfig = temp.BleConfig
    s.ConfigAutoRevert = temp.ConfigAutoRevert
    s.ConfigPushPolicy = temp.ConfigPushPolicy
    s.CreatedTime = temp.CreatedTime
    s.CriticalUrlMonitoring = temp.CriticalUrlMonitoring
    s.DeviceUpdownThreshold = temp.DeviceUpdownThreshold
    s.DhcpSnooping = temp.DhcpSnooping
    s.DisabledSystemDefinedPortUsages = temp.DisabledSystemDefinedPortUsages
    s.DnsServers = temp.DnsServers
    s.DnsSuffix = temp.DnsSuffix
    s.Engagement = temp.Engagement
    s.EvpnOptions = temp.EvpnOptions
    s.ExtraRoutes = temp.ExtraRoutes
    s.ExtraRoutes6 = temp.ExtraRoutes6
    s.Flags = temp.Flags
    s.ForSite = temp.ForSite
    s.Gateway = temp.Gateway
    s.GatewayAdditionalConfigCmds = temp.GatewayAdditionalConfigCmds
    s.GatewayMgmt = temp.GatewayMgmt
    s.GatewayUpdownThreshold = temp.GatewayUpdownThreshold
    s.Id = temp.Id
    s.Led = temp.Led
    s.MistNac = temp.MistNac
    s.ModifiedTime = temp.ModifiedTime
    s.Mxedge = temp.Mxedge
    s.MxedgeMgmt = temp.MxedgeMgmt
    s.Mxtunnels = temp.Mxtunnels
    s.Networks = temp.Networks
    s.NtpServers = temp.NtpServers
    s.Occupancy = temp.Occupancy
    s.OrgId = temp.OrgId
    s.OspfAreas = temp.OspfAreas
    s.PaloaltoNetworks = temp.PaloaltoNetworks
    s.PersistConfigOnDevice = temp.PersistConfigOnDevice
    s.PortMirroring = temp.PortMirroring
    s.PortUsages = temp.PortUsages
    s.Proxy = temp.Proxy
    s.RadioConfig = temp.RadioConfig
    s.RadiusConfig = temp.RadiusConfig
    s.RemoteSyslog = temp.RemoteSyslog
    s.ReportGatt = temp.ReportGatt
    s.Rogue = temp.Rogue
    s.Rtsa = temp.Rtsa
    s.SimpleAlert = temp.SimpleAlert
    s.SiteId = temp.SiteId
    s.Skyatp = temp.Skyatp
    s.SnmpConfig = temp.SnmpConfig
    s.SrxApp = temp.SrxApp
    s.SshKeys = temp.SshKeys
    s.Ssr = temp.Ssr
    s.StatusPortal = temp.StatusPortal
    s.Switch = temp.Switch
    s.SwitchMatching = temp.SwitchMatching
    s.SwitchMgmt = temp.SwitchMgmt
    s.SwitchUpdownThreshold = temp.SwitchUpdownThreshold
    s.SyntheticTest = temp.SyntheticTest
    s.TrackAnonymousDevices = temp.TrackAnonymousDevices
    s.TuntermMonitoring = temp.TuntermMonitoring
    s.TuntermMonitoringDisabled = temp.TuntermMonitoringDisabled
    s.TuntermMulticastConfig = temp.TuntermMulticastConfig
    s.Vars = temp.Vars
    s.Vna = temp.Vna
    s.VrfConfig = temp.VrfConfig
    s.VrfInstances = temp.VrfInstances
    s.VrrpGroups = temp.VrrpGroups
    s.VsInstance = temp.VsInstance
    s.WanVna = temp.WanVna
    s.WatchedStationUrl = temp.WatchedStationUrl
    s.WhitelistUrl = temp.WhitelistUrl
    s.Wids = temp.Wids
    s.Wifi = temp.Wifi
    s.WiredVna = temp.WiredVna
    s.ZoneOccupancyAlert = temp.ZoneOccupancyAlert
    return nil
}

// siteSetting is a temporary struct used for validating the fields of SiteSetting.
type siteSetting  struct {
    AclPolicies                     []AclPolicy                            `json:"acl_policies,omitempty"`
    AclTags                         map[string]AclTag                      `json:"acl_tags,omitempty"`
    AdditionalConfigCmds            []string                               `json:"additional_config_cmds,omitempty"`
    Analytic                        *SiteSettingAnalytic                   `json:"analytic,omitempty"`
    ApMatching                      *SiteSettingApMatching                 `json:"ap_matching,omitempty"`
    ApPortConfig                    *SiteSettingApPortConfig               `json:"ap_port_config,omitempty"`
    ApUpdownThreshold               Optional[int]                          `json:"ap_updown_threshold"`
    AutoPlacement                   *SiteSettingAutoPlacement              `json:"auto_placement,omitempty"`
    AutoUpgrade                     *SiteSettingAutoUpgrade                `json:"auto_upgrade,omitempty"`
    BlacklistUrl                    *string                                `json:"blacklist_url,omitempty"`
    BleConfig                       *BleConfig                             `json:"ble_config,omitempty"`
    ConfigAutoRevert                *bool                                  `json:"config_auto_revert,omitempty"`
    ConfigPushPolicy                *SiteSettingConfigPushPolicy           `json:"config_push_policy,omitempty"`
    CreatedTime                     *float64                               `json:"created_time,omitempty"`
    CriticalUrlMonitoring           *SiteSettingCriticalUrlMonitoring      `json:"critical_url_monitoring,omitempty"`
    DeviceUpdownThreshold           *int                                   `json:"device_updown_threshold,omitempty"`
    DhcpSnooping                    *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
    DisabledSystemDefinedPortUsages []string                               `json:"disabled_system_defined_port_usages,omitempty"`
    DnsServers                      []string                               `json:"dns_servers,omitempty"`
    DnsSuffix                       []string                               `json:"dns_suffix,omitempty"`
    Engagement                      *SiteEngagement                        `json:"engagement,omitempty"`
    EvpnOptions                     *EvpnOptions                           `json:"evpn_options,omitempty"`
    ExtraRoutes                     map[string]ExtraRouteProperties        `json:"extra_routes,omitempty"`
    ExtraRoutes6                    map[string]ExtraRoute6Properties       `json:"extra_routes6,omitempty"`
    Flags                           map[string]string                      `json:"flags,omitempty"`
    ForSite                         *bool                                  `json:"for_site,omitempty"`
    Gateway                         *GatewayTemplate                       `json:"gateway,omitempty"`
    GatewayAdditionalConfigCmds     []string                               `json:"gateway_additional_config_cmds,omitempty"`
    GatewayMgmt                     *SiteSettingGatewayMgmt                `json:"gateway_mgmt,omitempty"`
    GatewayUpdownThreshold          Optional[int]                          `json:"gateway_updown_threshold"`
    Id                              *uuid.UUID                             `json:"id,omitempty"`
    Led                             *ApLed                                 `json:"led,omitempty"`
    MistNac                         *SwitchMistNac                         `json:"mist_nac,omitempty"`
    ModifiedTime                    *float64                               `json:"modified_time,omitempty"`
    Mxedge                          *SiteSettingMxedge                     `json:"mxedge,omitempty"`
    MxedgeMgmt                      *MxedgeMgmt                            `json:"mxedge_mgmt,omitempty"`
    Mxtunnels                       *SiteMxtunnel                          `json:"mxtunnels,omitempty"`
    Networks                        map[string]SwitchNetwork               `json:"networks,omitempty"`
    NtpServers                      []string                               `json:"ntp_servers,omitempty"`
    Occupancy                       *SiteOccupancyAnalytics                `json:"occupancy,omitempty"`
    OrgId                           *uuid.UUID                             `json:"org_id,omitempty"`
    OspfAreas                       map[string]OspfAreas                   `json:"ospf_areas,omitempty"`
    PaloaltoNetworks                *SiteSettingPaloaltoNetworks           `json:"paloalto_networks,omitempty"`
    PersistConfigOnDevice           *bool                                  `json:"persist_config_on_device,omitempty"`
    PortMirroring                   map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
    PortUsages                      map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
    Proxy                           *Proxy                                 `json:"proxy,omitempty"`
    RadioConfig                     *ApRadio                               `json:"radio_config,omitempty"`
    RadiusConfig                    *RadiusConfig                          `json:"radius_config,omitempty"`
    RemoteSyslog                    *RemoteSyslog                          `json:"remote_syslog,omitempty"`
    ReportGatt                      *bool                                  `json:"report_gatt,omitempty"`
    Rogue                           *SiteRogue                             `json:"rogue,omitempty"`
    Rtsa                            *SiteSettingRtsa                       `json:"rtsa,omitempty"`
    SimpleAlert                     *SimpleAlert                           `json:"simple_alert,omitempty"`
    SiteId                          *uuid.UUID                             `json:"site_id,omitempty"`
    Skyatp                          *SiteSettingSkyatp                     `json:"skyatp,omitempty"`
    SnmpConfig                      *SnmpConfig                            `json:"snmp_config,omitempty"`
    SrxApp                          *SiteSettingSrxApp                     `json:"srx_app,omitempty"`
    SshKeys                         []string                               `json:"ssh_keys,omitempty"`
    Ssr                             *SiteSettingSsr                        `json:"ssr,omitempty"`
    StatusPortal                    *SiteSettingStatusPortal               `json:"status_portal,omitempty"`
    Switch                          *NetworkTemplate                       `json:"switch,omitempty"`
    SwitchMatching                  *SwitchMatching                        `json:"switch_matching,omitempty"`
    SwitchMgmt                      *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
    SwitchUpdownThreshold           Optional[int]                          `json:"switch_updown_threshold"`
    SyntheticTest                   *SynthetictestConfig                   `json:"synthetic_test,omitempty"`
    TrackAnonymousDevices           *bool                                  `json:"track_anonymous_devices,omitempty"`
    TuntermMonitoring               []TuntermMonitoringItem                `json:"tunterm_monitoring,omitempty"`
    TuntermMonitoringDisabled       *bool                                  `json:"tunterm_monitoring_disabled,omitempty"`
    TuntermMulticastConfig          *SiteSettingTuntermMulticastConfig     `json:"tunterm_multicast_config,omitempty"`
    Vars                            map[string]string                      `json:"vars,omitempty"`
    Vna                             *SiteSettingVna                        `json:"vna,omitempty"`
    VrfConfig                       *VrfConfig                             `json:"vrf_config,omitempty"`
    VrfInstances                    map[string]VrfInstance                 `json:"vrf_instances,omitempty"`
    VrrpGroups                      map[string]VrrpGroup                   `json:"vrrp_groups,omitempty"`
    VsInstance                      *VsInstance                            `json:"vs_instance,omitempty"`
    WanVna                          *SiteSettingWanVna                     `json:"wan_vna,omitempty"`
    WatchedStationUrl               *string                                `json:"watched_station_url,omitempty"`
    WhitelistUrl                    *string                                `json:"whitelist_url,omitempty"`
    Wids                            *SiteWids                              `json:"wids,omitempty"`
    Wifi                            *SiteWifi                              `json:"wifi,omitempty"`
    WiredVna                        *SiteSettingWiredVna                   `json:"wired_vna,omitempty"`
    ZoneOccupancyAlert              *SiteZoneOccupancyAlert                `json:"zone_occupancy_alert,omitempty"`
}