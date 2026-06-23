
# Response Auto Orientation Device

Per-device validation result for auto orientation

## Structure

`ResponseAutoOrientationDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `*string` | Optional | Provides the reason for the status if the AP is invalid. |
| `Valid` | `*bool` | Optional | Indicates whether the auto orient request is valid for the device. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoOrientationDevice := models.ResponseAutoOrientationDevice{
        Reason:               models.ToPointer("reason2"),
        Valid:                models.ToPointer(false),
    }

}
```

