
# Response Inventory Inventory Duplicated Items

Inventory device already present during the claim operation

## Structure

`ResponseInventoryInventoryDuplicatedItems`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Device MAC address for the duplicate inventory item |
| `Magic` | `string` | Required | Activation code or claim code associated with the duplicate inventory item |
| `Model` | `string` | Required | Device model for the duplicate inventory item |
| `Serial` | `string` | Required | Device serial number for the duplicate inventory item |
| `Type` | `string` | Required | Device type for the duplicate inventory item |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseInventoryInventoryDuplicatedItems := models.ResponseInventoryInventoryDuplicatedItems{
        Mac:                  "5c5b35000012",
        Magic:                "DVH4VSNMSZPDXBR",
        Model:                "AP41",
        Serial:               "FXLH2015150027",
        Type:                 "ap",
    }

}
```

