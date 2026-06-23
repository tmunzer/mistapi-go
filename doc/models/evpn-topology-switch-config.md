
# Evpn Topology Switch Config

Per-switch configuration used by an EVPN topology member

## Structure

`EvpnTopologySwitchConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DhcpdConfig` | [`*models.EvpnTopologySwitchConfigDhcpdConfig`](../../doc/models/evpn-topology-switch-config-dhcpd-config.md) | Optional | DHCP server enablement for an EVPN topology switch override |
| `Networks` | [`map[string]models.SwitchNetwork`](../../doc/models/switch-network.md) | Optional | Property key is network name |
| `OtherIpConfigs` | [`map[string]models.JunosOtherIpConfig`](../../doc/models/junos-other-ip-config.md) | Optional | Additional IP addresses configured on the switch. Property key is the port network name |
| `PortConfig` | [`map[string]models.JunosPortConfig`](../../doc/models/junos-port-config.md) | Optional | Property key is the port name or range (e.g. "ge-0/0/0-10") |
| `PortUsages` | [`map[string]models.SwitchPortUsage`](../../doc/models/switch-port-usage.md) | Optional | Property key is the port usage name. Defines the profiles of port configuration configured on the switch |
| `RouterId` | `*string` | Optional | Used for OSPF / BGP / EVPN |
| `VrfConfig` | [`*models.EvpnTopologySwitchConfigVrfConfig`](../../doc/models/evpn-topology-switch-config-vrf-config.md) | Optional | VRF enablement for an EVPN topology switch override |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnTopologySwitchConfig := models.EvpnTopologySwitchConfig{
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
            "key2": models.SwitchNetwork{
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
            "key2": models.JunosPortConfig{
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
        },
        RouterId:             models.ToPointer("10.2.1.10"),
    }

}
```

