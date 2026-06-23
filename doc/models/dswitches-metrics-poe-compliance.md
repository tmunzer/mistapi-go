
# Dswitches Metrics Poe Compliance

PoE compliance metric for APs connected to discovered switches

## Structure

`DswitchesMetricsPoeCompliance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`models.DswitchesMetricsPoeComplianceDetails`](../../doc/models/dswitches-metrics-poe-compliance-details.md) | Required | Detail values used by the discovered-switch PoE compliance metric |
| `Score` | `float64` | Required | Compliance score for the PoE compliance metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesMetricsPoeCompliance := models.DswitchesMetricsPoeCompliance{
        Details:              models.DswitchesMetricsPoeComplianceDetails{
            TotalAps:             132,
            TotalPower:           float64(137.54),
        },
        Score:                float64(19.76),
    }

}
```

