
# Account Oauth Config

OAuth linked apps (zoom/teams/intune) account details

## Structure

`AccountOauthConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `string` | Required | Linked app(zoom/teams/intune) account id |
| `MaxDailyApiRequests` | `*int` | Optional | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |

## Example (as JSON)

```json
{
  "account_id": "iojzXIJWEuiD73ZvydOfg",
  "max_daily_api_requests": 5000
}
```

