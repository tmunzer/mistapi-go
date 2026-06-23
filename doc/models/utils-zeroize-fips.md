
# Utils Zeroize Fips

Request body for confirming FIPS AP zeroize operations

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsZeroizeFips`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Password` | `string` | Required | Confirmation password for the FIPS zeroize operation |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsZeroizeFips := models.UtilsZeroizeFips{
        Password:             "password4",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

