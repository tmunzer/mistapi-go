
# Gateway Path Preferences Path

## Structure

`GatewayPathPreferencesPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cost` | `*int` | Optional | - |
| `Disabled` | `*bool` | Optional | For SSR Only. `true`, if this specific path is undesired |
| `GatewayIp` | `*string` | Optional | only if `type`==`local`, if a different gateway is desired |
| `InternetAccess` | `*bool` | Optional | only if `type`==`vpn`, if this vpn path can be used for internet |
| `Name` | `*string` | Optional | required when * `type`==`vpn`: the name of the VPN Path to use * `type`==`wan`: the name of the WAN interface to use |
| `Networks` | `[]string` | Optional | required when `type`==`local` |
| `TargetIps` | `[]string` | Optional | if `type`==`local`, if destination IP is to be replaced |
| `Type` | [`*models.GatewayPathTypeEnum`](../../doc/models/gateway-path-type-enum.md) | Optional | - |
| `WanName` | `*string` | Optional | required when`type`==`tunnel`, optional if `type`==`vpn` wan |

## Example (as JSON)

```json
{
  "wan_name": "wan0",
  "cost": 128,
  "disabled": false,
  "gateway_ip": "gateway_ip2",
  "internet_access": false,
  "name": "name6"
}
```

