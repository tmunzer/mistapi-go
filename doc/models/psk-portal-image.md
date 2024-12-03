
# Psk Portal Image

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortalImage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | Binary file |
| `Json` | `*string` | Optional | JSON string describing the upload |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": "json0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

