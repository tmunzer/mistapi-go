
# Zone Vertex

Zone polygon vertex expressed in map pixels

## Structure

`ZoneVertex`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `X` | `float64` | Required | Horizontal pixel coordinate of the zone vertex |
| `Y` | `float64` | Required | Vertical pixel coordinate of the zone vertex |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    zoneVertex := models.ZoneVertex{
        X:                    float64(48.32),
        Y:                    float64(82.96),
    }

}
```

