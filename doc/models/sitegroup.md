
# Sitegroup

Group of sites within an organization

*This model accepts additional fields of type interface{}.*

## Structure

`Sitegroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the site group |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteIds` | `[]uuid.UUID` | Optional | Site identifiers included in a site group |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    sitegroup := models.Sitegroup{
        CreatedTime:          models.ToPointer(float64(2.38)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(76.58)),
        Name:                 "name8",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteIds:              []uuid.UUID{
            uuid.MustParse("00000322-0000-0000-0000-000000000000"),
            uuid.MustParse("00000323-0000-0000-0000-000000000000"),
            uuid.MustParse("00000324-0000-0000-0000-000000000000"),
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

