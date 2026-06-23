
# Response Async License Detail

Per-device asynchronous license claim status

## Structure

`ResponseAsyncLicenseDetail`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Device MAC address for this license claim detail |
| `Status` | `*string` | Optional | Claim processing state for this device |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseAsyncLicenseDetail := models.ResponseAsyncLicenseDetail{
        Mac:                  models.ToPointer("mac6"),
        Status:               models.ToPointer("status4"),
        Timestamp:            models.ToPointer(float64(45.4)),
    }

}
```

