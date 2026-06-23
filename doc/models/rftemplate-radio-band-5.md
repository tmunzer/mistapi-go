
# Rftemplate Radio Band 5

Radio Band AP settings

## Structure

`RftemplateRadioBand5`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `*bool` | Optional | Whether RRM may disable the 5 GHz radio when optimizing RF settings<br><br>**Default**: `false` |
| `AntGain` | `models.Optional[int]` | Optional | External antenna gain for the 5 GHz radio<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `AntennaMode` | [`*models.RadioBandAntennaModeEnum`](../../doc/models/radio-band-antenna-mode-enum.md) | Optional | enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`<br><br>**Default**: `"default"` |
| `Bandwidth` | [`*models.Dot11Bandwidth5Enum`](../../doc/models/dot-11-bandwidth-5-enum.md) | Optional | channel width for the 5GHz band. enum: `0`(disabled, response only), `20`, `40`, `80`<br><br>**Default**: `40` |
| `Channels` | `models.Optional[[]int]` | Optional | For RFTemplates. List of channels, null or empty array means auto |
| `Disabled` | `*bool` | Optional | Whether to disable the radio<br><br>**Default**: `false` |
| `Power` | `models.Optional[int]` | Optional | Radio Tx power, in dBm. Can be an integer 0-25 for static power configuration, or `null` or unset for auto power mode<br><br>**Constraints**: `>= 0`, `<= 25` |
| `PowerMax` | `models.Optional[int]` | Optional | When power=null/unset, max tx power to use, HW-specific values will be used if not set<br><br>**Constraints**: `>= 5`, `<= 17` |
| `PowerMin` | `models.Optional[int]` | Optional | When power=null/unset, min tx power to use, HW-specific values will be used if not set<br><br>**Constraints**: `>= 5`, `<= 17` |
| `Preamble` | [`*models.RadioBandPreambleEnum`](../../doc/models/radio-band-preamble-enum.md) | Optional | enum: `auto`, `long`, `short`<br><br>**Default**: `"short"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rftemplateRadioBand5 := models.RftemplateRadioBand5{
        AllowRrmDisable:      models.ToPointer(false),
        AntGain:              models.NewOptional(models.ToPointer(0)),
        AntennaMode:          models.ToPointer(models.RadioBandAntennaModeEnum_ENUMDEFAULT),
        Bandwidth:            models.ToPointer(models.Dot11Bandwidth5Enum_ENUM40),
        Channels:             models.NewOptional[int](nil),
        Disabled:             models.ToPointer(false),
        Power:                models.NewOptional(models.ToPointer(6)),
        Preamble:             models.ToPointer(models.RadioBandPreambleEnum_SHORT),
    }

}
```

