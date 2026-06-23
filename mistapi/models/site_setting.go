// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// SiteSetting represents a SiteSetting struct.
// Configuration settings applied at the site level
type SiteSetting struct {
	// List of ACL policy definitions
	AclPolicies []AclPolicy `json:"acl_policies,omitempty"`
	// ACL Tags to identify traffic source or destination. Key name is the tag name
	AclTags map[string]AclTag `json:"acl_tags,omitempty"`
	// additional CLI commands to append to the generated Junos config. **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	// whether to allow Mist to look at this org
	AllowMist *bool `json:"allow_mist,omitempty"`
	// Advanced analytics feature settings for a site
	Analytic *SiteSettingAnalytic `json:"analytic,omitempty"`
	// Rules for applying AP port configuration by AP model or name
	ApMatching *SiteSettingApMatching `json:"ap_matching,omitempty"`
	// AP Ethernet port configuration overrides by model
	ApPortConfig *SiteSettingApPortConfig `json:"ap_port_config,omitempty"`
	// AP Synthetic Test configuration
	ApSyntheticTest *SiteSettingApSyntheticTest `json:"ap_synthetic_test,omitempty"`
	// Enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.
	ApUpdownThreshold Optional[int] `json:"ap_updown_threshold"`
	// Automatically determined AP placement coordinates and orientation
	AutoPlacement *SiteSettingAutoPlacement `json:"auto_placement,omitempty"`
	// Automatic AP firmware upgrade policy
	AutoUpgrade *SiteSettingAutoUpgrade `json:"auto_upgrade,omitempty"`
	// Automatic AP ESL firmware upgrade policy. When both firmware and ESL auto-upgrade are enabled, ESL upgrade will be done only after firmware upgrade
	AutoUpgradeEsl *SiteSettingAutoUpgradeEsl `json:"auto_upgrade_esl,omitempty"`
	// Whether line cards are included in automatic switch upgrades
	AutoUpgradeLinecard *bool `json:"auto_upgrade_linecard,omitempty"`
	// enable threshold-based bgp neighbor down delivery.
	BgpNeighborUpdownThreshold Optional[int] `json:"bgp_neighbor_updown_threshold"`
	// Read-only URL for the site blacklist file
	BlacklistUrl *string `json:"blacklist_url,omitempty"`
	// Bluetooth Low Energy beacon and asset advertising settings for an AP
	BleConfig *BleConfig `json:"ble_config,omitempty"`
	// Whether to enable ap auto config revert
	ConfigAutoRevert *bool `json:"config_auto_revert,omitempty"`
	// Mist also uses some heuristic rules to prevent destructive configs from being pushed
	ConfigPushPolicy *SiteSettingConfigPushPolicy `json:"config_push_policy,omitempty"`
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	// Critical URLs whose latency is measured and included in site health
	CriticalUrlMonitoring *SiteSettingCriticalUrlMonitoring `json:"critical_url_monitoring,omitempty"`
	// By default, device_updown_threshold, if set, will apply to all devices types if different values for specific device type is desired, use the following
	DeviceUpdownThreshold Optional[int] `json:"device_updown_threshold"`
	// DHCP snooping security settings
	DhcpSnooping *DhcpSnooping `json:"dhcp_snooping,omitempty"`
	// If some system-default port usages are not desired - namely, ap / iot / uplink
	DisabledSystemDefinedPortUsages []SystemDefinedPortUsagesEnum `json:"disabled_system_defined_port_usages,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// Whether UNII-4 channels are enabled for the site
	EnableUnii4 *bool `json:"enable_unii_4,omitempty"`
	// Engagement analytics dwell-time rules for classifying site visits. If hours is omitted, rules apply every day from 00:00 to 23:59. Multiple ranges for the same day are not supported.
	Engagement *SiteEngagement `json:"engagement,omitempty"`
	// EVPN topology generation options for campus fabric configuration
	EvpnOptions *EvpnOptions `json:"evpn_options,omitempty"`
	// Property key is the destination CIDR (e.g. "10.0.0.0/8")
	ExtraRoutes map[string]ExtraRoute `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64")
	ExtraRoutes6 map[string]ExtraRoute6 `json:"extra_routes6,omitempty"`
	// Name/val pair objects for location engine to use
	Flags map[string]string `json:"flags,omitempty"`
	// Whether this settings object is scoped to a site
	ForSite *bool `json:"for_site,omitempty"`
	// Gateway Template is applied to a site for gateway(s) in a site.
	Gateway *GatewayTemplate `json:"gateway,omitempty"`
	// additional CLI commands to append to the generated Junos config. **Note**: no check is done
	GatewayAdditionalConfigCmds []string `json:"gateway_additional_config_cmds,omitempty"`
	// Gateway management-plane and access settings
	GatewayMgmt *GatewayMgmt `json:"gateway_mgmt,omitempty"`
	// enable threshold-based gateway tunnel (secure edge tunnels) up-down delivery.
	GatewayTunnelUpdownThreshold Optional[int] `json:"gateway_tunnel_updown_threshold"`
	// Enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.
	GatewayUpdownThreshold Optional[int] `json:"gateway_updown_threshold"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// IoT proxy configuration for the site
	Iotproxy *Iotproxy `json:"iotproxy,omitempty"`
	// Site-level Juniper SRX integration settings
	JuniperSrx *SiteSettingJuniperSrx `json:"juniper_srx,omitempty"`
	// Indicator light settings for an access point
	Led *ApLed `json:"led,omitempty"`
	// Marvis automation and client settings
	Marvis *Marvis `json:"marvis,omitempty"`
	// Mist NAC RadSec settings for a switch
	MistNac *SwitchMistNac `json:"mist_nac,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	// Service settings for the site Mist Edge cluster
	Mxedge *SiteSettingMxedge `json:"mxedge,omitempty"`
	// Management settings for a Mist Edge appliance
	MxedgeMgmt *MxedgeMgmt `json:"mxedge_mgmt,omitempty"`
	// Site Mist Tunnel configuration for tunneling AP user VLANs to Mist Edge tunnel peers
	Mxtunnels *SiteMxtunnel `json:"mxtunnels,omitempty"`
	// Property key is network name
	Networks map[string]SwitchNetwork `json:"networks,omitempty"`
	// List of NTP servers
	NtpServers []string `json:"ntp_servers,omitempty"`
	// Analytics settings for site occupancy
	Occupancy *SiteOccupancyAnalytics `json:"occupancy,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Junos OSPF areas. Property key is the OSPF Area (Area should be a number (0-255) / IP address)
	OspfAreas map[string]OspfArea `json:"ospf_areas,omitempty"`
	// Palo Alto Networks integration settings for the site
	PaloaltoNetworks *SiteSettingPaloaltoNetworks `json:"paloalto_networks,omitempty"`
	// Whether to store the config on AP
	PersistConfigOnDevice *bool `json:"persist_config_on_device,omitempty"`
	// Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed
	PortMirroring map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
	// Property key is the port usage name. Defines the profiles of port configuration configured on the switch
	PortUsages map[string]SwitchPortUsage `json:"port_usages,omitempty"`
	// Proxy Configuration to talk to Mist
	Proxy *Proxy `json:"proxy,omitempty"`
	// Radio configuration settings for an access point
	RadioConfig *ApRadio `json:"radio_config,omitempty"`
	// Switch RADIUS authentication and accounting configuration
	RadiusConfig *SwitchRadiusConfig `json:"radius_config,omitempty"`
	// Remote syslog forwarding settings
	RemoteSyslog *RemoteSyslog `json:"remote_syslog,omitempty"`
	// By default, only the configuration generated by Mist is cleaned up during the configuration process. If `true`, all the existing configuration will be removed.
	RemoveExistingConfigs *bool `json:"remove_existing_configs,omitempty"`
	// Whether AP should periodically connect to BLE devices and report GATT device info (device name, manufacturer name, serial number, battery %, temperature, humidity)
	ReportGatt *bool `json:"report_gatt,omitempty"`
	// Rogue AP detection settings for a site
	Rogue *SiteRogue `json:"rogue,omitempty"`
	// Switch routing policies keyed by routing policy name
	RoutingPolicies map[string]SwRoutingPolicy `json:"routing_policies,omitempty"`
	// Managed mobility and asset tracking settings
	Rtsa *SiteSettingRtsa `json:"rtsa,omitempty"`
	// Heuristic alert thresholds used when a Marvis subscription is unavailable
	SimpleAlert *SimpleAlert `json:"simple_alert,omitempty"`
	// Unique identifier of a Mist site
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// Sky ATP threat intelligence settings for the site
	Skyatp *SiteSettingSkyatp `json:"skyatp,omitempty"`
	// Site SLE threshold overrides for capacity, coverage, throughput, and time to connect
	SleThresholds *SleThresholds `json:"sle_thresholds,omitempty"`
	// SNMP configuration for managed network devices
	SnmpConfig *SnmpConfig `json:"snmp_config,omitempty"`
	// Juniper SRX application visibility settings for the site
	SrxApp *SiteSettingSrxApp `json:"srx_app,omitempty"`
	// When limit_ssh_access = true in Org Setting, list of SSH public keys provided by Mist Support to install onto APs (see Org:Setting)
	SshKeys []string `json:"ssh_keys,omitempty"`
	// SSR management settings for device onboarding and connectivity
	Ssr *SettingSsr `json:"ssr,omitempty"`
	// End-user status portal settings for the site
	StatusPortal *SiteSettingStatusPortal `json:"status_portal,omitempty"`
	// Site switch settings combining a network template and auto-upgrade controls
	Switch *SiteSettingSwitch `json:"switch,omitempty"`
	// Defines custom switch configuration based on different criteria
	SwitchMatching *SwitchMatching `json:"switch_matching,omitempty"`
	// Switch management-plane access and proxy settings
	SwitchMgmt *SwitchMgmt `json:"switch_mgmt,omitempty"`
	// Enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.
	SwitchUpdownThreshold Optional[int] `json:"switch_updown_threshold"`
	// Synthetic test configuration for Marvis Minis
	SyntheticTest *SynthetictestConfig `json:"synthetic_test,omitempty"`
	// Whether to track anonymous BLE assets (requires ‘track_asset’  enabled)
	TrackAnonymousDevices *bool `json:"track_anonymous_devices,omitempty"`
	// Tunnel termination monitoring checks
	TuntermMonitoring []TuntermMonitoringItem `json:"tunterm_monitoring,omitempty"`
	// Whether tunnel termination monitoring is disabled for the site
	TuntermMonitoringDisabled *bool `json:"tunterm_monitoring_disabled,omitempty"`
	// Multicast forwarding settings for tunnel termination at the site
	TuntermMulticastConfig *SiteSettingTuntermMulticastConfig `json:"tunterm_multicast_config,omitempty"`
	// AP Uplink port configuration
	UplinkPortConfig *ApUplinkPortConfig `json:"uplink_port_config,omitempty"`
	// by default, we only honor description provided in port_config. This allows fallback to those defined in port_usages
	UsesDescriptionFromPortUsage *bool `json:"uses_description_from_port_usage,omitempty"`
	// Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars
	Vars map[string]string `json:"vars,omitempty"`
	// Optional annotations for vars defined in this site. Keys match var names; values describe the var purpose and type for UI auto-complete.
	VarsAnnotations map[string]VarsAnnotation `json:"vars_annotations,omitempty"`
	// Virtual Network Assistant settings for AP, switch, and gateway experiences at a site
	Vna *SiteSettingVna `json:"vna,omitempty"`
	// enable threshold-based vpn path down delivery.
	VpnPathUpdownThreshold Optional[int] `json:"vpn_path_updown_threshold"`
	// enable threshold-based vpn peer down delivery.
	VpnPeerUpdownThreshold Optional[int] `json:"vpn_peer_updown_threshold"`
	// VRF enablement settings applied when supported on the device
	VrfConfig *VrfConfig `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances map[string]SwitchVrfInstance `json:"vrf_instances,omitempty"`
	// Property key is the vrrp group
	VrrpGroups map[string]VrrpGroup `json:"vrrp_groups,omitempty"`
	// Optional, for EX9200 only to segregate virtual-switches. Property key is the instance name
	VsInstance map[string]VsInstanceProperty `json:"vs_instance,omitempty"`
	// WAN Virtual Network Assistant settings for the site
	WanVna *SiteSettingWanVna `json:"wan_vna,omitempty"`
	// Read-only URL for the watched station list file
	WatchedStationUrl *string `json:"watched_station_url,omitempty"`
	// Read-only URL for the site whitelist file
	WhitelistUrl *string `json:"whitelist_url,omitempty"`
	// Wireless intrusion detection settings for a site
	Wids *SiteWids `json:"wids,omitempty"`
	// Wi-Fi configuration settings for a site
	Wifi *SiteWifi `json:"wifi,omitempty"`
	// Wired Virtual Network Assistant settings for the site
	WiredVna *SiteSettingWiredVna `json:"wired_vna,omitempty"`
	// Zone occupancy alert settings for a site
	ZoneOccupancyAlert   *SiteZoneOccupancyAlert `json:"zone_occupancy_alert,omitempty"`
	AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSetting,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSetting) String() string {
	return fmt.Sprintf(
		"SiteSetting[AclPolicies=%v, AclTags=%v, AdditionalConfigCmds=%v, AllowMist=%v, Analytic=%v, ApMatching=%v, ApPortConfig=%v, ApSyntheticTest=%v, ApUpdownThreshold=%v, AutoPlacement=%v, AutoUpgrade=%v, AutoUpgradeEsl=%v, AutoUpgradeLinecard=%v, BgpNeighborUpdownThreshold=%v, BlacklistUrl=%v, BleConfig=%v, ConfigAutoRevert=%v, ConfigPushPolicy=%v, CreatedTime=%v, CriticalUrlMonitoring=%v, DeviceUpdownThreshold=%v, DhcpSnooping=%v, DisabledSystemDefinedPortUsages=%v, DnsServers=%v, DnsSuffix=%v, EnableUnii4=%v, Engagement=%v, EvpnOptions=%v, ExtraRoutes=%v, ExtraRoutes6=%v, Flags=%v, ForSite=%v, Gateway=%v, GatewayAdditionalConfigCmds=%v, GatewayMgmt=%v, GatewayTunnelUpdownThreshold=%v, GatewayUpdownThreshold=%v, Id=%v, Iotproxy=%v, JuniperSrx=%v, Led=%v, Marvis=%v, MistNac=%v, ModifiedTime=%v, Mxedge=%v, MxedgeMgmt=%v, Mxtunnels=%v, Networks=%v, NtpServers=%v, Occupancy=%v, OrgId=%v, OspfAreas=%v, PaloaltoNetworks=%v, PersistConfigOnDevice=%v, PortMirroring=%v, PortUsages=%v, Proxy=%v, RadioConfig=%v, RadiusConfig=%v, RemoteSyslog=%v, RemoveExistingConfigs=%v, ReportGatt=%v, Rogue=%v, RoutingPolicies=%v, Rtsa=%v, SimpleAlert=%v, SiteId=%v, Skyatp=%v, SleThresholds=%v, SnmpConfig=%v, SrxApp=%v, SshKeys=%v, Ssr=%v, StatusPortal=%v, Switch=%v, SwitchMatching=%v, SwitchMgmt=%v, SwitchUpdownThreshold=%v, SyntheticTest=%v, TrackAnonymousDevices=%v, TuntermMonitoring=%v, TuntermMonitoringDisabled=%v, TuntermMulticastConfig=%v, UplinkPortConfig=%v, UsesDescriptionFromPortUsage=%v, Vars=%v, VarsAnnotations=%v, Vna=%v, VpnPathUpdownThreshold=%v, VpnPeerUpdownThreshold=%v, VrfConfig=%v, VrfInstances=%v, VrrpGroups=%v, VsInstance=%v, WanVna=%v, WatchedStationUrl=%v, WhitelistUrl=%v, Wids=%v, Wifi=%v, WiredVna=%v, ZoneOccupancyAlert=%v, AdditionalProperties=%v]",
		s.AclPolicies, s.AclTags, s.AdditionalConfigCmds, s.AllowMist, s.Analytic, s.ApMatching, s.ApPortConfig, s.ApSyntheticTest, s.ApUpdownThreshold, s.AutoPlacement, s.AutoUpgrade, s.AutoUpgradeEsl, s.AutoUpgradeLinecard, s.BgpNeighborUpdownThreshold, s.BlacklistUrl, s.BleConfig, s.ConfigAutoRevert, s.ConfigPushPolicy, s.CreatedTime, s.CriticalUrlMonitoring, s.DeviceUpdownThreshold, s.DhcpSnooping, s.DisabledSystemDefinedPortUsages, s.DnsServers, s.DnsSuffix, s.EnableUnii4, s.Engagement, s.EvpnOptions, s.ExtraRoutes, s.ExtraRoutes6, s.Flags, s.ForSite, s.Gateway, s.GatewayAdditionalConfigCmds, s.GatewayMgmt, s.GatewayTunnelUpdownThreshold, s.GatewayUpdownThreshold, s.Id, s.Iotproxy, s.JuniperSrx, s.Led, s.Marvis, s.MistNac, s.ModifiedTime, s.Mxedge, s.MxedgeMgmt, s.Mxtunnels, s.Networks, s.NtpServers, s.Occupancy, s.OrgId, s.OspfAreas, s.PaloaltoNetworks, s.PersistConfigOnDevice, s.PortMirroring, s.PortUsages, s.Proxy, s.RadioConfig, s.RadiusConfig, s.RemoteSyslog, s.RemoveExistingConfigs, s.ReportGatt, s.Rogue, s.RoutingPolicies, s.Rtsa, s.SimpleAlert, s.SiteId, s.Skyatp, s.SleThresholds, s.SnmpConfig, s.SrxApp, s.SshKeys, s.Ssr, s.StatusPortal, s.Switch, s.SwitchMatching, s.SwitchMgmt, s.SwitchUpdownThreshold, s.SyntheticTest, s.TrackAnonymousDevices, s.TuntermMonitoring, s.TuntermMonitoringDisabled, s.TuntermMulticastConfig, s.UplinkPortConfig, s.UsesDescriptionFromPortUsage, s.Vars, s.VarsAnnotations, s.Vna, s.VpnPathUpdownThreshold, s.VpnPeerUpdownThreshold, s.VrfConfig, s.VrfInstances, s.VrrpGroups, s.VsInstance, s.WanVna, s.WatchedStationUrl, s.WhitelistUrl, s.Wids, s.Wifi, s.WiredVna, s.ZoneOccupancyAlert, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSetting.
// It customizes the JSON marshaling process for SiteSetting objects.
func (s SiteSetting) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"acl_policies", "acl_tags", "additional_config_cmds", "allow_mist", "analytic", "ap_matching", "ap_port_config", "ap_synthetic_test", "ap_updown_threshold", "auto_placement", "auto_upgrade", "auto_upgrade_esl", "auto_upgrade_linecard", "bgp_neighbor_updown_threshold", "blacklist_url", "ble_config", "config_auto_revert", "config_push_policy", "created_time", "critical_url_monitoring", "device_updown_threshold", "dhcp_snooping", "disabled_system_defined_port_usages", "dns_servers", "dns_suffix", "enable_unii_4", "engagement", "evpn_options", "extra_routes", "extra_routes6", "flags", "for_site", "gateway", "gateway_additional_config_cmds", "gateway_mgmt", "gateway_tunnel_updown_threshold", "gateway_updown_threshold", "id", "iotproxy", "juniper_srx", "led", "marvis", "mist_nac", "modified_time", "mxedge", "mxedge_mgmt", "mxtunnels", "networks", "ntp_servers", "occupancy", "org_id", "ospf_areas", "paloalto_networks", "persist_config_on_device", "port_mirroring", "port_usages", "proxy", "radio_config", "radius_config", "remote_syslog", "remove_existing_configs", "report_gatt", "rogue", "routing_policies", "rtsa", "simple_alert", "site_id", "skyatp", "sle_thresholds", "snmp_config", "srx_app", "ssh_keys", "ssr", "status_portal", "switch", "switch_matching", "switch_mgmt", "switch_updown_threshold", "synthetic_test", "track_anonymous_devices", "tunterm_monitoring", "tunterm_monitoring_disabled", "tunterm_multicast_config", "uplink_port_config", "uses_description_from_port_usage", "vars", "vars_annotations", "vna", "vpn_path_updown_threshold", "vpn_peer_updown_threshold", "vrf_config", "vrf_instances", "vrrp_groups", "vs_instance", "wan_vna", "watched_station_url", "whitelist_url", "wids", "wifi", "wired_vna", "zone_occupancy_alert"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSetting object to a map representation for JSON marshaling.
func (s SiteSetting) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AclPolicies != nil {
		structMap["acl_policies"] = s.AclPolicies
	}
	if s.AclTags != nil {
		structMap["acl_tags"] = s.AclTags
	}
	if s.AdditionalConfigCmds != nil {
		structMap["additional_config_cmds"] = s.AdditionalConfigCmds
	}
	if s.AllowMist != nil {
		structMap["allow_mist"] = s.AllowMist
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
	if s.ApSyntheticTest != nil {
		structMap["ap_synthetic_test"] = s.ApSyntheticTest.toMap()
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
	if s.AutoUpgradeEsl != nil {
		structMap["auto_upgrade_esl"] = s.AutoUpgradeEsl.toMap()
	}
	if s.AutoUpgradeLinecard != nil {
		structMap["auto_upgrade_linecard"] = s.AutoUpgradeLinecard
	}
	if s.BgpNeighborUpdownThreshold.IsValueSet() {
		if s.BgpNeighborUpdownThreshold.Value() != nil {
			structMap["bgp_neighbor_updown_threshold"] = s.BgpNeighborUpdownThreshold.Value()
		} else {
			structMap["bgp_neighbor_updown_threshold"] = nil
		}
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
	if s.DeviceUpdownThreshold.IsValueSet() {
		if s.DeviceUpdownThreshold.Value() != nil {
			structMap["device_updown_threshold"] = s.DeviceUpdownThreshold.Value()
		} else {
			structMap["device_updown_threshold"] = nil
		}
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
	if s.EnableUnii4 != nil {
		structMap["enable_unii_4"] = s.EnableUnii4
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
	if s.GatewayTunnelUpdownThreshold.IsValueSet() {
		if s.GatewayTunnelUpdownThreshold.Value() != nil {
			structMap["gateway_tunnel_updown_threshold"] = s.GatewayTunnelUpdownThreshold.Value()
		} else {
			structMap["gateway_tunnel_updown_threshold"] = nil
		}
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
	if s.Iotproxy != nil {
		structMap["iotproxy"] = s.Iotproxy.toMap()
	}
	if s.JuniperSrx != nil {
		structMap["juniper_srx"] = s.JuniperSrx.toMap()
	}
	if s.Led != nil {
		structMap["led"] = s.Led.toMap()
	}
	if s.Marvis != nil {
		structMap["marvis"] = s.Marvis.toMap()
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
	if s.RemoveExistingConfigs != nil {
		structMap["remove_existing_configs"] = s.RemoveExistingConfigs
	}
	if s.ReportGatt != nil {
		structMap["report_gatt"] = s.ReportGatt
	}
	if s.Rogue != nil {
		structMap["rogue"] = s.Rogue.toMap()
	}
	if s.RoutingPolicies != nil {
		structMap["routing_policies"] = s.RoutingPolicies
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
	if s.SleThresholds != nil {
		structMap["sle_thresholds"] = s.SleThresholds.toMap()
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
	if s.UplinkPortConfig != nil {
		structMap["uplink_port_config"] = s.UplinkPortConfig.toMap()
	}
	if s.UsesDescriptionFromPortUsage != nil {
		structMap["uses_description_from_port_usage"] = s.UsesDescriptionFromPortUsage
	}
	if s.Vars != nil {
		structMap["vars"] = s.Vars
	}
	if s.VarsAnnotations != nil {
		structMap["vars_annotations"] = s.VarsAnnotations
	}
	if s.Vna != nil {
		structMap["vna"] = s.Vna.toMap()
	}
	if s.VpnPathUpdownThreshold.IsValueSet() {
		if s.VpnPathUpdownThreshold.Value() != nil {
			structMap["vpn_path_updown_threshold"] = s.VpnPathUpdownThreshold.Value()
		} else {
			structMap["vpn_path_updown_threshold"] = nil
		}
	}
	if s.VpnPeerUpdownThreshold.IsValueSet() {
		if s.VpnPeerUpdownThreshold.Value() != nil {
			structMap["vpn_peer_updown_threshold"] = s.VpnPeerUpdownThreshold.Value()
		} else {
			structMap["vpn_peer_updown_threshold"] = nil
		}
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
		structMap["vs_instance"] = s.VsInstance
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
	var temp tempSiteSetting
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acl_policies", "acl_tags", "additional_config_cmds", "allow_mist", "analytic", "ap_matching", "ap_port_config", "ap_synthetic_test", "ap_updown_threshold", "auto_placement", "auto_upgrade", "auto_upgrade_esl", "auto_upgrade_linecard", "bgp_neighbor_updown_threshold", "blacklist_url", "ble_config", "config_auto_revert", "config_push_policy", "created_time", "critical_url_monitoring", "device_updown_threshold", "dhcp_snooping", "disabled_system_defined_port_usages", "dns_servers", "dns_suffix", "enable_unii_4", "engagement", "evpn_options", "extra_routes", "extra_routes6", "flags", "for_site", "gateway", "gateway_additional_config_cmds", "gateway_mgmt", "gateway_tunnel_updown_threshold", "gateway_updown_threshold", "id", "iotproxy", "juniper_srx", "led", "marvis", "mist_nac", "modified_time", "mxedge", "mxedge_mgmt", "mxtunnels", "networks", "ntp_servers", "occupancy", "org_id", "ospf_areas", "paloalto_networks", "persist_config_on_device", "port_mirroring", "port_usages", "proxy", "radio_config", "radius_config", "remote_syslog", "remove_existing_configs", "report_gatt", "rogue", "routing_policies", "rtsa", "simple_alert", "site_id", "skyatp", "sle_thresholds", "snmp_config", "srx_app", "ssh_keys", "ssr", "status_portal", "switch", "switch_matching", "switch_mgmt", "switch_updown_threshold", "synthetic_test", "track_anonymous_devices", "tunterm_monitoring", "tunterm_monitoring_disabled", "tunterm_multicast_config", "uplink_port_config", "uses_description_from_port_usage", "vars", "vars_annotations", "vna", "vpn_path_updown_threshold", "vpn_peer_updown_threshold", "vrf_config", "vrf_instances", "vrrp_groups", "vs_instance", "wan_vna", "watched_station_url", "whitelist_url", "wids", "wifi", "wired_vna", "zone_occupancy_alert")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AclPolicies = temp.AclPolicies
	s.AclTags = temp.AclTags
	s.AdditionalConfigCmds = temp.AdditionalConfigCmds
	s.AllowMist = temp.AllowMist
	s.Analytic = temp.Analytic
	s.ApMatching = temp.ApMatching
	s.ApPortConfig = temp.ApPortConfig
	s.ApSyntheticTest = temp.ApSyntheticTest
	s.ApUpdownThreshold = temp.ApUpdownThreshold
	s.AutoPlacement = temp.AutoPlacement
	s.AutoUpgrade = temp.AutoUpgrade
	s.AutoUpgradeEsl = temp.AutoUpgradeEsl
	s.AutoUpgradeLinecard = temp.AutoUpgradeLinecard
	s.BgpNeighborUpdownThreshold = temp.BgpNeighborUpdownThreshold
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
	s.EnableUnii4 = temp.EnableUnii4
	s.Engagement = temp.Engagement
	s.EvpnOptions = temp.EvpnOptions
	s.ExtraRoutes = temp.ExtraRoutes
	s.ExtraRoutes6 = temp.ExtraRoutes6
	s.Flags = temp.Flags
	s.ForSite = temp.ForSite
	s.Gateway = temp.Gateway
	s.GatewayAdditionalConfigCmds = temp.GatewayAdditionalConfigCmds
	s.GatewayMgmt = temp.GatewayMgmt
	s.GatewayTunnelUpdownThreshold = temp.GatewayTunnelUpdownThreshold
	s.GatewayUpdownThreshold = temp.GatewayUpdownThreshold
	s.Id = temp.Id
	s.Iotproxy = temp.Iotproxy
	s.JuniperSrx = temp.JuniperSrx
	s.Led = temp.Led
	s.Marvis = temp.Marvis
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
	s.RemoveExistingConfigs = temp.RemoveExistingConfigs
	s.ReportGatt = temp.ReportGatt
	s.Rogue = temp.Rogue
	s.RoutingPolicies = temp.RoutingPolicies
	s.Rtsa = temp.Rtsa
	s.SimpleAlert = temp.SimpleAlert
	s.SiteId = temp.SiteId
	s.Skyatp = temp.Skyatp
	s.SleThresholds = temp.SleThresholds
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
	s.UplinkPortConfig = temp.UplinkPortConfig
	s.UsesDescriptionFromPortUsage = temp.UsesDescriptionFromPortUsage
	s.Vars = temp.Vars
	s.VarsAnnotations = temp.VarsAnnotations
	s.Vna = temp.Vna
	s.VpnPathUpdownThreshold = temp.VpnPathUpdownThreshold
	s.VpnPeerUpdownThreshold = temp.VpnPeerUpdownThreshold
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

// tempSiteSetting is a temporary struct used for validating the fields of SiteSetting.
type tempSiteSetting struct {
	AclPolicies                     []AclPolicy                            `json:"acl_policies,omitempty"`
	AclTags                         map[string]AclTag                      `json:"acl_tags,omitempty"`
	AdditionalConfigCmds            []string                               `json:"additional_config_cmds,omitempty"`
	AllowMist                       *bool                                  `json:"allow_mist,omitempty"`
	Analytic                        *SiteSettingAnalytic                   `json:"analytic,omitempty"`
	ApMatching                      *SiteSettingApMatching                 `json:"ap_matching,omitempty"`
	ApPortConfig                    *SiteSettingApPortConfig               `json:"ap_port_config,omitempty"`
	ApSyntheticTest                 *SiteSettingApSyntheticTest            `json:"ap_synthetic_test,omitempty"`
	ApUpdownThreshold               Optional[int]                          `json:"ap_updown_threshold"`
	AutoPlacement                   *SiteSettingAutoPlacement              `json:"auto_placement,omitempty"`
	AutoUpgrade                     *SiteSettingAutoUpgrade                `json:"auto_upgrade,omitempty"`
	AutoUpgradeEsl                  *SiteSettingAutoUpgradeEsl             `json:"auto_upgrade_esl,omitempty"`
	AutoUpgradeLinecard             *bool                                  `json:"auto_upgrade_linecard,omitempty"`
	BgpNeighborUpdownThreshold      Optional[int]                          `json:"bgp_neighbor_updown_threshold"`
	BlacklistUrl                    *string                                `json:"blacklist_url,omitempty"`
	BleConfig                       *BleConfig                             `json:"ble_config,omitempty"`
	ConfigAutoRevert                *bool                                  `json:"config_auto_revert,omitempty"`
	ConfigPushPolicy                *SiteSettingConfigPushPolicy           `json:"config_push_policy,omitempty"`
	CreatedTime                     *float64                               `json:"created_time,omitempty"`
	CriticalUrlMonitoring           *SiteSettingCriticalUrlMonitoring      `json:"critical_url_monitoring,omitempty"`
	DeviceUpdownThreshold           Optional[int]                          `json:"device_updown_threshold"`
	DhcpSnooping                    *DhcpSnooping                          `json:"dhcp_snooping,omitempty"`
	DisabledSystemDefinedPortUsages []SystemDefinedPortUsagesEnum          `json:"disabled_system_defined_port_usages,omitempty"`
	DnsServers                      []string                               `json:"dns_servers,omitempty"`
	DnsSuffix                       []string                               `json:"dns_suffix,omitempty"`
	EnableUnii4                     *bool                                  `json:"enable_unii_4,omitempty"`
	Engagement                      *SiteEngagement                        `json:"engagement,omitempty"`
	EvpnOptions                     *EvpnOptions                           `json:"evpn_options,omitempty"`
	ExtraRoutes                     map[string]ExtraRoute                  `json:"extra_routes,omitempty"`
	ExtraRoutes6                    map[string]ExtraRoute6                 `json:"extra_routes6,omitempty"`
	Flags                           map[string]string                      `json:"flags,omitempty"`
	ForSite                         *bool                                  `json:"for_site,omitempty"`
	Gateway                         *GatewayTemplate                       `json:"gateway,omitempty"`
	GatewayAdditionalConfigCmds     []string                               `json:"gateway_additional_config_cmds,omitempty"`
	GatewayMgmt                     *GatewayMgmt                           `json:"gateway_mgmt,omitempty"`
	GatewayTunnelUpdownThreshold    Optional[int]                          `json:"gateway_tunnel_updown_threshold"`
	GatewayUpdownThreshold          Optional[int]                          `json:"gateway_updown_threshold"`
	Id                              *uuid.UUID                             `json:"id,omitempty"`
	Iotproxy                        *Iotproxy                              `json:"iotproxy,omitempty"`
	JuniperSrx                      *SiteSettingJuniperSrx                 `json:"juniper_srx,omitempty"`
	Led                             *ApLed                                 `json:"led,omitempty"`
	Marvis                          *Marvis                                `json:"marvis,omitempty"`
	MistNac                         *SwitchMistNac                         `json:"mist_nac,omitempty"`
	ModifiedTime                    *float64                               `json:"modified_time,omitempty"`
	Mxedge                          *SiteSettingMxedge                     `json:"mxedge,omitempty"`
	MxedgeMgmt                      *MxedgeMgmt                            `json:"mxedge_mgmt,omitempty"`
	Mxtunnels                       *SiteMxtunnel                          `json:"mxtunnels,omitempty"`
	Networks                        map[string]SwitchNetwork               `json:"networks,omitempty"`
	NtpServers                      []string                               `json:"ntp_servers,omitempty"`
	Occupancy                       *SiteOccupancyAnalytics                `json:"occupancy,omitempty"`
	OrgId                           *uuid.UUID                             `json:"org_id,omitempty"`
	OspfAreas                       map[string]OspfArea                    `json:"ospf_areas,omitempty"`
	PaloaltoNetworks                *SiteSettingPaloaltoNetworks           `json:"paloalto_networks,omitempty"`
	PersistConfigOnDevice           *bool                                  `json:"persist_config_on_device,omitempty"`
	PortMirroring                   map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
	PortUsages                      map[string]SwitchPortUsage             `json:"port_usages,omitempty"`
	Proxy                           *Proxy                                 `json:"proxy,omitempty"`
	RadioConfig                     *ApRadio                               `json:"radio_config,omitempty"`
	RadiusConfig                    *SwitchRadiusConfig                    `json:"radius_config,omitempty"`
	RemoteSyslog                    *RemoteSyslog                          `json:"remote_syslog,omitempty"`
	RemoveExistingConfigs           *bool                                  `json:"remove_existing_configs,omitempty"`
	ReportGatt                      *bool                                  `json:"report_gatt,omitempty"`
	Rogue                           *SiteRogue                             `json:"rogue,omitempty"`
	RoutingPolicies                 map[string]SwRoutingPolicy             `json:"routing_policies,omitempty"`
	Rtsa                            *SiteSettingRtsa                       `json:"rtsa,omitempty"`
	SimpleAlert                     *SimpleAlert                           `json:"simple_alert,omitempty"`
	SiteId                          *uuid.UUID                             `json:"site_id,omitempty"`
	Skyatp                          *SiteSettingSkyatp                     `json:"skyatp,omitempty"`
	SleThresholds                   *SleThresholds                         `json:"sle_thresholds,omitempty"`
	SnmpConfig                      *SnmpConfig                            `json:"snmp_config,omitempty"`
	SrxApp                          *SiteSettingSrxApp                     `json:"srx_app,omitempty"`
	SshKeys                         []string                               `json:"ssh_keys,omitempty"`
	Ssr                             *SettingSsr                            `json:"ssr,omitempty"`
	StatusPortal                    *SiteSettingStatusPortal               `json:"status_portal,omitempty"`
	Switch                          *SiteSettingSwitch                     `json:"switch,omitempty"`
	SwitchMatching                  *SwitchMatching                        `json:"switch_matching,omitempty"`
	SwitchMgmt                      *SwitchMgmt                            `json:"switch_mgmt,omitempty"`
	SwitchUpdownThreshold           Optional[int]                          `json:"switch_updown_threshold"`
	SyntheticTest                   *SynthetictestConfig                   `json:"synthetic_test,omitempty"`
	TrackAnonymousDevices           *bool                                  `json:"track_anonymous_devices,omitempty"`
	TuntermMonitoring               []TuntermMonitoringItem                `json:"tunterm_monitoring,omitempty"`
	TuntermMonitoringDisabled       *bool                                  `json:"tunterm_monitoring_disabled,omitempty"`
	TuntermMulticastConfig          *SiteSettingTuntermMulticastConfig     `json:"tunterm_multicast_config,omitempty"`
	UplinkPortConfig                *ApUplinkPortConfig                    `json:"uplink_port_config,omitempty"`
	UsesDescriptionFromPortUsage    *bool                                  `json:"uses_description_from_port_usage,omitempty"`
	Vars                            map[string]string                      `json:"vars,omitempty"`
	VarsAnnotations                 map[string]VarsAnnotation              `json:"vars_annotations,omitempty"`
	Vna                             *SiteSettingVna                        `json:"vna,omitempty"`
	VpnPathUpdownThreshold          Optional[int]                          `json:"vpn_path_updown_threshold"`
	VpnPeerUpdownThreshold          Optional[int]                          `json:"vpn_peer_updown_threshold"`
	VrfConfig                       *VrfConfig                             `json:"vrf_config,omitempty"`
	VrfInstances                    map[string]SwitchVrfInstance           `json:"vrf_instances,omitempty"`
	VrrpGroups                      map[string]VrrpGroup                   `json:"vrrp_groups,omitempty"`
	VsInstance                      map[string]VsInstanceProperty          `json:"vs_instance,omitempty"`
	WanVna                          *SiteSettingWanVna                     `json:"wan_vna,omitempty"`
	WatchedStationUrl               *string                                `json:"watched_station_url,omitempty"`
	WhitelistUrl                    *string                                `json:"whitelist_url,omitempty"`
	Wids                            *SiteWids                              `json:"wids,omitempty"`
	Wifi                            *SiteWifi                              `json:"wifi,omitempty"`
	WiredVna                        *SiteSettingWiredVna                   `json:"wired_vna,omitempty"`
	ZoneOccupancyAlert              *SiteZoneOccupancyAlert                `json:"zone_occupancy_alert,omitempty"`
}
