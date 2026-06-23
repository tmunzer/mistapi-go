
# Msp

Managed service provider account

*This model accepts additional fields of type interface{}.*

## Structure

`Msp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowMist` | `*bool` | Optional | Whether Mist support access is allowed for this MSP account |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `LogoUrl` | `*string` | Optional | For advanced tier (uMSPs) only |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the MSP account |
| `Tier` | [`*models.MspTierEnum`](../../doc/models/msp-tier-enum.md) | Optional | Service tier for the MSP account. enum: `advanced`, `base`<br><br>**Default**: `"base"` |
| `Url` | `*string` | Optional | For advanced tier (uMSPs) only |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    msp := models.Msp{
        AllowMist:            models.ToPointer(false),
        CreatedTime:          models.ToPointer(float64(109.3)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        LogoUrl:              models.ToPointer("logo_url0"),
        ModifiedTime:         models.ToPointer(float64(225.66)),
        Tier:                 models.ToPointer(models.MspTierEnum_BASE),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

