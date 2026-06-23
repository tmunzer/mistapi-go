
# Response Msp Inventory Device

Inventory device record returned by MSP device MAC lookup

## Structure

`ResponseMspInventoryDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForSite` | `*bool` | Optional, Read-only | Whether this inventory device is associated with a site |
| `Mac` | `string` | Required, Read-only | Device MAC address for the inventory device |
| `Model` | `string` | Required, Read-only | Device model for the inventory device |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `Serial` | `string` | Required, Read-only | Device serial number for the inventory device |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Type` | `string` | Required, Read-only | Device type for the inventory device |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseMspInventoryDevice := models.ResponseMspInventoryDevice{
        ForSite:              models.ToPointer(false),
        Mac:                  "mac2",
        Model:                "model6",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        Serial:               "serial8",
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Type:                 "type2",
    }

}
```

