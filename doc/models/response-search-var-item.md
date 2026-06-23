
# Response Search Var Item

Organization variable record returned by search

## Structure

`ResponseSearchVarItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Src` | `*string` | Optional | Source scope for the variable, such as `site` or `deviceprofile` |
| `Var` | `*string` | Optional | Name of the variable matched by the search |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSearchVarItem := models.ResponseSearchVarItem{
        CreatedTime:          models.ToPointer(float64(134.26)),
        ModifiedTime:         models.ToPointer(float64(200.7)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Src:                  models.ToPointer("src2"),
    }

}
```

