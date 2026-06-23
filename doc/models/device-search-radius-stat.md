
# Device Search Radius Stat

RADIUS authentication counters and server status for a device search result

## Structure

`DeviceSearchRadiusStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAccepts` | `*int` | Optional | Number of accepted authentication requests |
| `AuthRejects` | `*int` | Optional | Number of rejected authentication requests |
| `AuthServerStatus` | [`*models.DeviceSearchRadiusFilterStatusEnum`](../../doc/models/device-search-radius-filter-status-enum.md) | Optional | Status of the device search RADIUS filter. enum: `up`, `down`, `unreachable` |
| `AuthTimeouts` | `*int` | Optional | Number of authentication timeouts |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    deviceSearchRadiusStat := models.DeviceSearchRadiusStat{
        AuthAccepts:          models.ToPointer(106),
        AuthRejects:          models.ToPointer(92),
        AuthServerStatus:     models.ToPointer(models.DeviceSearchRadiusFilterStatusEnum_UNREACHABLE),
        AuthTimeouts:         models.ToPointer(230),
    }

}
```

