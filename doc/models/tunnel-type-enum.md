
# Tunnel Type Enum

Tunnel category, either WAN tunnel or WxLAN tunnel. enum: `wan`, `wxtunnel`

## Enumeration

`TunnelTypeEnum`

## Fields

| Name |
|  --- |
| `wan` |
| `wxtunnel` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelType := models.TunnelTypeEnum_WAN

}
```

