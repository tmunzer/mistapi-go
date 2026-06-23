
# Response Async Claim Create

Response to an async inventory claim request

## Structure

`ResponseAsyncClaimCreate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClaimId` | `*uuid.UUID` | Optional | Unique identifier for the async claim job, used to poll status |
| `InventoryPending` | [`[]models.ResponseClaimLicenseInventoryPendingItem`](../../doc/models/response-claim-license-inventory-pending-item.md) | Optional | Inventory devices pending asynchronous claim processing<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseAsyncClaimCreate := models.ResponseAsyncClaimCreate{
        ClaimId:              models.ToPointer(uuid.MustParse("00000fbe-0000-0000-0000-000000000000")),
        InventoryPending:     []models.ResponseClaimLicenseInventoryPendingItem{
            models.ResponseClaimLicenseInventoryPendingItem{
                Mac:                  models.ToPointer("mac6"),
            },
            models.ResponseClaimLicenseInventoryPendingItem{
                Mac:                  models.ToPointer("mac6"),
            },
            models.ResponseClaimLicenseInventoryPendingItem{
                Mac:                  models.ToPointer("mac6"),
            },
        },
    }

}
```

