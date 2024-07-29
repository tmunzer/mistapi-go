
# Events Client

client events

## Structure

`EventsClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `6` |
| `Bssid` | `*string` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `Proto` | [`models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Required | enum: `a`, `ac`, `ax`, `b`, `g`, `n` |
| `Ssid` | `*string` | Optional | - |
| `Text` | `*string` | Optional | - |
| `Timestamp` | `float64` | Required | - |
| `Type` | `*string` | Optional | event type, e.g. MARVIS_EVENT_CLIENT_FBT_FAILURE |
| `TypeCode` | `*int` | Optional | for assoc/disassoc events |
| `WlanId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "ap8",
  "band": "5",
  "bssid": "bssid6",
  "channel": 108,
  "proto": "ax",
  "ssid": "ssid0",
  "text": "text2",
  "timestamp": 204.1
}
```

