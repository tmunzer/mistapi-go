
# Const Mxedge Model

Mist Edge model definition returned by the constants API

## Structure

`ConstMxedgeModel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomPorts` | `*bool` | Optional | Whether this Mist Edge model supports custom port definitions |
| `Display` | `*string` | Optional | User-facing model name shown for the Mist Edge |
| `Model` | `*string` | Optional | Mist Edge model identifier |
| `Ports` | [`map[string]models.ConstMxedgeModelPort`](../../doc/models/const-mxedge-model-port.md) | Optional | Port metadata keyed by interface number for this Mist Edge model |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constMxedgeModel := models.ConstMxedgeModel{
        CustomPorts:          models.ToPointer(false),
        Display:              models.ToPointer("X10"),
        Model:                models.ToPointer("ME-X10"),
        Ports:                map[string]models.ConstMxedgeModelPort{
            "key0": models.ConstMxedgeModelPort{
                Display:              models.ToPointer("display2"),
                Speed:                models.ToPointer(14),
            },
        },
    }

}
```

