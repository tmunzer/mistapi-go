
# Map Import Json

*This model accepts additional fields of type interface{}.*

## Structure

`MapImportJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ImportAllFloorplans` | `*bool` | Optional | **Default**: `false` |
| `ImportHeight` | `*bool` | Optional | **Default**: `true` |
| `ImportOrientation` | `*bool` | Optional | **Default**: `true` |
| `VendorName` | [`models.MapImportJsonVendorNameEnum`](../../doc/models/map-import-json-vendor-name-enum.md) | Required | enum: `ekahau`, `ibwave` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "import_all_floorplans": false,
  "import_height": true,
  "import_orientation": true,
  "vendor_name": "ekahau",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

