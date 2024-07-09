
# Ap Stats Radio Stat

## Structure

`ApStatsRadioStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.ApRadioStats`](../../doc/models/ap-radio-stats.md) | Optional | radio stat |
| `Band5` | [`*models.ApRadioStats`](../../doc/models/ap-radio-stats.md) | Optional | radio stat |
| `Band6` | [`*models.ApRadioStats`](../../doc/models/ap-radio-stats.md) | Optional | radio stat |

## Example (as JSON)

```json
{
  "band_24": {
    "bandwidth": 20,
    "channel": 80,
    "dynamic_chaining_enalbed": false,
    "mac": "mac4",
    "num_clients": 190
  },
  "band_5": {
    "bandwidth": 20,
    "channel": 132,
    "dynamic_chaining_enalbed": false,
    "mac": "mac6",
    "num_clients": 138
  },
  "band_6": {
    "bandwidth": 20,
    "channel": 200,
    "dynamic_chaining_enalbed": false,
    "mac": "mac8",
    "num_clients": 70
  }
}
```

