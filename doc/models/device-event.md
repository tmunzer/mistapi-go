
# Device Event

Device event payload returned by search and webhook APIs

## Structure

`DeviceEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | Deprecated AP MAC address field; use `mac` instead |
| `ApName` | `*string` | Optional | Deprecated AP name field; use `device_name` instead |
| `Apfw` | `*string` | Optional | AP firmware version associated with the device event |
| `AuditId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Bandwidth` | `*int` | Optional | Channel bandwidth associated with a radio event, in MHz |
| `Channel` | `*int` | Optional | RF channel associated with the device event |
| `ChassisMac` | `*string` | Optional | Chassis MAC address associated with the device event |
| `Count` | `*int` | Optional | Event count reported in the device event payload |
| `DeviceName` | `*string` | Optional | Name of the device associated with the event |
| `DeviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `EvType` | [`*models.WebhookDeviceEventsEventEvTypeEnum`](../../doc/models/webhook-device-events-event-ev-type-enum.md) | Optional | (optional) event advisory. enum: `notice`, `warn` |
| `ExtIp` | `*string` | Optional | External IP address reported for the device event |
| `Mac` | `*string` | Optional | Device MAC address associated with the event |
| `Model` | `*string` | Optional | Device model associated with the event |
| `Node` | `*string` | Optional | Cluster node identifier associated with the device event |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `PortId` | `*string` | Optional | Port identifier associated with the device event |
| `Power` | `*int` | Optional | Transmit power associated with a radio event |
| `PreBandwidth` | `*int` | Optional | Previous channel bandwidth before an RRM change, in MHz |
| `PreChannel` | `*int` | Optional | Previous RF channel before an RRM change |
| `PrePower` | `*int` | Optional | Previous transmit power before an RRM change |
| `PreUsage` | `*int` | Optional | Previous radio usage band before an RRM change |
| `Reason` | `*string` | Optional | Optional reason text reported for the device event |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SiteName` | `*string` | Optional | Name of the site associated with the event |
| `Text` | `*string` | Optional | Optional human-readable text for the device event |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required | Device event type key |
| `Usage` | `*int` | Optional | Current radio usage band for an RRM event |
| `Version` | `*string` | Optional | Firmware or software version associated with the device event |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    deviceEvent := models.DeviceEvent{
        Ap:                   models.ToPointer("ap2"),
        ApName:               models.ToPointer("ap_name0"),
        Apfw:                 models.ToPointer("apfw2"),
        AuditId:              models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Bandwidth:            models.ToPointer(188),
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Timestamp:            float64(122.9),
        Type:                 "type8",
    }

}
```

