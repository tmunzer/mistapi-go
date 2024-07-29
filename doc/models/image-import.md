
# Image Import

## Structure

`ImageImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `[]byte` | Required | binary file |
| `Json` | `*string` | Optional | JSON string describing your upload |

## Example (as JSON)

```json
{
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": "json6"
}
```

