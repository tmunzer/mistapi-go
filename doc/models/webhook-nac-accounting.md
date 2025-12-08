
# Webhook Nac Accounting

Sample of the `nac-accounting` webhook payload.

## Structure

`WebhookNacAccounting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookNacAccountingEvent`](../../doc/models/webhook-nac-accounting-event.md) | Optional | - |
| `Topic` | [`*models.WebhookNacAccountingTopicEnum`](../../doc/models/webhook-nac-accounting-topic-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6"
    },
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6"
    },
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6"
    }
  ],
  "topic": "nac-accounting"
}
```

