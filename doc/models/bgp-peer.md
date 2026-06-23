
# Bgp Peer

Only present when `bgp_peers` in `fields` query parameter

## Structure

`BgpPeer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EvpnOverlay` | `*bool` | Optional | If this is created for evpn overlay |
| `ForOverlay` | `*bool` | Optional | If this is created for overlay |
| `LocalAs` | [`*models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Neighbor` | `*string` | Optional | IP address of the BGP neighbor |
| `NeighborAs` | [`*models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `NeighborMac` | `*string` | Optional | MAC address of the BGP neighbor when it is another device in the same organization |
| `Node` | `*string` | Optional | HA node reporting this BGP peer, such as node0 or node1 |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `RxRoutes` | `*int` | Optional | Number of received routes |
| `State` | [`*models.BgpStatsStateEnum`](../../doc/models/bgp-stats-state-enum.md) | Optional | enum: `active`, `connect`, `established`, `idle`, `open_config`, `open_sent` |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `TxRoutes` | `*int` | Optional | Number of routes advertised to this BGP neighbor |
| `Up` | `*bool` | Optional | Whether the BGP session is currently up |
| `Uptime` | `*int` | Optional | Number of seconds the BGP session has been up |
| `VrfName` | `*string` | Optional | VRF name associated with this BGP session |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    bgpPeer := models.BgpPeer{
        EvpnOverlay:          models.ToPointer(false),
        ForOverlay:           models.ToPointer(false),
        LocalAs:              models.ToPointer(models.BgpAsContainer.FromNumber(65000)),
        Neighbor:             models.ToPointer("15.8.3.5"),
        NeighborAs:           models.ToPointer(models.BgpAsContainer.FromNumber(65000)),
        NeighborMac:          models.ToPointer("020001c04600"),
        Node:                 models.ToPointer("node0"),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        RxRoutes:             models.ToPointer(60),
        State:                models.ToPointer(models.BgpStatsStateEnum_ESTABLISHED),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        TxRoutes:             models.ToPointer(60),
        Uptime:               models.ToPointer(31355),
        VrfName:              models.ToPointer("default"),
    }

}
```

