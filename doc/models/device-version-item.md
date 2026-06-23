
# Device Version Item

Available firmware version for a specific device model and release tag

## Structure

`DeviceVersionItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Model` | `string` | Required | Device model (as seen in the device stats) |
| `Tag` | `*string` | Optional | Annotation, stable / beta / alpha. Or it can be empty or nothing which is likely a dev build |
| `Version` | `string` | Required | Available firmware version for this device model |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    deviceVersionItem := models.DeviceVersionItem{
        Model:                "model0",
        Tag:                  models.ToPointer("tag6"),
        Version:              "version8",
    }

}
```

