
# Inventory Search

Paginated inventory search response

## Structure

`InventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*float64` | Optional | End of the inventory search time window, in epoch seconds |
| `Limit` | `*int` | Optional | Maximum number of inventory search results requested |
| `Next` | `*string` | Optional | URL for the next page of inventory search results |
| `Results` | [`[]models.InventorySearchResult`](../../doc/models/inventory-search-result.md) | Optional | Inventory records returned by an inventory search |
| `Start` | `*float64` | Optional | Start of the inventory search time window, in epoch seconds |
| `Total` | `*int` | Optional | Number of inventory records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    inventorySearch := models.InventorySearch{
        End:                  models.ToPointer(float64(1784062455.7383447)),
        Limit:                models.ToPointer(1000),
        Next:                 models.ToPointer("next0"),
        Results:              []models.InventorySearchResult{
            models.InventorySearchResult{
                LastDisconnected:     models.ToPointer(float64(106.86)),
                LastNameChange:       models.ToPointer(float64(80.82)),
                Mac:                  models.ToPointer("mac0"),
                Magic:                models.ToPointer("magic6"),
                Master:               models.ToPointer(false),
            },
            models.InventorySearchResult{
                LastDisconnected:     models.ToPointer(float64(106.86)),
                LastNameChange:       models.ToPointer(float64(80.82)),
                Mac:                  models.ToPointer("mac0"),
                Magic:                models.ToPointer("magic6"),
                Master:               models.ToPointer(false),
            },
            models.InventorySearchResult{
                LastDisconnected:     models.ToPointer(float64(106.86)),
                LastNameChange:       models.ToPointer(float64(80.82)),
                Mac:                  models.ToPointer("mac0"),
                Magic:                models.ToPointer("magic6"),
                Master:               models.ToPointer(false),
            },
        },
        Start:                models.ToPointer(float64(1784058855.7383447)),
        Total:                models.ToPointer(1),
    }

}
```

