
# Mxedge Event

Event reported for a Mist Edge appliance or service

## Structure

`MxedgeEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuditId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Component` | `models.Optional[string]` | Optional | Mist Edge component that reported the event, such as a power supply or fan |
| `DeviceId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Identifier of the device associated with the event, when available |
| `DeviceType` | `*string` | Optional | Device type associated with the event |
| `FromVersion` | `*string` | Optional | Service version before an upgrade-related event |
| `Mac` | `*string` | Optional | Mist Edge MAC address associated with the event |
| `MxclusterId` | `*string` | Optional | Mist Edge cluster identifier associated with the event |
| `MxedgeId` | `*string` | Optional | Mist Edge identifier associated with the event |
| `MxedgeName` | `*string` | Optional | Display name of the Mist Edge associated with the event |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Package` | `*string` | Optional | Software package associated with the event |
| `Service` | `*string` | Optional | Mist Edge service associated with the event |
| `Severity` | [`*models.EventSeverityEnum`](../../doc/models/event-severity-enum.md) | Optional | Severity level for an event. enum: `normal`, `critical`, `high`, `warning` |
| `SysInfoUsage` | [`*models.MxedgeEventSysInfo`](../../doc/models/mxedge-event-sys-info.md) | Optional | System resource details for a Mist Edge event |
| `Text` | `*string` | Optional | Human-readable message describing the Mist Edge event |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `ToVersion` | `*string` | Optional | Service version after an upgrade-related event |
| `Type` | `*string` | Optional | Mist Edge event type code |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxedgeEvent := models.MxedgeEvent{
        AuditId:              models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Component:            models.NewOptional(models.ToPointer("PS1")),
        DeviceId:             models.NewOptional(models.ToPointer(uuid.MustParse("00000f40-0000-0000-0000-000000000000"))),
        DeviceType:           models.ToPointer("device_type2"),
        FromVersion:          models.ToPointer("from_version0"),
        MxclusterId:          models.ToPointer("2815c917-58e7-472f-a190-bfd44fb58d05"),
        MxedgeId:             models.ToPointer("00000000-0000-0000-1000-020000dc585c"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Service:              models.ToPointer("tunterm"),
        Type:                 models.ToPointer("ME_SERVICE_STOPPED"),
    }

}
```

