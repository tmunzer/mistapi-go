
# Response Claim License Inventory Item

Inventory device returned by a license claim

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClaimLicenseInventoryItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Device MAC address for the claimed inventory item |
| `Magic` | `string` | Required | Activation code or claim code associated with the inventory item |
| `Model` | `string` | Required | Device model for the claimed inventory item |
| `Serial` | `string` | Required | Device serial number for the claimed inventory item |
| `Type` | `string` | Required | Device type for the claimed inventory item |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseClaimLicenseInventoryItem := models.ResponseClaimLicenseInventoryItem{
        Mac:                  "mac4",
        Magic:                "magic0",
        Model:                "model8",
        Serial:               "serial0",
        Type:                 "type0",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

