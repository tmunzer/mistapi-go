
# Snmp Vacm Security to Group Content Item

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpVacmSecurityToGroupContentItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Group` | `*string` | Optional | Refer to group_name under access |
| `SecurityName` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "group": "group4",
  "security_name": "security_name6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

