
# Response Synthetictest Search

Paginated response for site synthetic test search results

## Structure

`ResponseSynthetictestSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp, in seconds, for the end of the synthetic test search window |
| `Limit` | `int` | Required | Maximum number of synthetic test results returned in this page |
| `Next` | `*string` | Optional | URL for retrieving the next page of synthetic test results |
| `Results` | [`[]models.SynthetictestInfo`](../../doc/models/synthetictest-info.md) | Required | Synthetic test result records returned by search |
| `Start` | `int` | Required | Epoch timestamp, in seconds, for the start of the synthetic test search window |
| `Total` | `int` | Required | Number of synthetic test results matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSynthetictestSearch := models.ResponseSynthetictestSearch{
        End:                  228,
        Limit:                198,
        Next:                 models.ToPointer("next2"),
        Results:              []models.SynthetictestInfo{
            models.SynthetictestInfo{
                By:                   models.ToPointer("user"),
                DeviceType:           models.ToPointer(models.DeviceTypeEnum_GATEWAY),
                Failed:               models.ToPointer(false),
                Latency:              models.ToPointer(40),
                Mac:                  models.ToPointer("mac0"),
                PortId:               models.ToPointer("ge-0/0/2"),
                Reason:               models.ToPointer("interface not ready to perform test"),
                RxMbps:               models.ToPointer(322),
                StartTime:            models.ToPointer(1675718807),
                TxMbps:               models.ToPointer(199),
                VlanId:               models.ToPointer(20),
            },
        },
        Start:                186,
        Total:                36,
    }

}
```

