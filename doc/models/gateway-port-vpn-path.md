
# Gateway Port Vpn Path

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayPortVpnPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BfdProfile` | [`*models.GatewayPortVpnPathBfdProfileEnum`](../../doc/models/gateway-port-vpn-path-bfd-profile-enum.md) | Optional | Only if the VPN `type`==`hub_spoke`. enum: `broadband`, `lte`<br>**Default**: `"broadband"` |
| `BfdUseTunnelMode` | `*bool` | Optional | Only if the VPN `type`==`hub_spoke`. Whether to use tunnel mode. SSR only<br>**Default**: `false` |
| `Preference` | `*int` | Optional | Only if the VPN `type`==`hub_spoke`. For a given VPN, when `path_selection.strategy`==`simple`, the preference for a path (lower is preferred) |
| `Role` | [`*models.GatewayPortVpnPathRoleEnum`](../../doc/models/gateway-port-vpn-path-role-enum.md) | Optional | If the VPN `type`==`hub_spoke`, enum: `hub`, `spoke`. If the VPN `type`==`mesh`, enum: `mesh`<br>**Default**: `"spoke"` |
| `TrafficShaping` | [`*models.GatewayPortVpnPathTrafficShaping`](../../doc/models/gateway-port-vpn-path-traffic-shaping.md) | Optional | Only if the VPN `type`==`hub_spoke` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
    "enabled": false,
    "max_tx_kbps": 212,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

