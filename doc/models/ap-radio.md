
# Ap Radio

Radio configuration settings for an access point

## Structure

`ApRadio`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `*bool` | Optional | Whether RRM can be disabled for individual radio-band settings<br><br>**Default**: `false` |
| `AntGain24` | `*int` | Optional | Antenna gain for 2.4G - for models with external antenna only<br><br>**Constraints**: `>= 0` |
| `AntGain5` | `*int` | Optional | Antenna gain for 5G - for models with external antenna only<br><br>**Constraints**: `>= 0` |
| `AntGain6` | `*int` | Optional | Antenna gain for 6G - for models with external antenna only<br><br>**Constraints**: `>= 0` |
| `AntennaMode` | [`*models.ApRadioAntennaModeEnum`](../../doc/models/ap-radio-antenna-mode-enum.md) | Optional | enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`<br><br>**Default**: `"default"` |
| `AntennaSelect` | [`*models.AntennaSelectEnum`](../../doc/models/antenna-select-enum.md) | Optional | Antenna Mode for AP which supports selectable antennas. enum: `""` (default), `external`, `internal` |
| `Band24` | [`*models.ApRadioBand24`](../../doc/models/ap-radio-band-24.md) | Optional | 2.4 GHz radio settings for an access point |
| `Band24Usage` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | enum: `24`, `5`, `6`, `auto` |
| `Band5` | [`*models.ApRadioBand5`](../../doc/models/ap-radio-band-5.md) | Optional | 5 GHz radio settings for an access point |
| `Band5On24Radio` | [`*models.ApRadioBand5`](../../doc/models/ap-radio-band-5.md) | Optional | 5 GHz radio settings for an access point |
| `Band6` | [`*models.ApRadioBand6`](../../doc/models/ap-radio-band-6.md) | Optional | 6 GHz radio settings for an access point |
| `FullAutomaticRrm` | `*bool` | Optional | Let RRM control everything, only the `channels` and `ant_gain` will be honored (i.e. disabled/bandwidth/power/band_24_usage are all controlled by RRM)<br><br>**Default**: `false` |
| `IndoorUse` | `*bool` | Optional | To make an outdoor operate indoor. For an outdoor-ap, some channels are disallowed by default, this allows the user to use it as an indoor-ap<br><br>**Default**: `false` |
| `RrmManaged` | `*bool` | Optional | Enable RRM to manage all radio settings (ignores all band_xxx configs) |
| `ScanningEnabled` | `*bool` | Optional | Whether scanning radio is enabled |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apRadio := models.ApRadio{
        AllowRrmDisable:      models.ToPointer(false),
        AntGain24:            models.ToPointer(4),
        AntGain5:             models.ToPointer(5),
        AntGain6:             models.ToPointer(5),
        AntennaMode:          models.ToPointer(models.ApRadioAntennaModeEnum_ENUMDEFAULT),
        AntennaSelect:        models.ToPointer(models.AntennaSelectEnum_EXTERNAL),
        FullAutomaticRrm:     models.ToPointer(false),
        IndoorUse:            models.ToPointer(false),
        ScanningEnabled:      models.ToPointer(true),
    }

}
```

