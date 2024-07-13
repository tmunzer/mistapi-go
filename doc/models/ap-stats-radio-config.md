
# Ap Stats Radio Config

## Structure

`ApStatsRadioConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band24` | [`*models.ApStatsRadioConfigBand`](../../doc/models/ap-stats-radio-config-band.md) | Optional | - |
| `Band24Usage` | `models.Optional[string]` | Optional | - |
| `Band5` | [`*models.ApStatsRadioConfigBand`](../../doc/models/ap-stats-radio-config-band.md) | Optional | - |
| `Band6` | [`*models.ApStatsRadioConfigBand`](../../doc/models/ap-stats-radio-config-band.md) | Optional | - |
| `ScanningEnabled` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "band_24_usage": "5",
  "band_24": {
    "allow_rrm_disable": false,
    "bandwidth": 4.04,
    "channel": 80,
    "disabled": false,
    "dynamic_chaining_enabled": false
  },
  "band_5": {
    "allow_rrm_disable": false,
    "bandwidth": 218.56,
    "channel": 132,
    "disabled": false,
    "dynamic_chaining_enabled": false
  },
  "band_6": {
    "allow_rrm_disable": false,
    "bandwidth": 77.08,
    "channel": 200,
    "disabled": false,
    "dynamic_chaining_enabled": false
  },
  "scanning_enabled": false
}
```

