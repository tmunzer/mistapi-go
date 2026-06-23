
# Response Insight Rogue Client

Paginated response for rogue client insight results

## Structure

`ResponseInsightRogueClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp for the end of the rogue client insight window |
| `Limit` | `int` | Required | Maximum number of rogue client observations returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of rogue client observations |
| `Results` | [`[]models.InsightRogueClient`](../../doc/models/insight-rogue-client.md) | Required | Rogue client insight observations returned by a query<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | Epoch timestamp for the start of the rogue client insight window |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseInsightRogueClient := models.ResponseInsightRogueClient{
        End:                  148,
        Limit:                22,
        Next:                 models.ToPointer("next4"),
        Results:              []models.InsightRogueClient{
            models.InsightRogueClient{
                Annotation:           "annotation0",
                ApMac:                "ap_mac8",
                AvgRssi:              float64(170.84),
                Band:                 "band8",
                Bssid:                "bssid0",
                ClientMac:            "client_mac4",
                NumAps:               140,
            },
        },
        Start:                106,
    }

}
```

