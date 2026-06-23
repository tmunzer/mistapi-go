
# Synthetictest Device Protocol Enum

If `type`==`lan_connectivity`. Protocol or protocol combination used by the LAN connectivity test. enum: `ping`, `traceroute`, `ping+traceroute`

## Enumeration

`SynthetictestDeviceProtocolEnum`

## Fields

| Name |
|  --- |
| `ping` |
| `ping+traceroute` |
| `traceroute` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestDeviceProtocol := models.SynthetictestDeviceProtocolEnum_PING

}
```

