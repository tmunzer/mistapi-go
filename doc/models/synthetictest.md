
# Synthetictest

Request body for triggering a site synthetic test

*This model accepts additional fields of type interface{}.*

## Structure

`Synthetictest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `*string` | Optional | Contact email address supplied when triggering the site synthetic test |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictest := models.Synthetictest{
        Email:                models.ToPointer("test@mist.com"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

