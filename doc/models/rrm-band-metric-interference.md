
# Rrm Band Metric Interference

Interference metrics observed for one channel

## Structure

`RrmBandMetricInterference`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Radar` | `*float64` | Optional | Interference value attributed to radar on the channel |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmBandMetricInterference := models.RrmBandMetricInterference{
        Radar:                models.ToPointer(float64(75.48)),
    }

}
```

