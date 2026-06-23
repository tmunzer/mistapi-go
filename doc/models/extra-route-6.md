
# Extra Route 6

IPv6 static route settings for a destination prefix

## Structure

`ExtraRoute6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `*bool` | Optional | Whether to install a discard route; this takes precedence over next-hop settings<br><br>**Default**: `false` |
| `Metric` | `models.Optional[int]` | Optional | Route metric for the IPv6 static route<br><br>**Constraints**: `>= 0`, `<= 2147483647` |
| `NextQualified` | [`map[string]models.ExtraRoute6NextQualifiedProperties`](../../doc/models/extra-route-6-next-qualified-properties.md) | Optional | Qualified next-hop settings keyed by IPv6 next-hop address |
| `NoResolve` | `*bool` | Optional | Whether to prevent recursive next-hop resolution for the IPv6 static route<br><br>**Default**: `false` |
| `Preference` | `models.Optional[int]` | Optional | Route preference for the IPv6 static route<br><br>**Constraints**: `>= 0`, `<= 2147483647` |
| `Via` | [`*models.NextHopVia`](../../doc/models/containers/next-hop-via.md) | Optional | Next-hop IP address. Can be a single IP address or an array of IP addresses for ECMP (Equal-Cost Multi-Path) load balancing across multiple next-hops. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    extraRoute6 := models.ExtraRoute6{
        Discard:              models.ToPointer(false),
        Metric:               models.NewOptional(models.ToPointer(66)),
        NextQualified:        map[string]models.ExtraRoute6NextQualifiedProperties{
            "2a02:1234:200a::100": models.ExtraRoute6NextQualifiedProperties{
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

