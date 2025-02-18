
# Webhook Site Sle

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookSiteSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookSiteSleEvent`](../../doc/models/webhook-site-sle-event.md) | Optional | - |
| `Topic` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "topic": "site_sle",
  "events": [
    {
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "sle": {
        "ap-availability": 199.22,
        "successful-connect": 14.8,
        "time-to-connect": 125.56,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "timestamp": 130,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "sle": {
        "ap-availability": 199.22,
        "successful-connect": 14.8,
        "time-to-connect": 125.56,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "timestamp": 130,
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

