
# Evpn Topology Switch

Switch member of an EVPN topology, including role and link relationships

## Structure

`EvpnTopologySwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Config` | [`*models.EvpnTopologySwitchConfig`](../../doc/models/evpn-topology-switch-config.md) | Optional | Per-switch configuration used by an EVPN topology member |
| `DeviceprofileId` | `*uuid.UUID` | Optional, Read-only | Associated device profile identifier for the switch. Use the [Assign Org Device Profile](../../doc/controllers/orgs-device-profiles.md#assign-org-device-profile) endpoint to assign a Device Profile to the switch. |
| `DownlinkIps` | `[]string` | Optional, Read-only | Downlink IP addresses used by an EVPN topology switch |
| `Downlinks` | `[]string` | Optional | Downlink switch MAC addresses for an EVPN topology switch |
| `Esilaglinks` | `[]string` | Optional | ESI-LAG switch MAC addresses for an EVPN topology switch |
| `EvpnId` | `*int` | Optional, Read-only | Topology identifier number for this EVPN switch member<br><br>**Constraints**: `>= 1` |
| `Mac` | `string` | Required | Switch MAC address used to identify the topology member<br><br>**Constraints**: *Minimum Length*: `1` |
| `Model` | `*string` | Optional, Read-only | Switch model for this topology member |
| `Pod` | `*int` | Optional | Optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.<br><br>* for CLOS, to group dist / access switches into pods<br>* for ERB/CRB, to group dist / esilag-access into pods<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1`, `<= 255` |
| `Pods` | `[]int` | Optional | By default, core switches are assumed to be connecting all pods.<br>if you want to limit the pods, you can specify pods. |
| `Role` | [`models.EvpnTopologySwitchRoleEnum`](../../doc/models/evpn-topology-switch-role-enum.md) | Required | use `role`==`none` to remove a switch from the topology. enum: `access`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`<br><br>**Constraints**: *Minimum Length*: `1` |
| `RouterId` | `*string` | Optional, Read-only | Routing identifier used by this switch for EVPN routing |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SuggestedDownlinks` | `[]string` | Optional, Read-only | Builder-suggested downlink switch MAC addresses for an EVPN topology switch |
| `SuggestedEsilaglinks` | `[]string` | Optional, Read-only | Builder-suggested ESI-LAG switch MAC addresses for an EVPN topology switch |
| `SuggestedUplinks` | `[]string` | Optional, Read-only | Builder-suggested uplink switch MAC addresses for an EVPN topology switch |
| `Uplinks` | `[]string` | Optional | Uplink switch MAC addresses for an EVPN topology switch |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    evpnTopologySwitch := models.EvpnTopologySwitch{
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
            "downlink_ips0",
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
        Role:                 models.EvpnTopologySwitchRoleEnum_ESILAGACCESS,
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
    }

}
```

