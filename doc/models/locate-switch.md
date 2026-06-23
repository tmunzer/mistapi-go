
# Locate Switch

Request body for locating a switch or virtual chassis member by MAC address for a limited duration

*This model accepts additional fields of type interface{}.*

## Structure

`LocateSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Minutes the leds should keep flashing<br><br>**Default**: `5`<br><br>**Constraints**: `>= 1`, `<= 120` |
| `Mac` | `*string` | Optional | For virtual chassis, the MAC of the member |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    locateSwitch := models.LocateSwitch{
        Duration:             models.ToPointer(5),
        Mac:                  models.ToPointer("f01c2d4ff760"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

