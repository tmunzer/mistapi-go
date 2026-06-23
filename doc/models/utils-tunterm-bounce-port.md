
# Utils Tunterm Bounce Port

Request to bounce Mist Edge TunTerm data ports

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsTuntermBouncePort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `HoldTime` | `*int` | Optional | In milli seconds, hold time between multiple port bounces |
| `Ports` | `[]string` | Required | TunTerm data port identifiers to bounce |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsTuntermBouncePort := models.UtilsTuntermBouncePort{
        HoldTime:             models.ToPointer(140),
        Ports:                []string{
            "ports5",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

