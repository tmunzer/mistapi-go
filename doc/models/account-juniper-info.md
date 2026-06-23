
# Account Juniper Info

Linked Juniper account information returned by the integration

*This model accepts additional fields of type interface{}.*

## Structure

`AccountJuniperInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accounts` | [`[]models.JuniperAccount`](../../doc/models/juniper-account.md) | Optional | Linked Juniper accounts available to the organization |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountJuniperInfo := models.AccountJuniperInfo{
        Accounts:             []models.JuniperAccount{
            models.JuniperAccount{
                LinkedBy:             models.ToPointer("linked_by8"),
                Name:                 models.ToPointer("name0"),
            },
            models.JuniperAccount{
                LinkedBy:             models.ToPointer("linked_by8"),
                Name:                 models.ToPointer("name0"),
            },
            models.JuniperAccount{
                LinkedBy:             models.ToPointer("linked_by8"),
                Name:                 models.ToPointer("name0"),
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

