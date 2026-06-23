
# Gateway Path Preferences Path

Candidate path within a gateway path preference

## Structure

`GatewayPathPreferencesPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cost` | `*int` | Optional | Relative cost assigned to this path for gateway path selection |
| `Disabled` | `*bool` | Optional | For SSR Only. `true`, if this specific path is undesired |
| `GatewayIp` | `*string` | Optional | Only if `type`==`local`, if a different gateway is desired |
| `InternetAccess` | `*bool` | Optional | Only if `type`==`vpn`, if this vpn path can be used for internet |
| `Name` | `*string` | Optional | Required when<br><br>* `type`==`vpn`: the name of the VPN Path to use<br>* `type`==`wan`: the name of the WAN interface to use |
| `Networks` | `[]string` | Optional | Network names used by a local path; required when `type`==`local` |
| `TargetIps` | `[]string` | Optional | Destination IP addresses to replace when `type`==`local` |
| `Type` | [`models.GatewayPathTypeEnum`](../../doc/models/gateway-path-type-enum.md) | Required | enum: `local`, `tunnel`, `vpn`, `wan` |
| `WanName` | `*string` | Optional | Optional if `type`==`vpn`; WAN interface name associated with the VPN path |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPathPreferencesPath := models.GatewayPathPreferencesPath{
        Cost:                 models.ToPointer(50),
        Disabled:             models.ToPointer(false),
        GatewayIp:            models.ToPointer("gateway_ip0"),
        InternetAccess:       models.ToPointer(false),
        Name:                 models.ToPointer("name4"),
        Type:                 models.GatewayPathTypeEnum_LOCAL,
        WanName:              models.ToPointer("wan0"),
    }

}
```

