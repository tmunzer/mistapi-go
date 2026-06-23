
# Response Auto Orientation Info

Auto orientation status response

## Structure

`ResponseAutoOrientationInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EstTimeLeft` | `*float64` | Optional | Only when `status`==`inprogress`, estimate of the time to completion |
| `StartTime` | `*float64` | Optional | Timestamp when auto orient process was last queued for this map |
| `Status` | [`*models.ResponseAutoOrientationInfoStatusEnum`](../../doc/models/response-auto-orientation-info-status-enum.md) | Optional | The status of auto orient for a given map. enum:<br><br>* `pending`: Auto orient has not been requested for this map<br>* `inprogress`: Auto orient is currently processing<br>* `done`: The auto orient process has completed<br>* `error`: There was an error in the auto orient process |
| `StopTime` | `*float64` | Optional | Timestamp when auto orient completed or was manually stopped |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoOrientationInfo := models.ResponseAutoOrientationInfo{
        EstTimeLeft:          models.ToPointer(float64(3.56)),
        StartTime:            models.ToPointer(float64(223.64)),
        Status:               models.ToPointer(models.ResponseAutoOrientationInfoStatusEnum_PENDING),
        StopTime:             models.ToPointer(float64(20.42)),
    }

}
```

