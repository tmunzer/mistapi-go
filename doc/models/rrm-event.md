
# Rrm Event

## Structure

`RrmEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApId` | `uuid.UUID` | Required | - |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `6` |
| `Bandwidth` | [`models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Required | channel width for the band.enum: `0`(disabled, response only), `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6) |
| `Channel` | `int` | Required | Channel for the band from rrm |
| `Event` | [`models.RrmEventTypeEnum`](../../doc/models/rrm-event-type-enum.md) | Required | enum: `interference-ap-co-channel`, `interference-ap-non-wifi`, `neighbor-ap-down`, `neighbor-ap-recovered`, `radar-detected`, `rrm-radar`, `scheduled-site_rrm`, `triggered-site_rrm` |
| `Power` | `int` | Required | Tx power of the radio |
| `PreBandwidth` | [`models.RrmEventPreBandwidthEnum`](../../doc/models/rrm-event-pre-bandwidth-enum.md) | Required | (previously) channel width for the band , 0 means no previously available. enum: `0`, `20`, `40`, `80`, `160` |
| `PreChannel` | `int` | Required | (previously) channel for the band, 0 means no previously available |
| `PrePower` | `float64` | Required | (previously) tx power of the radio, 0 means no previously available |
| `PreUsage` | `string` | Required | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Usage` | `string` | Required | - |

## Example (as JSON)

```json
{
  "ap_id": "00000c18-0000-0000-0000-000000000000",
  "band": "5",
  "bandwidth": 20,
  "channel": 130,
  "event": "interference-ap-co-channel",
  "power": 186,
  "pre_bandwidth": 20,
  "pre_channel": 64,
  "pre_power": 187.92,
  "pre_usage": "pre_usage8",
  "timestamp": 162.92,
  "usage": "usage2"
}
```

