
# Response Oauth App Link Item

## Structure

`ResponseOauthAppLinkItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `*string` | Optional | customer account client id |
| `Error` | `*string` | Optional | This error is provided when the Jamf account fails to fetch token/data |
| `InstanceUrl` | `*string` | Optional | customer account Jamf instance URL |
| `LastStatus` | `*string` | Optional | Is the last data pull for Jamf account is successful or not |
| `LastSync` | `*int64` | Optional | Last data pull timestamp, background jobs that pull Jamf account data |
| `LinkedBy` | `*string` | Optional | First name of the user who linked the Jamf account |
| `Name` | `*string` | Optional | Name of the company whose Jamf account mist has subscribed to |
| `SmartgroupName` | `*string` | Optional | smart group membership for determining compliance status |
| `AccountId` | `*string` | Optional | Linked app(zoom/teams/intune) account id |
| `Company` | `*string` | Optional | Name of the company whose account mist has subscribed to |
| `Errors` | `[]string` | Optional | - |
| `MaxDailyApiRequests` | `*int` | Optional | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |
| `LinkedTimestamp` | `*int` | Optional | This error is provided when the VMware account fails to fetch token/data |

## Example (as JSON)

```json
{
  "error": "OAuth token refresh failed, please re-link your account",
  "instance_url": "junipertest.jamfcloud.com",
  "last_status": "failed",
  "last_sync": 1665465339000,
  "linked_by": "Testname1",
  "name": "Test Compay1 Ltd",
  "smartgroup_name": "CompliantGroup1",
  "account_id": "iojzXIJWEuiD73ZvydOfg",
  "company": "Test Compay1 Ltd",
  "errors": [
    "OAuth token refresh failed, please re-link your account",
    "API daily rate limit reached for your account"
  ],
  "max_daily_api_requests": 5000,
  "client_id": "client_id0"
}
```

