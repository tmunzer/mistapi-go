
# Device Event

## Structure

`DeviceEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | (will be deprecated soon; please use mac instead) ap mac |
| `ApName` | `*string` | Optional | (will be deprecated soon; please use device_name instead) ap name |
| `Apfw` | `*string` | Optional | - |
| `AuditId` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Bandwidth` | `*int` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `ChassisMac` | `*string` | Optional | - |
| `Count` | `*int` | Optional | - |
| `DeviceName` | `*string` | Optional | Device name |
| `DeviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `EvType` | [`*models.WebhookDeviceEventsEventEvTypeEnum`](../../doc/models/webhook-device-events-event-ev-type-enum.md) | Optional | (optional) event advisory. enum: `notice`, `warn` |
| `ExtIp` | `*string` | Optional | - |
| `Mac` | `*string` | Optional | Device mac |
| `Model` | `*string` | Optional | - |
| `Node` | `*string` | Optional | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `PortId` | `*string` | Optional | - |
| `Power` | `*int` | Optional | - |
| `PreBandwidth` | `*int` | Optional | - |
| `PreChannel` | `*int` | Optional | - |
| `PrePower` | `*int` | Optional | - |
| `PreUsage` | `*int` | Optional | - |
| `Reason` | `*string` | Optional | (optional) event reason |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SiteName` | `*string` | Optional | Site name |
| `Text` | `*string` | Optional | (optional) event description |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Type` | `string` | Required | Event type |
| `Usage` | `*int` | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "audit_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 63.58,
  "type": "type6",
  "ap": "ap0",
  "ap_name": "ap_name8",
  "apfw": "apfw0",
  "bandwidth": 232
}
```

