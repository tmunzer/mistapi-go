
# Map Org Import File Json

*This model accepts additional fields of type interface{}.*

## Structure

`MapOrgImportFileJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ImportAllFloorplans` | `*bool` | Optional | **Default**: `false` |
| `ImportHeight` | `*bool` | Optional | **Default**: `true` |
| `ImportOrientation` | `*bool` | Optional | **Default**: `true` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `VendorName` | [`models.MapOrgImportFileJsonVendorNameEnum`](../../doc/models/map-org-import-file-json-vendor-name-enum.md) | Required | enum: `ekahau`, `ibwave` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "import_all_floorplans": false,
  "import_height": true,
  "import_orientation": true,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "vendor_name": "ekahau",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

