
# Sle Impact Summary Wlan Item

SLE impact summary row for a WLAN

## Structure

`SleImpactSummaryWlanItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `float64` | Required | Portion of the SLE total that was degraded for this WLAN |
| `Duration` | `float64` | Required | Observation time represented by this WLAN impact row |
| `Name` | `string` | Required | Display name for the WLAN impact row<br><br>**Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | Overall SLE total measured for this WLAN impact row |
| `WlanId` | `string` | Required | Identifier of the WLAN represented by this impact row<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactSummaryWlanItem := models.SleImpactSummaryWlanItem{
        Degraded:             float64(133.6),
        Duration:             float64(6.66),
        Name:                 "name0",
        Total:                float64(94.4),
        WlanId:               "wlan_id2",
    }

}
```

