
# Response Inventory Inventory Added Items

Inventory device added by the claim operation

## Structure

`ResponseInventoryInventoryAddedItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Device MAC address for the claimed inventory item |
| `Magic` | `string` | Required | Activation code or claim code associated with the inventory item |
| `Model` | `string` | Required | Device model for the claimed inventory item |
| `Serial` | `string` | Required | Device serial number for the claimed inventory item |
| `Type` | `string` | Required | Device type for the claimed inventory item |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseInventoryInventoryAddedItems := models.ResponseInventoryInventoryAddedItems{
        Mac:                  "5c5b35000018",
        Magic:                "6JG8EPTFV2A9Z2N",
        Model:                "AP41",
        Serial:               "FXLH2015150025",
        Type:                 "ap",
    }

}
```

