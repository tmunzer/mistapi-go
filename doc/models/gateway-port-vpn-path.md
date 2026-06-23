
# Gateway Port Vpn Path

VPN path settings for traffic that uses a gateway port

## Structure

`GatewayPortVpnPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BfdProfile` | [`*models.GatewayPortVpnPathBfdProfileEnum`](../../doc/models/gateway-port-vpn-path-bfd-profile-enum.md) | Optional | Only if the VPN `type`==`hub_spoke`. enum: `broadband`, `lte`<br><br>**Default**: `"broadband"` |
| `BfdUseTunnelMode` | `*bool` | Optional | Only if the VPN `type`==`hub_spoke`. Whether to use tunnel mode. SSR only<br><br>**Default**: `false` |
| `Preference` | `*int` | Optional | Only if the VPN `type`==`hub_spoke`. For a given VPN, when `path_selection.strategy`==`simple`, the preference for a path (lower is preferred) |
| `Role` | [`*models.GatewayPortVpnPathRoleEnum`](../../doc/models/gateway-port-vpn-path-role-enum.md) | Optional | If the VPN `type`==`hub_spoke`, enum: `hub`, `spoke`. If the VPN `type`==`mesh`, enum: `mesh`<br><br>**Default**: `"spoke"` |
| `TrafficShaping` | [`*models.GatewayTrafficShaping`](../../doc/models/gateway-traffic-shaping.md) | Optional | Traffic shaping settings for a gateway interface or VPN path |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortVpnPath := models.GatewayPortVpnPath{
        BfdProfile:           models.ToPointer(models.GatewayPortVpnPathBfdProfileEnum_BROADBAND),
        BfdUseTunnelMode:     models.ToPointer(false),
        Preference:           models.ToPointer(172),
        Role:                 models.ToPointer(models.GatewayPortVpnPathRoleEnum_SPOKE),
        TrafficShaping:       models.ToPointer(models.GatewayTrafficShaping{
            ClassPercentages:     []int{
                93,
                94,
                95,
            },
            Enabled:              models.ToPointer(false),
            MaxTxKbps:            models.ToPointer(212),
        }),
    }

}
```

