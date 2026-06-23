
# Response Auto Map Assignment Device

Per-device validation result for auto map assignment

## Structure

`ResponseAutoMapAssignmentDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `*string` | Optional | Provides the reason for the status if the AP is invalid |
| `Valid` | `*bool` | Optional | Indicates whether the device meets requirements for auto map assignment |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoMapAssignmentDevice := models.ResponseAutoMapAssignmentDevice{
        Reason:               models.ToPointer("reason8"),
        Valid:                models.ToPointer(false),
    }

}
```

