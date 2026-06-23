
# Gw Routing Policy Term Matching Route Exists

Route-existence match condition for a gateway routing policy term

## Structure

`GwRoutingPolicyTermMatchingRouteExists`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Route` | `*string` | Optional | Prefix that must exist for this condition to match |
| `VrfName` | `*string` | Optional | Name of the VRF instance where the route is checked; can also be a VPN or WAN name when applicable<br><br>**Default**: `"default"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gwRoutingPolicyTermMatchingRouteExists := models.GwRoutingPolicyTermMatchingRouteExists{
        Route:                models.ToPointer("192.168.0.0/24"),
        VrfName:              models.ToPointer("default"),
    }

}
```

