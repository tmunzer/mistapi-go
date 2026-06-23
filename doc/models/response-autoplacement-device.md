
# Response Autoplacement Device

Per-AP validation result for autoplacement

## Structure

`ResponseAutoplacementDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `*string` | Optional, Read-only | Provides the reason for the status if the AP is invalid. |
| `Valid` | `*bool` | Optional, Read-only | Indicates whether the ap is valid. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoplacementDevice := models.ResponseAutoplacementDevice{
        Reason:               models.ToPointer("reason6"),
        Valid:                models.ToPointer(false),
    }

}
```

