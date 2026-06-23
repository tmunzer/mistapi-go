
# Synthetictest Radius Server

Request body for testing RADIUS server availability from a switch

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestRadiusServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | Specify the password associated with the username |
| `Profile` | `*string` | Optional | Specify the access profile associated with the subscriber<br><br>**Default**: `"dot1x"` |
| `User` | `string` | Required | Specify the subscriber username to test |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestRadiusServer := models.SynthetictestRadiusServer{
        Password:             "password0",
        Profile:              models.ToPointer("dot1x"),
        User:                 "user6",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

