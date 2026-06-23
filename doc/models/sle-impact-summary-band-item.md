
# Sle Impact Summary Band Item

SLE impact summary row for a radio band

## Structure

`SleImpactSummaryBandItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `string` | Required | Radio band represented by this impact row<br><br>**Constraints**: *Minimum Length*: `1` |
| `Degraded` | `float64` | Required | Portion of the SLE total that was degraded for this radio band |
| `Duration` | `float64` | Required | Observation time represented by this radio-band impact row |
| `Name` | `string` | Required | Display name for the radio-band impact row<br><br>**Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | Overall SLE total measured for this radio-band impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactSummaryBandItem := models.SleImpactSummaryBandItem{
        Band:                 "band4",
        Degraded:             float64(133.02),
        Duration:             float64(6.08),
        Name:                 "name2",
        Total:                float64(94.98),
    }

}
```

