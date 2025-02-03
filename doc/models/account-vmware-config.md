
# Account Vmware Config

*This model accepts additional fields of type interface{}.*

## Structure

`AccountVmwareConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account Client ID |
| `ClientSecret` | `string` | Required | Customer account Client Secret |
| `InstanceUrl` | `string` | Required | Customer account VMware instance URL |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_id": "client_id8",
  "client_secret": "client_secret4",
  "instance_url": "instance_url4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

