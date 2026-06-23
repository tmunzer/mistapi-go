
# Utils Bounce Port

Request to bounce one or more device ports

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsBouncePort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | Device port identifiers to bounce |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsBouncePort := models.UtilsBouncePort{
        Ports:                []string{
            "ports7",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

