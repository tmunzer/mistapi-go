
# Response Inventory

Result of adding device claim codes to organization inventory

## Structure

`ResponseInventory`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Added` | `[]string` | Optional | Claim codes accepted into organization inventory |
| `Duplicated` | `[]string` | Optional | Claim codes already present in organization inventory |
| `Error` | `[]string` | Optional | Claim codes rejected by the inventory add operation |
| `InventoryAdded` | [`[]models.ResponseInventoryInventoryAddedItems`](../../doc/models/response-inventory-inventory-added-items.md) | Optional | Detailed inventory records added by the claim operation<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `InventoryDuplicated` | [`[]models.ResponseInventoryInventoryDuplicatedItems`](../../doc/models/response-inventory-inventory-duplicated-items.md) | Optional | Detailed inventory records already present during the claim operation<br><br>**Constraints**: *Unique Items Required* |
| `Reason` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseInventory := models.ResponseInventory{
        Added:                []string{
            "added8",
        },
        Duplicated:           []string{
            "duplicated7",
            "duplicated8",
            "duplicated9",
        },
        Error:                []string{
            "error1",
            "error2",
        },
        InventoryAdded:       []models.ResponseInventoryInventoryAddedItems{
            models.ResponseInventoryInventoryAddedItems{
                Mac:                  "mac0",
                Magic:                "magic6",
                Model:                "model4",
                Serial:               "serial6",
                Type:                 "type6",
            },
            models.ResponseInventoryInventoryAddedItems{
                Mac:                  "mac0",
                Magic:                "magic6",
                Model:                "model4",
                Serial:               "serial6",
                Type:                 "type6",
            },
            models.ResponseInventoryInventoryAddedItems{
                Mac:                  "mac0",
                Magic:                "magic6",
                Model:                "model4",
                Serial:               "serial6",
                Type:                 "type6",
            },
        },
        InventoryDuplicated:  []models.ResponseInventoryInventoryDuplicatedItems{
            models.ResponseInventoryInventoryDuplicatedItems{
                Mac:                  "mac0",
                Magic:                "magic6",
                Model:                "model4",
                Serial:               "serial6",
                Type:                 "type6",
            },
            models.ResponseInventoryInventoryDuplicatedItems{
                Mac:                  "mac0",
                Magic:                "magic6",
                Model:                "model4",
                Serial:               "serial6",
                Type:                 "type6",
            },
        },
    }

}
```

