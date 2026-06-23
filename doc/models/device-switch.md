
# Device Switch

You can configure `port_usages` and `networks` settings at the device level, but most of the time it's better use the Site Setting to achieve better consistency and be able to re-use the same settings across switches entries defined here will "replace" those defined in Site Setting/Network Template
In addition it is possible to use the `port_config_overwrite` to overwrite some attributes of the port_usage without having to create a new port_usage.

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | List of ACL policy definitions |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `AggregateRoutes` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "172.16.3.0/24") |
| `AggregateRoutes6` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64") |
| `BgpConfig` | [`map[string]models.SwitchBgpConfig`](../../doc/models/switch-bgp-config.md) | Optional | BGP routing configuration for this switch. Property key is the BGP session name |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DefaultPortUsage` | `*string` | Optional | Port usage to assign to switch ports without any port usage assigned. Default: `default` to preserve default behavior<br><br>**Default**: `"default"` |
| `DeviceprofileId` | `*uuid.UUID` | Optional, Read-only | Device profile associated with this switch |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | DHCP snooping security settings |
| `DhcpdConfig` | [`*models.SwitchDhcpdConfig`](../../doc/models/switch-dhcpd-config.md) | Optional | Switch DHCP server or relay configuration keyed by network name |
| `DisableAutoConfig` | `*bool` | Optional | This disables the default behavior of a cloud-ready switch/gateway being managed/configured by Mist. Setting this to `true` means you want to disable the default behavior and do not want the device to be Mist-managed.<br><br>**Default**: `false` |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `EvpnConfig` | [`*models.EvpnConfig`](../../doc/models/evpn-config.md) | Optional, Read-only | EVPN configuration settings applied to a Junos switch |
| `ExtraRoutes` | [`map[string]models.ExtraRoute`](../../doc/models/extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.ExtraRoute6`](../../doc/models/extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Image1Url` | `models.Optional[string]` | Optional | First custom image URL associated with the switch |
| `Image2Url` | `models.Optional[string]` | Optional | Second custom image URL associated with the switch |
| `Image3Url` | `models.Optional[string]` | Optional | Third custom image URL associated with the switch |
| `IotConfig` | [`map[string]models.SwitchIotPort`](../../doc/models/switch-iot-port.md) | Optional | Property Key is the IOT port name, e.g.:<br><br>* `IN0` or `IN1` for the FPC0 input port with 5V triggered inputs<br>* `OUT1` for the FPC0 output port (can only be triggered by either IN0 or IN1)<br>* "X/IN0`, `X/IN1`and`X/OUT` are used to define IOT ports on VC members |
| `IpConfig` | [`*models.JunosIpConfig`](../../doc/models/junos-ip-config.md) | Optional | Junos management IP configuration |
| `LocalPortConfig` | [`map[string]models.JunosLocalPortConfig`](../../doc/models/junos-local-port-config.md) | Optional | Per-port Switch Port Operator (SPO) override, overriding the port configuration from `port_config`. Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `Mac` | `*string` | Optional, Read-only | Switch MAC address used to identify the device |
| `Managed` | `*bool` | Optional | An adopted switch/gateway will not be managed/configured by Mist by default. Setting this parameter to `true` enables the adopted switch/gateway to be managed/configured by Mist. Deprecated in favour of mist_configured, which is more intuitive and can be used for both adopted and claimed devices.<br><br>**Default**: `false` |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `MistConfigured` | `*bool` | Optional | whether the device can be configured by Mist or not. This deprecates `managed` (for adopted device) and `disable_auto_config` for claimed device) |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | Mist NAC RadSec settings for a switch |
| `Model` | `*string` | Optional, Read-only | Switch model reported for the device |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Friendly display name assigned to the switch |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `Notes` | `*string` | Optional | Free-form administrative notes for this switch |
| `NtpServers` | `[]string` | Optional | List of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OobIpConfig` | [`*models.SwitchOobIpConfig`](../../doc/models/switch-oob-ip-config.md) | Optional | Switch OOB IP Config:<br><br>- If HA configuration: key parameter will be nodeX (eg: node1)<br>- If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `OspfAreas` | [`map[string]models.OspfArea`](../../doc/models/ospf-area.md) | Optional | Junos OSPF areas. Property key is the OSPF Area (Area should be a number (0-255) / IP address) |
| `OspfConfig` | [`*models.SwitchOspfConfig`](../../doc/models/switch-ospf-config.md) | Optional | OSPF routing configuration for a Junos switch |
| `OtherIpConfigs` | [`map[string]models.JunosOtherIpConfig`](../../doc/models/junos-other-ip-config.md) | Optional | Property key is the network name. Defines the additional IP Addresses configured on the device. |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortConfigOverwrite` | [`map[string]models.SwitchPortConfigOverwrite`](../../doc/models/switch-port-config-overwrite.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10"). This can be used to override some attributes of the port_usage without having to create a new port_usage. |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Switch RADIUS authentication and accounting configuration |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | Remote syslog forwarding settings |
| `Role` | `*string` | Optional | Deployment role label for this switch |
| `RouterId` | `*string` | Optional | Used for OSPF / BGP / EVPN |
| `RoutingPolicies` | [`map[string]models.SwRoutingPolicy`](../../doc/models/sw-routing-policy.md) | Optional | Switch routing policies keyed by routing policy name |
| `Serial` | `*string` | Optional, Read-only | Manufacturer serial number for the switch |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | SNMP configuration for managed network devices |
| `StpConfig` | [`*models.SwitchStpConfig`](../../doc/models/switch-stp-config.md) | Optional | Switch spanning-tree protocol configuration |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch management-plane access and proxy settings |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `switch`<br><br>**Value**: `"switch"` |
| `UseRouterIdAsSourceIp` | `*bool` | Optional | Whether to use it for snmp / syslog / tacplus / radius<br><br>**Default**: `false` |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `VirtualChassis` | [`*models.SwitchVirtualChassis`](../../doc/models/switch-virtual-chassis.md) | Optional | Required for preprovisioned Virtual Chassis |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | VRF enablement settings applied when supported on the device |
| `VrfInstances` | [`map[string]models.SwitchVrfInstance`](../../doc/models/switch-vrf-instance.md) | Optional | Property key is the network name |
| `VrrpConfig` | [`*models.VrrpConfig`](../../doc/models/vrrp-config.md) | Optional | Junos VRRP configuration applied to a switch or switch profile |
| `X` | `*float64` | Optional | Horizontal map position of the switch, in pixels |
| `Y` | `*float64` | Optional | Vertical map position of the switch, in pixels |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    deviceSwitch := models.DeviceSwitch{
        AclPolicies:           []models.AclPolicy{
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
        AclTags:               map[string]models.AclTag{
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
        AdditionalConfigCmds:  []string{
            "additional_config_cmds2",
            "additional_config_cmds1",
        },
        AggregateRoutes:       map[string]models.AggregateRoute{
            "172.16.3.0/24": models.AggregateRoute{
                Discard:              models.ToPointer(false),
                Metric:               models.NewOptional(models.ToPointer(28)),
                Preference:           models.NewOptional(models.ToPointer(30)),
            },
        },
        AggregateRoutes6:      map[string]models.AggregateRoute{
            "2a02:1234:420a:10c9::/64": models.AggregateRoute{
                Discard:              models.ToPointer(false),
                Metric:               models.NewOptional(models.ToPointer(126)),
                Preference:           models.NewOptional(models.ToPointer(30)),
            },
        },
        DefaultPortUsage:      models.ToPointer("default"),
        DisableAutoConfig:     models.ToPointer(false),
        ExtraRoutes:           map[string]models.ExtraRoute{
            "0.0.0.0/0": models.ExtraRoute{
                Via:                  models.ToPointer(),
            },
        },
        ExtraRoutes6:          map[string]models.ExtraRoute6{
            "2a02:1234:420a:10c9::/64": models.ExtraRoute6{
                Via:                  models.ToPointer(),
            },
        },
        Id:                    models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Managed:               models.ToPointer(false),
        MapId:                 models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
        OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RouterId:              models.ToPointer("10.2.1.10"),
        SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                  "switch",
        UseRouterIdAsSourceIp: models.ToPointer(false),
        Vars:                  map[string]string{
            "RADIUS_IP1": "172.31.2.5",
            "RADIUS_SECRET": "11s64632d",
        },
        VrfInstances:          map[string]models.SwitchVrfInstance{
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
        X:                     models.ToPointer(float64(53.5)),
        Y:                     models.ToPointer(float64(173.1)),
        AdditionalProperties:  map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

