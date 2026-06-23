
# Sdktemplate

Visual customization template for the mobile SDK experience

*This model accepts additional fields of type interface{}.*

## Structure

`Sdktemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BgImage` | `*string` | Optional | Background image URL displayed by the SDK template |
| `BtnFlrBgcolor` | `*string` | Optional | Floor button background color used by the SDK template, as a hex color |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Default` | `*bool` | Optional | Whether this template is the default among available SDK templates |
| `ForSite` | `*bool` | Optional, Read-only | Whether the SDK template is scoped to a site |
| `HeaderTxt` | `*string` | Optional | Header text displayed by the SDK template |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name used to identify the SDK template |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SearchTxtcolor` | `*string` | Optional | Text color used for search controls in the SDK template, as a hex color |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `WelcomeMsg` | `*string` | Optional | Welcome message displayed by the SDK template |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    sdktemplate := models.Sdktemplate{
        BgImage:              models.ToPointer("bg_image0"),
        BtnFlrBgcolor:        models.ToPointer("btn_flr_bgcolor4"),
        CreatedTime:          models.ToPointer(float64(41.34)),
        Default:              models.ToPointer(false),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 "name4",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

