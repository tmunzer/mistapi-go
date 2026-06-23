
# Tunnel Config Protocol Enum

Only if `provider`==`custom-ipsec`. enum: `gre`, `ipsec`

## Enumeration

`TunnelConfigProtocolEnum`

## Fields

| Name |
|  --- |
| `gre` |
| `ipsec` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigProtocol := models.TunnelConfigProtocolEnum_GRE

}
```

