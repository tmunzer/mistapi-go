
# Webhook Nac Accounting

nac-accounting webhook sample

## Structure

`WebhookNacAccounting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookNacAccountingEvent`](../../doc/models/webhook-nac-accounting-event.md) | Optional | - |
| `Topic` | `*string` | Optional | **Default**: `"nac-accounting"` |

## Example (as JSON)

```json
{
  "topic": "nac-accounting",
  "events": [
    {
      "ap": "ap6",
      "auth_type": "auth_type0",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6"
    },
    {
      "ap": "ap6",
      "auth_type": "auth_type0",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6"
    },
    {
      "ap": "ap6",
      "auth_type": "auth_type0",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6"
    }
  ]
}
```

