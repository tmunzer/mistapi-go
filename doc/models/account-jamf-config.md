
# Account Jamf Config

OAuth linked Jamf apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountJamfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account api client id. Required if `app_name`==`crowdstrike` |
| `ClientSecret` | `string` | Required | Customer account api client secret |
| `InstanceUrl` | `string` | Required | Customer account Jamf instance URL |
| `SmartgroupName` | `string` | Required | Smart group membership for determining compliance status |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_id": "client_id0",
  "client_secret": "client_secret6",
  "instance_url": "junipertest.jamfcloud.com",
  "smartgroup_name": "CompliantGroup1",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

