
# Response Call Troubleshoot Summary

Site call troubleshooting summary response

## Structure

`ResponseCallTroubleshootSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Client MAC address for the summarized call troubleshooting result |
| `MeetingId` | `*uuid.UUID` | Optional | Meeting identifier for the summarized call troubleshooting result |
| `Results` | [`[]models.CallTroubleshootSummary`](../../doc/models/call-troubleshoot-summary.md) | Optional | List of call troubleshooting summary records |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseCallTroubleshootSummary := models.ResponseCallTroubleshootSummary{
        Mac:                  models.ToPointer("983a78ea4a44"),
        MeetingId:            models.ToPointer(uuid.MustParse("b784d744-9a7c-4fad-9af0-f78858a319b1")),
        Results:              []models.CallTroubleshootSummary{
            models.CallTroubleshootSummary{
                ApNumClients:           models.ToPointer(float64(15.88)),
                ApRtt:                  models.ToPointer(float64(252.92)),
                AudioIn:                models.ToPointer(models.CallTroubleshootSummaryData{
                    ApNumClients:           models.ToPointer(float64(152.32)),
                    ApRtt:                  models.ToPointer(float64(133.36)),
                    ClientCpu:              models.ToPointer(float64(164.78)),
                    ClientNStreams:         models.ToPointer(float64(206.36)),
                    ClientRadioBand:        models.ToPointer(float64(43.4)),
                }),
                AudioOut:               models.ToPointer(models.CallTroubleshootSummaryData{
                    ApNumClients:           models.ToPointer(float64(71.16)),
                    ApRtt:                  models.ToPointer(float64(52.2)),
                    ClientCpu:              models.ToPointer(float64(245.94)),
                    ClientNStreams:         models.ToPointer(float64(125.2)),
                    ClientRadioBand:        models.ToPointer(float64(218.24)),
                }),
                ClientCpu:              models.ToPointer(float64(45.22)),
            },
        },
    }

}
```

