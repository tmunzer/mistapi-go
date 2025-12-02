
# Map Org Import File

## Structure

`MapOrgImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | Whether to auto assign device to deviceprofile by name |
| `Csv` | `*string` | Optional | CSV file for ap name mapping, optional |
| `File` | `*string` | Optional | Ekahau or ibwave file |
| `Json` | [`*models.MapOrgImportFileJson`](../../doc/models/map-org-import-file-json.md) | Optional | - |

## Example (as JSON)

```json
{
  "auto_deviceprofile_assignment": true,
  "csv": "csv8",
  "file": "file8",
  "json": {
    "import_all_floorplans": false,
    "import_height": false,
    "import_orientation": false,
    "site_id": "00001f12-0000-0000-0000-000000000000",
    "vendor_name": "ekahau"
  }
}
```

