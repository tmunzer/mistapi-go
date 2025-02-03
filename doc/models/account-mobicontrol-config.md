
# Account Mobicontrol Config

*This model accepts additional fields of type interface{}.*

## Structure

`AccountMobicontrolConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account Client ID |
| `ClientSecret` | `string` | Required | Customer account Client Secret |
| `InstanceUrl` | `string` | Required | Customer account MobiControl instance URL |
| `Password` | `string` | Required | Customer account password instance URL |
| `Username` | `string` | Required | Customer account username |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_id": "client_id0",
  "client_secret": "client_secret6",
  "instance_url": "instance_url2",
  "password": "password2",
  "username": "username2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

