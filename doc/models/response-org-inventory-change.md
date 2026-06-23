
# Response Org Inventory Change

Result of an organization inventory assignment or deletion operation

## Structure

`ResponseOrgInventoryChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Op` | [`models.ResponseOrgInventoryChangeOpEnum`](../../doc/models/response-org-inventory-change-op-enum.md) | Required | enum: `assign`, `delete`, `downgrade_to_jsi`, `unassign`, `upgrade_to_mist` |
| `Reason` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Success` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseOrgInventoryChange := models.ResponseOrgInventoryChange{
        Error:                []string{
            "error3",
            "error4",
            "error5",
        },
        Op:                   models.ResponseOrgInventoryChangeOpEnum_DELETE,
        Reason:               []string{
            "reason6",
            "reason7",
        },
        Success:              []string{
            "success2",
        },
    }

}
```

