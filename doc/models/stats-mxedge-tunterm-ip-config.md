
# Stats Mxedge Tunterm Ip Config

Tunnel termination IP configuration reported by a Mist Edge

## Structure

`StatsMxedgeTuntermIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `*string` | Optional | IPv4 default gateway for tunnel termination traffic |
| `Ip` | `*string` | Optional | Tunnel termination IPv4 address used by the Mist Edge |
| `Netmask` | `*string` | Optional | IPv4 netmask for the tunnel termination address |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeTuntermIpConfig := models.StatsMxedgeTuntermIpConfig{
        Gateway:              models.ToPointer("192.168.11.1"),
        Ip:                   models.ToPointer("192.168.11.91"),
        Netmask:              models.ToPointer("255.255.255.0"),
    }

}
```

