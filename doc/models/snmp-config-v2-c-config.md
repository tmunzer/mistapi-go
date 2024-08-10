
# Snmp Config V2 C Config

## Structure

`SnmpConfigV2cConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Authorization` | `*string` | Optional | - |
| `ClientListName` | `*string` | Optional | client_list_name here should refer to client_list above |
| `CommunityName` | `*string` | Optional | - |
| `View` | `*string` | Optional | view name here should be defined in views above |

## Example (as JSON)

```json
{
  "authorization": "read-only",
  "client_list_name": "clist-1",
  "community_name": "abc123",
  "view": "all"
}
```

