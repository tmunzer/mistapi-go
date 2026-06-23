
# Sle Impacted Chassis Chassis Item

SLE impact row for a chassis

## Structure

`SleImpactedChassisChassisItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Chassis` | `*string` | Optional | Identifier of the chassis represented by this impacted row |
| `Degraded` | `*float64` | Optional | Portion of the SLE total that was degraded for this chassis |
| `Duration` | `*float64` | Optional | Observation time represented by this chassis impact row |
| `Role` | `*string` | Optional | Virtual Chassis role for the switch member |
| `SwitchMac` | `*string` | Optional | MAC address of the switch represented by this chassis row |
| `SwitchName` | `*string` | Optional | Display name of the switch represented by this chassis row |
| `Total` | `*float64` | Optional | Overall SLE total measured for this chassis impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedChassisChassisItem := models.SleImpactedChassisChassisItem{
        Chassis:              models.ToPointer("chassis8"),
        Degraded:             models.ToPointer(float64(156.04)),
        Duration:             models.ToPointer(float64(29.1)),
        Role:                 models.ToPointer("role8"),
        SwitchMac:            models.ToPointer("switch_mac2"),
    }

}
```

