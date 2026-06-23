
# Map Org Import File

Multipart payload for importing map files at organization scope

## Structure

`MapOrgImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoDeviceprofileAssignment` | `*bool` | Optional | Whether to auto assign device to deviceprofile by name |
| `Csv` | `*[]byte` | Optional | Optional AP name-mapping CSV file |
| `File` | `*[]byte` | Optional | Ekahau or iBwave floorplan file to import |
| `Json` | [`*models.MapOrgImportFileJson`](../../doc/models/map-org-import-file-json.md) | Optional | Options for importing map files at organization scope |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mapOrgImportFile := models.MapOrgImportFile{
        AutoDeviceprofileAssignment: models.ToPointer(true),
        Csv:                         models.ToPointer(nil),
        File:                        models.ToPointer(nil),
        Json:                        models.ToPointer(models.MapOrgImportFileJson{
            ImportAllFloorplans:  models.ToPointer(false),
            ImportHeight:         models.ToPointer(false),
            ImportOrientation:    models.ToPointer(false),
            SiteId:               models.ToPointer(uuid.MustParse("00001f12-0000-0000-0000-000000000000")),
            VendorName:           models.MapOrgImportFileJsonVendorNameEnum_EKAHAU,
        }),
    }

}
```

