
# Const Insight Metrics Property Report Duration

Report duration definition for an insight metric

## Structure

`ConstInsightMetricsPropertyReportDuration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Report duration length, in seconds |
| `Interval` | `*int` | Optional | Sampling interval used for this report duration, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constInsightMetricsPropertyReportDuration := models.ConstInsightMetricsPropertyReportDuration{
        Duration:             models.ToPointer(32),
        Interval:             models.ToPointer(84),
    }

}
```

