
# Rf Template

RF template used by the current RRM calculation

*This model accepts additional fields of type interface{}.*

## Structure

`RfTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntGain24` | `*int` | Optional | External antenna gain for the 2.4 GHz radio |
| `AntGain5` | `*int` | Optional | External antenna gain for the 5 GHz radio |
| `AntGain6` | `*int` | Optional | External antenna gain for the 6 GHz radio |
| `Band24` | [`*models.RftemplateRadioBand24`](../../doc/models/rftemplate-radio-band-24.md) | Optional | Radio Band AP settings |
| `Band24Usage` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | enum: `24`, `5`, `6`, `auto` |
| `Band5` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band5On24Radio` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band6` | [`*models.RftemplateRadioBand6`](../../doc/models/rftemplate-radio-band-6.md) | Optional | Radio Band AP settings |
| `CountryCode` | `*string` | Optional | Optional, country code to use. If specified, this gets applied to all sites using the RF Template |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether the RF template is scoped to a site rather than the organization |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModelSpecific` | [`map[string]models.RfTemplateModelSpecificProperty`](../../doc/models/rf-template-model-specific-property.md) | Optional | overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. "AP63") |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | The name of the RF template |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `ScanningEnabled` | `*bool` | Optional | Whether scanning radio is enabled |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    rfTemplate := models.RfTemplate{
        AntGain24:            models.ToPointer(32),
        AntGain5:             models.ToPointer(200),
        AntGain6:             models.ToPointer(60),
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
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 "name2",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

