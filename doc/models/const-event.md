
# Const Event

*This model accepts additional fields of type interface{}.*

## Structure

`ConstEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | - |
| `Display` | `string` | Required | - |
| `Example` | `*interface{}` | Optional | - |
| `Key` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "description": "description4",
  "display": "display6",
  "example": {
    "key1": "val1",
    "key2": "val2"
  },
  "key": "key4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

