
# Snmp Vacm Security to Group

*This model accepts additional fields of type interface{}.*

## Structure

`SnmpVacmSecurityToGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Content` | [`[]models.SnmpVacmSecurityToGroupContentItem`](../../doc/models/snmp-vacm-security-to-group-content-item.md) | Optional | - |
| `SecurityModel` | [`*models.SnmpVacmSecurityModelEnum`](../../doc/models/snmp-vacm-security-model-enum.md) | Optional | enum: `usm`, `v1`, `v2c` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "content": [
    {
      "group": "group2",
      "security_name": "security_name6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "security_model": "usm",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

