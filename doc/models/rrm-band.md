
# Rrm Band

*This model accepts additional fields of type interface{}.*

## Structure

`RrmBand`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `*int` | Optional | proposed channel |
| `CurrBandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `CurrChannel` | `*int` | Optional | Current channel |
| `CurrPower` | `*int` | Optional | Current tx power |
| `CurrUsage` | `*string` | Optional | Current radio band<br><br>**Constraints**: *Minimum Length*: `1` |
| `Power` | `*int` | Optional | proposed tx power |
| `Usage` | `*string` | Optional | proposed radio band<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bandwidth": 20,
  "curr_bandwidth": 20,
  "channel": 56,
  "curr_channel": 176,
  "curr_power": 140,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

