
# Flow Capture Protocol Enum

Flow capture protocol filter. enum: `tcp`, `udp`, `icmp`, `icmp6`

## Enumeration

`FlowCaptureProtocolEnum`

## Fields

| Name |
|  --- |
| `tcp` |
| `udp` |
| `icmp` |
| `icmp6` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    flowCaptureProtocol := models.FlowCaptureProtocolEnum_ICMP

}
```

