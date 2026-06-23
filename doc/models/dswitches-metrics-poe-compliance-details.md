
# Dswitches Metrics Poe Compliance Details

Detail values used by the discovered-switch PoE compliance metric

## Structure

`DswitchesMetricsPoeComplianceDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalAps` | `int` | Required | Number of APs evaluated for PoE compliance |
| `TotalPower` | `float64` | Required | Aggregate AP PoE power demand across evaluated APs, in mW |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesMetricsPoeComplianceDetails := models.DswitchesMetricsPoeComplianceDetails{
        TotalAps:             118,
        TotalPower:           float64(60.6),
    }

}
```

