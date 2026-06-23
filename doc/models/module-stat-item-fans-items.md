
# Module Stat Item Fans Items

Cooling fan telemetry for a device module

## Structure

`ModuleStatItemFansItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Airflow` | `*string` | Optional | Direction of airflow reported for the fan |
| `Name` | `*string` | Optional | Fan label reported by the device |
| `Rpm` | `*int` | Optional | Current fan speed in revolutions per minute |
| `Status` | `*string` | Optional | Operational status reported for the fan |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    moduleStatItemFansItems := models.ModuleStatItemFansItems{
        Airflow:              models.ToPointer("out"),
        Name:                 models.ToPointer("Fan 0"),
        Rpm:                  models.ToPointer(154),
        Status:               models.ToPointer("ok"),
    }

}
```

