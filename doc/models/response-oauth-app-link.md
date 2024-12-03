
# Response Oauth App Link

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseOauthAppLink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accounts` | [`[]models.ResponseOauthAppLinkItem`](../../doc/models/containers/response-oauth-app-link-item.md) | Optional | List of linked account details |
| `Linked` | `*bool` | Optional | Basic Auth application linked status in mist portal enabled for VMware |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "accounts": [
    {
      "client_id": "client_id6",
      "error": "error8",
      "instance_url": "instance_url6",
      "last_status": "last_status8",
      "last_sync": 168,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "client_id": "client_id6",
      "error": "error8",
      "instance_url": "instance_url6",
      "last_status": "last_status8",
      "last_sync": 168,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "linked": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

