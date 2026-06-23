
# Module Stat Item Psus Item

Power supply status for a device module

## Structure

`ModuleStatItemPsusItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Power supply label reported by the device |
| `Status` | `*string` | Optional | Operational status reported for the power supply |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemPsusItem := models.ModuleStatItemPsusItem{
        Name:                 models.ToPointer("Power Supply 0"),
        Status:               models.ToPointer("ok"),
    }

}
```

