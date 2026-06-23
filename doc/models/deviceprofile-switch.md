
# Deviceprofile Switch

Switch Device Profiles can be applied to one or multiple switches. The settings from the Device Profile will override the settings from the Switch Template and the Site Settings.

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceprofileSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AclPolicies` | [`[]models.AclPolicy`](../../doc/models/acl-policy.md) | Optional | List of ACL policy definitions |
| `AclTags` | [`map[string]models.AclTag`](../../doc/models/acl-tag.md) | Optional | ACL Tags to identify traffic source or destination. Key name is the tag name |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `AggregateRoutes` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "172.16.3.0/24") |
| `AggregateRoutes6` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64") |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DhcpSnooping` | [`*models.DhcpSnooping`](../../doc/models/dhcp-snooping.md) | Optional | DHCP snooping security settings |
| `DhcpdConfig` | [`*models.SwitchDhcpdConfig`](../../doc/models/switch-dhcpd-config.md) | Optional | Switch DHCP server or relay configuration keyed by network name |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `EvpnConfig` | [`*models.EvpnConfig`](../../doc/models/evpn-config.md) | Optional, Read-only | EVPN configuration settings applied to a Junos switch |
| `ExtraRoutes` | [`map[string]models.ExtraRoute`](../../doc/models/extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.ExtraRoute6`](../../doc/models/extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IotConfig` | [`map[string]models.SwitchIotPort`](../../doc/models/switch-iot-port.md) | Optional | Property Key is the IOT port name, e.g.:<br><br>* `IN0` or `IN1` for the FPC0 input port with 5V triggered inputs<br>* `OUT1` for the FPC0 output port (can only be triggered by either IN0 or IN1)<br>* "X/IN0`, `X/IN1`and`X/OUT` are used to define IOT ports on VC members |
| `IpConfig` | [`*models.JunosIpConfig`](../../doc/models/junos-ip-config.md) | Optional | Junos management IP configuration |
| `MistNac` | [`*models.SwitchMistNac`](../../doc/models/switch-mist-nac.md) | Optional | Mist NAC RadSec settings for a switch |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the switch profile |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `NtpServers` | `[]string` | Optional | List of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OobIpConfig` | [`*models.SwitchOobIpConfig`](../../doc/models/switch-oob-ip-config.md) | Optional | Switch OOB IP Config:<br><br>- If HA configuration: key parameter will be nodeX (eg: node1)<br>- If there are 2 routing engines, re1 mgmt IP has to be set separately (if desired): key parameter = `re1` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `OspfAreas` | [`map[string]models.OspfArea`](../../doc/models/ospf-area.md) | Optional | Junos OSPF areas. Property key is the OSPF Area (Area should be a number (0-255) / IP address) |
| `OtherIpConfigs` | [`map[string]models.JunosOtherIpConfig`](../../doc/models/junos-other-ip-config.md) | Optional | Property key is the network name. Defines the additional IP Addresses configured on the device. |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortMirroring` | [`map[string]models.SwitchPortMirroringProperty`](../../doc/models/switch-port-mirroring-property.md) | Optional | Property key is the port mirroring instance name. `port_mirroring` can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 mirroring ports is allowed |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Switch RADIUS authentication and accounting configuration |
| `RemoteSyslog` | [`*models.RemoteSyslog`](../../doc/models/remote-syslog.md) | Optional | Remote syslog forwarding settings |
| `RoutingPolicies` | [`map[string]models.SwRoutingPolicy`](../../doc/models/sw-routing-policy.md) | Optional | Switch routing policies keyed by routing policy name |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SnmpConfig` | [`*models.SnmpConfig`](../../doc/models/snmp-config.md) | Optional | SNMP configuration for managed network devices |
| `StpConfig` | [`*models.SwitchStpConfig`](../../doc/models/switch-stp-config.md) | Optional | Switch spanning-tree protocol configuration |
| `SwitchMgmt` | [`*models.SwitchMgmt`](../../doc/models/switch-mgmt.md) | Optional | Switch management-plane access and proxy settings |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `switch`<br><br>**Value**: `"switch"` |
| `UseRouterIdAsSourceIp` | `*bool` | Optional | Whether to use it for snmp / syslog / tacplus / radius<br><br>**Default**: `false` |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | VRF enablement settings applied when supported on the device |
| `VrfInstances` | [`map[string]models.SwitchVrfInstance`](../../doc/models/switch-vrf-instance.md) | Optional | Property key is the network name |
| `VrrpConfig` | [`*models.VrrpConfig`](../../doc/models/vrrp-config.md) | Optional | Junos VRRP configuration applied to a switch or switch profile |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    deviceprofileSwitch := models.DeviceprofileSwitch{
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
            "key1": models.AclTag{
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
            "key2": models.AclTag{
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
            "additional_config_cmds4",
            "additional_config_cmds5",
            "additional_config_cmds6",
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
        Name:                  "name2",
        OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:                models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                  "switch",
        UseRouterIdAsSourceIp: models.ToPointer(false),
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
        AdditionalProperties:  map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

