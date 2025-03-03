
# Const Webhook Topic

*This model accepts additional fields of type interface{}.*

## Structure

`ConstWebhookTopic`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowsSingleEventPerMessage` | `*bool` | Optional | supports single event per message results |
| `ForOrg` | `*bool` | Optional | Can be used in org webhooks, optional |
| `HasDeliveryResults` | `*bool` | Optional | Supports webhook delivery results /api/v1/:scope/:scope_id/webhooks/:webhook_id/events/search |
| `Internal` | `*bool` | Optional | Internal topic (not selectable in site/org webhooks) |
| `Key` | `*string` | Optional | Webhook topic name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "key": "alarms",
  "allows_single_event_per_message": false,
  "for_org": false,
  "has_delivery_results": false,
  "internal": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

