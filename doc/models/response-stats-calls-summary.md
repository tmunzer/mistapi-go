
# Response Stats Calls Summary

Aggregated site call statistics summary

## Structure

`ResponseStatsCallsSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BadMinutes` | `*float64` | Optional | Total call minutes classified as bad across all quality categories |
| `BadMinutesClient` | `*float64` | Optional | Call minutes classified as bad due to client-side issues |
| `BadMinutesSiteWan` | `*float64` | Optional | Call minutes classified as bad due to site WAN issues |
| `BadMinutesWireless` | `*float64` | Optional | Call minutes classified as bad due to wireless issues |
| `NumAps` | `*int` | Optional | Number of APs represented in the call statistics summary |
| `NumUsers` | `*int` | Optional | Number of users represented in the call statistics summary |
| `TotalMinutes` | `*float64` | Optional | Total call minutes represented in the summary |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseStatsCallsSummary := models.ResponseStatsCallsSummary{
        BadMinutes:           models.ToPointer(float64(5566)),
        BadMinutesClient:     models.ToPointer(float64(526)),
        BadMinutesSiteWan:    models.ToPointer(float64(3612)),
        BadMinutesWireless:   models.ToPointer(float64(1428)),
        NumAps:               models.ToPointer(1),
        NumUsers:             models.ToPointer(3),
        TotalMinutes:         models.ToPointer(float64(5566)),
    }

}
```

