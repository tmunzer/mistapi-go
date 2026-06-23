
# Response Claim License License Item

License entitlement returned by a claim operation

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClaimLicenseLicenseItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | Epoch timestamp when the license entitlement ends |
| `Quantity` | `int` | Required | Number of license units included in the entitlement |
| `Start` | `int` | Required | Epoch timestamp when the license entitlement starts |
| `Type` | `string` | Required | License SKU or subscription type |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseClaimLicenseLicenseItem := models.ResponseClaimLicenseLicenseItem{
        End:                  14,
        Quantity:             84,
        Start:                228,
        Type:                 "type0",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

