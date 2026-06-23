
# Psk Portal Image

PSK portal image upload payload

## Structure

`PskPortalImage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*[]byte` | Optional | Image binary payload to upload for the PSK portal |
| `Json` | `*string` | Optional | Metadata JSON string describing the PSK portal image upload |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pskPortalImage := models.PskPortalImage{
        File:                 models.ToPointer(nil),
        Json:                 models.ToPointer("json6"),
    }

}
```

