
# Response Async Claim Status

Async inventory claim job status

## Structure

`ResponseAsyncClaimStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClaimId` | `*uuid.UUID` | Optional | Unique identifier of the async claim job |
| `Completed` | `[]string` | Optional | Device MAC addresses that completed asynchronous license claim processing |
| `Details` | [`[]models.ResponseAsyncLicenseDetail`](../../doc/models/response-async-license-detail.md) | Optional | Per-device asynchronous license claim status details |
| `Failed` | `*int` | Optional | Number of devices that failed claim processing |
| `Incompleted` | `[]string` | Optional | Device MAC addresses not yet completed in asynchronous license claim processing |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Processed` | `*int` | Optional | Number of devices processed so far |
| `ScheduledAt` | `*int` | Optional | Epoch timestamp when the async claim was scheduled |
| `Status` | [`*models.ResponseAsyncLicenseStatusEnum`](../../doc/models/response-async-license-status-enum.md) | Optional | Processing state for an asynchronous license claim. enum: `prepared`, `ongoing`, `done` |
| `Succeed` | `*int` | Optional | Number of devices that successfully completed claim processing |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Total` | `*int` | Optional | Total number of devices included in the claim |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseAsyncClaimStatus := models.ResponseAsyncClaimStatus{
        ClaimId:              models.ToPointer(uuid.MustParse("00002674-0000-0000-0000-000000000000")),
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
        Failed:               models.ToPointer(76),
        Incompleted:          []string{
            "incompleted0",
            "incompleted1",
            "incompleted2",
        },
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    }

}
```

