
# Insight Rogue Client

Rogue client observation returned by site insights

## Structure

`InsightRogueClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Annotation` | `string` | Required | Rogue client annotation or classification |
| `ApMac` | `string` | Required | Reporting AP MAC address that observed the rogue client |
| `AvgRssi` | `float64` | Required | Average RSSI for the rogue client as heard by reporting APs |
| `Band` | `string` | Required | Radio band on which the rogue client was observed |
| `Bssid` | `string` | Required | Rogue BSSID associated with the client |
| `ClientMac` | `string` | Required | Observed client MAC address on the rogue BSSID |
| `NumAps` | `int` | Required | Number of APs that observed the rogue client |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    insightRogueClient := models.InsightRogueClient{
        Annotation:           "annotation2",
        ApMac:                "ap_mac6",
        AvgRssi:              float64(196.66),
        Band:                 "band6",
        Bssid:                "bssid8",
        ClientMac:            "client_mac2",
        NumAps:               162,
    }

}
```

