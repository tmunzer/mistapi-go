
# Map Org Import File Json

Options for importing map files at organization scope

## Structure

`MapOrgImportFileJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ImportAllFloorplans` | `*bool` | Optional | Whether to import all floorplans from the uploaded map file<br><br>**Default**: `false` |
| `ImportHeight` | `*bool` | Optional | Whether to import height metadata from the uploaded map file<br><br>**Default**: `true` |
| `ImportOrientation` | `*bool` | Optional | Whether to import orientation metadata from the uploaded map file<br><br>**Default**: `true` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `VendorName` | [`models.MapOrgImportFileJsonVendorNameEnum`](../../doc/models/map-org-import-file-json-vendor-name-enum.md) | Required | Map import vendor for the uploaded file. enum: `ekahau`, `ibwave` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mapOrgImportFileJson := models.MapOrgImportFileJson{
        ImportAllFloorplans:  models.ToPointer(false),
        ImportHeight:         models.ToPointer(true),
        ImportOrientation:    models.ToPointer(true),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        VendorName:           models.MapOrgImportFileJsonVendorNameEnum_EKAHAU,
    }

}
```

