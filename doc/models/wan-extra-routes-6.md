
# Wan Extra Routes 6

Additional IPv6 route for a WAN interface

## Structure

`WanExtraRoutes6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | IPv6 next-hop address for this WAN extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wanExtraRoutes6 := models.WanExtraRoutes6{
        Via:                  models.ToPointer("via6"),
    }

}
```

