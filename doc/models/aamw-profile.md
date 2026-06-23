
# Aamw Profile

Advanced Anti Malware profile that controls Sky ATP file verdict handling

*This model accepts additional fields of type interface{}.*

## Structure

`AamwProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Categories` | [`[]models.AamwProfileCategory`](../../doc/models/aamw-profile-category.md) | Optional | File category rules included in an Advanced Anti Malware profile |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `FallbackAction` | [`*models.AamwProfileActionEnum`](../../doc/models/aamw-profile-action-enum.md) | Optional | Action applied to files by an Advanced Anti Malware profile. enum: `block`, `permit`<br><br>**Default**: `"block"` |
| `FileAction` | [`*models.AamwProfileActionEnum`](../../doc/models/aamw-profile-action-enum.md) | Optional | Action applied to files by an Advanced Anti Malware profile. enum: `block`, `permit`<br><br>**Default**: `"block"` |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the Advanced Anti Malware profile |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `VerdictThreshold` | `*int` | Optional | Minimum Sky ATP verdict score that triggers the configured file action<br><br>**Default**: `8`<br><br>**Constraints**: `>= 1`, `<= 10` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    aamwProfile := models.AamwProfile{
        Categories:           []models.AamwProfileCategory{
            models.AamwProfileCategory{
                Category:             models.ToPointer(models.AamwProfileCategoryCategoryEnum_EXECUTABLE),
                HashLookupOnly:       models.ToPointer(false),
            },
        },
        CreatedTime:          models.ToPointer(float64(3.48)),
        FallbackAction:       models.ToPointer(models.AamwProfileActionEnum_BLOCK),
        FileAction:           models.ToPointer(models.AamwProfileActionEnum_BLOCK),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 models.ToPointer("aamw-custom"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        VerdictThreshold:     models.ToPointer(8),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

