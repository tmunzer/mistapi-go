
# Template

WLAN template configuration applied across sites or site groups

*This model accepts additional fields of type interface{}.*

## Structure

`Template`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Applies` | [`*models.TemplateApplies`](../../doc/models/template-applies.md) | Optional | Where this template should be applied to, can be org_id, site_ids, sitegroup_ids |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceprofileIds` | `[]uuid.UUID` | Optional | List of Device Profile ids |
| `Exceptions` | [`*models.TemplateExceptions`](../../doc/models/template-exceptions.md) | Optional | Where this template should not be applied to (takes precedence) |
| `FilterByDeviceprofile` | `*bool` | Optional | Whether to further filter by Device Profile |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the WLAN template |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    template := models.Template{
        Applies:               models.ToPointer(models.TemplateApplies{
            OrgId:                models.ToPointer(uuid.MustParse("00001bdc-0000-0000-0000-000000000000")),
            SiteIds:              []uuid.UUID{
                uuid.MustParse("00001706-0000-0000-0000-000000000000"),
                uuid.MustParse("00001707-0000-0000-0000-000000000000"),
                uuid.MustParse("00001708-0000-0000-0000-000000000000"),
            },
            SitegroupIds:         []uuid.UUID{
                uuid.MustParse("00000634-0000-0000-0000-000000000000"),
            },
        }),
        CreatedTime:           models.ToPointer(float64(247.46)),
        DeviceprofileIds:      []uuid.UUID{
            uuid.MustParse("0000104d-0000-0000-0000-000000000000"),
            uuid.MustParse("0000104e-0000-0000-0000-000000000000"),
            uuid.MustParse("0000104f-0000-0000-0000-000000000000"),
        },
        Exceptions:            models.ToPointer(models.TemplateExceptions{
            SiteIds:              []uuid.UUID{
                uuid.MustParse("00001604-0000-0000-0000-000000000000"),
                uuid.MustParse("00001603-0000-0000-0000-000000000000"),
                uuid.MustParse("00001602-0000-0000-0000-000000000000"),
            },
            SitegroupIds:         []uuid.UUID{
                uuid.MustParse("00000c2e-0000-0000-0000-000000000000"),
            },
        }),
        FilterByDeviceprofile: models.ToPointer(false),
        Id:                    models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                  "name6",
        OrgId:                 models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        AdditionalProperties:  map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

