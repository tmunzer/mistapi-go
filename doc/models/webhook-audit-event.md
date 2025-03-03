
# Webhook Audit Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookAuditEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminName` | `string` | Required | - |
| `DeviceId` | `uuid.UUID` | Required | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Message` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `SrcIp` | `string` | Required | - |
| `Timestamp` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "admin_name": "admin_name8",
  "device_id": "00000000-0000-0000-1000-d8695a0f9e61",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "message": "message0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "src_ip": "src_ip6",
  "timestamp": 72.08,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

