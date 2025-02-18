
# Snmp Config V2 C Config

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpConfigV2cConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Authorization` | `*string` | Optional | - |
| `ClientListName` | `*string` | Optional | Client_list_name here should refer to client_list above |
| `CommunityName` | `*string` | Optional | - |
| `View` | `*string` | Optional | View name here should be defined in views above |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "authorization": "read-only",
  "client_list_name": "clist-1",
  "community_name": "abc123",
  "view": "all",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

