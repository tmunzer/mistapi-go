
# Inventory Search Result Member

Virtual Chassis member in an inventory search result

## Structure

`InventorySearchResultMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Member device MAC address in this inventory search result |
| `Model` | `*string` | Optional | Member device model in this inventory search result |
| `Serial` | `*string` | Optional | Member device serial number in this inventory search result |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    inventorySearchResultMember := models.InventorySearchResultMember{
        Mac:                  models.ToPointer("f01c2df166e0"),
        Model:                models.ToPointer("EX4300-48P"),
        Serial:               models.ToPointer("PD3714460200"),
    }

}
```

