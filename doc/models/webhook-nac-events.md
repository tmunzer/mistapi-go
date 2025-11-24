
# Webhook Nac Events

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookNacEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookNacEventsEvent`](../../doc/models/webhook-nac-events-event.md) | Optional | - |
| `Topic` | `*string` | Optional | **Default**: `"nac_events"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "topic": "nac_events",
  "events": [
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_type": "wired",
      "device_mac": "device_mac4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_type": "wired",
      "device_mac": "device_mac4",
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

