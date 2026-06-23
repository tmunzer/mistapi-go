
# Gateway Port Usage Enum

port usage name. enum: `ha_control`, `ha_data`, `lan`, `wan`

## Enumeration

`GatewayPortUsageEnum`

## Fields

| Name |
|  --- |
| `ha_control` |
| `ha_data` |
| `lan` |
| `wan` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortUsage := models.GatewayPortUsageEnum_LAN

}
```

