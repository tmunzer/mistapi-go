
# Response Async License

Asynchronous license claim progress response

## Structure

`ResponseAsyncLicense`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Completed` | `[]string` | Optional | Device MAC addresses that completed asynchronous license claim processing |
| `Details` | [`[]models.ResponseAsyncLicenseDetail`](../../doc/models/response-async-license-detail.md) | Optional | Per-device asynchronous license claim status details |
| `Failed` | `*int` | Optional | Number of devices that failed license claim processing |
| `Incompleted` | `[]string` | Optional | Device MAC addresses not yet completed in asynchronous license claim processing |
| `Processed` | `*int` | Optional | Number of devices processed so far by asynchronous license claim |
| `ScheduledAt` | `*int` | Optional | Epoch timestamp when the asynchronous license claim was scheduled |
| `Status` | [`*models.ResponseAsyncLicenseStatusEnum`](../../doc/models/response-async-license-status-enum.md) | Optional | Processing state for an asynchronous license claim. enum: `prepared`, `ongoing`, `done` |
| `Succeed` | `*int` | Optional | Number of devices that successfully completed license claim processing |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Total` | `*int` | Optional | Number of devices included in the license claim |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAsyncLicense := models.ResponseAsyncLicense{
        Completed:            []string{
            "completed0",
            "completed1",
            "completed2",
        },
        Details:              []models.ResponseAsyncLicenseDetail{
            models.ResponseAsyncLicenseDetail{
                Mac:                  models.ToPointer("mac4"),
                Status:               models.ToPointer("status2"),
                Timestamp:            models.ToPointer(float64(235.48)),
            },
        },
        Failed:               models.ToPointer(166),
        Incompleted:          []string{
            "incompleted0",
            "incompleted1",
            "incompleted2",
        },
        Processed:            models.ToPointer(72),
    }

}
```

