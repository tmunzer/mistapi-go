
# Response Switch Metrics Active Ports Summary

Active-port summary metric for switches in the requested scope

## Structure

`ResponseSwitchMetricsActivePortsSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`*models.SwitchMetricsActivePortsSummaryDetails`](../../doc/models/switch-metrics-active-ports-summary-details.md) | Optional | Port counts used by the active-port switch metric |
| `Score` | `*int` | Optional | Reported metric score for active switch ports |
| `TotalSwitchCount` | `*int` | Optional | Number of switches included in the active-port summary metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSwitchMetricsActivePortsSummary := models.ResponseSwitchMetricsActivePortsSummary{
        Details:              models.ToPointer(models.SwitchMetricsActivePortsSummaryDetails{
            ActivePortCount:      models.ToPointer(136),
            TotalPortCount:       models.ToPointer(42),
        }),
        Score:                models.ToPointer(114),
        TotalSwitchCount:     models.ToPointer(158),
    }

}
```

