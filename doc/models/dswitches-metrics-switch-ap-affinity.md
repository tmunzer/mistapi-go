
# Dswitches Metrics Switch Ap Affinity

Switch/AP affinity metric for discovered switches

## Structure

`DswitchesMetricsSwitchApAffinity`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | [`models.DswitchesMetricsSwitchApAffinityDetails`](../../doc/models/dswitches-metrics-switch-ap-affinity-details.md) | Required | Detail values used by the switch/AP affinity metric |
| `Score` | `float64` | Required | Compliance score for the switch/AP affinity metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesMetricsSwitchApAffinity := models.DswitchesMetricsSwitchApAffinity{
        Details:              models.DswitchesMetricsSwitchApAffinityDetails{
            SystemName:           []string{
                "system_name0",
                "system_name1",
            },
            Threshold:            float64(56.78),
        },
        Score:                float64(8.72),
    }

}
```

