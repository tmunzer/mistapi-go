
# License Sub

Read-only license subscription record

## Structure

`LicenseSub`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `EndTime` | `*int` | Optional, Read-only | End date of the license term |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `OrderId` | `*string` | Optional, Read-only | Order identifier for this license subscription |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Quantity` | `*int` | Optional, Read-only | Number of devices entitled for this license |
| `RemainingQuantity` | `*int` | Optional | Number of licenses left in this subscription |
| `StartTime` | `*int` | Optional, Read-only | Start date of the license term |
| `SubscriptionId` | `*string` | Optional, Read-only | Subscription identifier for this license subscription |
| `Type` | `*string` | Optional, Read-only | Type of license. The list of supported license type can be retrieve with the [List License Type](../../doc/controllers/constants-definitions.md#list-license-types) API request. |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    licenseSub := models.LicenseSub{
        CreatedTime:          models.ToPointer(float64(200.08)),
        EndTime:              models.ToPointer(18),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(134.88)),
        OrderId:              models.ToPointer("order_id2"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    }

}
```

