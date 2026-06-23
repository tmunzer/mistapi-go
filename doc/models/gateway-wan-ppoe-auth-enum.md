
# Gateway Wan Ppoe Auth Enum

if `type`==`pppoe`. enum: `chap`, `none`, `pap`

## Enumeration

`GatewayWanPpoeAuthEnum`

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
    gatewayWanPpoeAuth := models.GatewayWanPpoeAuthEnum_CHAP

}
```

