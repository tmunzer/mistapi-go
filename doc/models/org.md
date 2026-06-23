
# Org

Mist organization containing sites, devices, users, and organization-level settings

*This model accepts additional fields of type interface{}.*

## Structure

`Org`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmtemplateId` | `models.Optional[uuid.UUID]` | Optional | Org-level alarm template ID used as the default for sites |
| `AllowMist` | `*bool` | Optional | Whether Mist support access is allowed for this organization<br><br>**Default**: `true` |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `MspLogoUrl` | `*string` | Optional, Read-only | logo uploaded by the MSP with advanced tier, only present if provided |
| `MspName` | `*string` | Optional, Read-only | Name of the msp the org belongs to |
| `Name` | `string` | Required | Display name of the organization |
| `OrggroupIds` | `[]uuid.UUID` | Optional | List of organization group identifiers |
| `SessionExpiry` | `*int` | Optional | Admin session lifetime for the organization, in minutes<br><br>**Default**: `1440`<br><br>**Constraints**: `>= 10`, `<= 20160` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    org := models.Org{
        AlarmtemplateId:      models.NewOptional(models.ToPointer(uuid.MustParse("00000682-0000-0000-0000-000000000000"))),
        AllowMist:            models.ToPointer(true),
        CreatedTime:          models.ToPointer(float64(36.26)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(42.7)),
        MspId:                models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        MspLogoUrl:           models.ToPointer("https://example.com/logo/b9d42c2e-88ee-41f8-b798-f009ce7fe909.jpeg"),
        MspName:              models.ToPointer("MSP"),
        Name:                 "Org",
        SessionExpiry:        models.ToPointer(1440),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

