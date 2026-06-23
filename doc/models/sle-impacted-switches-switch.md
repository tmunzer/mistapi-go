
# Sle Impacted Switches Switch

SLE impact row for a switch

## Structure

`SleImpactedSwitchesSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | Portion of the SLE total that was degraded for this switch |
| `Duration` | `*float64` | Optional | Observation time represented by this switch impact row |
| `Interface` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Name` | `*string` | Optional | Display name for the switch impact row |
| `SwitchMac` | `*string` | Optional | MAC address of the switch represented by this impacted row |
| `SwitchModel` | `*string` | Optional | Model of the switch represented by this impacted row |
| `SwitchVersion` | `*string` | Optional | Firmware version of the switch represented by this impacted row |
| `Total` | `*float64` | Optional | Overall SLE total measured for this switch impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedSwitchesSwitch := models.SleImpactedSwitchesSwitch{
        Degraded:             models.ToPointer(float64(192.6)),
        Duration:             models.ToPointer(float64(65.66)),
        Interface:            []string{
            "interface8",
            "interface9",
        },
        Name:                 models.ToPointer("name0"),
        SwitchMac:            models.ToPointer("switch_mac8"),
    }

}
```

