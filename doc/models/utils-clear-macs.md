
# Utils Clear Macs

Request to clear learned MAC addresses from device ports

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearMacs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | Device ports on which to clear learned MAC addresses. Must include logical unit number |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsClearMacs := models.UtilsClearMacs{
        Ports:                []string{
            "ports3",
            "ports4",
            "ports5",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

