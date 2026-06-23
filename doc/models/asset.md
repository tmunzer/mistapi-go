
# Asset

Named Bluetooth Low Energy asset record

*This model accepts additional fields of type interface{}.*

## Structure

`Asset`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether this asset is scoped directly to a site |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Mac` | `string` | Required | Bluetooth MAC address used to identify the BLE asset |
| `MapId` | `*uuid.UUID` | Optional | Identifier of the map associated with this asset, when specified |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name or label for the BLE asset |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TagId` | `*uuid.UUID` | Optional | BLE tag identifier associated with this asset |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    asset := models.Asset{
        CreatedTime:          models.ToPointer(float64(89.5)),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Mac:                  "mac4",
        MapId:                models.ToPointer(uuid.MustParse("0000057c-0000-0000-0000-000000000000")),
        ModifiedTime:         models.ToPointer(float64(245.46)),
        Name:                 "name0",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

