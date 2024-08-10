
# Account Vmware Info

## Structure

`AccountVmwareInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccountId` | `*string` | Optional | - |
| `InstanceUrl` | `*string` | Optional | Linked VMware Instance URL |
| `LastStatus` | `*string` | Optional | Is the last data pull for VMware account is successful or not |
| `LastSync` | `*int` | Optional | Last data pull timestamp, background jobs that pull VMware account data |
| `LinkedBy` | `*string` | Optional | First name of the user who linked the VMware account |
| `LinkedTimestamp` | `*int` | Optional | This error is provided when the VMware account fails to fetch token/data |
| `Name` | `*string` | Optional | Name of the company whose VMware account mist has subscribed to |

## Example (as JSON)

```json
{
  "account_id": "account_id0",
  "instance_url": "instance_url2",
  "last_status": "last_status2",
  "last_sync": 142,
  "linked_by": "linked_by4"
}
```

