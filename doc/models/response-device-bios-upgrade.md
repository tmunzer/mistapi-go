
# Response Device Bios Upgrade

Device BIOS upgrade status response

## Structure

`ResponseDeviceBiosUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | `*string` | Optional | Current BIOS upgrade status for the device |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseDeviceBiosUpgrade := models.ResponseDeviceBiosUpgrade{
        Status:               models.ToPointer("status0"),
        Timestamp:            models.ToPointer(float64(153.66)),
    }

}
```

