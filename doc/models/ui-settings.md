
# Ui Settings

Site UI settings databoard

*This model accepts additional fields of type interface{}.*

## Structure

`UiSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DefaultScopeId` | `*string` | Optional | Scope identifier used by default for this databoard |
| `DefaultScopeType` | `*string` | Optional | Scope type used by default for this databoard |
| `DefaultTimeRange` | [`*models.UiSettingsDefaultTimeRange`](../../doc/models/ui-settings-default-time-range.md) | Optional | Default time range applied to a site UI databoard |
| `Description` | `string` | Required | Text describing the databoard |
| `ForSite` | `*bool` | Optional, Read-only | Whether this databoard is scoped to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IsCustomDataboard` | `*bool` | Optional | Whether this databoard is custom-created |
| `IsScopeLinked` | `*bool` | Optional | Whether tile scopes are linked to the databoard default scope |
| `IsTimeRangeLinked` | `*bool` | Optional | Whether tile time ranges are linked to the databoard default time range |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the databoard |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Purpose` | `string` | Required | UI surface or purpose for this databoard |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Tiles` | [`[]models.UiSettingsTile`](../../doc/models/ui-settings-tile.md) | Optional | List of tiles in a site UI databoard<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    uiSettings := models.UiSettings{
        CreatedTime:          models.ToPointer(float64(50.84)),
        DefaultScopeId:       models.ToPointer("67970e46-4e12-11e6-9188-0242ad112847"),
        DefaultScopeType:     models.ToPointer("site"),
        DefaultTimeRange:     models.ToPointer(models.UiSettingsDefaultTimeRange{
            End:                  models.ToPointer(224),
            EndDate:              models.ToPointer("endDate8"),
            Interval:             models.ToPointer("interval4"),
            Name:                 models.ToPointer("name6"),
            ShortName:            models.ToPointer("shortName4"),
        }),
        Description:          "Description of the databoard",
        ForSite:              models.ToPointer(true),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 models.ToPointer("New Databoard"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Purpose:              "databoard",
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

