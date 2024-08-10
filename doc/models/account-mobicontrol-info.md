
# Account Mobicontrol Info

OAuth linked Mobicontrol apps account details

## Structure

`AccountMobicontrolInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `*string` | Optional | customer account client id |
| `ClientId` | `*string` | Optional | Linked MobiControl Client Id |
| `Error` | `*string` | Optional | This error is provided when the MobiControl account fails to fetch token/data |
| `InstanceUrl` | `*string` | Optional | Linked MobiControl Instance URL |
| `LastStatus` | `*string` | Optional | Is the last data pull for MobiControl account is successful or not |
| `LastSync` | `*int64` | Optional | Last data pull timestamp, background jobs that pull MobiControl account data |
| `LinkedBy` | `*string` | Optional | First name of the user who linked the MobiControl account |
| `LinkedTimestamp` | `*int64` | Optional | - |
| `Name` | `*string` | Optional | Name of the company whose MobiControl account mist has subscribed to |

## Example (as JSON)

```json
{
  "error": "Get token failed, please re-link MobiControl account",
  "instance_url": "https://a0032314.mobicontrol.cloud",
  "last_status": "failed",
  "last_sync": 1665465339000,
  "linked_by": "Testname1",
  "linked_timestamp": 1665465339000,
  "name": "Test Compay1 Ltd",
  "account_id": "account_id4",
  "client_id": "client_id4"
}
```

