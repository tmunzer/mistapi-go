
# Stats Ap Radio Stat

## Structure

`StatsApRadioStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Radio stat |
| `Band5` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Radio stat |
| `Band6` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | Radio stat |

## Example (as JSON)

```json
{
  "band_24": {
    "bandwidth": 160,
    "channel": 80,
    "dynamic_chaining_enabled": false,
    "mac": "mac4",
    "noise_floor": 180
  },
  "band_5": {
    "bandwidth": 20,
    "channel": 132,
    "dynamic_chaining_enabled": false,
    "mac": "mac6",
    "noise_floor": 128
  },
  "band_6": {
    "bandwidth": 80,
    "channel": 200,
    "dynamic_chaining_enabled": false,
    "mac": "mac8",
    "noise_floor": 60
  }
}
```

