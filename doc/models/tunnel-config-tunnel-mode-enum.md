
# Tunnel Config Tunnel Mode Enum

Required if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`. enum: `active-active`, `active-standby`

## Enumeration

`TunnelConfigTunnelModeEnum`

## Fields

| Name |
|  --- |
| `active-active` |
| `active-standby` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigTunnelMode := models.TunnelConfigTunnelModeEnum_ACTIVEACTIVE

}
```

