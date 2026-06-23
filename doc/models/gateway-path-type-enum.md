
# Gateway Path Type Enum

enum: `local`, `tunnel`, `vpn`, `wan`

## Enumeration

`GatewayPathTypeEnum`

## Fields

| Name |
|  --- |
| `local` |
| `tunnel` |
| `vpn` |
| `wan` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPathType := models.GatewayPathTypeEnum_VPN

}
```

