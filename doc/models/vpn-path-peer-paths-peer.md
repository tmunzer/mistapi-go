
# Vpn Path Peer Paths Peer

Peer path preference settings for mesh VPN routing

## Structure

`VpnPathPeerPathsPeer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Preference` | `*int` | Optional | Lower numeric value makes this outgoing WAN path more preferred |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vpnPathPeerPathsPeer := models.VpnPathPeerPathsPeer{
        Preference:           models.ToPointer(10),
    }

}
```

