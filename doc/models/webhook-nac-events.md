
# Webhook Nac Events

Sample of the `nac-events` webhook payload.

## Structure

`WebhookNacEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.NacClientEvent`](../../doc/models/nac-client-event.md) | Optional | - |
| `Topic` | [`*models.WebhookNacEventsTopicEnum`](../../doc/models/webhook-nac-events-topic-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "events": [
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_type": "wired",
      "device_mac": "device_mac4"
    },
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_type": "wired",
      "device_mac": "device_mac4"
    }
  ],
  "topic": "nac-events"
}
```

