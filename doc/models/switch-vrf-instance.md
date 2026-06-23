
# Switch Vrf Instance

Switch VRF instance routing and network membership settings

## Structure

`SwitchVrfInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AggregateRoutes` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "172.16.3.0/24") |
| `AggregateRoutes6` | [`map[string]models.AggregateRoute`](../../doc/models/aggregate-route.md) | Optional | Property key is the destination subnet (e.g. "2a02:1234:420a:10c9::/64") |
| `EvpnAutoLoopbackSubnet` | `*string` | Optional | IPv4 subnet used for automatic EVPN loopback addresses in this VRF instance |
| `EvpnAutoLoopbackSubnet6` | `*string` | Optional | IPv6 subnet used for automatic EVPN loopback addresses in this VRF instance |
| `ExtraRoutes` | [`map[string]models.VrfExtraRoute`](../../doc/models/vrf-extra-route.md) | Optional | Property key is the destination CIDR (e.g. "10.0.0.0/8") |
| `ExtraRoutes6` | [`map[string]models.VrfExtraRoute6`](../../doc/models/vrf-extra-route-6.md) | Optional | Property key is the destination CIDR (e.g. "2a02:1234:420a:10c9::/64") |
| `Networks` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchVrfInstance := models.SwitchVrfInstance{
        AggregateRoutes:         map[string]models.AggregateRoute{
            "key0": nil,
        },
        AggregateRoutes6:        map[string]models.AggregateRoute{
            "key0": nil,
            "key1": models.AggregateRoute{
            },
        },
        EvpnAutoLoopbackSubnet:  models.ToPointer("evpn_auto_loopback_subnet0"),
        EvpnAutoLoopbackSubnet6: models.ToPointer("evpn_auto_loopback_subnet64"),
        ExtraRoutes:             map[string]models.VrfExtraRoute{
            "0.0.0.0/0": models.VrfExtraRoute{
                Via:                  models.ToPointer("192.168.31.1"),
            },
        },
        Networks:                []string{
            "guest",
        },
    }

}
```

