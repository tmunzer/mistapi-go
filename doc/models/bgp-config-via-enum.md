
# Bgp Config via Enum

enum: `lan`, `tunnel`, `vpn`, `wan`

## Enumeration

`BgpConfigViaEnum`

## Fields

| Name |
|  --- |
| `lan` |
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
    bgpConfigVia := models.BgpConfigViaEnum_LAN

}
```

