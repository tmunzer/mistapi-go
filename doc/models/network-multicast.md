
# Network Multicast

Whether to enable multicast support (only PIM-sparse mode is supported)

## Structure

`NetworkMulticast`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableIgmp` | `*bool` | Optional | If the network will only be the source of the multicast traffic, IGMP can be disabled<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | Whether multicast support is enabled for this network<br><br>**Default**: `false` |
| `Groups` | [`map[string]models.NetworkMulticastGroup`](../../doc/models/network-multicast-group.md) | Optional | Group address to RP (rendezvous point) mapping. Property Key is the CIDR (example "225.1.0.3/32") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkMulticast := models.NetworkMulticast{
        DisableIgmp:          models.ToPointer(false),
        Enabled:              models.ToPointer(false),
        Groups:               map[string]models.NetworkMulticastGroup{
            "key0": models.NetworkMulticastGroup{
                RpIp:                 models.ToPointer("rp_ip4"),
            },
        },
    }

}
```

