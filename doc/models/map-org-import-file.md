
# Map Org Import File

*This model accepts additional fields of type interface{}.*

## Structure

`MapOrgImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | whether to auto assign device to deviceprofile by name |
| `Csv` | `*[]byte` | Optional | csv file for ap name mapping, optional |
| `File` | `*[]byte` | Optional | ekahau or ibwave file |
| `Json` | [`*models.MapOrgImportFileJson`](../../doc/models/map-org-import-file-json.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
    "site_id": "00001f12-0000-0000-0000-000000000000",
    "vendor_name": "ekahau",
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

