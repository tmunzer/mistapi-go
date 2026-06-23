
# Rrm Event

RRM event record for radio setting changes or RF conditions

## Structure

`RrmEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | MAC address of the AP associated with the RRM event |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `5-dedicated`, `5-selectable`, `6`, `6-dedicated`, `6-selectable` |
| `Bandwidth` | [`models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Required | channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `int` | Required | RF channel after the RRM event |
| `Event` | [`models.RrmEventTypeEnum`](../../doc/models/rrm-event-type-enum.md) | Required | enum: `interference-ap-co-channel`, `interference-ap-non-wifi`, `neighbor-ap-down`, `neighbor-ap-recovered`, `radar-detected`, `rrm-radar`, `scheduled-site_rrm`, `triggered-site_rrm` |
| `Power` | `int` | Required | Transmit power after the RRM event |
| `PreBandwidth` | [`models.RrmEventPreBandwidthEnum`](../../doc/models/rrm-event-pre-bandwidth-enum.md) | Required | (previously) channel width for the band , 0 means no previously available. enum: `0`, `20`, `40`, `80`, `160` |
| `PreChannel` | `int` | Required | RF channel before the RRM event; 0 means no previous value was available |
| `PrePower` | `float64` | Required | Transmit power before the RRM event; 0 means no previous value was available |
| `PreUsage` | `string` | Required | Radio usage band before the RRM event |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Usage` | `string` | Required | Radio usage band after the RRM event |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmEvent := models.RrmEvent{
        Ap:                   "5c5b350e0001",
        Band:                 models.Dot11BandEnum_ENUM5,
        Bandwidth:            models.Dot11BandwidthEnum_ENUM20,
        Channel:              74,
        Event:                models.RrmEventTypeEnum_INTERFERENCEAPCOCHANNEL,
        Power:                242,
        PreBandwidth:         models.RrmEventPreBandwidthEnum_ENUM0,
        PreChannel:           120,
        PrePower:             float64(141.28),
        PreUsage:             "pre_usage6",
        Timestamp:            float64(209.56),
        Usage:                "usage6",
    }

}
```

