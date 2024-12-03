
# Map Site Replace File

*This model accepts additional fields of type interface{}.*

## Structure

`MapSiteReplaceFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `[]byte` | Required | - |
| `Json` | [`*models.MapSiteReplaceFileJson`](../../doc/models/map-site-replace-file-json.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": {
    "transform": {
      "rotation": 130.62,
      "scale": 195.12,
      "x": 28.38,
      "y": 102.9,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

