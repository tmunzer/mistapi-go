
# Ap Stats Radio Config

## Structure

`ApStatsRadioConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.ApStatsRadioConfigBand`](../../doc/models/ap-stats-radio-config-band.md) | Optional | - |
| `Band5` | [`*models.ApStatsRadioConfigBand`](../../doc/models/ap-stats-radio-config-band.md) | Optional | - |
| `Band6` | [`*models.ApStatsRadioConfigBand`](../../doc/models/ap-stats-radio-config-band.md) | Optional | - |
| `ScanningEnabled` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "band_24": {
    "bandwidth": 4.04,
    "channel": 80,
    "dynamic_chaining_enabled": false,
    "power": 240.44,
    "rx_chain": 8
  },
  "band_5": {
    "bandwidth": 218.56,
    "channel": 132,
    "dynamic_chaining_enabled": false,
    "power": 198.96,
    "rx_chain": 212
  },
  "band_6": {
    "bandwidth": 77.08,
    "channel": 200,
    "dynamic_chaining_enabled": false,
    "power": 198.52,
    "rx_chain": 144
  },
  "scanning_enabled": false
}
```

