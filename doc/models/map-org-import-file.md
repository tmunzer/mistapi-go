
# Map Org Import File

## Structure

`MapOrgImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | whether to auto assign device to deviceprofile by name |
| `Csv` | `*[]byte` | Optional | csv file for ap name mapping, optional |
| `File` | `*[]byte` | Optional | ekahau or ibwave file |
| `Json` | [`*models.MapOrgImportFileJson`](../../doc/models/map-org-import-file-json.md) | Optional | - |

## Example (as JSON)

```json
{
  "auto_deviceprofile_assignment": true,
  "csv": "csv0",
  "file": "file6",
  "json": {
    "import_all_floorplans": false,
    "import_height": false,
    "import_orientation": false,
    "site_id": "00001f12-0000-0000-0000-000000000000",
    "vendor_name": "ekahau"
  }
}
```

