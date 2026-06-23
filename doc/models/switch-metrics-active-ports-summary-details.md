
# Switch Metrics Active Ports Summary Details

Port counts used by the active-port switch metric

## Structure

`SwitchMetricsActivePortsSummaryDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ActivePortCount` | `*int` | Optional | Number of active ports observed across evaluated switches |
| `TotalPortCount` | `*int` | Optional | Total number of ports evaluated for the active-port metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMetricsActivePortsSummaryDetails := models.SwitchMetricsActivePortsSummaryDetails{
        ActivePortCount:      models.ToPointer(146),
        TotalPortCount:       models.ToPointer(204),
    }

}
```

