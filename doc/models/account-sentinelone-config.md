
# Account Sentinelone Config

OAuth linked CrowdStrike apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountSentineloneConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiToken` | `string` | Required | Customer account api_token |
| `InstanceUrl` | `string` | Required | Customer account SentinelOne instance URL |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "api_token": "api_token6",
  "instance_url": "instance_url4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

