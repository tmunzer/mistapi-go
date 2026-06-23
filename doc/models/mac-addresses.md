
# Mac Addresses

Request containing one or more MAC addresses

*This model accepts additional fields of type interface{}.*

## Structure

`MacAddresses`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Macs` | `[]string` | Required | Unique MAC addresses included in a request<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    macAddresses := models.MacAddresses{
        Macs:                 []string{
            "683b679ac024",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

