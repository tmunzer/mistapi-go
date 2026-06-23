
# Rf Template Model Specific Property

AP model-specific RF overrides for an RF template

## Structure

`RfTemplateModelSpecificProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntGain24` | `*int` | Optional | Model-specific external antenna gain for the 2.4 GHz radio<br><br>**Default**: `0` |
| `AntGain5` | `*int` | Optional | Model-specific external antenna gain for the 5 GHz radio<br><br>**Default**: `0` |
| `AntGain6` | `*int` | Optional | Model-specific external antenna gain for the 6 GHz radio<br><br>**Default**: `0` |
| `Band24` | [`*models.RftemplateRadioBand24`](../../doc/models/rftemplate-radio-band-24.md) | Optional | Radio Band AP settings |
| `Band24Usage` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | enum: `24`, `5`, `6`, `auto` |
| `Band5` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band5On24Radio` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band6` | [`*models.RftemplateRadioBand6`](../../doc/models/rftemplate-radio-band-6.md) | Optional | Radio Band AP settings |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rfTemplateModelSpecificProperty := models.RfTemplateModelSpecificProperty{
        AntGain24:            models.ToPointer(0),
        AntGain5:             models.ToPointer(0),
        AntGain6:             models.ToPointer(0),
        Band24:               models.ToPointer(models.RftemplateRadioBand24{
            AllowRrmDisable:      models.ToPointer(false),
            AntGain:              models.NewOptional(models.ToPointer(0)),
            AntennaMode:          models.ToPointer(models.RadioBandAntennaModeEnum_ENUM1X1),
            Bandwidth:            models.ToPointer(models.Dot11Bandwidth24Enum_ENUM0),
            Channels:             models.NewOptional(models.ToPointer([]int{
                221,
            })),
        }),
        Band24Usage:          models.ToPointer(models.RadioBand24UsageEnum_ENUM24),
    }

}
```

