
# Mxedge Event

## Structure

`MxedgeEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuditId` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Component` | `*string` | Optional | - |
| `DeviceType` | `*string` | Optional | - |
| `FromVersion` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `MxclusterId` | `*string` | Optional | - |
| `MxedgeId` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Package` | `*string` | Optional | - |
| `Service` | `*string` | Optional | - |
| `Severity` | [`*models.EventSeverityEnum`](../../doc/models/event-severity-enum.md) | Optional | - |
| `SysInfoUsage` | [`*models.MxedgeEventSysInfo`](../../doc/models/mxedge-event-sys-info.md) | Optional | - |
| `Text` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `ToVersion` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "audit_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "component": "PS1",
  "mxcluster_id": "2815c917-58e7-472f-a190-bfd44fb58d05",
  "mxedge_id": "00000000-0000-0000-1000-020000dc585c",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "service": "tunterm",
  "type": "ME_SERVICE_STOPPED",
  "device_type": "device_type2",
  "from_version": "from_version0",
  "mac": "mac2"
}
```

