
# Webhook Audits

audit webhook sample

## Structure

`WebhookAudits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookAuditEvent`](../../doc/models/webhook-audit-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"audits"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "admin_name": "admin_name8",
      "device_id": "0000254a-0000-0000-0000-000000000000",
      "id": "00000ce4-0000-0000-0000-000000000000",
      "message": "message0",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "src_ip": "src_ip4",
      "timestamp": 188.18
    }
  ],
  "topic": "audits"
}
```

