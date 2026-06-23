
# Response Troubleshoot

Organization troubleshooting response for the requested time window

## Structure

`ResponseTroubleshoot`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp, in seconds, for the end of the troubleshooting window |
| `Results` | [`[]models.ResponseTroubleshootItem`](../../doc/models/response-troubleshoot-item.md) | Optional | Troubleshooting findings returned by the organization troubleshoot API |
| `Start` | `*int` | Optional | Epoch timestamp, in seconds, for the start of the troubleshooting window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseTroubleshoot := models.ResponseTroubleshoot{
        End:                  models.ToPointer(1655151856),
        Results:              []models.ResponseTroubleshootItem{
            models.ResponseTroubleshootItem{
                Category:             models.ToPointer("category4"),
                Reason:               models.ToPointer("reason8"),
                Recommendation:       models.ToPointer("recommendation8"),
                Text:                 models.ToPointer("text4"),
            },
        },
        Start:                models.ToPointer(1655065456),
    }

}
```

