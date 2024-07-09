
# Rf Template Model Specific Property

## Structure

`RfTemplateModelSpecificProperty`

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

## Example (as JSON)

```json
{
  "ant_gain_24": 20,
  "ant_gain_5": 188,
  "ant_gain_6": 48,
  "band_24": {
    "allow_rrm_disable": false,
    "ant_gain": 0,
    "antenna_mode": "default",
    "bandwidth": 20,
    "channels": [
      221
    ]
  },
  "band_24_usage": "24"
}
```

