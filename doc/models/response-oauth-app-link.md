
# Response Oauth App Link

## Structure

`ResponseOauthAppLink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accounts` | [`[]models.ResponseOauthAppLinkItem2`](../../doc/models/containers/response-oauth-app-link-item-2.md) | Optional | List of linked account details |
| `Linked` | `*bool` | Optional | Basic Auth application linked status in mist portal enabled for VMware |

## Example (as JSON)

```json
{
  "accounts": [
    {
      "client_id": "client_id6",
      "error": "error8",
      "instance_url": "instance_url6",
      "last_status": "last_status8",
      "last_sync": 168
    },
    {
      "client_id": "client_id6",
      "error": "error8",
      "instance_url": "instance_url6",
      "last_status": "last_status8",
      "last_sync": 168
    }
  ],
  "linked": false
}
```

