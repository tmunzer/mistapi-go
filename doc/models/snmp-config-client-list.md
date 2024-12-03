
# Snmp Config Client List

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpConfigClientList`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientListName` | `*string` | Optional | - |
| `Clients` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "client_list_name": "clist-1",
  "clients": [
    "clients0",
    "clients9",
    "clients8"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

