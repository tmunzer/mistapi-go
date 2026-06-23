
# Module Stat Item Poe

Power over Ethernet telemetry for a device module

## Structure

`ModuleStatItemPoe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxPower` | `*float64` | Optional | Total PoE power budget available to the module, in watts |
| `PowerDraw` | `*float64` | Optional | Current PoE power draw on the module, in watts |
| `Status` | `*string` | Optional | Operational status of PoE on the module |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemPoe := models.ModuleStatItemPoe{
        MaxPower:             models.ToPointer(float64(250)),
        PowerDraw:            models.ToPointer(float64(120.3)),
        Status:               models.ToPointer("status4"),
    }

}
```

