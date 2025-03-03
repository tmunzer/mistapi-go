
# Last Trouble

Last trouble code of switch

*This model accepts additional fields of type interface{}.*

## Structure

`LastTrouble`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | Code definitions list at [List Ap Led Definition](../../doc/controllers/constants-definitions.md#list-ap-led-definition) |
| `Timestamp` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "code": "03",
  "timestamp": 1428949501,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

