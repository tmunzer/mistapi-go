
# Stats Ap Radio Config

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApRadioConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.StatsApRadioConfigBand`](../../doc/models/stats-ap-radio-config-band.md) | Optional | - |
| `Band24Usage` | `models.Optional[string]` | Optional | - |
| `Band5` | [`*models.StatsApRadioConfigBand`](../../doc/models/stats-ap-radio-config-band.md) | Optional | - |
| `Band6` | [`*models.StatsApRadioConfigBand`](../../doc/models/stats-ap-radio-config-band.md) | Optional | - |
| `ScanningEnabled` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band_24_usage": "5",
  "band_24": {
    "allow_rrm_disable": false,
    "bandwidth": 4.04,
    "channel": 80,
    "disabled": false,
    "dynamic_chaining_enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band_5": {
    "allow_rrm_disable": false,
    "bandwidth": 218.56,
    "channel": 132,
    "disabled": false,
    "dynamic_chaining_enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band_6": {
    "allow_rrm_disable": false,
    "bandwidth": 77.08,
    "channel": 200,
    "disabled": false,
    "dynamic_chaining_enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "scanning_enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

