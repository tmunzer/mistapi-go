
# Sle Impact Summary Device Type Item

SLE impact summary row for a client device type

## Structure

`SleImpactSummaryDeviceTypeItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `float64` | Required | Portion of the SLE total that was degraded for this device type |
| `DeviceType` | `string` | Required | Client device type represented by this impact row |
| `Duration` | `float64` | Required | Observation time represented by this device-type impact row |
| `Name` | `string` | Required | Display name for the device-type impact row<br><br>**Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | Overall SLE total measured for this device-type impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactSummaryDeviceTypeItem := models.SleImpactSummaryDeviceTypeItem{
        Degraded:             float64(51.88),
        DeviceType:           "device_type2",
        Duration:             float64(180.94),
        Name:                 "name8",
        Total:                float64(79.88),
    }

}
```

