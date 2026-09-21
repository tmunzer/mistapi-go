
# Response Client Sessions Search Item

Wireless client session record

## Structure

`ResponseClientSessionsSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required, Read-only | Access point MAC address associated with the client session |
| `Band` | `string` | Required, Read-only | Radio band used by the client session |
| `ClientManufacture` | `string` | Required, Read-only | Manufacturer reported for the client device |
| `Connect` | `float64` | Required, Read-only | Epoch timestamp when the client session connected |
| `Disconnect` | `float64` | Required, Read-only | Epoch timestamp when the client session disconnected |
| `Duration` | `float64` | Required, Read-only | Length of the client session, in seconds |
| `Mac` | `string` | Required, Read-only | Client MAC address for the session |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Ssid` | `string` | Required, Read-only | WLAN SSID used by the client session |
| `Tags` | `[]string` | Optional, Read-only | Tags attached to a wireless client session |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `WlanId` | `uuid.UUID` | Required | WLAN identifier associated with the client session |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseClientSessionsSearchItem := models.ResponseClientSessionsSearchItem{
        Ap:                   "",
        Band:                 "",
        ClientManufacture:    "",
        Connect:              0.0,
        Disconnect:           0.0,
        Duration:             0.0,
        Mac:                  "",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Ssid:                 "",
        Timestamp:            0.0,
        WlanId:               uuid.MustParse("0000154c-0000-0000-0000-000000000000"),
    }

}
```

