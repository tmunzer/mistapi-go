
# Last Trouble

Last trouble code of switch

*This model accepts additional fields of type interface{}.*

## Structure

`LastTrouble`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | Code definitions list at [List Ap Led Definition](/#operations/listApLedDefinition) |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "code": "03",
  "timestamp": 160.78,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

