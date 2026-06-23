
# Tunnel Config Probe Type Enum

Protocol used by the custom IPsec tunnel health probe. enum: `http`, `icmp`

## Enumeration

`TunnelConfigProbeTypeEnum`

## Fields

| Name |
|  --- |
| `http` |
| `icmp` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigProbeType := models.TunnelConfigProbeTypeEnum_HTTP

}
```

