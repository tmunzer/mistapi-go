
# Map Wall Path

JSON blob for wall definition (same format as wayfinding_path)

## Structure

`MapWallPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Coordinate` | `*string` | Optional | Wall path coordinate space |
| `Nodes` | [`[]models.MapNode`](../../doc/models/map-node.md) | Optional | Map nodes included in a wall path<br><br>**Constraints**: *Minimum Items*: `0` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapWallPath := models.MapWallPath{
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

