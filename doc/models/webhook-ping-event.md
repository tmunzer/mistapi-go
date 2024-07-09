
# Webhook Ping Event

## Structure

`WebhookPingEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `SiteId` | `uuid.UUID` | Required | - |
| `Timestamp` | `float64` | Required | - |

## Example (as JSON)

```json
{
  "id": "0000020e-0000-0000-0000-000000000000",
  "name": "name6",
  "site_id": "43e9c864-a7e4-4310-8031-d9817d2c5a43",
  "timestamp": 83.56
}
```

