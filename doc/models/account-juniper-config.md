
# Account Juniper Config

Juniper account credentials used to link the integration

*This model accepts additional fields of type interface{}.*

## Structure

`AccountJuniperConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | Authentication password for the Juniper account |
| `Username` | `string` | Required | Customer account user name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountJuniperConfig := models.AccountJuniperConfig{
        Password:             "password",
        Username:             "john@nmo.com",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

