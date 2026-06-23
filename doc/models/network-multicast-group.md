
# Network Multicast Group

Multicast group rendezvous point mapping

## Structure

`NetworkMulticastGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RpIp` | `*string` | Optional | RP (rendezvous point) IP address |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkMulticastGroup := models.NetworkMulticastGroup{
        RpIp:                 models.ToPointer("rp_ip0"),
    }

}
```

