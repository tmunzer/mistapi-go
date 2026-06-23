
# Wan Extra Routes

Additional IPv4 route for a WAN interface

## Structure

`WanExtraRoutes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | IPv4 next-hop address for this WAN extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wanExtraRoutes := models.WanExtraRoutes{
        Via:                  models.ToPointer("via6"),
    }

}
```

