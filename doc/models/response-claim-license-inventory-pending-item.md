
# Response Claim License Inventory Pending Item

Inventory device pending asynchronous claim processing

## Structure

`ResponseClaimLicenseInventoryPendingItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Device MAC address pending asynchronous inventory claim |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseClaimLicenseInventoryPendingItem := models.ResponseClaimLicenseInventoryPendingItem{
        Mac:                  models.ToPointer("mac8"),
    }

}
```

