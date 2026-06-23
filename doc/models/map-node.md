
# Map Node

Node in a map path graph

## Structure

`MapNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Edges` | `map[string]string` | Optional | Adjacent node IDs and path weights for this map node |
| `Name` | `string` | Required | Map node identifier or display name |
| `Position` | [`*models.MapNodePosition`](../../doc/models/map-node-position.md) | Optional | Position of a map path node |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapNode := models.MapNode{
        Edges:                map[string]string{
            "N1": "1",
        },
        Name:                 "N1",
        Position:             models.ToPointer(models.MapNodePosition{
            X:                    float64(224.72),
            Y:                    float64(100),
        }),
    }

}
```

