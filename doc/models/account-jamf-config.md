
# Account Jamf Config

OAuth linked Jamf apps account details

## Structure

`AccountJamfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account api client id. Required if `app_name`==`crowdstrike` |
| `ClientSecret` | `string` | Required | Customer account api client secret |
| `InstanceUrl` | `string` | Required | Customer account Jamf instance URL |
| `SmartgroupName` | `string` | Required | Smart group membership for determining compliance status |

## Example (as JSON)

```json
{
  "client_id": "client_id0",
  "client_secret": "client_secret6",
  "instance_url": "junipertest.jamfcloud.com",
  "smartgroup_name": "CompliantGroup1"
}
```

