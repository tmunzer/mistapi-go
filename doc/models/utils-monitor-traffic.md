
# Utils Monitor Traffic

Request body for monitoring traffic on one port or all ports

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsMonitorTraffic`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Port` | `*string` | Optional | Interface name to monitor; omit this field to monitor all ports |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsMonitorTraffic := models.UtilsMonitorTraffic{
        Port:                 models.ToPointer("ge-0/0/1"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

