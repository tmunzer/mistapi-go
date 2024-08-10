
# Account Jamf Config

OAuth linked Jamf apps account details

## Structure

`AccountJamfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | customer account api client id |
| `ClientSecret` | `string` | Required | customer account api client secret |
| `InstanceUrl` | `string` | Required | customer account Jamf instance URL |
| `SmartgroupName` | `string` | Required | smart group membership for determining compliance status |

## Example (as JSON)

```json
{
  "client_id": "client_id0",
  "client_secret": "client_secret6",
  "instance_url": "junipertest.jamfcloud.com",
  "smartgroup_name": "CompliantGroup1"
}
```

