
# Tunnel Config Ike Dh Group Enum

Diffie-Hellman group for IKE phase 1. enum: `1`, `14`, `15`, `16`, `19`, `2`, `20`, `21`, `24`, `5`. `14` is the default 2048-bit group; `19`, `20`, `21`, and `24` are ECP groups

## Enumeration

`TunnelConfigIkeDhGroupEnum`

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
    tunnelConfigIkeDhGroup := models.TunnelConfigIkeDhGroupEnum_ENUM15

}
```

