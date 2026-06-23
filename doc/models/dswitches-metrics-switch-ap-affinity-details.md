
# Dswitches Metrics Switch Ap Affinity Details

Detail values used by the switch/AP affinity metric

## Structure

`DswitchesMetricsSwitchApAffinityDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SystemName` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Threshold` | `float64` | Required | APs-per-switch threshold used for the affinity metric |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dswitchesMetricsSwitchApAffinityDetails := models.DswitchesMetricsSwitchApAffinityDetails{
        SystemName:           []string{
            "system_name0",
        },
        Threshold:            float64(224.08),
    }

}
```

