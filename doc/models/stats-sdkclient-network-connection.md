
# Stats Sdkclient Network Connection

Various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as

*This model accepts additional fields of type interface{}.*

## Structure

`StatsSdkclientNetworkConnection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | - |
| `Rssi` | `float64` | Required | - |
| `SignalLevel` | `float64` | Required | - |
| `Type` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "mac6",
  "rssi": 116.74,
  "signal_level": 184.76,
  "type": "type8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

