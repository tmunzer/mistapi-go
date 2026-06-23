
# Response Device Upgrade

Single-device upgrade status response

## Structure

`ResponseDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`models.UpgradeInfoStatusEnum`](../../doc/models/upgrade-info-status-enum.md) | Required | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseDeviceUpgrade := models.ResponseDeviceUpgrade{
        Status:               models.UpgradeInfoStatusEnum_SUCCESS,
        Timestamp:            float64(150.2),
    }

}
```

