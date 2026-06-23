
# Extra Route

IPv4 static route settings for a destination prefix

## Structure

`ExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | Whether to install a discard route; this takes precedence over next-hop settings<br><br>**Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | Route metric for the IPv4 static route<br><br>**Constraints**: `>= 0`, `<= 2147483647` |
| `NextQualified` | [`map[string]models.ExtraRouteNextQualifiedProperties`](../../doc/models/extra-route-next-qualified-properties.md) | Optional | Qualified next-hop settings keyed by IPv4 next-hop address |
| `NoResolve` | `*bool` | Optional | Whether to prevent recursive next-hop resolution for the IPv4 static route<br><br>**Default**: `false` |
| `Preference` | `models.Optional[int]` | Optional | Route preference for the IPv4 static route<br><br>**Constraints**: `>= 0`, `<= 2147483647` |
| `Via` | [`*models.NextHopVia`](../../doc/models/containers/next-hop-via.md) | Optional | Next-hop IP address. Can be a single IP address or an array of IP addresses for ECMP (Equal-Cost Multi-Path) load balancing across multiple next-hops. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    extraRoute := models.ExtraRoute{
        Discard:              models.ToPointer(false),
        Metric:               models.NewOptional(models.ToPointer(110)),
        NextQualified:        map[string]models.ExtraRouteNextQualifiedProperties{
            "10.3.1.1": models.ExtraRouteNextQualifiedProperties{
                Metric:               models.NewOptional(models.ToPointer(170)),
                Preference:           models.NewOptional(models.ToPointer(40)),
            },
        },
        NoResolve:            models.ToPointer(false),
        Preference:           models.NewOptional(models.ToPointer(30)),
        Via:                  models.ToPointer(),
    }

}
```

