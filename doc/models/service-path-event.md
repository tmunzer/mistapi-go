
# Service Path Event

Event describing a service path state change reported by a gateway or SSR device

## Structure

`ServicePathEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Device MAC address that reported the service path event |
| `Model` | `*string` | Optional | Device model that reported the service path event |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Policy` | `*string` | Optional | Service policy name associated with the path event |
| `PortId` | `*string` | Optional | Network interface associated with the service path event |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Text` | `*string` | Optional | Human-readable message for the service path event |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | Event type for the service path change |
| `Version` | `*string` | Optional | Device firmware version that reported the service path event |
| `VpnName` | `*string` | Optional | Peer name associated with the service path event |
| `VpnPath` | `*string` | Optional | Peer path name associated with the service path event |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    servicePathEvent := models.ServicePathEvent{
        Mac:                  models.ToPointer("90ec7734b374"),
        Model:                models.ToPointer("SSR120"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Policy:               models.ToPointer("INTERNET"),
        PortId:               models.ToPointer("ge-1/0/6"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Text:                 models.ToPointer("Peer Path Down"),
        Type:                 models.ToPointer("GW_SERVICE_PATH_REMOVE"),
        Version:              models.ToPointer("6.1.5-14.lts"),
        VpnName:              models.ToPointer("Syracuse_HUB"),
        VpnPath:              models.ToPointer("Syracuse_HUB-Wan0"),
    }

}
```

