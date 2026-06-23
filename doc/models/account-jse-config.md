
# Account Jse Config

Juniper Security Exchange account credentials used for integration

*This model accepts additional fields of type interface{}.*

## Structure

`AccountJseConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `*string` | Optional | JSE cloud hostname used for the integration |
| `Password` | `string` | Required | Credential password for the JSE integration user |
| `Username` | `string` | Required | JSE integration username with access to service location, site, and IPsec profile APIs |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountJseConfig := models.AccountJseConfig{
        CloudName:            models.ToPointer("devcentral.juniperclouds.net"),
        Password:             "foryoureyesonly",
        Username:             "john@abc.com",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

