
# Gateway Port Wan Type Enum

Only if `usage`==`wan`. enum: `broadband`, `dsl`, `lte`

## Enumeration

`GatewayPortWanTypeEnum`

## Fields

| Name |
|  --- |
| `broadband` |
| `dsl` |
| `lte` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortWanType := models.GatewayPortWanTypeEnum_DSL

}
```

