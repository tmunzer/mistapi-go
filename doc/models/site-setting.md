
# Site Setting

Site Settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | - |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `Analytic` | [`*models.SiteSettingAnalytic`](../../doc/models/site-setting-analytic.md) | Optional | - |
| `ApMatching` | [`*models.SiteSettingApMatching`](../../doc/models/site-setting-ap-matching.md) | Optional | - |
| `ApPortConfig` | [`*models.SiteSettingApPortConfig`](../../doc/models/site-setting-ap-port-config.md) | Optional | - |
| `ApUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `AutoPlacement` | [`*models.SiteSettingAutoPlacement`](../../doc/models/site-setting-auto-placement.md) | Optional | If we're able to determine its x/y/orientation, this will be populated |
| `AutoUpgrade` | [`*models.SiteSettingAutoUpgrade`](../../doc/models/site-setting-auto-upgrade.md) | Optional | Auto Upgrade Settings |
| `AutoUpgradeLinecard` | `*bool` | Optional | **Default**: `false` |
| `BlacklistUrl` | `*string` | Optional | - |
| `BleConfig` | [`*models.BleConfig`](../../doc/models/ble-config.md) | Optional | BLE AP settings |
| `ConfigAutoRevert` | `*bool` | Optional | Whether to enable ap auto config revert<br><br>**Default**: `false` |
| `ConfigPushPolicy` | [`*models.SiteSettingConfigPushPolicy`](../../doc/models/site-setting-config-push-policy.md) | Optional | Mist also uses some heuristic rules to prevent destructive configs from being pushed |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `CriticalUrlMonitoring` | [`*models.SiteSettingCriticalUrlMonitoring`](../../doc/models/site-setting-critical-url-monitoring.md) | Optional | You can define some URLs that's critical to site operations the latency will be captured and considered for site health |
| `DefaultPortUsage` | `*string` | Optional | Port usage to assign to switch ports without any port usage assngied. Default: `default` to preserve default behavior<br><br>**Default**: `"default"` |
| `DeviceUpdownThreshold` | `models.Optional[int]` | Optional | By default, device_updown_threshold, if set, will apply to all devices types if different values for specific device type is desired, use the following<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | - |
| `DisabledSystemDefinedPortUsages` | [`[]models.SystemDefinedPortUsagesEnum`](../../doc/models/system-defined-port-usages-enum.md) | Optional | If some system-default port usages are not desired - namely, ap / iot / uplink |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `EnableUnii4` | `*bool` | Optional | **Default**: `false` |
| `Engagement` | [`*models.SiteEngagement`](../../doc/models/site-engagement.md) | Optional | **Note**: if hours does not exist, it's treated as everyday of the week, 00:00-23:59. Currently, we don't allow multiple ranges for the same day |
| `EvpnOptions` | [`*models.EvpnOptions`](../../doc/models/evpn-options.md) | Optional | EVPN Options |
| `ExtraRoutes` | [`map[string]models.ExtraRoute`](../../doc/models/extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.ExtraRoute6`](../../doc/models/extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Flags` | `map[string]string` | Optional | Name/val pair objects for location engine to use |
| `ForSite` | `*bool` | Optional | - |
| `Gateway` | [`*models.GatewayTemplate`](../../doc/models/gateway-template.md) | Optional | Gateway Template is applied to a site for gateway(s) in a site. |
| `GatewayAdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `GatewayMgmt` | [`*models.SiteSettingGatewayMgmt`](../../doc/models/site-setting-gateway-mgmt.md) | Optional | Gateway Site settings |
| `GatewayUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `JuniperSrx` | [`*models.SiteSettingJuniperSrx`](../../doc/models/site-setting-juniper-srx.md) | Optional | - |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `Marvis` | [`*models.Marvis`](../../doc/models/marvis.md) | Optional | - |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | Enable mist_nac to use RadSec |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Mxedge` | [`*models.SiteSettingMxedge`](../../doc/models/site-setting-mxedge.md) | Optional | Site Mist Edges form a cluster of RadSec Proxy servers |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | - |
| `Mxtunnels` | [`*models.SiteMxtunnel`](../../doc/models/site-mxtunnel.md) | Optional | Site MxTunnel |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `NtpServers` | `[]string` | Optional | List of NTP servers |
| `Occupancy` | [`*models.SiteOccupancyAnalytics`](../../doc/models/site-occupancy-analytics.md) | Optional | Occupancy Analytics settings |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `OspfAreas` | [`map[string]models.OspfArea`](../../doc/models/ospf-area.md) | Optional | Junos OSPF areas |
| `PaloaltoNetworks` | [`*models.SiteSettingPaloaltoNetworks`](../../doc/models/site-setting-paloalto-networks.md) | Optional | - |
| `PersistConfigOnDevice` | `*bool` | Optional | Whether to store the config on AP<br><br>**Default**: `false` |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `Proxy` | [`*models.Proxy`](../../doc/models/proxy.md) | Optional | Proxy Configuration to talk to Mist |
| `RadioConfig` | [`*models.ApRadio`](../../doc/models/ap-radio.md) | Optional | Radio AP settings |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Junos Radius config |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | - |
| `RemoveExistingConfigs` | `*bool` | Optional | By default, only the configuration generated by Mist is cleaned up during the configuration process. If `true`, all the existing configuration will be removed.<br><br>**Default**: `false` |
| `ReportGatt` | `*bool` | Optional | Whether AP should periodically connect to BLE devices and report GATT device info (device name, manufacturer name, serial number, battery %, temperature, humidity)<br><br>**Default**: `false` |
| `Rogue` | [`*models.SiteRogue`](../../doc/models/site-rogue.md) | Optional | Rogue site settings |
| `Rtsa` | [`*models.SiteSettingRtsa`](../../doc/models/site-setting-rtsa.md) | Optional | Managed mobility |
| `SimpleAlert` | [`*models.SimpleAlert`](../../doc/models/simple-alert.md) | Optional | Set of heuristic rules will be enabled when marvis subscription is not available. It triggers when, in a Z minute window, there are more than Y distinct client encountering over X failures |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Skyatp` | [`*models.SiteSettingSkyatp`](../../doc/models/site-setting-skyatp.md) | Optional | - |
| `SleThresholds` | [`*models.SleThresholds`](../../doc/models/sle-thresholds.md) | Optional | - |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | - |
| `SrxApp` | [`*models.SiteSettingSrxApp`](../../doc/models/site-setting-srx-app.md) | Optional | - |
| `SshKeys` | `[]string` | Optional | When limit_ssh_access = true in Org Setting, list of SSH public keys provided by Mist Support to install onto APs (see Org:Setting) |
| `Ssr` | [`*models.SettingSsr`](../../doc/models/setting-ssr.md) | Optional | - |
| `StatusPortal` | [`*models.SiteSettingStatusPortal`](../../doc/models/site-setting-status-portal.md) | Optional | - |
| `Switch` | [`*models.SiteSettingSwitch`](../../doc/models/site-setting-switch.md) | Optional | - |
| `SwitchMatching` | [`*models.SwitchMatching`](../../doc/models/switch-matching.md) | Optional | Defines custom switch configuration based on different criteria |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch settings |
| `SwitchUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `SyntheticTest` | [`*models.SynthetictestConfig`](../../doc/models/synthetictest-config.md) | Optional | - |
| `TrackAnonymousDevices` | `*bool` | Optional | Whether to track anonymous BLE assets (requires ‘track_asset’  enabled)<br><br>**Default**: `false` |
| `TuntermMonitoring` | [`[]models.TuntermMonitoringItem`](../../doc/models/tunterm-monitoring-item.md) | Optional | - |
| `TuntermMonitoringDisabled` | `*bool` | Optional | **Default**: `false` |
| `TuntermMulticastConfig` | [`*models.SiteSettingTuntermMulticastConfig`](../../doc/models/site-setting-tunterm-multicast-config.md) | Optional | - |
| `UplinkPortConfig` | [`*models.ApUplinkPortConfig`](../../doc/models/ap-uplink-port-config.md) | Optional | AP Uplink port configuration |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `Vna` | [`*models.SiteSettingVna`](../../doc/models/site-setting-vna.md) | Optional | - |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | - |
| `VrfInstances` | [`map[string]models.SwitchVrfInstance`](../../doc/models/switch-vrf-instance.md) | Optional | Property key is the network name |
| `VrrpGroups` | [`map[string]models.VrrpGroup`](../../doc/models/vrrp-group.md) | Optional | Property key is the vrrp group |
| `VsInstance` | [`map[string]models.VsInstanceProperty`](../../doc/models/vs-instance-property.md) | Optional | Optional, for EX9200 only to segregate virtual-switches. Property key is the instance name |
| `WanVna` | [`*models.SiteSettingWanVna`](../../doc/models/site-setting-wan-vna.md) | Optional | - |
| `WatchedStationUrl` | `*string` | Optional | - |
| `WhitelistUrl` | `*string` | Optional | - |
| `Wids` | [`*models.SiteWids`](../../doc/models/site-wids.md) | Optional | WIDS site settings |
| `Wifi` | [`*models.SiteWifi`](../../doc/models/site-wifi.md) | Optional | Wi-Fi site settings |
| `WiredVna` | [`*models.SiteSettingWiredVna`](../../doc/models/site-setting-wired-vna.md) | Optional | - |
| `ZoneOccupancyAlert` | [`*models.SiteZoneOccupancyAlert`](../../doc/models/site-zone-occupancy-alert.md) | Optional | Zone Occupancy alert site settings |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_updown_threshold": 0,
  "auto_upgrade_linecard": false,
  "blacklist_url": "https://papi.s3.amazonaws.com/blacklist/xxx...",
  "config_auto_revert": false,
  "default_port_usage": "default",
  "device_updown_threshold": 0,
  "enable_unii_4": false,
  "extra_routes": {
    "0.0.0.0/0": {
      "via": "192.168.1.10"
    }
  },
  "extra_routes6": {
    "2a02:1234:420a:10c9::/64": {
      "via": "2a02:1234:200a::100"
    }
  },
  "gateway_updown_threshold": 0,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "persist_config_on_device": false,
  "remove_existing_configs": false,
  "report_gatt": false,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "switch_updown_threshold": 0,
  "track_anonymous_devices": false,
  "tunterm_monitoring_disabled": false,
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "vrf_instances": {
    "guest": {
      "extra_routes": {
        "0.0.0.0/0": {
          "via": "192.168.31.1"
        }
      },
      "networks": [
        "guest"
      ]
    }
  },
  "watched_station_url": "https://papi.s3.amazonaws.com/watched_station/xxx...",
  "whitelist_url": "https://papi.s3.amazonaws.com/whitelist/xxx...",
  "acl_policies": [
    {
      "actions": [
        {
          "action": "allow",
          "dst_tag": "dst_tag0",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "name": "name2",
      "src_tags": [
        "src_tags1",
        "src_tags0"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "acl_tags": {
    "key0": {
      "gbp_tag": 14,
      "macs": [
        "macs1"
      ],
      "network": "network2",
      "port_usage": "port_usage0",
      "radius_group": "radius_group8",
      "type": "network",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "gbp_tag": 14,
      "macs": [
        "macs1"
      ],
      "network": "network2",
      "port_usage": "port_usage0",
      "radius_group": "radius_group8",
      "type": "network",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "gbp_tag": 14,
      "macs": [
        "macs1"
      ],
      "network": "network2",
      "port_usage": "port_usage0",
      "radius_group": "radius_group8",
      "type": "network",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "additional_config_cmds": [
    "additional_config_cmds0",
    "additional_config_cmds1",
    "additional_config_cmds2"
  ],
  "analytic": {
    "enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "ap_matching": {
    "enabled": false,
    "rules": [
      {
        "match_model": "match_model0",
        "name": "name8",
        "port_config": {
          "key0": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        },
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "match_model": "match_model0",
        "name": "name8",
        "port_config": {
          "key0": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        },
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "match_model": "match_model0",
        "name": "name8",
        "port_config": {
          "key0": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          },
          "key1": {
            "disabled": false,
            "dynamic_vlan": {
              "default_vlan_id": 34,
              "enabled": false,
              "type": "type6",
              "vlans": {
                "key0": "vlans1"
              },
              "exampleAdditionalProperty": {
                "key1": "val1",
                "key2": "val2"
              }
            },
            "enable_mac_auth": false,
            "forwarding": "site_mxedge",
            "mac_auth_preferred": false,
            "exampleAdditionalProperty": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        },
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

