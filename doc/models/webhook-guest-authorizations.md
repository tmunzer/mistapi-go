
# Webhook Guest Authorizations

Sample of the `guest-authorizations` webhook payload.

## Structure

`WebhookGuestAuthorizations`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookGuestAuthorizationsEvent`](../../doc/models/webhook-guest-authorizations-event.md) | Optional | List of events |
| `Topic` | [`*models.WebhookGuestAuthorizationsTopicEnum`](../../doc/models/webhook-guest-authorizations-topic-enum.md) | Optional | enum: `guest-authorizations` |

## Example (as JSON)

```json
{
  "events": [
    {
      "ap": "ap6",
      "auth_method": "auth_method4",
      "authorized_expiring_time": 42,
      "authorized_time": 252,
      "carrier": "carrier2"
    }
  ],
  "topic": "guest-authorizations"
}
```

