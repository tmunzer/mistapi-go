
# Ap Radio Band 6

6 GHz radio settings for an access point

## Structure

`ApRadioBand6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `*bool` | Optional | Whether RRM may disable the 6 GHz radio when optimizing RF settings<br><br>**Default**: `false` |
| `AntGain` | `models.Optional[int]` | Optional | External antenna gain for the 6 GHz radio<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `AntennaBeamPattern` | [`*models.RadioBandAntennaBeamPatternEnum`](../../doc/models/radio-band-antenna-beam-pattern-enum.md) | Optional | enum: `narrow`, `medium`, `wide` |
| `AntennaMode` | [`*models.RadioBandAntennaModeEnum`](../../doc/models/radio-band-antenna-mode-enum.md) | Optional | enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`<br><br>**Default**: `"default"` |
| `Bandwidth` | [`*models.Dot11Bandwidth6Enum`](../../doc/models/dot-11-bandwidth-6-enum.md) | Optional | channel width for the 6GHz band. enum: `0`(disabled, response only), `20`, `40`, `80`, `160`<br><br>**Default**: `80` |
| `Channel` | `models.Optional[int]` | Optional | For Device. (primary) channel for the band, 0 means using the Site Setting |
| `Channels` | `models.Optional[[]int]` | Optional | For RFTemplates. List of channels, null or empty array means auto |
| `Disabled` | `*bool` | Optional | Whether to disable the radio<br><br>**Default**: `false` |
| `Power` | `models.Optional[int]` | Optional | Radio Tx power, in dBm. Can be an integer 0-25 for static power configuration, or `null` or unset for auto power mode<br><br>**Constraints**: `>= 0`, `<= 25` |
| `PowerMax` | `models.Optional[int]` | Optional | When power=null/unset, max tx power to use, HW-specific values will be used if not set<br><br>**Constraints**: `>= 5`, `<= 18` |
| `PowerMin` | `models.Optional[int]` | Optional | When power=null/unset, min tx power to use, HW-specific values will be used if not set<br><br>**Constraints**: `>= 5`, `<= 18` |
| `Preamble` | [`*models.RadioBandPreambleEnum`](../../doc/models/radio-band-preamble-enum.md) | Optional | enum: `auto`, `long`, `short`<br><br>**Default**: `"short"` |
| `StandardPower` | `*bool` | Optional | For 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed, and we'll fall back to Low Power Indoor if AFC failed<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apRadioBand6 := models.ApRadioBand6{
        AllowRrmDisable:      models.ToPointer(false),
        AntGain:              models.NewOptional(models.ToPointer(0)),
        AntennaBeamPattern:   models.ToPointer(models.RadioBandAntennaBeamPatternEnum_NARROW),
        AntennaMode:          models.ToPointer(models.RadioBandAntennaModeEnum_ENUMDEFAULT),
        Bandwidth:            models.ToPointer(models.Dot11Bandwidth6Enum_ENUM80),
        Channel:              models.NewOptional(models.ToPointer(0)),
        Channels:             models.NewOptional[int](nil),
        Disabled:             models.ToPointer(false),
        Power:                models.NewOptional(models.ToPointer(7)),
        Preamble:             models.ToPointer(models.RadioBandPreambleEnum_SHORT),
        StandardPower:        models.ToPointer(false),
    }

}
```

