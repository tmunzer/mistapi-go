
# Inventory Search Result

Inventory record returned by inventory search

## Structure

`InventorySearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LastDisconnected` | `*float64` | Optional | Time the device was last disconnected, in epoch seconds |
| `LastNameChange` | `*float64` | Optional | Time the inventory entry name was last changed, in epoch seconds |
| `Mac` | `*string` | Optional | Device MAC address for this inventory search result |
| `Magic` | `*string` | Optional | Device claim code (magic) for this inventory search result |
| `Master` | `*bool` | Optional | Whether this search result represents the master member of a Virtual Chassis |
| `Members` | [`[]models.InventorySearchResultMember`](../../doc/models/inventory-search-result-member.md) | Optional | Virtual Chassis members included in an inventory search result |
| `Model` | `*string` | Optional | Device model for this inventory search result |
| `Name` | `*string` | Optional | Configured device name in this inventory search result |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Serial` | `*string` | Optional | Device serial number for this inventory search result |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Sku` | `*string` | Optional | Device SKU for this inventory search result |
| `Status` | `*string` | Optional | Current inventory status for the device |
| `Timestamp` | `*float64` | Optional | Time the inventory entry was created, in epoch seconds |
| `Type` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `VcMac` | `*string` | Optional | Virtual Chassis MAC address for this inventory search result |
| `Version` | `*string` | Optional | Software version reported for this inventory search result |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    inventorySearchResult := models.InventorySearchResult{
        LastDisconnected:     models.ToPointer(float64(1616109044)),
        LastNameChange:       models.ToPointer(float64(1784007045.409)),
        Mac:                  models.ToPointer("f01c2df166e0"),
        Magic:                models.ToPointer("WVTFBLTNPXD23H2"),
        Master:               models.ToPointer(true),
        Model:                models.ToPointer("EX4300-48P"),
        Name:                 models.ToPointer("mist-wa-ex4300-VC"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Serial:               models.ToPointer("PD3714460200"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Sku:                  models.ToPointer("EX4300-48P"),
        Status:               models.ToPointer("disconnected"),
        Timestamp:            models.ToPointer(float64(1616108044)),
        Type:                 models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
        VcMac:                models.ToPointer("f01c2df166e0"),
        Version:              models.ToPointer("21.4R3.5"),
    }

}
```

