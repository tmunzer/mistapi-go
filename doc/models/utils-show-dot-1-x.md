
# Utils Show Dot 1 X

802.1X table lookup request for device command output

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowDot1x`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Refresh duration in seconds; set only when `interval` is nonzero<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | Refresh interval in seconds for repeated command output<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `PortId` | `*string` | Optional | Device port identifier filter for the 802.1X table lookup |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowDot1x := models.UtilsShowDot1x{
        Duration:             models.ToPointer(0),
        Interval:             models.ToPointer(0),
        PortId:               models.ToPointer("ge-0/0/0.0"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

