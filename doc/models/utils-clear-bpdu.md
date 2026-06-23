
# Utils Clear Bpdu

Request to clear detected BPDU errors on switch ports

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearBpdu`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | List of ports on which to clear the detected BPDU error, or `all` for all ports |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsClearBpdu := models.UtilsClearBpdu{
        Ports:                []string{
            "ports3",
            "ports4",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

