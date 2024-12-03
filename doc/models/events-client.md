
# Events Client

client events

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "ap2",
  "band": "5",
  "bssid": "bssid6",
  "channel": 2,
  "proto": "g",
  "ssid": "ssid0",
  "text": "text8",
  "timestamp": 36.2,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

