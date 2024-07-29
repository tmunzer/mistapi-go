
# Map Site Import File

## Structure

`MapSiteImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | whether to auto assign device to deviceprofile by name |
| `Csv` | `*[]byte` | Optional | csv file for ap name mapping, optional |
| `File` | `*[]byte` | Optional | ekahau or ibwave file |
| `Json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Optional | - |

## Example (as JSON)

```json
{
  "auto_deviceprofile_assignment": true,
  "csv": "data:text/plain;name=dummy_file;base64,",
  "file": "data:text/plain;name=dummy_file;base64,",
  "json": {
    "import_all_floorplans": false,
    "import_height": false,
    "import_orientation": false,
    "vendor_name": "ekahau"
  }
}
```

