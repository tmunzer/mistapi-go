
# Device Gateway

Gateway configuration and placement data

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalConfigCmds` | `[]string` | Optional | additional CLI commands to append to the generated Junos config. **Note**: no check is done |
| `BgpConfig` | [`map[string]models.BgpConfig`](../../doc/models/bgp-config.md) | Optional | BGP routing configuration for this gateway. Property key is the BGP session name |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceprofileId` | `*uuid.UUID` | Optional, Read-only | Device profile associated with this gateway |
| `DhcpdConfig` | [`*models.DhcpdConfig`](../../doc/models/dhcpd-config.md) | Optional | DHCP server configuration map with a global enable flag |
| `DnsServers` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `DnsSuffix` | `[]string` | Optional | Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting |
| `ExtraRoutes` | [`map[string]models.GatewayExtraRoute`](../../doc/models/gateway-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `ExtraRoutes6` | [`map[string]models.GatewayExtraRoute6`](../../doc/models/gateway-extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64"), the destination Network name or a variable (e.g. "{{myvar}}") |
| `ForSite` | `*bool` | Optional, Read-only | Whether the gateway configuration is scoped directly to a site |
| `GatewayMgmt` | [`*models.GatewayMgmt`](../../doc/models/gateway-mgmt.md) | Optional | Gateway management-plane and access settings |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IdpProfiles` | [`map[string]models.IdpProfile`](../../doc/models/idp-profile.md) | Optional | Property key is the profile name |
| `Image1Url` | `models.Optional[string]` | Optional | First custom image URL associated with the gateway |
| `Image2Url` | `models.Optional[string]` | Optional | Second custom image URL associated with the gateway |
| `Image3Url` | `models.Optional[string]` | Optional | Third custom image URL associated with the gateway |
| `IpConfigs` | [`map[string]models.GatewayIpConfigProperty`](../../doc/models/gateway-ip-config-property.md) | Optional | Property key is the network name |
| `Mac` | `*string` | Optional, Read-only | Gateway MAC address used to identify the device |
| `Managed` | `*bool` | Optional | Whether the device is managed by Mist. Deprecated in favour of mist_configured. |
| `MapId` | `*uuid.UUID` | Optional | Map where the device belongs to |
| `MistConfigured` | `*bool` | Optional | whether the device can be configured by Mist or not. This deprecates `managed` for adopted devices. |
| `Model` | `*string` | Optional, Read-only | Gateway model reported for the device |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `Name` | `*string` | Optional | Friendly display name assigned to the gateway |
| `Networks` | [`[]models.Network`](../../doc/models/network.md) | Optional | List of organization network definitions |
| `Notes` | `*string` | Optional | Free-form administrative notes for this gateway |
| `NtpServers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `OobIpConfig` | [`*models.GatewayOobIpConfig`](../../doc/models/gateway-oob-ip-config.md) | Optional | Out-of-band management IP configuration for gateway interfaces such as vme, em0, or fxp0 |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PathPreferences` | [`map[string]models.GatewayPathPreferences`](../../doc/models/gateway-path-preferences.md) | Optional | Property key is the path name |
| `PortConfig` | [`map[string]models.GatewayPortConfig`](../../doc/models/gateway-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortMirroring` | [`*models.GatewayPortMirroring`](../../doc/models/gateway-port-mirroring.md) | Optional | Port mirroring settings for a gateway interface |
| `RouterId` | `*string` | Optional | Auto assigned if not set |
| `RoutingPolicies` | [`map[string]models.GwRoutingPolicy`](../../doc/models/gw-routing-policy.md) | Optional | Property key is the routing policy name |
| `Serial` | `*string` | Optional, Read-only | Manufacturer serial number for the gateway |
| `ServicePolicies` | [`[]models.ServicePolicy`](../../doc/models/service-policy.md) | Optional | Service policies returned by derived site configuration |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TunnelConfigs` | [`map[string]models.TunnelConfig`](../../doc/models/tunnel-config.md) | Optional | Property key is the tunnel name |
| `TunnelProviderOptions` | [`*models.TunnelProviderOptions`](../../doc/models/tunnel-provider-options.md) | Optional | Provider-specific options for gateway tunnel auto provisioning |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `gateway`<br><br>**Value**: `"gateway"` |
| `UrlFilteringDenyMsg` | `*string` | Optional | When a service policy denies a app_category, what message to show in user's browser<br><br>**Default**: `"Access to this URL Category has been blocked"` |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `VrfConfig` | [`*models.VrfConfig`](../../doc/models/vrf-config.md) | Optional | VRF enablement settings applied when supported on the device |
| `VrfInstances` | [`map[string]models.GatewayVrfInstance`](../../doc/models/gateway-vrf-instance.md) | Optional | Property key is the VRF instance name |
| `X` | `*float64` | Optional | Horizontal map position of the gateway, in pixels |
| `Y` | `*float64` | Optional | Vertical map position of the gateway, in pixels |
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
    deviceGateway := models.DeviceGateway{
        AdditionalConfigCmds:    []string{
            "additional_config_cmds4",
            "additional_config_cmds5",
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
            "key1": models.BgpConfig{
                AuthKey:                models.ToPointer("auth_key8"),
                BfdMinimumInterval:     models.NewOptional(models.ToPointer(212)),
                BfdMultiplier:          models.NewOptional(models.ToPointer(90)),
                DisableBfd:             models.ToPointer(false),
                Export:                 models.ToPointer("export6"),
                Via:                    models.BgpConfigViaEnum_VPN,
            },
        },
        CreatedTime:             models.ToPointer(float64(85.12)),
        DeviceprofileId:         models.ToPointer(uuid.MustParse("00001aec-0000-0000-0000-000000000000")),
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
        ExtraRoutes6:            map[string]models.GatewayExtraRoute6{
            "2a02:1234:420a:10c9::/64": models.GatewayExtraRoute6{
                Via:                  models.ToPointer("2a02:1234:200a::100"),
            },
        },
        Id:                      models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MapId:                   models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
        MspId:                   models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        OrgId:                   models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RouterId:                models.ToPointer("10.2.1.10"),
        SiteId:                  models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                    "gateway",
        UrlFilteringDenyMsg:     models.ToPointer("Access to this URL Category has been blocked"),
        Vars:                    map[string]string{
            "RADIUS_IP1": "172.31.2.5",
            "RADIUS_SECRET": "11s64632d",
        },
        VrfInstances:            map[string]models.GatewayVrfInstance{
            "CORP_VRF": models.GatewayVrfInstance{
                Networks:             []string{
                    "CORP_NET",
                    "MGMT_NET",
                },
            },
        },
        X:                       models.ToPointer(float64(53.5)),
        Y:                       models.ToPointer(float64(173.1)),
        AdditionalProperties:    map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

