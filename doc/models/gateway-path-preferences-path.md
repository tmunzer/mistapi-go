
# Gateway Path Preferences Path

## Structure

`GatewayPathPreferencesPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cost` | `*int` | Optional | - |
| `Disabled` | `*bool` | Optional | For SSR Only. `true`` if this specific path is undesired |
| `GatewayIp` | `*string` | Optional | if `type`==`local`, if a different gateway is desired |
| `InternetAccess` | `*bool` | Optional | when `type`==`vpn`, if this vpn path can be used for internet<br>**Default**: `false` |
| `Name` | `*string` | Optional | - |
| `Networks` | `[]string` | Optional | if `type`==`local` |
| `TargetIps` | `[]string` | Optional | if `type`==`local`, if destination IP is to be replaced |
| `Type` | [`*models.GatewayPathTypeEnum`](../../doc/models/gateway-path-type-enum.md) | Optional | - |
| `WanName` | `*string` | Optional | Spoke's outgoing wan |

## Example (as JSON)

```json
{
  "internet_access": false,
  "wan_name": "wan0",
  "cost": 128,
  "disabled": false,
  "gateway_ip": "gateway_ip2",
  "name": "name6"
}
```

