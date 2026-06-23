
# Response Rrm Neighbors

Paginated response for current RRM neighbor observations

## Structure

`ResponseRrmNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the RRM neighbor observation window |
| `Limit` | `int` | Required | Maximum number of RRM neighbor records returned in this page |
| `Next` | `*string` | Optional | Link to query next set of results. value is null if no next page exists. |
| `Results` | [`[]models.RrmNeighbors`](../../doc/models/rrm-neighbors.md) | Required | RRM neighbor records returned for the requested site and band<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the RRM neighbor observation window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseRrmNeighbors := models.ResponseRrmNeighbors{
        End:                  62,
        Limit:                148,
        Next:                 models.ToPointer("next0"),
        Results:              []models.RrmNeighbors{
            models.RrmNeighbors{
                Mac:                  models.ToPointer("5c5b35000001"),
                Neighbors:            []models.RrmNeighborsNeighbor{
                    models.RrmNeighborsNeighbor{
                        Mac:                  models.ToPointer("mac4"),
                        Rssi:                 models.ToPointer(56),
                    },
                },
            },
        },
        Start:                20,
    }

}
```

