
# Webhook Site Sle Event

## Structure

`WebhookSiteSleEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Sle` | [`*models.WebhookSiteSleEventSle`](../../doc/models/webhook-site-sle-event-sle.md) | Optional | - |
| `Timestamp` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "52b50564-8821-4c3e-97be-5061c7760002",
  "timestamp": 1694620800,
  "site_id": "00002394-0000-0000-0000-000000000000",
  "sle": {
    "ap-availability": 199.22,
    "successful-connect": 14.8,
    "time-to-connect": 125.56
  }
}
```

