
# Response Map Import Floorplan

Floorplan result from a map import

## Structure

`ResponseMapImportFloorplan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | `string` | Required | Import action applied to this floorplan |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `MapId` | `uuid.UUID` | Required | Map identifier associated with the imported floorplan |
| `Name` | `string` | Required | Floorplan name imported from the map file |
| `Reason` | `*string` | Optional | Explanation of why the floorplan import action was not completed, when provided |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseMapImportFloorplan := models.ResponseMapImportFloorplan{
        Action:               "action6",
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        MapId:                uuid.MustParse("00001bba-0000-0000-0000-000000000000"),
        Name:                 "name6",
        Reason:               models.ToPointer("reason2"),
    }

}
```

