
# Response Http 400 Webhook Exception

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp400WebhookException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | - |
| `Reason` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "detail": "invalid field: assetfilter_ids",
  "reason": "contains duplicate uuids",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

