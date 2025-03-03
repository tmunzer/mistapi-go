
# Response Autoplacement Device

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAutoplacementDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `*string` | Optional | Provides the reason for the status if the AP is invalid. |
| `Valid` | `*bool` | Optional | Indicates whether the ap is valid. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "reason": "reason0",
  "valid": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

