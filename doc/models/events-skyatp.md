
# Events Skyatp

SkyATP threat event returned by SkyATP event search APIs

## Structure

`EventsSkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `string` | Required, Read-only | Network device MAC address that reported the SkyATP event |
| `ForSite` | `*bool` | Optional, Read-only | Whether the SkyATP event is scoped to a site rather than only the organization |
| `Ip` | `string` | Required, Read-only | Client IP address associated with the SkyATP event |
| `Mac` | `string` | Required, Read-only | Client MAC address associated with the SkyATP event |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `ThreatLevel` | `int` | Required, Read-only | Numeric SkyATP threat level reported for the event |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required, Read-only | SkyATP event type, such as `cc`, `fs`, or `mw` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    eventsSkyatp := models.EventsSkyatp{
        DeviceMac:            "device_mac8",
        ForSite:              models.ToPointer(false),
        Ip:                   "ip8",
        Mac:                  "mac8",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        ThreatLevel:          144,
        Timestamp:            float64(235.52),
        Type:                 "type6",
    }

}
```

