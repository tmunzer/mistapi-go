
# Map Site Replace File

## Structure

`MapSiteReplaceFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `[]byte` | Required | - |
| `Json` | [`*models.MapSiteReplaceFileJson`](../../doc/models/map-site-replace-file-json.md) | Optional | - |

## Example (as JSON)

```json
{
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": {
    "transform": {
      "rotation": 130.62,
      "scale": 195.12,
      "x": 28.38,
      "y": 102.9
    }
  }
}
```

