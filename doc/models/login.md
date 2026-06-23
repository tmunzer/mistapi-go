
# Login

Login request using email/password and an optional two-factor code

*This model accepts additional fields of type interface{}.*

## Structure

`Login`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `string` | Required | Administrator email address used to log in |
| `Password` | `string` | Required | Administrator password used to log in |
| `TwoFactor` | `*string` | Optional | Optional two-factor authentication code for the login request |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    login := models.Login{
        Email:                "test@mistsys.com",
        Password:             "foryoureyesonly",
        TwoFactor:            models.ToPointer("123456"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

