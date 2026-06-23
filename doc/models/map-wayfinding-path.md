
# Map Wayfinding Path

JSON blob for wayfinding (using Dijkstra’s algorithm)

## Structure

`MapWayfindingPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coordinate` | `*string` | Optional | Wayfinding path coordinate space |
| `Nodes` | [`[]models.MapNode`](../../doc/models/map-node.md) | Optional | Map nodes included in a wayfinding path<br><br>**Constraints**: *Minimum Items*: `0` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapWayfindingPath := models.MapWayfindingPath{
        Coordinate:           models.ToPointer("actual"),
        Nodes:                []models.MapNode{
            models.MapNode{
                Edges:                map[string]string{
                    "key0": "edges6",
                },
                Name:                 "name6",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(224.72),
                    Y:                    float64(100),
                }),
            },
            models.MapNode{
                Edges:                map[string]string{
                    "key0": "edges6",
                },
                Name:                 "name6",
                Position:             models.ToPointer(models.MapNodePosition{
                    X:                    float64(224.72),
                    Y:                    float64(100),
                }),
            },
        },
    }

}
```

