
# Map Site Replace File

Multipart payload for replacing a site map image

## Structure

`MapSiteReplaceFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `[]byte` | Required | Map image file used to replace the existing site map |
| `Json` | [`*models.MapSiteReplaceFileJson`](../../doc/models/map-site-replace-file-json.md) | Optional | Options for replacing a site map image |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mapSiteReplaceFile := models.MapSiteReplaceFile{
        File:                 nil,
        Json:                 models.ToPointer(models.MapSiteReplaceFileJson{
            Transform:            models.ToPointer(models.MapSiteReplaceFileJsonTransform{
                Rotation:             models.ToPointer(float64(130.62)),
                Scale:                models.ToPointer(float64(195.12)),
                X:                    models.ToPointer(float64(28.38)),
                Y:                    models.ToPointer(float64(102.9)),
            }),
        }),
    }

}
```

