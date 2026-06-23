
# Stats Ap Port Stat

Ethernet port statistics reported by an AP

## Structure

`StatsApPortStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullDuplex` | `models.Optional[bool]` | Optional, Read-only | Whether the AP Ethernet port is operating in full-duplex mode |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxErrors` | `models.Optional[int]` | Optional, Read-only | Receive error count reported for the AP Ethernet port |
| `RxPeakBps` | `models.Optional[int]` | Optional, Read-only | Peak receive throughput on the AP Ethernet port, in bits per second |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `Speed` | `models.Optional[int]` | Optional, Read-only | Negotiated Ethernet link speed for the AP port, in Mbps |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPeakBps` | `models.Optional[int]` | Optional, Read-only | Peak transmit throughput on the AP Ethernet port, in bits per second |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Up` | `models.Optional[bool]` | Optional, Read-only | Whether the AP Ethernet port link is up |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApPortStat := models.StatsApPortStat{
        FullDuplex:           models.NewOptional(models.ToPointer(true)),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxErrors:             models.NewOptional(models.ToPointer(0)),
        RxPeakBps:            models.NewOptional(models.ToPointer(22185)),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        Speed:                models.NewOptional(models.ToPointer(1000)),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPeakBps:            models.NewOptional(models.ToPointer(43922)),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        Up:                   models.NewOptional(models.ToPointer(true)),
    }

}
```

