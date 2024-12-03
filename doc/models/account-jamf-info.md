
# Account Jamf Info

OAuth linked Jamf apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountJamfInfo`

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "client_id": "client_id6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

