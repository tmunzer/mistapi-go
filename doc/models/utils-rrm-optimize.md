
# Utils Rrm Optimize

Request to optimize RRM for selected AP radio bands

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsRrmOptimize`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | `[]string` | Required | Radio bands to include in the RRM optimization |
| `Macs` | `[]string` | Optional | AP MAC addresses to target for optimization. Neighboring APs may also change; omitted or empty targets all APs |
| `TxpowerOnly` | `*bool` | Optional | When true, adjust only transmit power so clients are not disconnected<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsRrmOptimize := models.UtilsRrmOptimize{
        Bands:                []string{
            "bands6",
        },
        Macs:                 []string{
            "macs9",
            "macs0",
            "macs1",
        },
        TxpowerOnly:          models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

