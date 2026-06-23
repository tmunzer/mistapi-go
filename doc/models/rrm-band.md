
# Rrm Band

RRM proposed and current radio settings for an AP band

## Structure

`RrmBand`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `*int` | Optional | Proposed RF channel for the radio band |
| `CurrBandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `CurrChannel` | `*int` | Optional | Current RF channel for the radio band |
| `CurrPower` | `*int` | Optional | Current transmit power for the radio band |
| `CurrUsage` | `*string` | Optional | Current radio usage band for the AP radio<br><br>**Constraints**: *Minimum Length*: `1` |
| `Power` | `*int` | Optional | Proposed transmit power for the radio band |
| `Usage` | `*string` | Optional | Proposed radio usage band for the AP radio<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmBand := models.RrmBand{
        Bandwidth:            models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
        Channel:              models.ToPointer(210),
        CurrBandwidth:        models.ToPointer(models.Dot11BandwidthEnum_ENUM20),
        CurrChannel:          models.ToPointer(74),
        CurrPower:            models.ToPointer(242),
    }

}
```

