
# Response Auto Zone Zone

Suggested zone returned by the auto zones service

## Structure

`ResponseAutoZoneZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Human-readable name of the suggested zone |
| `Vertices` | [`[]models.ResponseAutoZoneZoneVertex`](../../doc/models/response-auto-zone-zone-vertex.md) | Optional | List of points comprising the suggested zone polygon in map pixels |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAutoZoneZone := models.ResponseAutoZoneZone{
        Name:                 models.ToPointer("zone1"),
        Vertices:             []models.ResponseAutoZoneZoneVertex{
            models.ResponseAutoZoneZoneVertex{
                X:                    models.ToPointer(16),
                Y:                    models.ToPointer(88),
            },
        },
    }

}
```

