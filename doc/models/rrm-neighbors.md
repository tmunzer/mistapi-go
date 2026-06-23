
# Rrm Neighbors

RRM neighbor observations for one AP

## Structure

`RrmNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | AP MAC address whose neighbors are reported |
| `Neighbors` | [`[]models.RrmNeighborsNeighbor`](../../doc/models/rrm-neighbors-neighbor.md) | Optional | Neighbor AP observations heard by an AP |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmNeighbors := models.RrmNeighbors{
        Mac:                  models.ToPointer("5c5b35000001"),
        Neighbors:            []models.RrmNeighborsNeighbor{
            models.RrmNeighborsNeighbor{
                Mac:                  models.ToPointer("mac4"),
                Rssi:                 models.ToPointer(56),
            },
            models.RrmNeighborsNeighbor{
                Mac:                  models.ToPointer("mac4"),
                Rssi:                 models.ToPointer(56),
            },
            models.RrmNeighborsNeighbor{
                Mac:                  models.ToPointer("mac4"),
                Rssi:                 models.ToPointer(56),
            },
        },
    }

}
```

