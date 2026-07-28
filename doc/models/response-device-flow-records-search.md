
# Response Device Flow Records Search

Paginated response for device flow record search results

## Structure

`ResponseDeviceFlowRecordsSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp, in seconds, for the end of the flow record search window |
| `Limit` | `*int` | Optional | Maximum number of flow records returned in this page |
| `Results` | [`[]models.FlowRecord`](../../doc/models/flow-record.md) | Optional | Flow records matching the search filters |
| `SearchAfter` | `*string` | Optional | Cursor token for retrieving the next page of flow records |
| `Start` | `*int` | Optional | Epoch timestamp, in seconds, for the start of the flow record search window |
| `Total` | `*int` | Optional | Number of flow records matching the search filters |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseDeviceFlowRecordsSearch := models.ResponseDeviceFlowRecordsSearch{
        End:                  models.ToPointer(220),
        Limit:                models.ToPointer(50),
        Results:              []models.FlowRecord{
            models.FlowRecord{
                DeviceMac:            models.ToPointer("device_mac0"),
                Direction:            models.ToPointer(models.FlowRecordDirectionEnum_EGRESS),
                DstIp:                models.ToPointer("dst_ip4"),
                DstPort:              models.ToPointer(240),
                Duration:             models.ToPointer(int64(202)),
            },
        },
        SearchAfter:          models.ToPointer("search_after4"),
        Start:                models.ToPointer(178),
    }

}
```

