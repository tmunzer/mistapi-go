
# Note String

*This model accepts additional fields of type interface{}.*

## Structure

`NoteString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Note` | `*string` | Optional | Some text note describing the intent |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "note": "maintenance window",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

