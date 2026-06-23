
# Map Site Import File

Multipart payload for importing map files at site scope

## Structure

`MapSiteImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | Whether to auto assign device to deviceprofile by name |
| `Csv` | `*[]byte` | Optional | Optional AP name-mapping CSV file |
| `File` | `*[]byte` | Optional | Ekahau or iBwave floorplan file to import |
| `Json` | [`*models.MapImportJson`](../../doc/models/map-import-json.md) | Optional | Options for importing map data from Ekahau or iBwave JSON |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapSiteImportFile := models.MapSiteImportFile{
        AutoDeviceprofileAssignment: models.ToPointer(true),
        Csv:                         models.ToPointer(nil),
        File:                        models.ToPointer(nil),
        Json:                        models.ToPointer(models.MapImportJson{
            ImportAllFloorplans:  models.ToPointer(false),
            ImportHeight:         models.ToPointer(false),
            ImportOrientation:    models.ToPointer(false),
            VendorName:           models.MapImportJsonVendorNameEnum_EKAHAU,
        }),
    }

}
```

