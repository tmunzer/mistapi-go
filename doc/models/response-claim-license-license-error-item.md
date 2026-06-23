
# Response Claim License License Error Item

License order that could not be claimed

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClaimLicenseLicenseErrorItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Order` | `string` | Required | License order number or claim identifier that failed |
| `Reason` | `string` | Required | Explanation of why the license order could not be claimed |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseClaimLicenseLicenseErrorItem := models.ResponseClaimLicenseLicenseErrorItem{
        Order:                "order0",
        Reason:               "reason2",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

