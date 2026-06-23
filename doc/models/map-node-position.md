
# Map Node Position

Position of a map path node

## Structure

`MapNodePosition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `X` | `float64` | Required | Horizontal coordinate of the map node |
| `Y` | `float64` | Required | Vertical coordinate of the map node |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapNodePosition := models.MapNodePosition{
        X:                    float64(746),
        Y:                    float64(104),
    }

}
```

