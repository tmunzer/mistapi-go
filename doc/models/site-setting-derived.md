
# Site Setting Derived

Site settings with derived OAuth account integration data

*This model accepts additional fields of type [models.AccountOauthInfoAccount](../../doc/models/account-oauth-info-account.md).*

## Structure

`SiteSettingDerived`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | List of ACL policy definitions |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `AllowMist` | `*bool` | Optional | whether to allow Mist to look at this org<br><br>**Default**: `false` |
| `Analytic` | [`*models.SiteSettingAnalytic`](../../doc/models/site-setting-analytic.md) | Optional | Advanced analytics feature settings for a site |
| `ApMatching` | [`*models.SiteSettingApMatching`](../../doc/models/site-setting-ap-matching.md) | Optional | Rules for applying AP port configuration by AP model or name |
| `ApPortConfig` | [`*models.SiteSettingApPortConfig`](../../doc/models/site-setting-ap-port-config.md) | Optional | AP Ethernet port configuration overrides by model |
| `ApSyntheticTest` | [`*models.SiteSettingApSyntheticTest`](../../doc/models/site-setting-ap-synthetic-test.md) | Optional | AP Synthetic Test configuration |
| `ApUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `AutoPlacement` | [`*models.SiteSettingAutoPlacement`](../../doc/models/site-setting-auto-placement.md) | Optional | Automatically determined AP placement coordinates and orientation |
| `AutoUpgrade` | [`*models.SiteSettingAutoUpgrade`](../../doc/models/site-setting-auto-upgrade.md) | Optional | Automatic AP firmware upgrade policy |
| `AutoUpgradeEsl` | [`*models.SiteSettingAutoUpgradeEsl`](../../doc/models/site-setting-auto-upgrade-esl.md) | Optional | Automatic AP ESL firmware upgrade policy. When both firmware and ESL auto-upgrade are enabled, ESL upgrade will be done only after firmware upgrade |
| `AutoUpgradeLinecard` | `*bool` | Optional | Whether line cards are included in automatic switch upgrades<br><br>**Default**: `true` |
| `BgpNeighborUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based bgp neighbor down delivery.<br><br>**Constraints**: `>= 0` |
| `BlacklistUrl` | `*string` | Optional, Read-only | Read-only URL for the site blacklist file |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | Bluetooth Low Energy beacon and asset advertising settings for an AP |
| `ConfigAutoRevert` | `*bool` | Optional | Whether to enable ap auto config revert<br><br>**Default**: `false` |
| `ConfigPushPolicy` | [`*models.SiteSettingConfigPushPolicy`](../../doc/models/site-setting-config-push-policy.md) | Optional | Mist also uses some heuristic rules to prevent destructive configs from being pushed |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `CriticalUrlMonitoring` | [`*models.SiteSettingCriticalUrlMonitoring`](../../doc/models/site-setting-critical-url-monitoring.md) | Optional | Critical URLs whose latency is measured and included in site health |
| `DeviceUpdownThreshold` | `models.Optional[int]` | Optional | By default, device_updown_threshold, if set, will apply to all devices types if different values for specific device type is desired, use the following<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | DHCP snooping security settings |
| `DisabledSystemDefinedPortUsages` | [`[]models.SystemDefinedPortUsagesEnum`](../../doc/models/system-defined-port-usages-enum.md) | Optional | If some system-default port usages are not desired - namely, ap / iot / uplink |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `EnableUnii4` | `*bool` | Optional | Whether UNII-4 channels are enabled for the site<br><br>**Default**: `false` |
| `Engagement` | [`*models.SiteEngagement`](../../doc/models/site-engagement.md) | Optional | Engagement analytics dwell-time rules for classifying site visits. If hours is omitted, rules apply every day from 00:00 to 23:59. Multiple ranges for the same day are not supported. |
| `EvpnOptions` | [`*models.EvpnOptions`](../../doc/models/evpn-options.md) | Optional | EVPN topology generation options for campus fabric configuration |
| `ExtraRoutes` | [`map[string]models.ExtraRoute`](../../doc/models/extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.ExtraRoute6`](../../doc/models/extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Flags` | `map[string]string` | Optional | Name/val pair objects for location engine to use |
| `ForSite` | `*bool` | Optional, Read-only | Whether this settings object is scoped to a site |
| `Gateway` | [`*models.GatewayTemplate`](../../doc/models/gateway-template.md) | Optional | Gateway Template is applied to a site for gateway(s) in a site. |
| `GatewayAdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `GatewayMgmt` | [`*models.GatewayMgmt`](../../doc/models/gateway-mgmt.md) | Optional | Gateway management-plane and access settings |
| `GatewayTunnelUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based gateway tunnel (secure edge tunnels) up-down delivery.<br><br>**Constraints**: `>= 0` |
| `GatewayUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Iotproxy` | [`*models.Iotproxy`](../../doc/models/iotproxy.md) | Optional | IoT proxy configuration for the site |
| `JuniperSrx` | [`*models.SiteSettingJuniperSrx`](../../doc/models/site-setting-juniper-srx.md) | Optional | Site-level Juniper SRX integration settings |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | Indicator light settings for an access point |
| `Marvis` | [`*models.Marvis`](../../doc/models/marvis.md) | Optional | Marvis automation and client settings |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | Mist NAC RadSec settings for a switch |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Mxedge` | [`*models.SiteSettingMxedge`](../../doc/models/site-setting-mxedge.md) | Optional | Service settings for the site Mist Edge cluster |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | Management settings for a Mist Edge appliance |
| `Mxtunnels` | [`*models.SiteMxtunnel`](../../doc/models/site-mxtunnel.md) | Optional | Site Mist Tunnel configuration for tunneling AP user VLANs to Mist Edge tunnel peers |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `NtpServers` | `[]string` | Optional | List of NTP servers |
| `Occupancy` | [`*models.SiteOccupancyAnalytics`](../../doc/models/site-occupancy-analytics.md) | Optional | Analytics settings for site occupancy |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `OspfAreas` | [`map[string]models.OspfArea`](../../doc/models/ospf-area.md) | Optional | Junos OSPF areas. Property key is the OSPF Area (Area should be a number (0-255) / IP address) |
| `PaloaltoNetworks` | [`*models.SiteSettingPaloaltoNetworks`](../../doc/models/site-setting-paloalto-networks.md) | Optional | Palo Alto Networks integration settings for the site |
| `PersistConfigOnDevice` | `*bool` | Optional | Whether to store the config on AP<br><br>**Default**: `false` |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `Proxy` | [`*models.Proxy`](../../doc/models/proxy.md) | Optional | Proxy Configuration to talk to Mist |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio configuration settings for an access point |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Switch RADIUS authentication and accounting configuration |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | Remote syslog forwarding settings |
| `RemoveExistingConfigs` | `*bool` | Optional | By default, only the configuration generated by Mist is cleaned up during the configuration process. If `true`, all the existing configuration will be removed.<br><br>**Default**: `false` |
| `ReportGatt` | `*bool` | Optional | Whether AP should periodically connect to BLE devices and report GATT device info (device name, manufacturer name, serial number, battery %, temperature, humidity)<br><br>**Default**: `false` |
| `Rogue` | [`*models.SiteRogue`](../../doc/models/site-rogue.md) | Optional | Rogue AP detection settings for a site |
| `RoutingPolicies` | [`map[string]models.SwRoutingPolicy`](../../doc/models/sw-routing-policy.md) | Optional | Switch routing policies keyed by routing policy name |
| `Rtsa` | [`*models.SiteSettingRtsa`](../../doc/models/site-setting-rtsa.md) | Optional | Managed mobility and asset tracking settings |
| `SimpleAlert` | [`*models.SimpleAlert`](../../doc/models/simple-alert.md) | Optional | Heuristic alert thresholds used when a Marvis subscription is unavailable |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Skyatp` | [`*models.SiteSettingSkyatp`](../../doc/models/site-setting-skyatp.md) | Optional | Sky ATP threat intelligence settings for the site |
| `SleThresholds` | [`*models.SleThresholds`](../../doc/models/sle-thresholds.md) | Optional | Site SLE threshold overrides for capacity, coverage, throughput, and time to connect |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | SNMP configuration for managed network devices |
| `SrxApp` | [`*models.SiteSettingSrxApp`](../../doc/models/site-setting-srx-app.md) | Optional | Juniper SRX application visibility settings for the site |
| `SshKeys` | `[]string` | Optional | When limit_ssh_access = true in Org Setting, list of SSH public keys provided by Mist Support to install onto APs (see Org:Setting) |
| `Ssr` | [`*models.SettingSsr`](../../doc/models/setting-ssr.md) | Optional | SSR management settings for device onboarding and connectivity |
| `StatusPortal` | [`*models.SiteSettingStatusPortal`](../../doc/models/site-setting-status-portal.md) | Optional | End-user status portal settings for the site |
| `Switch` | [`*models.SiteSettingSwitch`](../../doc/models/site-setting-switch.md) | Optional | Site switch settings combining a network template and auto-upgrade controls |
| `SwitchMatching` | [`*models.SwitchMatching`](../../doc/models/switch-matching.md) | Optional | Defines custom switch configuration based on different criteria |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch management-plane access and proxy settings |
| `SwitchUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `SyntheticTest` | [`*models.SynthetictestConfig`](../../doc/models/synthetictest-config.md) | Optional | Synthetic test configuration for Marvis Minis |
| `TrackAnonymousDevices` | `*bool` | Optional | Whether to track anonymous BLE assets (requires ‘track_asset’  enabled)<br><br>**Default**: `false` |
| `TuntermMonitoring` | [`[]models.TuntermMonitoringItem`](../../doc/models/tunterm-monitoring-item.md) | Optional | Tunnel termination monitoring checks |
| `TuntermMonitoringDisabled` | `*bool` | Optional | Whether tunnel termination monitoring is disabled for the site<br><br>**Default**: `false` |
| `TuntermMulticastConfig` | [`*models.SiteSettingTuntermMulticastConfig`](../../doc/models/site-setting-tunterm-multicast-config.md) | Optional | Multicast forwarding settings for tunnel termination at the site |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | AP Uplink port configuration |
| `UsesDescriptionFromPortUsage` | `*bool` | Optional | by default, we only honor description provided in port_config. This allows fallback to those defined in port_usages<br><br>**Default**: `false` |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `VarsAnnotations` | [`map[string]models.VarsAnnotation`](../../doc/models/vars-annotation.md) | Optional | Optional annotations for vars defined in this site. Keys match var names; values describe the var purpose and type for UI auto-complete. |
| `Vna` | [`*models.SiteSettingVna`](../../doc/models/site-setting-vna.md) | Optional | Virtual Network Assistant settings for AP, switch, and gateway experiences at a site |
| `VpnPathUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based vpn path down delivery.<br><br>**Constraints**: `>= 0` |
| `VpnPeerUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based vpn peer down delivery.<br><br>**Constraints**: `>= 0` |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | VRF enablement settings applied when supported on the device |
| `VrfInstances` | [`map[string]models.SwitchVrfInstance`](../../doc/models/switch-vrf-instance.md) | Optional | Property key is the network name |
| `VrrpGroups` | [`map[string]models.VrrpGroup`](../../doc/models/vrrp-group.md) | Optional | Property key is the vrrp group |
| `VsInstance` | [`map[string]models.VsInstanceProperty`](../../doc/models/vs-instance-property.md) | Optional | Optional, for EX9200 only to segregate virtual-switches. Property key is the instance name |
| `WanVna` | [`*models.SiteSettingWanVna`](../../doc/models/site-setting-wan-vna.md) | Optional | WAN Virtual Network Assistant settings for the site |
| `WatchedStationUrl` | `*string` | Optional, Read-only | Read-only URL for the watched station list file |
| `WhitelistUrl` | `*string` | Optional, Read-only | Read-only URL for the site whitelist file |
| `Wids` | [`*models.SiteWids`](../../doc/models/site-wids.md) | Optional | Wireless intrusion detection settings for a site |
| `Wifi` | [`*models.SiteWifi`](../../doc/models/site-wifi.md) | Optional | Wi-Fi configuration settings for a site |
| `WiredVna` | [`*models.SiteSettingWiredVna`](../../doc/models/site-setting-wired-vna.md) | Optional | Wired Virtual Network Assistant settings for the site |
| `ZoneOccupancyAlert` | [`*models.SiteZoneOccupancyAlert`](../../doc/models/site-zone-occupancy-alert.md) | Optional | Zone occupancy alert settings for a site |
| `AccountId` | `*string` | Optional, Read-only | Linked app account id |
| `AutoProbeSubnet` | `*string` | Optional, Read-only | For Prisma accounts only, tunnel auto probe subnet |
| `ClientId` | `*string` | Optional, Read-only | Customer account Client ID |
| `CloudName` | `*string` | Optional, Read-only | Name of the company whose account mist has subscribed to |
| `Company` | `*string` | Optional, Read-only | Name of the company whose account mist has subscribed to |
| `EnableProbe` | `*bool` | Optional, Read-only | For Prisma accounts only, tunnel probe enable/disable |
| `Error` | `*string` | Optional, Read-only | This error is provided when the account fails to fetch token/data |
| `Errors` | `[]string` | Optional, Read-only | Read-only OAuth account error messages |
| `InstanceUrl` | `*string` | Optional, Read-only | Customer account instance URL |
| `KeyId` | `*string` | Optional | For ZDX Account only, Customer account API key ID |
| `LastStatus` | `*string` | Optional, Read-only | Is the last data pull for account is successful or not |
| `LastSync` | `*int64` | Optional, Read-only | Last data pull timestamp, background jobs that pull account data |
| `LinkedBy` | `*string` | Optional, Read-only | First name of the user who linked the account |
| `LinkedTimestamp` | `*float64` | Optional, Read-only | Timestamp when this third-party account was linked |
| `MaxDailyApiRequests` | `*int` | Optional, Read-only | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |
| `Name` | `*string` | Optional, Read-only | Display name of the linked third-party account or company |
| `Password` | `*string` | Optional, Read-only | Customer account password instance URL |
| `Region` | `*string` | Optional, Read-only | For Prisma accounts only |
| `Regions` | [`map[string]models.AccountOauthInfoAccountRegion`](../../doc/models/account-oauth-info-account-region.md) | Optional | For Prisma accounts only, property key is the region name. Regions with allocated bandwidth |
| `ServiceAccountName` | `*string` | Optional, Read-only | For Prisma accounts only |
| `ServiceConnections` | [`map[string]models.AccountOauthInfoAccountServiceConnection`](../../doc/models/account-oauth-info-account-service-connection.md) | Optional | For Prisma accounts only, property key is the service connection name |
| `SmartgroupName` | `*string` | Optional, Read-only | Smart group membership for determining compliance status |
| `TsgId` | `*string` | Optional, Read-only | For Prisma accounts only, Prisma Tenant Service Group id |
| `Username` | `*string` | Optional, Read-only | Login name configured for the linked third-party account |
| `WebhookAuthType` | `*string` | Optional | For Crowdstrike, JAMF, SentinelOne and VMWare accounts only |
| `WebhookEnabled` | `*bool` | Optional | For Crowdstrike, JAMF, SentinelOne and VMWare accounts only |
| `WebhookPassword` | `*string` | Optional | For VMWare accounts only |
| `WebhookSecret` | `*string` | Optional | For Crowdstrike accounts only |
| `WebhookToken` | `*string` | Optional | For JAMF and SentinelOne accounts only |
| `WebhookUrl` | `*string` | Optional | For Crowdstrike, JAMF, SentinelOne and VMWare accounts only |
| `WebhookUsername` | `*string` | Optional | For VMWare accounts only |
| `ZdxOrgId` | `*string` | Optional | For ZDX Account only, ZDX organization id |
| `AdditionalProperties` | [`map[string]models.AccountOauthInfoAccount`](../../doc/models/account-oauth-info-account.md) | Optional | OAuth linked apps account info |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    siteSettingDerived := models.SiteSettingDerived{
        AclPolicies:                     []models.AclPolicy{
            models.AclPolicy{
                Actions:              []models.AclPolicyAction{
                    models.AclPolicyAction{
                        Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
                        DstTag:               "dst_tag0",
                    },
                },
                Name:                 models.ToPointer("name2"),
                SrcTags:              []string{
                    "src_tags1",
                    "src_tags0",
                },
            },
            models.AclPolicy{
                Actions:              []models.AclPolicyAction{
                    models.AclPolicyAction{
                        Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
                        DstTag:               "dst_tag0",
                    },
                },
                Name:                 models.ToPointer("name2"),
                SrcTags:              []string{
                    "src_tags1",
                    "src_tags0",
                },
            },
        },
        AclTags:                         map[string]models.AclTag{
            "key0": models.AclTag{
                EtherTypes:           []string{
                    "ether_types8",
                    "ether_types9",
                },
                GbpTag:               models.ToPointer(14),
                Macs:                 []string{
                    "macs1",
                },
                Network:              models.ToPointer("network2"),
                PortUsage:            models.ToPointer("port_usage0"),
                Type:                 models.AclTagTypeEnum_NETWORK,
            },
        },
        AdditionalConfigCmds:            []string{
            "additional_config_cmds6",
            "additional_config_cmds5",
        },
        AllowMist:                       models.ToPointer(false),
        Analytic:                        models.ToPointer(models.SiteSettingAnalytic{
            Enabled:              models.ToPointer(false),
        }),
        ApUpdownThreshold:               models.NewOptional(models.ToPointer(0)),
        AutoUpgradeLinecard:             models.ToPointer(true),
        BlacklistUrl:                    models.ToPointer("https://papi.s3.amazonaws.com/blacklist/xxx..."),
        ConfigAutoRevert:                models.ToPointer(false),
        DeviceUpdownThreshold:           models.NewOptional(models.ToPointer(0)),
        EnableUnii4:                     models.ToPointer(false),
        ExtraRoutes:                     map[string]models.ExtraRoute{
            "0.0.0.0/0": models.ExtraRoute{
                Via:                  models.ToPointer(),
            },
        },
        ExtraRoutes6:                    map[string]models.ExtraRoute6{
            "2a02:1234:420a:10c9::/64": models.ExtraRoute6{
                Via:                  models.ToPointer(),
            },
        },
        GatewayUpdownThreshold:          models.NewOptional(models.ToPointer(0)),
        Id:                              models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        OrgId:                           models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PersistConfigOnDevice:           models.ToPointer(false),
        RemoveExistingConfigs:           models.ToPointer(false),
        ReportGatt:                      models.ToPointer(false),
        SiteId:                          models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        SwitchUpdownThreshold:           models.NewOptional(models.ToPointer(0)),
        TrackAnonymousDevices:           models.ToPointer(false),
        TuntermMonitoringDisabled:       models.ToPointer(false),
        UsesDescriptionFromPortUsage:    models.ToPointer(false),
        Vars:                            map[string]string{
            "RADIUS_IP1": "172.31.2.5",
            "RADIUS_SECRET": "11s64632d",
        },
        VarsAnnotations:                 map[string]models.VarsAnnotation{
            "MXTUNNEL_GUEST": models.VarsAnnotation{
                Type:                 models.ToPointer("mxtunnel_id"),
            },
            "RADIUS_IP1": models.VarsAnnotation{
                Note:                 models.ToPointer("RADIUS server IP address for US East Campus"),
            },
        },
        VrfInstances:                    map[string]models.SwitchVrfInstance{
            "guest": models.SwitchVrfInstance{
                ExtraRoutes:             map[string]models.VrfExtraRoute{
                    "0.0.0.0/0": models.VrfExtraRoute{
                        Via:                  models.ToPointer("192.168.31.1"),
                    },
                },
                Networks:                []string{
                    "guest",
                },
            },
        },
        WatchedStationUrl:               models.ToPointer("https://papi.s3.amazonaws.com/watched_station/xxx..."),
        WhitelistUrl:                    models.ToPointer("https://papi.s3.amazonaws.com/whitelist/xxx..."),
        AccountId:                       models.ToPointer("iojzXIJWEuiD73ZvydOfg"),
        AutoProbeSubnet:                 models.ToPointer("11.0.0.0/8"),
        CloudName:                       models.ToPointer("Tapi.sase.paloaltonetworks.com"),
        Company:                         models.ToPointer("Test Company1 Ltd"),
        EnableProbe:                     models.ToPointer(false),
        Error:                           models.ToPointer("OAuth token refresh failed, please re-link your account"),
        Errors:                          []string{
            "OAuth token refresh failed, please re-link your account",
            "API daily rate limit reached for your account",
        },
        KeyId:                           models.ToPointer("L72frZcK3JvrZc"),
        LastStatus:                      models.ToPointer("failed"),
        LastSync:                        models.ToPointer(int64(1665465339000)),
        LinkedBy:                        models.ToPointer("Testname1"),
        LinkedTimestamp:                 models.ToPointer(float64(1665465339000)),
        MaxDailyApiRequests:             models.ToPointer(5000),
        Name:                            models.ToPointer("Test Compay1 Ltd"),
        Region:                          models.ToPointer("americas"),
        ServiceAccountName:              models.ToPointer("Corp SA"),
        SmartgroupName:                  models.ToPointer("CompliantGroup1"),
        TsgId:                           models.ToPointer("189953456"),
        WebhookAuthType:                 models.ToPointer("Basic"),
        WebhookPassword:                 models.ToPointer("password_1234"),
        WebhookSecret:                   models.ToPointer("secret-value"),
        WebhookToken:                    models.ToPointer("token-value"),
        WebhookUrl:                      models.ToPointer("https://websync.nac-staging.mistsys.com/v1/S_org-8dcbe9005/ae9dee49-69e7-4710-a114-5b827a777738/crowdstrike/edr"),
        WebhookUsername:                 models.ToPointer("username_1234"),
        ZdxOrgId:                        models.ToPointer("123456"),
        AdditionalProperties:            map[string]models.AccountOauthInfoAccount{
            "exampleAdditionalProperty": models.AccountOauthInfoAccount{
                AccountId:            models.ToPointer("account_id8"),
                AutoProbeSubnet:      models.ToPointer("auto_probe_subnet4"),
                ClientId:             models.ToPointer("client_id8"),
                CloudName:            models.ToPointer("cloud_name8"),
                Company:              models.ToPointer("company6"),
            },
        },
    }

}
```

