
# Webhook Audits

Sample of the `audits` webhook payload.

## Structure

`WebhookAudits`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.LogEvent`](../../doc/models/log-event.md) | Required | **Constraints**: *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `audits`<br><br>**Value**: `"audits"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "message": "message0",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "timestamp": 188.18,
      "admin_id": "0000104e-0000-0000-0000-000000000000",
      "admin_name": "admin_name8",
      "after": {
        "key1": "val1",
        "key2": "val2"
      },
      "before": {
        "key1": "val1",
        "key2": "val2"
      },
      "for_site": false
    }
  ],
  "topic": "audits"
}
```

