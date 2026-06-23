
# Response Auto Placement Info

Auto-placement status response with progress and validation details

## Structure

`ResponseAutoPlacementInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EndTime` | `*float64` | Optional | Timestamp when autoplacement completed or was manually stopped |
| `EstTimeLeft` | `*float64` | Optional | (Only when inprogress) estimate of the time to completion |
| `StartTime` | `*int` | Optional | Timestamp when autoplacement process was last queued for this map |
| `Status` | [`*models.AutoPlacementInfoStatusEnum`](../../doc/models/auto-placement-info-status-enum.md) | Optional | the status of autoplacement for a given map. enum: `done`, `error`, `inprogress`, `pending` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoPlacementInfo := models.ResponseAutoPlacementInfo{
        EndTime:              models.ToPointer(float64(91.48)),
        EstTimeLeft:          models.ToPointer(float64(40.6)),
        StartTime:            models.ToPointer(228),
        Status:               models.ToPointer(models.AutoPlacementInfoStatusEnum_DONE),
    }

}
```

