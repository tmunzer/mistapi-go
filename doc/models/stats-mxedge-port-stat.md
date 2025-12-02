
# Stats Mxedge Port Stat

## Structure

`StatsMxedgePortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `*bool` | Optional | - |
| `Lacp` | [`*models.StatsMxedgePortStatLacp`](../../doc/models/stats-mxedge-port-stat-lacp.md) | Optional | - |
| `LldpStats` | [`*models.StatsMxedgePortStatLldpStats`](../../doc/models/stats-mxedge-port-stat-lldp-stats.md) | Optional | - |
| `Mac` | `*string` | Optional | - |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxErrors` | `*int` | Optional | - |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `Sfp` | [`*models.StatsMxedgePortStatSfp`](../../doc/models/stats-mxedge-port-stat-sfp.md) | Optional | - |
| `Speed` | `*int` | Optional | - |
| `State` | `*string` | Optional | - |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxErrors` | `*int` | Optional | - |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "full_duplex": false,
  "lacp": {
    "mux_state": "mux_state0",
    "rx_lacpdu": 100,
    "rx_state": "rx_state6",
    "tx_lacpdu": 122
  },
  "lldp_stats": {
    "chassis_id": "chassis_id0",
    "mgmt_addr": "mgmt_addr8",
    "port_desc": "port_desc4",
    "port_id": "port_id4",
    "system_desc": "system_desc8"
  },
  "mac": "mac4"
}
```

