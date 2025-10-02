
# Account Oauth Config

OAuth linked apps (zoom/teams/intune) account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountOauthConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `string` | Required | Linked app(zoom/teams/intune) account id |
| `DiscardGuestInfo` | `*bool` | Optional | Optional, for Zoom/Teams. Whether to redact identifying information for call participants that are not part of the Zoom/Teams account identified by `account_id` |
| `MaxDailyApiRequests` | `*int` | Optional | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "account_id": "iojzXIJWEuiD73ZvydOfg",
  "max_daily_api_requests": 5000,
  "discard_guest_info": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

