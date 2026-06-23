
# Tunnel Config Dh Group Enum

Only if `provider`==`custom-ipsec`. Diffie-Hellman group for IPsec phase 2. enum: `1`, `14`, `15`, `16`, `19`, `2`, `20`, `21`, `24`, `5`. `14` is the default 2048-bit group; `19`, `20`, `21`, and `24` are ECP groups

## Enumeration

`TunnelConfigDhGroupEnum`

## Fields

| Name |
|  --- |
| `1` |
| `14` |
| `15` |
| `16` |
| `19` |
| `2` |
| `20` |
| `21` |
| `24` |
| `5` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigDhGroup := models.TunnelConfigDhGroupEnum_ENUM1

}
```

