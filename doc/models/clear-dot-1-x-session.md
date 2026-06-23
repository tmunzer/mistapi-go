
# Clear Dot 1 X Session

Request body for clearing dot1x sessions on switch ports

*This model accepts additional fields of type interface{}.*

## Structure

`ClearDot1xSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `[]string` | Optional | List of port IDs where the dot1x session must be cleared. Use `all` to clear sessions on all ports. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    clearDot1xSession := models.ClearDot1xSession{
        Ports:                []string{
            "ge-0/0/0",
            "ge-0/0/1",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

