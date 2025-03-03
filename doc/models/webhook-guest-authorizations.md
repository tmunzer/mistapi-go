
# Webhook Guest Authorizations

guest-authorizations webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookGuestAuthorizations`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookGuestAuthorizationsEvent`](../../doc/models/webhook-guest-authorizations-event.md) | Optional | List of events |
| `Topic` | `*string` | Optional | **Default**: `"guest-authorizations"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "topic": "guest-authorizations",
  "events": [
    {
      "ap": "ap6",
      "auth_method": "auth_method4",
      "authorized_expiring_time": 42,
      "authorized_time": 252,
      "carrier": "carrier2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

