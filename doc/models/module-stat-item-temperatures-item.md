
# Module Stat Item Temperatures Item

Temperature sensor reading for a device module

## Structure

`ModuleStatItemTemperaturesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Celsius` | `*float64` | Optional | Temperature reading for the sensor, in degrees Celsius |
| `Name` | `*string` | Optional | Temperature sensor label reported by the device |
| `Status` | `*string` | Optional | Operational status reported for the temperature sensor |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemTemperaturesItem := models.ModuleStatItemTemperaturesItem{
        Celsius:              models.ToPointer(float64(45)),
        Name:                 models.ToPointer("CPU"),
        Status:               models.ToPointer("ok"),
    }

}
```

