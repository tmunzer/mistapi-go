
# Const Marvisclient Event

A Marvis Client event type definition

## Structure

`ConstMarvisclientEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `*string` | Optional | Human-readable name for this Marvis Client event type |
| `Key` | `*string` | Optional | Event type key used in Marvis Client event search and count APIs |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constMarvisclientEvent := models.ConstMarvisclientEvent{
        Display:              models.ToPointer("Marvis Client Roamed"),
        Key:                  models.ToPointer("MARVISCLIENT_ROAMED"),
    }

}
```

