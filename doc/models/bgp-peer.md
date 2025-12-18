
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
| `Neighbor` | `*string` | Optional | - |
| `NeighborAs` | [`*models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `NeighborMac` | `*string` | Optional | If it's another device in the same org |
| `Node` | `*string` | Optional | Node0/node1 |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `RxRoutes` | `*int` | Optional | Number of received routes |
| `State` | [`*models.BgpStatsStateEnum`](../../doc/models/bgp-stats-state-enum.md) | Optional | enum: `active`, `connect`, `established`, `idle`, `open_config`, `open_sent` |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `TxRoutes` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `VrfName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "local_as": 65000,
  "neighbor": "15.8.3.5",
  "neighbor_as": 65000,
  "neighbor_mac": "020001c04600",
  "node": "node0",
  "rx_pkts": 57770567,
  "rx_routes": 60,
  "state": "established",
  "tx_pkts": 812204062,
  "tx_routes": 60,
  "uptime": 31355,
  "vrf_name": "default",
  "evpn_overlay": false,
  "for_overlay": false
}
```

