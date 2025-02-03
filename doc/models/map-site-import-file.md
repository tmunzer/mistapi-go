
# Map Site Import File

*This model accepts additional fields of type interface{}.*

## Structure

`MapSiteImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | Whether to auto assign device to deviceprofile by name |
| `Csv` | `*[]byte` | Optional | CSV file for ap name mapping, optional |
| `File` | `*[]byte` | Optional | Ekahau or ibwave file |
| `Json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Optional | - |
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

