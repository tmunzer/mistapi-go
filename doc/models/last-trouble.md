
# Last Trouble

Last trouble indicator reported by a switch

## Structure

`LastTrouble`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | Trouble code; definitions are listed at [List Ap Led Definition](../../doc/controllers/constants-definitions.md#list-ap-led-definition) |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    lastTrouble := models.LastTrouble{
        Code:                 models.ToPointer("03"),
        Timestamp:            models.ToPointer(float64(183.7)),
    }

}
```

