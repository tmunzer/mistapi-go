
# Ap Radio

Radio AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApRadio`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowRrmDisable` | `*bool` | Optional | **Default**: `false` |
| `AntGain24` | `*int` | Optional | antenna gain for 2.4G - for models with external antenna only<br>**Constraints**: `>= 0` |
| `AntGain5` | `*int` | Optional | antenna gain for 5G - for models with external antenna only<br>**Constraints**: `>= 0` |
| `AntGain6` | `*int` | Optional | antenna gain for 6G - for models with external antenna only<br>**Constraints**: `>= 0` |
| `AntennaMode` | [`*models.ApRadioAntennaModeEnum`](../../doc/models/ap-radio-antenna-mode-enum.md) | Optional | enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`<br>**Default**: `"default"` |
| `Band24` | [`*models.ApRadioBand24`](../../doc/models/ap-radio-band-24.md) | Optional | Radio Band AP settings |
| `Band24Usage` | [`*models.RadioBand24UsageEnum`](../../doc/models/radio-band-24-usage-enum.md) | Optional | enum: `24`, `5`, `6`, `auto` |
| `Band5` | [`*models.ApRadioBand5`](../../doc/models/ap-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band5On24Radio` | [`*models.ApRadioBand5`](../../doc/models/ap-radio-band-5.md) | Optional | Radio Band AP settings |
| `Band6` | [`*models.ApRadioBand6`](../../doc/models/ap-radio-band-6.md) | Optional | Radio Band AP settings |
| `IndoorUse` | `*bool` | Optional | to make an outdoor operate indoor.<br>for an outdoor-ap, some channels are disallowed by default, this allows the user to use it as an indoor-ap<br>**Default**: `false` |
| `ScanningEnabled` | `*bool` | Optional | whether scanning radio is enabled |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "allow_rrm_disable": false,
  "ant_gain_24": 4,
  "ant_gain_5": 5,
  "ant_gain_6": 5,
  "antenna_mode": "default",
  "indoor_use": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

