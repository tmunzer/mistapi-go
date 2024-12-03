
# Const Device Ap Band 5

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceApBand5`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxClients` | `*int` | Optional | - |
| `MaxPower` | `*int` | Optional | - |
| `MinPower` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "max_clients": 128,
  "max_power": 17,
  "min_power": 8,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

