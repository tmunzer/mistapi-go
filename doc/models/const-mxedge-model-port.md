
# Const Mxedge Model Port

Mist Edge model port metadata

## Structure

`ConstMxedgeModelPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `*string` | Optional | User-facing interface name for this Mist Edge port |
| `Speed` | `*int` | Optional | Port speed for the Mist Edge interface, in Mbps |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constMxedgeModelPort := models.ConstMxedgeModelPort{
        Display:              models.ToPointer("xe0"),
        Speed:                models.ToPointer(10000),
    }

}
```

