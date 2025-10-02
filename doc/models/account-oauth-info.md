
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
      "auto_probe_subnet": "11.0.0.0/8",
      "cloud_name": "Tapi.sase.paloaltonetworks.com",
      "company": "Test Company1 Ltd",
      "enable_probe": false,
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
      "region": "americas",
      "service_account_name": "Corp SA",
      "smartgroup_name": "CompliantGroup1",
      "tsg_id": "189953456",
      "webhook_auth_type": "Basic",
      "webhook_password": "password_1234",
      "webhook_url": "https://websync.nac-staging.mistsys.com/v1/S_41b2525af1d8dcbe9005/f43ea4c48f22/vmware/mdm",
      "webhook_username": "username_1234",
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

