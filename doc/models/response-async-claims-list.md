
# Response Async Claims List

List of async inventory claim jobs for the organization

## Structure

`ResponseAsyncClaimsList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Claims` | [`[]models.ResponseAsyncClaimStatus`](../../doc/models/response-async-claim-status.md) | Optional | Async claim job status records |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseAsyncClaimsList := models.ResponseAsyncClaimsList{
        Claims:               []models.ResponseAsyncClaimStatus{
            models.ResponseAsyncClaimStatus{
                ClaimId:              models.ToPointer(uuid.MustParse("00000d5e-0000-0000-0000-000000000000")),
                Completed:            []string{
                    "completed8",
                    "completed7",
                },
                Details:              []models.ResponseAsyncLicenseDetail{
                    models.ResponseAsyncLicenseDetail{
                        Mac:                  models.ToPointer("mac4"),
                        Status:               models.ToPointer("status2"),
                        Timestamp:            models.ToPointer(float64(235.48)),
                    },
                    models.ResponseAsyncLicenseDetail{
                        Mac:                  models.ToPointer("mac4"),
                        Status:               models.ToPointer("status2"),
                        Timestamp:            models.ToPointer(float64(235.48)),
                    },
                },
                Failed:               models.ToPointer(50),
                Incompleted:          []string{
                    "incompleted8",
                    "incompleted7",
                },
            },
            models.ResponseAsyncClaimStatus{
                ClaimId:              models.ToPointer(uuid.MustParse("00000d5e-0000-0000-0000-000000000000")),
                Completed:            []string{
                    "completed8",
                    "completed7",
                },
                Details:              []models.ResponseAsyncLicenseDetail{
                    models.ResponseAsyncLicenseDetail{
                        Mac:                  models.ToPointer("mac4"),
                        Status:               models.ToPointer("status2"),
                        Timestamp:            models.ToPointer(float64(235.48)),
                    },
                    models.ResponseAsyncLicenseDetail{
                        Mac:                  models.ToPointer("mac4"),
                        Status:               models.ToPointer("status2"),
                        Timestamp:            models.ToPointer(float64(235.48)),
                    },
                },
                Failed:               models.ToPointer(50),
                Incompleted:          []string{
                    "incompleted8",
                    "incompleted7",
                },
            },
        },
    }

}
```

