
# Rrm Event

## Structure

`RrmEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApId` | `uuid.UUID` | Required | - |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | - |
| `Bandwidth` | [`models.Dot11BandwidthEnum`](../../doc/models/dot-11-bandwidth-enum.md) | Required | channel width for the band<br><br>* `80` is only applicable for band_5 and band_6<br>* `160` is only for band_6 |
| `Channel` | `int` | Required | channel for the band from rrm |
| `Event` | [`models.RrmEventTypeEnum`](../../doc/models/rrm-event-type-enum.md) | Required | schedule-site_rrm / triggered-site_rrm / interference-ap-co-channel / rrm-radar |
| `Power` | `int` | Required | tx power of the radio |
| `PreBandwidth` | [`models.RrmEventPreBandwidthEnum`](../../doc/models/rrm-event-pre-bandwidth-enum.md) | Required | (previously) channel width for the band , 0 means no previously available |
| `PreChannel` | `int` | Required | (previously) channel for the band, 0 means no previously available |
| `PrePower` | `float64` | Required | (previously) tx power of the radio, 0 means no previously available |
| `PreUsage` | `string` | Required | - |
| `Timestamp` | `float64` | Required | timestamp of the event |
| `Usage` | `string` | Required | - |

## Example (as JSON)

```json
{
  "ap_id": "00001a68-0000-0000-0000-000000000000",
  "band": "24",
  "bandwidth": 20,
  "channel": 34,
  "event": "triggered-site_rrm",
  "power": 26,
  "pre_bandwidth": 0,
  "pre_channel": 160,
  "pre_power": 51.28,
  "pre_usage": "pre_usage6",
  "timestamp": 43.56,
  "usage": "usage6"
}
```

