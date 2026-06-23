
# Gateway Port Lte Auth Enum

if `wan_type`==`lte`. enum: `chap`, `none`, `pap`

## Enumeration

`GatewayPortLteAuthEnum`

## Fields

| Name |
|  --- |
| `chap` |
| `none` |
| `pap` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortLteAuth := models.GatewayPortLteAuthEnum_PAP

}
```

