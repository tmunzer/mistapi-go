
# Rogue Details

Rogue AP detail response

## Structure

`RogueDetails`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Manufacture` | `string` | Required | Vendor or manufacturer name reported for the rogue AP |
| `SeenAsClient` | `bool` | Required | Whether this rogue AP was also observed as a client |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rogueDetails := models.RogueDetails{
        Manufacture:          "manufacture2",
        SeenAsClient:         false,
    }

}
```

