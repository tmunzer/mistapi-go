
# Utils Cable Tests

Request body for running a switch cable test on a port

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsCableTests`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Port` | `string` | Required | The port to run the cable test |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsCableTests := models.UtilsCableTests{
        Port:                 "port0",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

