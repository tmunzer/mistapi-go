
# Map Site Replace File Json Transform

If `transform` is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)

## Structure

`MapSiteReplaceFileJsonTransform`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Rotation` | `*float64` | Optional | Whether to rotate the replacing image, in degrees<br><br>**Default**: `0` |
| `Scale` | `*float64` | Optional | Whether to scale the replacing image<br><br>**Default**: `1` |
| `X` | `*float64` | Optional | Where the (0, 0) of the new image is relative to the original map<br><br>**Default**: `0` |
| `Y` | `*float64` | Optional | Where the (0, 0) of the new image is relative to the original map<br><br>**Default**: `0` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapSiteReplaceFileJsonTransform := models.MapSiteReplaceFileJsonTransform{
        Rotation:             models.ToPointer(float64(0)),
        Scale:                models.ToPointer(float64(0.98)),
        X:                    models.ToPointer(float64(3.16)),
        Y:                    models.ToPointer(float64(12)),
    }

}
```

