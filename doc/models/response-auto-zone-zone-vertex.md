
# Response Auto Zone Zone Vertex

Pixel coordinate for a suggested zone vertex

## Structure

`ResponseAutoZoneZoneVertex`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `X` | `*int` | Optional | Horizontal pixel coordinate of the vertex |
| `Y` | `*int` | Optional | Vertical pixel coordinate of the vertex |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoZoneZoneVertex := models.ResponseAutoZoneZoneVertex{
        X:                    models.ToPointer(10),
        Y:                    models.ToPointer(42),
    }

}
```

