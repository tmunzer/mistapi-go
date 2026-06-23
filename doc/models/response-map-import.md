
# Response Map Import

Result of importing map files and matching AP placements

## Structure

`ResponseMapImport`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aps` | [`[]models.ResponseMapImportAp`](../../doc/models/response-map-import-ap.md) | Required | AP placement results produced by the map import<br><br>**Constraints**: *Unique Items Required* |
| `Floorplans` | [`[]models.ResponseMapImportFloorplan`](../../doc/models/response-map-import-floorplan.md) | Required | Floorplan import results produced by the map import<br><br>**Constraints**: *Unique Items Required* |
| `ForSite` | `*bool` | Optional, Read-only | Whether this map import response is scoped to a site |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Summary` | [`models.ResponseMapImportSummary`](../../doc/models/response-map-import-summary.md) | Required | Counts summarizing assignments made during the map import |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseMapImport := models.ResponseMapImport{
        Aps:                  []models.ResponseMapImportAp{
            models.ResponseMapImportAp{
                Action:               models.ResponseMapImportApActionEnum_NAMEDPLACED,
                FloorplanId:          uuid.MustParse("000013fe-0000-0000-0000-000000000000"),
                Height:               models.ToPointer(float64(205.3)),
                Mac:                  "mac8",
                MapId:                uuid.MustParse("00000c12-0000-0000-0000-000000000000"),
                Orientation:          126,
                Reason:               models.ToPointer("reason0"),
            },
        },
        Floorplans:           []models.ResponseMapImportFloorplan{
            models.ResponseMapImportFloorplan{
                Action:               "action6",
                Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
                MapId:                uuid.MustParse("00002380-0000-0000-0000-000000000000"),
                Name:                 "name6",
                Reason:               models.ToPointer("reason2"),
            },
        },
        ForSite:              models.ToPointer(false),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Summary:              models.ResponseMapImportSummary{
            NumApAssigned:        66,
            NumInvAssigned:       174,
            NumMapAssigned:       164,
        },
    }

}
```

