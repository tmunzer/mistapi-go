
# Rrm Band

## Structure

`RrmBand`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bandwidth` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `*int` | Optional | proposed channel |
| `CurrBandwidht` | [`*models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Optional | channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `CurrChannel` | `*int` | Optional | current channel |
| `CurrPower` | `*int` | Optional | current tx power |
| `CurrUsage` | `*string` | Optional | current radio band<br>**Constraints**: *Minimum Length*: `1` |
| `Power` | `*int` | Optional | proposed tx power |
| `Usage` | `*string` | Optional | proposed radio band<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "bandwidth": 20,
  "curr_bandwidht": 20,
  "channel": 56,
  "curr_channel": 176,
  "curr_power": 140
}
```

