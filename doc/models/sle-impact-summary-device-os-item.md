
# Sle Impact Summary Device Os Item

SLE impact summary row for a client device OS

## Structure

`SleImpactSummaryDeviceOsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `float64` | Required | Portion of the SLE total that was degraded for this device OS |
| `DeviceOs` | `string` | Required | Client device OS represented by this impact row |
| `Duration` | `float64` | Required | Observation time represented by this device-OS impact row |
| `Name` | `string` | Required | Display name for the device-OS impact row<br><br>**Constraints**: *Minimum Length*: `1` |
| `Total` | `float64` | Required | Overall SLE total measured for this device-OS impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactSummaryDeviceOsItem := models.SleImpactSummaryDeviceOsItem{
        Degraded:             float64(101.5),
        DeviceOs:             "device_os0",
        Duration:             float64(230.56),
        Name:                 "name0",
        Total:                float64(129.5),
    }

}
```

