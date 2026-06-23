
# License Amendment

Read-only change applied to a license subscription

## Structure

`LicenseAmendment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `EndTime` | `*int` | Optional, Read-only | End time of the license amendment |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Quantity` | `*int` | Optional, Read-only | License quantity associated with this amendment |
| `StartTime` | `*int` | Optional, Read-only | Start time of the license amendment |
| `SubscriptionId` | `*string` | Optional, Read-only | Subscription identifier associated with this amendment |
| `Type` | `*string` | Optional, Read-only | Type of license. The list of supported license type can be retrieve with the [List License Type](../../doc/controllers/constants-definitions.md#list-license-types) API request. |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    licenseAmendment := models.LicenseAmendment{
        CreatedTime:          models.ToPointer(float64(194.52)),
        EndTime:              models.ToPointer(230),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(140.44)),
        Quantity:             models.ToPointer(202),
    }

}
```

