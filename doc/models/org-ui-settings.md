
# Org Ui Settings

Organization UI settings databoard

*This model accepts additional fields of type interface{}.*

## Structure

`OrgUiSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Description` | `*string` | Optional | Text describing the databoard |
| `ForSite` | `*bool` | Optional, Read-only | Whether this databoard is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IsCustomDataboard` | `*bool` | Optional | Whether this is a custom databoard or not |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the databoard |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Purpose` | [`*models.OrgUiSettingsPurposeEnum`](../../doc/models/org-ui-settings-purpose-enum.md) | Optional | UI surface or purpose for this databoard. enum: `marvisdashboard` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Tiles` | [`[]models.OrgUiSettingsTile`](../../doc/models/org-ui-settings-tile.md) | Optional | List of tiles in the databoard |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgUiSettings := models.OrgUiSettings{
        CreatedTime:          models.ToPointer(float64(109.12)),
        Description:          models.ToPointer("This databoard shows AP stats"),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        IsCustomDataboard:    models.ToPointer(false),
        Name:                 models.ToPointer("AP Stats"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

