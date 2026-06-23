
# Wan Usages

WAN usage record returned by site WAN usage searches

## Structure

`WanUsages`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Local router MAC address for the WAN usage record |
| `PathType` | `*string` | Optional | WAN path type used by this usage record |
| `PathWeight` | `*int` | Optional | Configured weight for the WAN path |
| `PeerMac` | `*string` | Optional | Remote peer MAC address for the WAN path |
| `PeerPortId` | `*string` | Optional | Remote peer interface identifier for the WAN path |
| `Policy` | `*string` | Optional | WAN path policy that selected this path |
| `PortId` | `*string` | Optional | Local router interface identifier for the WAN path |
| `Tenant` | `*string` | Optional | Network tenant context for the WAN usage record |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wanUsages := models.WanUsages{
        Mac:                  models.ToPointer("5c5b35000001"),
        PathType:             models.ToPointer("vpn"),
        PathWeight:           models.ToPointer(10),
        PeerMac:              models.ToPointer("0200018c95e1"),
        PeerPortId:           models.ToPointer("ge-0/0/3"),
        Policy:               models.ToPointer("policy1"),
        PortId:               models.ToPointer("ge-0/0/0.0"),
        Tenant:               models.ToPointer("tenant1"),
    }

}
```

