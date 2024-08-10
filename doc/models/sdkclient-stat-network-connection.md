
# Sdkclient Stat Network Connection

various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as

## Structure

`SdkclientStatNetworkConnection`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | - |
| `Rssi` | `float64` | Required | - |
| `SignalLevel` | `float64` | Required | - |
| `Type` | `string` | Required | - |

## Example (as JSON)

```json
{
  "mac": "mac2",
  "rssi": 143.3,
  "signal_level": 211.32,
  "type": "type2"
}
```

