
# Account Oauth Info Account

OAuth linked apps account info

*This model accepts additional fields of type interface{}.*

## Structure

`AccountOauthInfoAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `*string` | Optional | Linked app account id |
| `AutoProbeSubnet` | `*string` | Optional | For Prisma accounts only, tunnel auto probe subnet |
| `ClientId` | `*string` | Optional | Customer account Client ID |
| `CloudName` | `*string` | Optional | Name of the company whose account mist has subscribed to |
| `Company` | `*string` | Optional | Name of the company whose account mist has subscribed to |
| `EnableProbe` | `*bool` | Optional | For Prisma accounts only, tunnel probe enable/disable |
| `Error` | `*string` | Optional | This error is provided when the account fails to fetch token/data |
| `Errors` | `[]string` | Optional | - |
| `InstanceUrl` | `*string` | Optional | Customer account instance URL |
| `LastStatus` | `*string` | Optional | Is the last data pull for account is successful or not |
| `LastSync` | `*int64` | Optional | Last data pull timestamp, background jobs that pull account data |
| `LinkedBy` | `*string` | Optional | First name of the user who linked the account |
| `LinkedTimestamp` | `*float64` | Optional | - |
| `MaxDailyApiRequests` | `*int` | Optional | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ |
| `Name` | `*string` | Optional | Name of the company whose account mist has subscribed to |
| `Password` | `*string` | Optional | Customer account password instance URL |
| `Region` | `*string` | Optional | For Prisma accounts only |
| `ServiceAccountName` | `*string` | Optional | For Prisma accounts only |
| `ServiceConnections` | `[]string` | Optional | For Prisma accounts only |
| `SmartgroupName` | `*string` | Optional | Smart group membership for determining compliance status |
| `TsgId` | `*string` | Optional | For Prisma accounts only, Prisma Tenant Service Group id |
| `Username` | `*string` | Optional | Customer account username |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
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
  "service_connections": [
    "sc1",
    "sc2"
  ],
  "smartgroup_name": "CompliantGroup1",
  "tsg_id": "189953456",
  "client_id": "client_id6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

