
# Gateway Port Vpn Path

## Structure

`GatewayPortVpnPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BfdProfile` | [`*models.GatewayPortVpnPathBfdProfileEnum`](../../doc/models/gateway-port-vpn-path-bfd-profile-enum.md) | Optional | enum: `broadband`, `lte`<br>**Default**: `"broadband"` |
| `BfdUseTunnelMode` | `*bool` | Optional | whether to use tunnel mode. SSR only<br>**Default**: `false` |
| `Preference` | `*int` | Optional | for a given VPN, when `path_selection.strategy`==`simple`, the preference for a path (lower is preferred) |
| `Role` | [`*models.GatewayPortVpnPathRoleEnum`](../../doc/models/gateway-port-vpn-path-role-enum.md) | Optional | enum: `hub`, `spoke`<br>**Default**: `"spoke"` |
| `TrafficShaping` | [`*models.GatewayTrafficShaping`](../../doc/models/gateway-traffic-shaping.md) | Optional | - |

## Example (as JSON)

```json
{
  "bfd_profile": "broadband",
  "bfd_use_tunnel_mode": false,
  "role": "spoke",
  "preference": 38,
  "traffic_shaping": {
    "class_percentages": [
      93,
      94,
      95
    ],
    "enabled": false
  }
}
```

