
# Webhook Audits

audit webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookAudits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookAuditEvent`](../../doc/models/webhook-audit-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required | **Default**: `"audits"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "admin_name": "admin_name8",
      "device_id": "00000000-0000-0000-1000-d8695a0f9e61",
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "message": "message0",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "src_ip": "src_ip4",
      "timestamp": 188.18,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "topic": "audits",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

