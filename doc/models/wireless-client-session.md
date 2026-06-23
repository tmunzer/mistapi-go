
# Wireless Client Session

Wireless client session record returned by a session search

## Structure

`WirelessClientSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `string` | Required, Read-only | MAC address of the access point associated with the client session |
| `Band` | `string` | Required, Read-only | Radio band used by the client session |
| `ClientManufacture` | `models.Optional[string]` | Optional, Read-only | Manufacturer reported for the client device, when available |
| `Connect` | `int` | Required, Read-only | Time when the client session connected, in epoch seconds |
| `Disconnect` | `int` | Required, Read-only | Time when the client session disconnected, in epoch seconds |
| `Duration` | `float64` | Required, Read-only | Length of the client session, in seconds |
| `ForSite` | `*bool` | Optional, Read-only | Whether this client session record is scoped to a site |
| `Mac` | `string` | Required, Read-only | Client MAC address for the session |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Ssid` | `string` | Required, Read-only | WLAN SSID used by the client session |
| `Tags` | `[]string` | Optional, Read-only | Tags attached to a wireless client session |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `WlanId` | `uuid.UUID` | Required, Read-only | Identifier of the WLAN associated with the client session |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    wirelessClientSession := models.WirelessClientSession{
        Ap:                   "ap2",
        Band:                 "band8",
        ClientManufacture:    models.NewOptional(models.ToPointer("client_manufacture0")),
        Connect:              50,
        Disconnect:           0,
        Duration:             float64(161.82),
        ForSite:              models.ToPointer(false),
        Mac:                  "mac0",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Ssid:                 "ssid4",
        Tags:                 []string{
            "tags1",
        },
        Timestamp:            float64(60.04),
        WlanId:               uuid.MustParse("00000c7e-0000-0000-0000-000000000000"),
    }

}
```

