
# Const Device Ap Band 24

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceApBand24`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band5ChannelsOp` | `*string` | Optional | - |
| `MaxClients` | `*int` | Optional | - |
| `MaxPower` | `*int` | Optional | - |
| `MinPower` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band5_channels_op": "low",
  "max_clients": 128,
  "max_power": 19,
  "min_power": 8,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

