
# Account Oauth Info

*This model accepts additional fields of type interface{}.*

## Structure

`AccountOauthInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accounts` | [`[]models.AccountOauthInfoAccount`](../../doc/models/account-oauth-info-account.md) | Required | List of linked account details |
| `AuthorizationUrl` | `*string` | Optional | - |
| `Linked` | `bool` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "accounts": [
    {
      "account_id": "iojzXIJWEuiD73ZvydOfg",
      "company": "Test Compay1 Ltd",
      "error": "OAuth token refresh failed, please re-link your account",
      "errors": [
        "OAuth token refresh failed, please re-link your account",
        "API daily rate limit reached for your account"
      ],
      "last_status": "failed",
      "last_sync": 1665465339000,
      "linked_by": "Testname1",
      "linked_timestamp": 1665465339000,
      "max_daily_api_requests": 5000,
      "name": "Test Compay1 Ltd",
      "smartgroup_name": "CompliantGroup1",
      "client_id": "client_id2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "linked": false,
  "authorization_url": "authorization_url0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

