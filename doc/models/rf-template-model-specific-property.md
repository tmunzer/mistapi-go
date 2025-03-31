
# Rf Template Model Specific Property

*This model accepts additional fields of type interface{}.*

## Structure

`RfTemplateModelSpecificProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntGain24` | `*int` | Optional | **Default**: `0` |
| `AntGain5` | `*int` | Optional | **Default**: `0` |
| `AntGain6` | `*int` | Optional | **Default**: `0` |
| `Band24` | [`*models.RftemplateRadioBand24`](../../doc/models/rftemplate-radio-band-24.md) | Optional | Radio Band AP settings |
| `Band24Usage` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | enum: `24`, `5`, `6`, `auto` |
| `Band5` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band5On24Radio` | [`*models.RftemplateRadioBand5`](../../doc/models/rftemplate-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band6` | [`*models.RftemplateRadioBand6`](../../doc/models/rftemplate-radio-band-6.md) | Optional | Radio Band AP settings |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ant_gain_24": 0,
  "ant_gain_5": 0,
  "ant_gain_6": 0,
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

