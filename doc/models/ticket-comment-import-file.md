
# Ticket Comment Import File

*This model accepts additional fields of type interface{}.*

## Structure

`TicketCommentImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Comment` | `*string` | Optional | - |
| `File` | `*[]byte` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "comment": "this is urgent",
  "file": "data:text/plain;name=dummy_file;base64,",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

