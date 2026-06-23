
# Stats Mxedge Port Stat

Mist Edge port traffic and link status statistics

## Structure

`StatsMxedgePortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `*bool` | Optional | Whether the port is operating in full-duplex mode |
| `Lacp` | [`*models.StatsMxedgePortStatLacp`](../../doc/models/stats-mxedge-port-stat-lacp.md) | Optional | LACP state and counters for a Mist Edge port |
| `LldpStats` | [`*models.StatsMxedgePortStatLldpStats`](../../doc/models/stats-mxedge-port-stat-lldp-stats.md) | Optional | LLDP neighbor information reported for a Mist Edge port |
| `Mac` | `*string` | Optional | Mist Edge port MAC address reported by Mist |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxErrors` | `*int` | Optional | Number of receive errors observed on the port |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `Sfp` | [`*models.StatsMxedgePortStatSfp`](../../doc/models/stats-mxedge-port-stat-sfp.md) | Optional | SFP transceiver details reported for a Mist Edge port |
| `Speed` | `*int` | Optional | Current link speed of the port, in Mbps |
| `State` | `*string` | Optional | Forwarding or operational state reported for the port |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxErrors` | `*int` | Optional | Number of transmit errors observed on the port |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Up` | `*bool` | Optional | Whether the Mist Edge port link is up |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgePortStat := models.StatsMxedgePortStat{
        FullDuplex:           models.ToPointer(false),
        Lacp:                 models.ToPointer(models.StatsMxedgePortStatLacp{
            MuxState:             models.ToPointer("mux_state0"),
            RxLacpdu:             models.ToPointer(100),
            RxState:              models.ToPointer("rx_state6"),
            TxLacpdu:             models.ToPointer(122),
        }),
        LldpStats:            models.ToPointer(models.StatsMxedgePortStatLldpStats{
            ChassisId:            models.ToPointer("chassis_id0"),
            MgmtAddr:             models.ToPointer("mgmt_addr8"),
            PortDesc:             models.ToPointer("port_desc4"),
            PortId:               models.ToPointer("port_id4"),
            SystemDesc:           models.ToPointer("system_desc8"),
        }),
        Mac:                  models.ToPointer("mac8"),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
    }

}
```

