
# Map Site Import File

## Structure

`MapSiteImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | Whether to auto assign device to deviceprofile by name |
| `Csv` | `*string` | Optional | CSV file for ap name mapping, optional |
| `File` | `*string` | Optional | Ekahau or ibwave file |
| `Json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Optional | - |

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
    "vendor_name": "ekahau"
  }
}
```

