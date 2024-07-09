
# Webhook Site Sle

## Structure

`WebhookSiteSle`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookSiteSleEvent`](../../doc/models/webhook-site-sle-event.md) | Optional | - |
| `Topic` | `*string` | Optional | - |

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
        "time-to-connect": 125.56
      },
      "timestamp": 130
    },
    {
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "sle": {
        "ap-availability": 199.22,
        "successful-connect": 14.8,
        "time-to-connect": 125.56
      },
      "timestamp": 130
    },
    {
      "org_id": "00000dbc-0000-0000-0000-000000000000",
      "site_id": "0000245a-0000-0000-0000-000000000000",
      "sle": {
        "ap-availability": 199.22,
        "successful-connect": 14.8,
        "time-to-connect": 125.56
      },
      "timestamp": 130
    }
  ]
}
```

