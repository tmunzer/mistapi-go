
# Events Client

Client events

## Structure

`EventsClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `6` |
| `Bssid` | `*string` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `KeyMgmt` | [`*models.ClientKeyMgmtEnum`](../../doc/models/client-key-mgmt-enum.md) | Optional | Key management protocol used for the latest authentication. enum: `WPA2-PSK`, `WPA2-PSK-FT`, `WPA2-PSK-SHA256`, `WPA3-EAP-SHA256`, `WPA3-SAE-FT`, `WPA3-SAE-PSK` |
| `Proto` | [`*models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Optional | enum: `a`, `ac`, `ax`, `b`, `be`, `g`, `n` |
| `Ssid` | `*string` | Optional | - |
| `Text` | `*string` | Optional | - |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Type` | `*string` | Optional | Event type, e.g. MARVIS_EVENT_CLIENT_FBT_FAILURE |
| `TypeCode` | `*int` | Optional | For assoc/disassoc events |
| `WlanId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "band": "5",
  "key_mgmt": "WPA2-PSK",
  "timestamp": 36.2,
  "ap": "ap2",
  "bssid": "bssid6",
  "channel": 2,
  "proto": "a"
}
```

