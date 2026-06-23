
# Vpn Path

VPN path settings used by an organization VPN

## Structure

`VpnPath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BfdProfile` | [`*models.VpnPathBfdProfileEnum`](../../doc/models/vpn-path-bfd-profile-enum.md) | Optional | BFD profile used for this VPN path. enum: `broadband`, `lte`<br><br>**Default**: `"broadband"` |
| `BfdUseTunnelMode` | `*bool` | Optional | If `type`==`mesh` and for SSR only, whether to use tunnel mode<br><br>**Default**: `false` |
| `Ip` | `*string` | Optional | Source IP address for this VPN path, if different from the WAN port IP |
| `PeerPaths` | [`map[string]models.VpnPathPeerPathsPeer`](../../doc/models/vpn-path-peer-paths-peer.md) | Optional | If `type`==`mesh`, property key is the peer interface name |
| `Pod` | `*int` | Optional | Grouping index used to place this VPN path into a pod<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1`, `<= 128` |
| `TrafficShaping` | [`*models.VpnPathTrafficShaping`](../../doc/models/vpn-path-traffic-shaping.md) | Optional | Traffic shaping settings for a VPN path |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vpnPath := models.VpnPath{
        BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_BROADBAND),
        BfdUseTunnelMode:     models.ToPointer(false),
        Ip:                   models.ToPointer("ip4"),
        PeerPaths:            map[string]models.VpnPathPeerPathsPeer{
            "key0": models.VpnPathPeerPathsPeer{
                Preference:           models.ToPointer(144),
            },
            "key1": models.VpnPathPeerPathsPeer{
                Preference:           models.ToPointer(144),
            },
            "key2": models.VpnPathPeerPathsPeer{
                Preference:           models.ToPointer(144),
            },
        },
        Pod:                  models.ToPointer(2),
    }

}
```

