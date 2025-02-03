
# Events Rogue

Rogue events

*This model accepts additional fields of type interface{}.*

## Structure

`EventsRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required | - |
| `Bssid` | `string` | Required | - |
| `Channel` | `int` | Required | - |
| `Rssi` | `int` | Required | - |
| `Ssid` | `string` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "ap2",
  "bssid": "bssid4",
  "channel": 26,
  "rssi": 72,
  "ssid": "ssid0",
  "timestamp": 43.4,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

