
# Response Map Import Ap

AP placement result from a map import

## Structure

`ResponseMapImportAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`models.ResponseMapImportApActionEnum`](../../doc/models/response-map-import-ap-action-enum.md) | Required | enum: `assigned-named-placed`, `assigned-placed`, `ignored`, `named-placed`, `placed` |
| `FloorplanId` | `uuid.UUID` | Required | Floorplan identifier where the AP placement was imported |
| `Height` | `*float64` | Optional | Mounting height for the AP on the imported floorplan |
| `Mac` | `string` | Required | AP MAC address matched from the import file |
| `MapId` | `uuid.UUID` | Required | Map identifier associated with the imported AP placement |
| `Orientation` | `int` | Required | AP orientation in degrees on the imported floorplan |
| `Reason` | `*string` | Optional | Explanation of why the AP import action was not completed, when provided |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseMapImportAp := models.ResponseMapImportAp{
        Action:               models.ResponseMapImportApActionEnum_ASSIGNEDNAMEDPLACED,
        FloorplanId:          uuid.MustParse("00001bcc-0000-0000-0000-000000000000"),
        Height:               models.ToPointer(float64(13.28)),
        Mac:                  "mac6",
        MapId:                uuid.MustParse("00000444-0000-0000-0000-000000000000"),
        Orientation:          124,
        Reason:               models.ToPointer("reason8"),
    }

}
```

