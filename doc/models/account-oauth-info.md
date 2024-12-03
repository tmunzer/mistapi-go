
# Account Oauth Info

OAuth linked apps (zoom/teams/intune) account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountOauthInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `*string` | Optional | Linked app(zoom/teams/intune) account id |
| `Company` | `*string` | Optional | Name of the company whose account mist has subscribed to |
| `Error` | `*string` | Optional | This error is provided when the account fails to fetch token/data |
| `Errors` | `[]string` | Optional | - |
| `LastStatus` | `*string` | Optional | Is the last data pull for account is successful or not |
| `LastSync` | `*int64` | Optional | Last data pull timestamp, background jobs that pull account data |
| `LinkedBy` | `*string` | Optional | First name of the user who linked the account |
| `MaxDailyApiRequests` | `*int` | Optional | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
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
  "max_daily_api_requests": 5000,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

