
# Const Webhook Topic

## Structure

`ConstWebhookTopic`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForOrg` | `*bool` | Optional | can be used in org webhooks, optional |
| `HasDeliveryResults` | `*bool` | Optional | supports webhook delivery results /api/v1/:scope/:scope_id/webhooks/:webhook_id/events/search |
| `Internal` | `*bool` | Optional | internal topic (not selectable in site/org webhooks) |
| `Key` | `*string` | Optional | webhook topic name |

## Example (as JSON)

```json
{
  "key": "alarms",
  "for_org": false,
  "has_delivery_results": false,
  "internal": false
}
```

