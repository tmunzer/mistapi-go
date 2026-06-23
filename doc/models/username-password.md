
# Username Password

Credential payload used to validate an identity provider login

*This model accepts additional fields of type interface{}.*

## Structure

`UsernamePassword`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `*string` | Optional | Credential password used for the validation attempt |
| `Username` | `*string` | Optional | Credential username used for the validation attempt |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    usernamePassword := models.UsernamePassword{
        Password:             models.ToPointer("password6"),
        Username:             models.ToPointer("username2"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

