
# Ap Client Bridge Auth Type Enum

wpa2-AES/CCMPp is assumed when `type`==`psk`. enum: `open`, `psk`

## Enumeration

`ApClientBridgeAuthTypeEnum`

## Fields

| Name |
|  --- |
| `open` |
| `psk` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apClientBridgeAuthType := models.ApClientBridgeAuthTypeEnum_OPEN

}
```

