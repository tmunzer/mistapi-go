
# Webhook Nac Accounting

nac-accounting webhook sample

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookNacAccounting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookNacAccountingEvent`](../../doc/models/webhook-nac-accounting-event.md) | Optional | - |
| `Topic` | `*string` | Optional | **Default**: `"nac-accounting"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "topic": "nac-accounting",
  "events": [
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "ap": "ap6",
      "auth_type": "cert",
      "bssid": "bssid4",
      "client_ip": "client_ip0",
      "client_type": "client_type6",
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

