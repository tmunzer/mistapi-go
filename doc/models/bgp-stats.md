
# Bgp Stats

BGP peer statistics reported by a network device

## Structure

`BgpStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EvpnOverlay` | `*bool` | Optional | If this is created for evpn overlay |
| `ForOverlay` | `*bool` | Optional | If this is created for overlay |
| `LocalAs` | [`*models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Mac` | `*string` | Optional | Router MAC address for the device reporting this BGP peer |
| `Model` | `*string` | Optional | Device model reporting this BGP peer statistic |
| `Neighbor` | `*string` | Optional | IP address of the BGP neighbor |
| `NeighborAs` | [`*models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `NeighborMac` | `*string` | Optional | MAC address of the BGP neighbor when it is another device in the same organization |
| `Node` | `*string` | Optional | HA node reporting this BGP peer, such as node0 or node1 |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `RxRoutes` | `*int` | Optional | Number of received routes |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
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
    "github.com/google/uuid"
)

func main() {
    bgpStats := models.BgpStats{
        EvpnOverlay:          models.ToPointer(false),
        ForOverlay:           models.ToPointer(false),
        LocalAs:              models.ToPointer(models.BgpAsContainer.FromNumber(65000)),
        Mac:                  models.ToPointer("020001c04668"),
        Model:                models.ToPointer("model6"),
        Neighbor:             models.ToPointer("15.8.3.5"),
        NeighborAs:           models.ToPointer(),
        NeighborMac:          models.ToPointer("020001c04600"),
        Node:                 models.ToPointer("node0"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        RxRoutes:             models.ToPointer(60),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        State:                models.ToPointer(models.BgpStatsStateEnum_ESTABLISHED),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        TxRoutes:             models.ToPointer(60),
        Uptime:               models.ToPointer(31355),
        VrfName:              models.ToPointer("default"),
    }

}
```

