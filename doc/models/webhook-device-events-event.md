
# Webhook Device Events Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookDeviceEventsEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | (will be deprecated soon; please use mac instead) ap mac |
| `ApName` | `*string` | Optional | (will be deprecated soon; please use device_name instead) ap name |
| `AuditId` | `*uuid.UUID` | Optional | (optional) audit id |
| `DeviceName` | `string` | Required | Device name |
| `DeviceType` | [`models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Required | enum: `ap`, `gateway`, `switch` |
| `EvType` | [`models.WebhookDeviceEventsEventEvTypeEnum`](../../doc/models/webhook-device-events-event-ev-type-enum.md) | Required | (optional) event advisory. enum: `notice`, `warn` |
| `Mac` | `string` | Required | Device mac |
| `OrgId` | `uuid.UUID` | Required | - |
| `Reason` | `*string` | Optional | (optional) event reason |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SiteName` | `*string` | Optional | Site name |
| `Text` | `*string` | Optional | (optional) event description |
| `Timestamp` | `float64` | Required | Epoch (seconds) |
| `Type` | `string` | Required | Event type |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "device_name": "device_name4",
  "device_type": "gateway",
  "ev_type": "notice",
  "mac": "mac2",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 184.26,
  "type": "type2",
  "ap": "ap4",
  "ap_name": "ap_name6",
  "audit_id": "000005d2-0000-0000-0000-000000000000",
  "reason": "reason6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

