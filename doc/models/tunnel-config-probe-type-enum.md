
# Tunnel Config Probe Type Enum

Protocol used by the custom IPsec tunnel health probe. `http` is deprecated — use `probe_ips`/`probe_hostnames` for ICMP probes and `probe_http` for HTTP probes instead. enum: `http`, `icmp`

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

