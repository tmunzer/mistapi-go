
# Response Oauth App Link

## Structure

`ResponseOauthAppLink`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accounts` | [`[]models.ResponseOauthAppLinkAccounts`](../../doc/models/containers/response-oauth-app-link-accounts.md) | Optional | This is Array of a container for any-of cases. |
| `Linked` | `*bool` | Optional | Basic Auth application linked status in mist portal enabled for VMware |

## Example (as JSON)

```json
{
  "accounts": [
    {
      "client_id": "client_id0",
      "error": "error2",
      "instance_url": "instance_url2",
      "last_status": "last_status2",
      "last_sync": 188
    },
    {
      "client_id": "client_id0",
      "error": "error2",
      "instance_url": "instance_url2",
      "last_status": "last_status2",
      "last_sync": 188
    }
  ],
  "linked": false
}
```

