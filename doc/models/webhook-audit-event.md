
# Webhook Audit Event

## Structure

`WebhookAuditEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminName` | `string` | Required | - |
| `DeviceId` | `uuid.UUID` | Required | - |
| `Id` | `uuid.UUID` | Required | - |
| `Message` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `SrcIp` | `string` | Required | - |
| `Timestamp` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "admin_name": "admin_name8",
  "device_id": "00000920-0000-0000-0000-000000000000",
  "id": "000017ca-0000-0000-0000-000000000000",
  "message": "message0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "src_ip": "src_ip6",
  "timestamp": 72.08
}
```

