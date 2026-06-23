
# Const Insight Metrics Property Interval

Interval definition for an insight metric

## Structure

`ConstInsightMetricsPropertyInterval`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `*int` | Optional | Sampling interval length, in seconds |
| `MaxAge` | `*int` | Optional | Maximum lookback age for this interval, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constInsightMetricsPropertyInterval := models.ConstInsightMetricsPropertyInterval{
        Interval:             models.ToPointer(158),
        MaxAge:               models.ToPointer(78),
    }

}
```

