
# Secpolicy

Security Policy is designed to audit / catch discrepancies between "what’s intended to be running" versus "what’s actually running" in a network. Many big organizations have separated Security and IT team (for good reasons). Each site can be assigned a security policy. Whenever an AP is provisioned, the configuration will be checked against the security policy. Any violations will be flagged in Device Config History where you can search for the when and where the violation occurs.

*This model accepts additional fields of type interface{}.*

## Structure

`Secpolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name of the security policy |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Wlans` | [`[]models.Wlan`](../../doc/models/wlan.md) | Optional | WLAN configurations audited by a security policy<br><br>**Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    secpolicy := models.Secpolicy{
        CreatedTime:          models.ToPointer(float64(67.52)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(11.44)),
        Name:                 models.ToPointer("name2"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

