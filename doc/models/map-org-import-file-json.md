
# Map Org Import File Json

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

## Example (as JSON)

```json
{
  "import_all_floorplans": false,
  "import_height": true,
  "import_orientation": true,
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "vendor_name": "ekahau"
}
```

