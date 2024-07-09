
# Rf Template

RF Template

## Structure

`RfTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntGain24` | `*int` | Optional | - |
| `AntGain5` | `*int` | Optional | - |
| `AntGain6` | `*int` | Optional | - |
| `Band24` | [`*models.RftemplateRadioBand24`](../../doc/models/rftemplate-radio-band-24.md) | Optional | Radio Band AP settings |
| `Band24Usage` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | - |
| `Band5` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band5On24Radio` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band6` | [`*models.RftemplateRadioBand6`](../../doc/models/rftemplate-radio-band-6.md) | Optional | Radio Band AP settings |
| `CountryCode` | `*string` | Optional | optional, country code to use. If specified, this gets applied to all sites using the RF Template |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `ModelSpecific` | [`map[string]models.RfTemplateModelSpecificProperty`](../../doc/models/rf-template-model-specific-property.md) | Optional | overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. "AP63") |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | The name of the RF template |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `ScanningEnabled` | `*bool` | Optional | whether scanning radio is enabled |

## Example (as JSON)

```json
{
  "name": "name2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "ant_gain_24": 158,
  "ant_gain_5": 70,
  "ant_gain_6": 186,
  "band_24": {
    "allow_rrm_disable": false,
    "ant_gain": 0,
    "antenna_mode": "default",
    "bandwidth": 20,
    "channels": [
      221
    ]
  },
  "band_24_usage": "6"
}
```

