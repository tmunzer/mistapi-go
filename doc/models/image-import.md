
# Image Import

*This model accepts additional fields of type interface{}.*

## Structure

`ImageImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `[]byte` | Required | Binary file |
| `Json` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": "json8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

