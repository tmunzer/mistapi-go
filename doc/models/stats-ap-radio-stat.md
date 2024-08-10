
# Stats Ap Radio Stat

## Structure

`StatsApRadioStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | radio stat |
| `Band5` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | radio stat |
| `Band6` | [`*models.ApRadioStat`](../../doc/models/ap-radio-stat.md) | Optional | radio stat |

## Example (as JSON)

```json
{
  "band_24": {
    "bandwidth": 20,
    "channel": 80,
    "dynamic_chaining_enalbed": false,
    "mac": "mac4",
    "noise_floor": 180
  },
  "band_5": {
    "bandwidth": 20,
    "channel": 132,
    "dynamic_chaining_enalbed": false,
    "mac": "mac6",
    "noise_floor": 128
  },
  "band_6": {
    "bandwidth": 20,
    "channel": 200,
    "dynamic_chaining_enalbed": false,
    "mac": "mac8",
    "noise_floor": 60
  }
}
```

