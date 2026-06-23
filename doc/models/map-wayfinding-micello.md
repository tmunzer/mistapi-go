
# Map Wayfinding Micello

Micello wayfinding integration settings

## Structure

`MapWayfindingMicello`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountKey` | `*string` | Optional | Micello account key used for wayfinding |
| `DefaultLevelId` | `*int` | Optional | Default Micello floor or level identifier for wayfinding |
| `MapId` | `*string` | Optional | Micello map identifier used for wayfinding |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapWayfindingMicello := models.MapWayfindingMicello{
        AccountKey:           models.ToPointer("adasdf"),
        DefaultLevelId:       models.ToPointer(5),
        MapId:                models.ToPointer("c660f81dd250c"),
    }

}
```

