
# Search Webhook Delivery

## Structure

`SearchWebhookDelivery`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | - |
| `Limit` | `*int` | Optional | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.WebhookDelivery`](../../doc/models/webhook-delivery.md) | Optional | - |
| `Start` | `*int` | Optional | - |
| `Total` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "end": 1688035193,
  "limit": 10,
  "start": 1687948793,
  "next": "next0",
  "results": [
    {
      "error": "error0",
      "id": "000023ba-0000-0000-0000-000000000000",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "req_headers": "req_headers6",
      "req_payload": "req_payload4"
    },
    {
      "error": "error0",
      "id": "000023ba-0000-0000-0000-000000000000",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "req_headers": "req_headers6",
      "req_payload": "req_payload4"
    },
    {
      "error": "error0",
      "id": "000023ba-0000-0000-0000-000000000000",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "req_headers": "req_headers6",
      "req_payload": "req_payload4"
    }
  ]
}
```

