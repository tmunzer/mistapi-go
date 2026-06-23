
# Ap Stat Mesh Downlink

Runtime statistics for a mesh downlink from this AP

## Structure

`ApStatMeshDownlink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Radio band used by this mesh downlink |
| `Channel` | `*int` | Optional | Radio channel used by this mesh downlink |
| `IdleTime` | `*int` | Optional | Seconds since traffic was last observed on this mesh downlink |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Proto` | `*string` | Optional | 802.11 protocol reported for this mesh downlink |
| `Rssi` | `*int` | Optional | Received signal strength for this mesh downlink, in dBm |
| `RxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of receiving traffic, bits/seconds, last known |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPackets` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `RxRate` | `models.Optional[float64]` | Optional, Read-only | Receive data rate reported for a wireless or mesh link, in Mbps |
| `RxRetries` | `models.Optional[int]` | Optional, Read-only | Amount of rx retries |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Snr` | `*int` | Optional | Signal-to-noise ratio for this mesh downlink, in dB |
| `TxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of transmitting traffic, bits/seconds, last known |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPackets` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `TxRate` | `models.Optional[float64]` | Optional, Read-only | Transmit data rate reported for a wireless or mesh link, in Mbps |
| `TxRetries` | `models.Optional[int]` | Optional, Read-only | Amount of tx retries |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    apStatMeshDownlink := models.ApStatMeshDownlink{
        Band:                 models.ToPointer("5"),
        Channel:              models.ToPointer(36),
        IdleTime:             models.ToPointer(3),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Proto:                models.ToPointer("n"),
        Rssi:                 models.ToPointer(-65),
        RxBps:                models.NewOptional(models.ToPointer(int64(60003))),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPackets:            models.NewOptional(models.ToPointer(int64(57770567))),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Snr:                  models.ToPointer(31),
        TxBps:                models.NewOptional(models.ToPointer(int64(634301))),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPackets:            models.NewOptional(models.ToPointer(int64(812204062))),
    }

}
```

