
# Sdkinvite

SDK invite configuration used to onboard mobile SDK clients to an organization

*This model accepts additional fields of type interface{}.*

## Structure

`Sdkinvite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | Whether the SDK invite can currently be used<br><br>**Default**: `true` |
| `ExpireTime` | `*int` | Optional | Expiration time for the SDK invite, in epoch seconds |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name shown for the SDK invite in the mobile experience |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Quota` | `*int` | Optional | Number of times this SDK invite can be used |
| `QuotaLimited` | `*bool` | Optional | Whether use of the SDK invite is limited by the quota value<br><br>**Default**: `false` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    sdkinvite := models.Sdkinvite{
        CreatedTime:          models.ToPointer(float64(203.14)),
        Enabled:              models.ToPointer(true),
        ExpireTime:           models.ToPointer(100),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(131.82)),
        Name:                 "name4",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        QuotaLimited:         models.ToPointer(false),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

