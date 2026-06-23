
# Gateway Path Strategy Enum

enum: `ecmp`, `ordered`, `weighted`

## Enumeration

`GatewayPathStrategyEnum`

## Fields

| Name |
|  --- |
| `ecmp` |
| `ordered` |
| `weighted` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPathStrategy := models.GatewayPathStrategyEnum_ORDERED

}
```

