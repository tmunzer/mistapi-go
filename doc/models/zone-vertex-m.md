
# Zone Vertex M

Zone polygon vertex expressed in meters

## Structure

`ZoneVertexM`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `X` | `float64` | Required | Horizontal coordinate of the zone vertex, in meters |
| `Y` | `float64` | Required | Vertical coordinate of the zone vertex, in meters |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    zoneVertexM := models.ZoneVertexM{
        X:                    float64(161.44),
        Y:                    float64(36.72),
    }

}
```

