
# Map Import Json

Options for importing map data from Ekahau or iBwave JSON

## Structure

`MapImportJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ImportAllFloorplans` | `*bool` | Optional | Whether to import all floorplans from the uploaded map JSON<br><br>**Default**: `false` |
| `ImportHeight` | `*bool` | Optional | Whether to import height metadata from the uploaded map JSON<br><br>**Default**: `true` |
| `ImportOrientation` | `*bool` | Optional | Whether to import orientation metadata from the uploaded map JSON<br><br>**Default**: `true` |
| `VendorName` | [`models.MapImportJsonVendorNameEnum`](../../doc/models/map-import-json-vendor-name-enum.md) | Required | Map import vendor for the uploaded JSON. enum: `ekahau`, `ibwave` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapImportJson := models.MapImportJson{
        ImportAllFloorplans:  models.ToPointer(false),
        ImportHeight:         models.ToPointer(true),
        ImportOrientation:    models.ToPointer(true),
        VendorName:           models.MapImportJsonVendorNameEnum_EKAHAU,
    }

}
```

