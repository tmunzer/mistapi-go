
# Idp Profile

Organization IDP profile with a base profile and targeted overwrite rules

*This model accepts additional fields of type interface{}.*

## Structure

`IdpProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BaseProfile` | [`*models.IdpProfileBaseProfileEnum`](../../doc/models/idp-profile-base-profile-enum.md) | Optional | enum: `critical`, `standard`, `strict` |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the IDP profile |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Overwrites` | [`[]models.IdpProfileOverwrite`](../../doc/models/idp-profile-overwrite.md) | Optional | IDP profile overwrite rules applied to the base profile |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    idpProfile := models.IdpProfile{
        BaseProfile:          models.ToPointer(models.IdpProfileBaseProfileEnum_STRICT),
        CreatedTime:          models.ToPointer(float64(44.6)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(34.36)),
        Name:                 models.ToPointer("relaxed"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

