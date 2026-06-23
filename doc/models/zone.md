
# Zone

Zone defined on a site map

*This model accepts additional fields of type interface{}.*

## Structure

`Zone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this zone is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MapId` | `*uuid.UUID` | Optional, Read-only | Map where this zone is defined |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the zone |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Vertices` | [`[]models.ZoneVertex`](../../doc/models/zone-vertex.md) | Optional | Vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    zone := models.Zone{
        CreatedTime:          models.ToPointer(float64(89.02)),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MapId:                models.ToPointer(uuid.MustParse("000005ac-0000-0000-0000-000000000000")),
        ModifiedTime:         models.ToPointer(float64(245.94)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Vertices:             []models.ZoneVertex{
            models.ZoneVertex{
                X:                    float64(732),
                Y:                    float64(1821),
            },
            models.ZoneVertex{
                X:                    float64(732.5),
                Y:                    float64(1731),
            },
            models.ZoneVertex{
                X:                    float64(837.5),
                Y:                    float64(1731.5),
            },
            models.ZoneVertex{
                X:                    float64(839),
                Y:                    float64(1821),
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

