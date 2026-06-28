
# Evpn Topology

EVPN topology create or update payload

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnTopology`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `EvpnOptions` | [`*models.EvpnOptions`](../../doc/models/evpn-options.md) | Optional | EVPN topology generation options for campus fabric configuration |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name for the EVPN topology |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Overwrite` | `*bool` | Optional | Whether to apply generated EVPN configuration changes automatically |
| `PodNames` | `map[string]string` | Optional | Property key is the pod number |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SwitchConfigs` | [`map[string]models.EvpnTopologySwitchConfig`](../../doc/models/evpn-topology-switch-config.md) | Optional | Property key is the switch MAC address |
| `Switches` | [`[]models.EvpnTopologySwitch`](../../doc/models/evpn-topology-switch.md) | Required | Switch members included in an EVPN topology<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    evpnTopology := models.EvpnTopology{
        CreatedTime:          models.ToPointer(float64(145.14)),
        EvpnOptions:          models.ToPointer(models.EvpnOptions{
            AutoLoopbackSubnet:   models.ToPointer("auto_loopback_subnet4"),
            AutoLoopbackSubnet6:  models.ToPointer("auto_loopback_subnet60"),
            AutoRouterIdSubnet:   models.ToPointer("auto_router_id_subnet8"),
            AutoRouterIdSubnet6:  models.ToPointer("auto_router_id_subnet60"),
            CoreAsBorder:         models.ToPointer(false),
        }),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(189.82)),
        Name:                 models.ToPointer("CC"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Switches:             []models.EvpnTopologySwitch{
            models.EvpnTopologySwitch{
                Config:               models.ToPointer(models.EvpnTopologySwitchConfig{
                    DhcpdConfig:          models.ToPointer(models.EvpnTopologySwitchConfigDhcpdConfig{
                        Enabled:              models.ToPointer(false),
                    }),
                    Networks:             map[string]models.SwitchNetwork{
                        "key0": models.SwitchNetwork{
                            Gateway:              models.ToPointer("gateway8"),
                            Gateway6:             models.ToPointer("gateway64"),
                            Isolation:            models.ToPointer(false),
                            IsolationVlanId:      models.ToPointer("isolation_vlan_id8"),
                            Subnet:               models.ToPointer("subnet6"),
                            VlanId:               models.VlanIdWithVariableContainer.FromString("String7"),
                        },
                        "key1": models.SwitchNetwork{
                            Gateway:              models.ToPointer("gateway8"),
                            Gateway6:             models.ToPointer("gateway64"),
                            Isolation:            models.ToPointer(false),
                            IsolationVlanId:      models.ToPointer("isolation_vlan_id8"),
                            Subnet:               models.ToPointer("subnet6"),
                            VlanId:               models.VlanIdWithVariableContainer.FromString("String7"),
                        },
                    },
                    OtherIpConfigs:       map[string]models.JunosOtherIpConfig{
                        "key0": models.JunosOtherIpConfig{
                            EvpnAnycast:          models.ToPointer(false),
                            Ip:                   models.ToPointer("ip4"),
                            Ip6:                  models.ToPointer("ip60"),
                            Netmask:              models.ToPointer("netmask0"),
                            Netmask6:             models.ToPointer("netmask60"),
                        },
                        "key1": models.JunosOtherIpConfig{
                            EvpnAnycast:          models.ToPointer(false),
                            Ip:                   models.ToPointer("ip4"),
                            Ip6:                  models.ToPointer("ip60"),
                            Netmask:              models.ToPointer("netmask0"),
                            Netmask6:             models.ToPointer("netmask60"),
                        },
                        "key2": models.JunosOtherIpConfig{
                            EvpnAnycast:          models.ToPointer(false),
                            Ip:                   models.ToPointer("ip4"),
                            Ip6:                  models.ToPointer("ip60"),
                            Netmask:              models.ToPointer("netmask0"),
                            Netmask6:             models.ToPointer("netmask60"),
                        },
                    },
                    PortConfig:           map[string]models.JunosPortConfig{
                        "key0": models.JunosPortConfig{
                            AeDisableLacp:        models.ToPointer(false),
                            AeIdx:                models.ToPointer(230),
                            AeLacpForceUp:        models.ToPointer(false),
                            AeLacpPassive:        models.ToPointer(false),
                            AeLacpSlow:           models.ToPointer(false),
                            Usage:                "usage6",
                        },
                        "key1": models.JunosPortConfig{
                            AeDisableLacp:        models.ToPointer(false),
                            AeIdx:                models.ToPointer(230),
                            AeLacpForceUp:        models.ToPointer(false),
                            AeLacpPassive:        models.ToPointer(false),
                            AeLacpSlow:           models.ToPointer(false),
                            Usage:                "usage6",
                        },
                    },
                    PortUsages:           map[string]models.SwitchPortUsage{
                        "key0": models.SwitchPortUsage{
                            AllNetworks:                              models.ToPointer(false),
                            AllowDhcpd:                               models.ToPointer(false),
                            AllowMultipleSupplicants:                 models.ToPointer(false),
                            BypassAuthWhenServerDown:                 models.ToPointer(false),
                            BypassAuthWhenServerDownForUnknownClient: models.ToPointer(false),
                        },
                        "key1": models.SwitchPortUsage{
                            AllNetworks:                              models.ToPointer(false),
                            AllowDhcpd:                               models.ToPointer(false),
                            AllowMultipleSupplicants:                 models.ToPointer(false),
                            BypassAuthWhenServerDown:                 models.ToPointer(false),
                            BypassAuthWhenServerDownForUnknownClient: models.ToPointer(false),
                        },
                        "key2": models.SwitchPortUsage{
                            AllNetworks:                              models.ToPointer(false),
                            AllowDhcpd:                               models.ToPointer(false),
                            AllowMultipleSupplicants:                 models.ToPointer(false),
                            BypassAuthWhenServerDown:                 models.ToPointer(false),
                            BypassAuthWhenServerDownForUnknownClient: models.ToPointer(false),
                        },
                    },
                }),
                DeviceprofileId:      models.ToPointer(uuid.MustParse("6a1deab1-96df-4fa2-8455-d5253f943d06")),
                DownlinkIps:          []string{
                    "downlink_ips6",
                },
                Downlinks:            []string{
                    "5c5b35000005",
                    "5c5b35000006",
                },
                Esilaglinks:          []string{
                    "5c5b35000005",
                    "5c5b35000006",
                },
                Mac:                  "5c5b35000003",
                Model:                models.ToPointer("QFX10002-36Q"),
                Pod:                  models.ToPointer(1),
                Role:                 models.EvpnTopologySwitchRoleEnum_NONE,
                RouterId:             models.ToPointer("172.16.254.4"),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                SuggestedDownlinks:   []string{
                    "5c5b35000005",
                    "5c5b35000006",
                },
                SuggestedEsilaglinks: []string{
                    "5c5b35000005",
                    "5c5b35000006",
                },
                SuggestedUplinks:     []string{
                    "5c5b35000005",
                    "5c5b35000006",
                },
                Uplinks:              []string{
                    "5c5b35000005",
                    "5c5b35000006",
                },
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

