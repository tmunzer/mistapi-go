
# Map Site Replace File Json

Options for replacing a site map image

## Structure

`MapSiteReplaceFileJson`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Transform` | [`*models.MapSiteReplaceFileJsonTransform`](../../doc/models/map-site-replace-file-json-transform.md) | Optional | If `transform` is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapSiteReplaceFileJson := models.MapSiteReplaceFileJson{
        Transform:            models.ToPointer(models.MapSiteReplaceFileJsonTransform{
            Rotation:             models.ToPointer(float64(130.62)),
            Scale:                models.ToPointer(float64(195.12)),
            X:                    models.ToPointer(float64(28.38)),
            Y:                    models.ToPointer(float64(102.9)),
        }),
    }

}
```

