
# Gateway Template

Gateway Template is applied to a site for gateway(s) in a site.

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | BGP routing defaults for this gateway template. Property key is the BGP session name |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | DHCP server configuration map with a global enable flag |
| `DnsOverride` | `*bool` | Optional | Whether DNS server and suffix settings in this template override inherited values<br><br>**Default**: `false` |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `ExtraRoutes6` | [`map[string]models.GatewayExtraRoute6`](../../doc/models/gateway-extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `GatewayMatching` | [`*models.GatewayMatching`](../../doc/models/gateway-matching.md) | Optional | Gateway matching configuration used to apply gateway-specific settings |
| `GatewayMgmt` | [`*models.GatewayMgmt`](../../doc/models/gateway-mgmt.md) | Optional | Gateway management-plane and access settings |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the gateway template |
| `Networks` | [`[]models.Network`](../../doc/models/network.md) | Optional | List of organization network definitions |
| `NtpOverride` | `*bool` | Optional | Whether NTP servers in this template override inherited values<br><br>**Default**: `false` |
| `NtpServers` | `[]string` | Optional | List of NTP servers specific to this device. By default, those in Site Settings will be used |
| `OobIpConfig` | [`*models.GatewayOobIpConfig`](../../doc/models/gateway-oob-ip-config.md) | Optional | Out-of-band management IP configuration for gateway interfaces such as vme, em0, or fxp0 |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `PortConfig` | [`map[string]models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Property key is the Port Name (i.e. "ge-0/0/0"), the Ports Range (i.e. "ge-0/0/0-10"), the List of Ports (i.e. "ge-0/0/0,ge-1/0/0", only allowed for Aggregated or Redundant interfaces) or a Variable (i.e. "{{myvar}}"). |
| `RouterId` | `*string` | Optional | Auto assigned if not set |
| `RoutingPolicies` | [`map[string]models.GwRoutingPolicy`](../../doc/models/gw-routing-policy.md) | Optional | Property key is the routing policy name |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | Service policies returned by derived site configuration |
| `TunnelConfigs` | [`map[string]models.TunnelConfig`](../../doc/models/tunnel-config.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | Provider-specific options for gateway tunnel auto provisioning |
| `Type` | [`*models.GatewayTemplateTypeEnum`](../../doc/models/gateway-template-type-enum.md) | Optional | Gateway template deployment type. enum: `spoke`, `standalone`<br><br>**Default**: `"standalone"` |
| `UrlFilteringDenyMsg` | `*string` | Optional | When a service policy denies a app_category, what message to show in user's browser<br><br>**Default**: `"Access to this URL Category has been blocked"` |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | VRF enablement settings applied when supported on the device |
| `VrfInstances` | [`map[string]models.GatewayVrfInstance`](../../doc/models/gateway-vrf-instance.md) | Optional | Property key is the VRF instance name |
| `SsrAdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated SSR config. **Note**: no check is done |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    gatewayTemplate := models.GatewayTemplate{
        AdditionalConfigCmds:    []string{
            "additional_config_cmds2",
        },
        BgpConfig:               map[string]models.BgpConfig{
            "key0": models.BgpConfig{
                AuthKey:                models.ToPointer("auth_key8"),
                BfdMinimumInterval:     models.NewOptional(models.ToPointer(212)),
                BfdMultiplier:          models.NewOptional(models.ToPointer(90)),
                DisableBfd:             models.ToPointer(false),
                Export:                 models.ToPointer("export6"),
                Via:                    models.BgpConfigViaEnum_VPN,
            },
        },
        CreatedTime:             models.ToPointer(float64(121.78)),
        DhcpdConfig:             models.ToPointer(models.DhcpdConfig{
            Enabled:              models.ToPointer(false),
            AdditionalProperties: map[string]models.DhcpdConfigProperty{
                "exampleAdditionalProperty": models.DhcpdConfigProperty{
                    DnsServers:           []string{
                        "dns_servers2",
                        "dns_servers3",
                        "dns_servers4",
                    },
                    DnsSuffix:            []string{
                        "dns_suffix1",
                        "dns_suffix0",
                        "dns_suffix9",
                    },
                    FixedBindings:        map[string]models.DhcpdConfigFixedBinding{
                        "key0": nil,
                    },
                    Gateway:              models.ToPointer("gateway8"),
                    Ip6End:               models.ToPointer("ip6_end8"),
                },
            },
        }),
        DnsOverride:             models.ToPointer(false),
        ExtraRoutes6:            map[string]models.GatewayExtraRoute6{
            "2a02:1234:420a:10c9::/64": models.GatewayExtraRoute6{
                Via:                  models.ToPointer("2a02:1234:200a::100"),
            },
        },
        Id:                      models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                    "gw_template",
        NtpOverride:             models.ToPointer(false),
        OrgId:                   models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RouterId:                models.ToPointer("10.2.1.10"),
        Type:                    models.ToPointer(models.GatewayTemplateTypeEnum_STANDALONE),
        UrlFilteringDenyMsg:     models.ToPointer("Access to this URL Category has been blocked"),
        VrfInstances:            map[string]models.GatewayVrfInstance{
            "CORP_VRF": models.GatewayVrfInstance{
                Networks:             []string{
                    "CORP_NET",
                    "MGMT_NET",
                },
            },
        },
        AdditionalProperties:    map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

