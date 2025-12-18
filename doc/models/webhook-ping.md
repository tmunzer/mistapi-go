
# Webhook Ping

Sample of the `ping` webhook payload.\n\nThe `ping` webhook can be manually sent with the following API calls:\n- for a Site level webhook with the [Ping Site Webhook](../../doc/controllers/orgs-webhooks.md#ping-org-webhook) endpoint\n- for an Org level webhook with the [Ping Org Webhook](../../doc/controllers/orgs-webhooks.md#ping-org-webhook) endpoint

## Structure

`WebhookPing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookPingEvent`](../../doc/models/webhook-ping-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `ping`<br><br>**Value**: `"ping"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "name": "name0",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 188.18
    }
  ],
  "topic": "ping"
}
```

