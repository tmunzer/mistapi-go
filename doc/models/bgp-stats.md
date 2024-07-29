
# Bgp Stats

## Structure

`BgpStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EvpnOverlay` | `*bool` | Optional | if this is created for evpn overlay |
| `ForOverlay` | `*bool` | Optional | if this is created for overlay |
| `LocalAs` | `*int` | Optional | AS |
| `Mac` | `*string` | Optional | router mac address |
| `Neighbor` | `*string` | Optional | - |
| `NeighborAs` | `*int` | Optional | - |
| `NeighborMac` | `*string` | Optional | if it's another device in the same org |
| `Node` | `*string` | Optional | node0/node1 |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `RxRoutes` | `*int` | Optional | number of received routes |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `State` | [`*models.BgpStatsStateEnum`](../../doc/models/bgp-stats-state-enum.md) | Optional | enum: `active`, `connect`, `established`, `idle`, `open_config`, `open_sent` |
| `Timestamp` | `*float64` | Optional | - |
| `TxPkts` | `*int` | Optional | - |
| `TxRoutes` | `*int` | Optional | - |
| `Up` | `*bool` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `VrfName` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "local_as": 65000,
  "mac": "020001c04668",
  "neighbor": "15.8.3.5",
  "neighbor_as": 65000,
  "neighbor_mac": "020001c04600",
  "node": "node0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "rx_pkts": 63366,
  "rx_routes": 60,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "state": "established",
  "timestamp": 1666251056.07,
  "tx_pkts": 1735,
  "tx_routes": 60,
  "uptime": 31355,
  "vrf_name": "default",
  "evpn_overlay": false,
  "for_overlay": false
}
```

