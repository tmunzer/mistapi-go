
# Gateway Wan Type Enum

enum: `dhcp`, `pppoe`, `static`

## Enumeration

`GatewayWanTypeEnum`

## Fields

| Name |
|  --- |
| `dhcp` |
| `pppoe` |
| `static` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayWanType := models.GatewayWanTypeEnum_DHCP

}
```

