
# Map Wayfinding

Properties related to wayfinding

## Structure

`MapWayfinding`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Micello` | [`*models.MapWayfindingMicello`](../../doc/models/map-wayfinding-micello.md) | Optional | Micello wayfinding integration settings |
| `SnapToPath` | `*bool` | Optional | Whether wayfinding should snap routes to configured paths |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapWayfinding := models.MapWayfinding{
        Micello:              models.ToPointer(models.MapWayfindingMicello{
            AccountKey:           models.ToPointer("account_key8"),
            DefaultLevelId:       models.ToPointer(220),
            MapId:                models.ToPointer("map_id8"),
        }),
        SnapToPath:           models.ToPointer(false),
    }

}
```

