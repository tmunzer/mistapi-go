
# Code String

Request body containing a single authorization or claim code

*This model accepts additional fields of type interface{}.*

## Structure

`CodeString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | Request-supplied authorization or claim code value |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    codeString := models.CodeString{
        Code:                 "code4",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

