
# Response Org Inventory Change Op Enum

enum: `assign`, `delete`, `downgrade_to_jsi`, `unassign`, `upgrade_to_mist`

## Enumeration

`ResponseOrgInventoryChangeOpEnum`

## Fields

| Name |
|  --- |
| `assign` |
| `delete` |
| `downgrade_to_jsi` |
| `unassign` |
| `upgrade_to_mist` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseOrgInventoryChangeOp := models.ResponseOrgInventoryChangeOpEnum_UPGRADETOMIST

}
```

