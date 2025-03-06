
# Response Zone Search Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseZoneSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enter` | `*int` | Optional | - |
| `Scope` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `User` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enter": 1541705254,
  "scope": "map",
  "user": "c4b301c81166",
  "timestamp": 133.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

