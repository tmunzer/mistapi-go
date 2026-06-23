
# Devices Gbp Tag

Request body for assigning a GBP tag to multiple devices

*This model accepts additional fields of type interface{}.*

## Structure

`DevicesGbpTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GbpTag` | `int` | Required | Group-Based Policy tag value to apply to the devices |
| `Macs` | `[]string` | Required | Unique MAC addresses included in a request<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    devicesGbpTag := models.DevicesGbpTag{
        GbpTag:               12,
        Macs:                 []string{
            "683b679ac024",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

