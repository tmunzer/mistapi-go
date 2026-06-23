
# Sle Impacted Interfaces Interface

SLE impact row for a switch interface

## Structure

`SleImpactedInterfacesInterface`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*float64` | Optional | Portion of the SLE total that was degraded for this interface |
| `Duration` | `*float64` | Optional | Observation time represented by this interface impact row |
| `InterfaceName` | `*string` | Optional | Name of the switch interface represented by this impacted row |
| `SwitchMac` | `*string` | Optional | MAC address of the switch associated with this interface |
| `SwitchName` | `*string` | Optional | Display name of the switch associated with this interface |
| `Total` | `*float64` | Optional | Overall SLE total measured for this interface impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedInterfacesInterface := models.SleImpactedInterfacesInterface{
        Degraded:             models.ToPointer(float64(246.44)),
        Duration:             models.ToPointer(float64(119.5)),
        InterfaceName:        models.ToPointer("interface_name4"),
        SwitchMac:            models.ToPointer("switch_mac2"),
        SwitchName:           models.ToPointer("switch_name6"),
    }

}
```

