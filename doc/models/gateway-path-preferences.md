
# Gateway Path Preferences

Gateway path preference that selects among one or more local, WAN, VPN, or tunnel paths

## Structure

`GatewayPathPreferences`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Paths` | [`[]models.GatewayPathPreferencesPath`](../../doc/models/gateway-path-preferences-path.md) | Optional | Candidate paths evaluated by a gateway path preference |
| `Strategy` | [`*models.GatewayPathStrategyEnum`](../../doc/models/gateway-path-strategy-enum.md) | Optional | enum: `ecmp`, `ordered`, `weighted`<br><br>**Default**: `"ordered"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPathPreferences := models.GatewayPathPreferences{
        Paths:                []models.GatewayPathPreferencesPath{
            models.GatewayPathPreferencesPath{
                Cost:                 models.ToPointer(118),
                Disabled:             models.ToPointer(false),
                GatewayIp:            models.ToPointer("gateway_ip0"),
                InternetAccess:       models.ToPointer(false),
                Name:                 models.ToPointer("name4"),
                Type:                 models.GatewayPathTypeEnum_LOCAL,
            },
            models.GatewayPathPreferencesPath{
                Cost:                 models.ToPointer(118),
                Disabled:             models.ToPointer(false),
                GatewayIp:            models.ToPointer("gateway_ip0"),
                InternetAccess:       models.ToPointer(false),
                Name:                 models.ToPointer("name4"),
                Type:                 models.GatewayPathTypeEnum_LOCAL,
            },
        },
        Strategy:             models.ToPointer(models.GatewayPathStrategyEnum_ORDERED),
    }

}
```

