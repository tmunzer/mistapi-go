
# Const Device Ap Vble

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceApVble`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BeaconRate` | `*int` | Optional | - |
| `Beams` | `*int` | Optional | - |
| `Power` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "beacon_rate": 4,
  "beams": 9,
  "power": 8,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

