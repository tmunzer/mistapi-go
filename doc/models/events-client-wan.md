
# Events Client Wan

WAN client event returned by WAN client event search APIs

## Structure

`EventsClientWan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `When` | `*string` | Optional | Event timestamp string reported by the WAN client event source |
| `EvType` | `*string` | Optional | WAN client event type identifier |
| `Metadata` | `*interface{}` | Optional | Additional attributes provided with the WAN client event |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RandomMac` | `*bool` | Optional | Whether the WAN client used a randomized MAC address |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Text` | `*string` | Optional | Human-readable message for the WAN client event |
| `Wcid` | `*uuid.UUID` | Optional | WAN client identifier associated with the event |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    eventsClientWan := models.EventsClientWan{
        When:                 models.ToPointer("2022-12-31 23:59:59.293000+00:00"),
        EvType:               models.ToPointer("CLIENT_IP_ASSIGNED"),
        Metadata:             models.ToPointer(interface{}("[key1, val1][key2, val2]")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        RandomMac:            models.ToPointer(false),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Text:                 models.ToPointer("DHCP Ack IP 192.168.88.216"),
        Wcid:                 models.ToPointer(uuid.MustParse("62bbfb75-10d8-49d1-dec7-d2df91624287")),
    }

}
```

