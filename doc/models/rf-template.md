
# Rf Template

RF Template

*This model accepts additional fields of type interface{}.*

## Structure

`RfTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntGain24` | `*int` | Optional | - |
| `AntGain5` | `*int` | Optional | - |
| `AntGain6` | `*int` | Optional | - |
| `Band24` | [`*models.RftemplateRadioBand24`](../../doc/models/rftemplate-radio-band-24.md) | Optional | Radio Band AP settings |
| `Band24Usage` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | enum: `24`, `5`, `6`, `auto` |
| `Band5` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band5On24Radio` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band6` | [`*models.RftemplateRadioBand6`](../../doc/models/rftemplate-radio-band-6.md) | Optional | Radio Band AP settings |
| `CountryCode` | `*string` | Optional | optional, country code to use. If specified, this gets applied to all sites using the RF Template |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModelSpecific` | [`map[string]models.RfTemplateModelSpecificProperty`](../../doc/models/rf-template-model-specific-property.md) | Optional | overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. "AP63") |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | The name of the RF template |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ScanningEnabled` | `*bool` | Optional | whether scanning radio is enabled |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "name8",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "ant_gain_24": 10,
  "ant_gain_5": 78,
  "ant_gain_6": 38,
  "band_24": {
    "allow_rrm_disable": false,
    "ant_gain": 0,
    "antenna_mode": "1x1",
    "bandwidth": 20,
    "channels": [
      221
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band_24_usage": "6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

