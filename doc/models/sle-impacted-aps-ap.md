
# Sle Impacted Aps Ap

SLE impact row for an AP

## Structure

`SleImpactedApsAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | MAC address of the AP represented by this impacted row<br><br>**Constraints**: *Minimum Length*: `1` |
| `Degraded` | `float64` | Required | Portion of the SLE total that was degraded for this AP |
| `Duration` | `float64` | Required | Observation time represented by this AP impact row |
| `Name` | `string` | Required | Display name for the AP impact row<br><br>**Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | Overall SLE total measured for this AP impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedApsAp := models.SleImpactedApsAp{
        ApMac:                "ap_mac6",
        Degraded:             float64(215.44),
        Duration:             float64(88.5),
        Name:                 "name4",
        Total:                float64(12.56),
    }

}
```

