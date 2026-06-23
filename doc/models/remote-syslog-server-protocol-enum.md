
# Remote Syslog Server Protocol Enum

Transport protocol used for this remote syslog server. enum: `tcp`, `udp`

## Enumeration

`RemoteSyslogServerProtocolEnum`

## Fields

| Name |
|  --- |
| `tcp` |
| `udp` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    remoteSyslogServerProtocol := models.RemoteSyslogServerProtocolEnum_TCP

}
```

