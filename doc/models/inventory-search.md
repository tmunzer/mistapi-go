
# Inventory Search

Paginated inventory search response

## Structure

`InventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Page ending index for this inventory search response |
| `Limit` | `*int` | Optional | Maximum number of inventory search results requested |
| `Next` | `*string` | Optional | URL for the next page of inventory search results |
| `Results` | [`[]models.InventorySearchResult`](../../doc/models/inventory-search-result.md) | Optional | Inventory records returned by an inventory search |
| `Start` | `*int` | Optional | Page starting index for this inventory search response |
| `Total` | `*int` | Optional | Number of inventory records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    inventorySearch := models.InventorySearch{
        End:                  models.ToPointer(36),
        Limit:                models.ToPointer(1000),
        Next:                 models.ToPointer("next0"),
        Results:              []models.InventorySearchResult{
            models.InventorySearchResult{
                Mac:                  models.ToPointer("mac0"),
                Master:               models.ToPointer(false),
                Members:              []models.InventorySearchResultMember{
                    models.InventorySearchResultMember{
                        Mac:                  models.ToPointer("mac2"),
                        Model:                models.ToPointer("model6"),
                        Serial:               models.ToPointer("serial8"),
                    },
                    models.InventorySearchResultMember{
                        Mac:                  models.ToPointer("mac2"),
                        Model:                models.ToPointer("model6"),
                        Serial:               models.ToPointer("serial8"),
                    },
                },
                Model:                models.ToPointer("model4"),
                Name:                 models.ToPointer("name6"),
            },
            models.InventorySearchResult{
                Mac:                  models.ToPointer("mac0"),
                Master:               models.ToPointer(false),
                Members:              []models.InventorySearchResultMember{
                    models.InventorySearchResultMember{
                        Mac:                  models.ToPointer("mac2"),
                        Model:                models.ToPointer("model6"),
                        Serial:               models.ToPointer("serial8"),
                    },
                    models.InventorySearchResultMember{
                        Mac:                  models.ToPointer("mac2"),
                        Model:                models.ToPointer("model6"),
                        Serial:               models.ToPointer("serial8"),
                    },
                },
                Model:                models.ToPointer("model4"),
                Name:                 models.ToPointer("name6"),
            },
            models.InventorySearchResult{
                Mac:                  models.ToPointer("mac0"),
                Master:               models.ToPointer(false),
                Members:              []models.InventorySearchResultMember{
                    models.InventorySearchResultMember{
                        Mac:                  models.ToPointer("mac2"),
                        Model:                models.ToPointer("model6"),
                        Serial:               models.ToPointer("serial8"),
                    },
                    models.InventorySearchResultMember{
                        Mac:                  models.ToPointer("mac2"),
                        Model:                models.ToPointer("model6"),
                        Serial:               models.ToPointer("serial8"),
                    },
                },
                Model:                models.ToPointer("model4"),
                Name:                 models.ToPointer("name6"),
            },
        },
        Start:                models.ToPointer(250),
        Total:                models.ToPointer(1),
    }

}
```

