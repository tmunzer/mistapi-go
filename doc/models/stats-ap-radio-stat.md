
# Stats Ap Radio Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApRadioStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Radio stat |
| `Band5` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Radio stat |
| `Band6` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Radio stat |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band_24": {
    "bandwidth": 20,
    "channel": 80,
    "dynamic_chaining_enalbed": false,
    "mac": "mac4",
    "noise_floor": 180,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band_5": {
    "bandwidth": 20,
    "channel": 132,
    "dynamic_chaining_enalbed": false,
    "mac": "mac6",
    "noise_floor": 128,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "band_6": {
    "bandwidth": 20,
    "channel": 200,
    "dynamic_chaining_enalbed": false,
    "mac": "mac8",
    "noise_floor": 60,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

