
# Webhook Site Sle Event

*This model accepts additional fields of type interface{}.*

## Structure

`WebhookSiteSleEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Sle` | [`*models.WebhookSiteSleEventSle`](../../doc/models/webhook-site-sle-event-sle.md) | Optional | - |
| `Timestamp` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "timestamp": 1694620800,
  "sle": {
    "ap-availability": 199.22,
    "successful-connect": 14.8,
    "time-to-connect": 125.56,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

