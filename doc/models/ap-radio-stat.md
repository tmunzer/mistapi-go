
# Ap Radio Stat

Runtime radio statistics for an access point radio

## Structure

`ApRadioStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `models.Optional[int]` | Optional, Read-only | Current channel the radio is running on |
| `DynamicChainingEnabled` | `models.Optional[bool]` | Optional, Read-only | Use dynamic chaining for downlink |
| `Mac` | `models.Optional[string]` | Optional, Read-only | Radio base MAC address; a base radio MAC can represent up to 16 BSSIDs (e.g. 5c5b350001a0-5c5b350001af) |
| `NoiseFloor` | `models.Optional[int]` | Optional, Read-only | Measured noise floor for this radio, in dBm |
| `NumClients` | `models.Optional[int]` | Optional, Read-only | Number of clients currently connected on this radio |
| `NumWlans` | `*int` | Optional | How many WLANs are applied to the radio |
| `Power` | `models.Optional[int]` | Optional, Read-only | Transmit power (in dBm) |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Usage` | `models.Optional[string]` | Optional, Read-only | Operating band reported for this radio, such as 24, 5, or 6 |
| `UtilAll` | `models.Optional[int]` | Optional, Read-only | All utilization in percentage |
| `UtilNonWifi` | `models.Optional[int]` | Optional, Read-only | Reception of "No Packets" utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise |
| `UtilRxInBss` | `models.Optional[int]` | Optional, Read-only | Reception of "In BSS" utilization in percentage, only frames that are received from AP/STAs within the BSS |
| `UtilRxOtherBss` | `models.Optional[int]` | Optional, Read-only | Reception of "Other BSS" utilization in percentage, all frames received from AP/STAs that are outside the BSS |
| `UtilTx` | `models.Optional[int]` | Optional, Read-only | Transmission utilization in percentage |
| `UtilUndecodableWifi` | `models.Optional[int]` | Optional, Read-only | Reception of "UnDecodable Wifi" utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio |
| `UtilUnknownWifi` | `models.Optional[int]` | Optional, Read-only | Reception of "No Category" utilization in percentage, all 802.11 frames that are corrupted at the receiver |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apRadioStat := models.ApRadioStat{
        Bandwidth:              models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
        Channel:                models.NewOptional(models.ToPointer(68)),
        DynamicChainingEnabled: models.NewOptional(models.ToPointer(false)),
        Mac:                    models.NewOptional(models.ToPointer("mac0")),
        NoiseFloor:             models.NewOptional(models.ToPointer(-90)),
        RxBytes:                models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:                 models.NewOptional(models.ToPointer(int64(57770567))),
        TxBytes:                models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:                 models.NewOptional(models.ToPointer(int64(812204062))),
        Usage:                  models.NewOptional(models.ToPointer("24")),
    }

}
```

