
# Account Zscaler Config

OAuth linked Zscaler apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountZscalerConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CloudName` | `string` | Required | Zscaler Internet Access cloud name used for the integration |
| `PartnerKey` | `string` | Required | Zscaler partner key generated for the Mist integration |
| `Password` | `string` | Required | Credential password for the Zscaler partner administrator |
| `Username` | `string` | Required | Zscaler partner administrator username used by Mist |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountZscalerConfig := models.AccountZscalerConfig{
        CloudName:            "zscalerbeta.net",
        PartnerKey:           "K35vrZcK3JvrZc",
        Password:             "password",
        Username:             "john@nmo.com",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

