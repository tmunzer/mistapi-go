
# Js Inventory Search

Paginated JSI inventory search response

## Structure

`JsInventorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Offset to end at |
| `Limit` | `*int` | Optional | Number of results to return |
| `Next` | `*string` | Optional | URL for the next page of JSI inventory results |
| `Results` | [`[]models.JsInventoryItem`](../../doc/models/js-inventory-item.md) | Optional | JSI inventory items returned for an organization |
| `Start` | `*int` | Optional | Offset to start from |
| `Total` | `*int` | Optional | Number of JSI inventory records matching the search |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    jsInventorySearch := models.JsInventorySearch{
        End:                  models.ToPointer(66),
        Limit:                models.ToPointer(104),
        Next:                 models.ToPointer("next2"),
        Results:              []models.JsInventoryItem{
            models.JsInventoryItem{
                Availability:         models.ToPointer("availability6"),
                Claimed:              models.ToPointer(false),
                ContractEndDate:      models.ToPointer("contract_end_date6"),
                ContractReseller:     models.ToPointer("contract_reseller4"),
                ContractStartDate:    models.ToPointer("contract_start_date2"),
            },
        },
        Start:                models.ToPointer(24),
    }

}
```

