
# Iotproxy

IoT proxy configuration for the site

## Structure

`Iotproxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether the site IoT proxy is enabled<br><br>**Default**: `false` |
| `Visionline` | [`*models.IotproxyVisionline`](../../doc/models/iotproxy-visionline.md) | Optional | Visionline integration settings for IoT proxy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    iotproxy := models.Iotproxy{
        Enabled:              models.ToPointer(false),
        Visionline:           models.ToPointer(models.IotproxyVisionline{
            AccessId:             models.ToPointer("access_id8"),
            Cacerts:              []string{
                "cacerts4",
                "cacerts3",
                "cacerts2",
            },
            Enabled:              models.ToPointer(false),
            Host:                 models.ToPointer("host4"),
            Password:             models.ToPointer("password8"),
        }),
    }

}
```

